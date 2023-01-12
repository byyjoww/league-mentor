//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // Fast Go linters runner
	_ "github.com/t-yuki/gocover-cobertura"                 // Go coverage report to cobertura for gitlab
	_ "golang.org/x/lint/golint"                            // Linter for Go source code
	_ "golang.org/x/tools/cmd/goimports"                    // Generates k8s manifests
)
