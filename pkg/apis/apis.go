// Package apis contains Kubernetes API Groups for DigitalOcean Cluster-API provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
)

//go:generate go run ../../vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go -O zz_generated.deepcopy -i ./...

// AddToSchemes is used to all resources used in the project to a Scheme.
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds a resource to the Schemes.
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}