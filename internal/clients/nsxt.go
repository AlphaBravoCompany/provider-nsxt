/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/AlphaBravoCompany/provider-nsxt/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal nsxt credentials as JSON"
)

var reqFields = []string{
	"host",
	"username",
	"password",
}

var optFields = []string{
	"client_auth_cert_file",
	"client_auth_key_file",
	"client_auth_cert",
	"client_auth_key",
	"allow_unverified_ssl",
	"ca_file",
	"ca",
	"max_retries",
	"retry_min_delay",
	"retry_max_delay",
	"retry_on_status_codes",
	"remote_auth",
	"session_auth",
	"tolerate_partial_success",
	"vmc_token",
	"vmc_client_id",
	"vmc_client_secret",
	"vmc_auth_host",
	"vmc_auth_mode",
	"enforcement_point",
	"global_manager",
	"license_keys",
	"on_demand_connection",
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]any{}
		// Required fields
		for _, req := range reqFields {
			ps.Configuration[req] = creds[req]
		}
		// Optional fields
		for _, opt := range optFields {
			if v, ok := creds[opt]; ok {
				ps.Configuration[opt] = v
			}
		}

		return ps, nil
	}
}
