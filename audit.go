package vaultpolicymanager

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func GetAuditBackends(c *vaultapi.Client) error {
	_, err := c.Logical().Read("sys/audit")
	if err != nil {
		return err
	}
	return nil
}

func SetFileAuditBackend(c *vaultapi.Client, path, file_path string) error {
	opts := map[string]interface{}{"path": file_path}
	data := map[string]interface{}{"type": "file", "options": opts}

	_, err := c.Logical().Write("sys/audit/"+path, data)
	if err != nil {
		return err
	}
	return nil
}

func DisableAuditBackend(c *vaultapi.Client, audit_path string) error {
	_, err := c.Logical().Delete("sys/audit/" + audit_path)
	if err != nil {
		return err
	}
	return nil
}

func GetAuditHash(c *vaultapi.Client, path, data string) error {
	input := map[string]interface{}{"input": data}
	_, err := c.Logical().Write("sys/audit/"+path, input)
	if err != nil {
		return err
	}
	return nil
}
