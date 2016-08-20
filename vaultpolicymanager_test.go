package vaultpolicymanager_test

import (
	. "vaultpolicymanager"

	vaultapi "github.com/hashicorp/vault/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	AdminClient *vaultapi.Client
	err         error
	UserClient  *vaultapi.Client
)

const ADMINISTRATOR_POLICY = "admin"
const USERPASS_TEST_USER = "test_user"
const USERPASS_TEST_PASSWORD = "password"

var _ = Describe("VaultPolicyManager", func() {

	BeforeSuite(func() {
		AdminClient, err = InitializeAdminClient("userpass")
		CheckError(err)
		err = CreateUser(AdminClient, USERPASS_TEST_USER, USERPASS_TEST_PASSWORD)
		CheckError(err)

		//General Admin Tests
		err = CreateUpdatePolicy(AdminClient, ADMINISTRATOR_POLICY, "policies/sys.hcl")
		CheckError(err)
		err = ApplyUserPolicy(AdminClient, USERPASS_TEST_USER, ADMINISTRATOR_POLICY)
		CheckError(err)
		UserClient, err = InitializeUserPassClient(USERPASS_TEST_USER, USERPASS_TEST_PASSWORD)
		CheckError(err)

	})

	var _ = AfterSuite(func() {
		err = DeleteUser(AdminClient, USERPASS_TEST_USER)
		err = DeletePolicy(AdminClient, ADMINISTRATOR_POLICY)
		CheckError(err)
	})

	Describe("Test Connection to Vault", func() {
		Context("With default environment variables", func() {
			It("should initialize admin vault client", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

})
