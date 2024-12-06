//go:build tools
// +build tools

package openhab

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
