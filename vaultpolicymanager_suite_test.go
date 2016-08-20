package vaultpolicymanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVaultPolicyManagerSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultPolicyManager Suite")
}
