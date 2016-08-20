package vaultpolicymanager

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func RenewLease(c *vaultapi.Client, lease_id, increment_duration string) error {
	data := map[string]interface{}{"increment": increment_duration}
	_, err := c.Logical().Write("sys/renew/"+lease_id, data)
	if err != nil {
		return err
	}
	return nil
}

func RevokeLease(c *vaultapi.Client, lease_id string) error {
	_, err := c.Logical().Write("sys/revoke/"+lease_id, nil)
	if err != nil {
		return err
	}
	return nil
}

func RevokePrefix(c *vaultapi.Client, path string) error {
	_, err := c.Logical().Write("sys/revoke-prefix/"+path, nil)
	if err != nil {
		return err
	}
	return nil
}

func RevokeForce(c *vaultapi.Client, path string) error {
	_, err := c.Logical().Write("sys/revoke-force/"+path, nil)
	if err != nil {
		return err
	}
	return nil
}
