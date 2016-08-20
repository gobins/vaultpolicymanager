package vaultpolicymanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "vaultpolicymanager"
)

var _ = Describe("Test Administrator Policy", func() {

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
		//  err = GetAuditHash(UserClient, "test", "test data")
		//  Expect(err).NotTo(HaveOccurred())
		// })

		It("should be able to disable an audit backend", func() {
			err = DisableAuditBackend(UserClient, "test")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to get leader node", func() {
			err = GetLeader(UserClient)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able step down a leader node", func() {
			err = StepDownLeader(UserClient)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to get current encryption key status", func() {
			err = GetKeyStatus(UserClient)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to rotate encryption key", func() {
			err = RotateKeys(UserClient)
			Expect(err).NotTo(HaveOccurred())
		})

		// It("should be able to renew a lease", func() {
		//  err = RenewLease(UserClient, lease_id, increment_duration)
		//  Expect(err).NotTo(HaveOccurred())
		// })

		// It("should be able to revoke a lease", func() {
		//  err = RevokeLease(UserClient, lease_id)
		//  Expect(err).NotTo(HaveOccurred())
		// })

		// It("should be able to revoke a prefix", func() {
		//  err = RevokePrefix(UserClient, prefix)
		//  Expect(err).NotTo(HaveOccurred())
		// })

		// It("should be able to force a revoke", func() {
		//  err = RevokeForce(UserClient, prefix)
		//  Expect(err).NotTo(HaveOccurred())
		// })

	})

})
