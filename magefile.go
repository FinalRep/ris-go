//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type (
	// Test namespace
	Test mg.Namespace
)

// Run - run tests
func (Test) Run() error {
	return sh.RunV("go", "test", "-v", "-cover", "./...", "-coverprofile=coverage.out")
}

// Cover - checking code coverage
func (t Test) Cover() error {
	mg.Deps(t.Run)
	return sh.RunV("go", "tool", "cover", "-html=coverage.out")
}
