package vaultpolicymanager

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func GetCapability(c *vaultapi.Client, token, path string) error {
	data := map[string]interface{}{"token": token, "path": path}

	_, err := c.Logical().Write("sys/capabilities", data)
	if err != nil {
		return err
	}
	return nil
}

func GetSelfCapability(c *vaultapi.Client, path string) error {
	data := map[string]interface{}{"path": path}

	_, err := c.Logical().Write("sys/capabilities-self", data)
	if err != nil {
		return err
	}
	return nil
}
