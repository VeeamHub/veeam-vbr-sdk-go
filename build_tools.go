//go:build build_tools

// While it's typically not valid to import a "main" package, it's being done
// here in a file that's never compiled (note the "build_tools" build tag above)
// in order to prevent `go mod tidy` from removing the module from the go.mod
// and go.sum files.
//
// This is to ensure that when the Makefile generates the client, the generator
// used is from the version of the module that is referenced by go.mod.

package main

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
