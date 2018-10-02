package apis

import "github.com/kubermatic/cluster-api-provider-digitalocean/pkg/apis/providerconfig/v1alpha1"

// Register ProviderConfig resource with the Scheme, so components can map objects to GroupVersionKind and back.
func init() {
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
