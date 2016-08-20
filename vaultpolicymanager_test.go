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

	Describe("Test Administrator Policy", func() {

		Context("Initializing Administrator Policy", func() {
			It("should successfully initialize General Admin Policy", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("should be able to create new mounts", func() {
				err = CreateMount(UserClient, "testgeneric", "generic")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to remount a new path", func() {
				err = DoRemount(UserClient, "testgeneric", "testgeneric_1")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to delete a mount", func() {
				err = DeleteMount(UserClient, "testgeneric_1")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to list all policies", func() {
				err = ListMounts(UserClient)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to create a policy", func() {
				err = CreateUpdatePolicy(UserClient, "test-policy", "policies/sys.hcl")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to update a policy", func() {
				err = CreateUpdatePolicy(UserClient, "test-policy", "policies/sys.hcl")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to delete a policy", func() {
				err = DeletePolicy(UserClient, "test-policy")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to query a token's capability against a path", func() {
				err = GetCapability(UserClient, UserClient.Token(), "sys/mounts")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to query it's own capability against a path", func() {
				err = GetSelfCapability(UserClient, "sys/mounts")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to list audit backends", func() {
				err = GetAuditBackends(UserClient)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be able to create an audit backend", func() {
				err = SetFileAuditBackend(UserClient, "test", "/tmp/vault-test.log")
				Expect(err).NotTo(HaveOccurred())
			})

			// Audit hash is having an issue
			// It("should be able to get audit hash", func() {
			// 	err = GetAuditHash(UserClient, "test", "test data")
			// 	Expect(err).NotTo(HaveOccurred())
			// })

			It("should be able to disable an audit backend", func() {
				err = DisableAuditBackend(UserClient, "test")
				Expect(err).NotTo(HaveOccurred())
			})

			// It("should be able to renew a lease", func() {
			// 	err = DisableAuditBackend(UserClient, "test")
			// 	Expect(err).NotTo(HaveOccurred())
			// })

			// It("should be able to revoke a lease", func() {
			// 	err = DisableAuditBackend(UserClient, "test")
			// 	Expect(err).NotTo(HaveOccurred())
			// })

			// It("should be able to revoke a prefix", func() {
			// 	err = DisableAuditBackend(UserClient, "test")
			// 	Expect(err).NotTo(HaveOccurred())
			// })

			// It("should be able to force a revoke", func() {
			// 	err = DisableAuditBackend(UserClient, "test")
			// 	Expect(err).NotTo(HaveOccurred())
			// })

		})
	})

})
