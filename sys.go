package vaultpolicymanager

import (
	vaultapi "github.com/hashicorp/vault/api"
)

func GetLeader(c *vaultapi.Client) error {
	_, err := c.Logical().Read("/sys/leader")
	if err != nil {
		return err
	}
	return nil
}

func StepDownLeader(c *vaultapi.Client) error {
	_, err := c.Logical().Write("/sys/step-down", nil)
	if err != nil {
		return err
	}
	return nil
}

func GetKeyStatus(c *vaultapi.Client) error {
	_, err := c.Logical().Read("/sys/key-status")
	if err != nil {
		return err
	}
	return nil
}

func RotateKeys(c *vaultapi.Client) error {
	_, err := c.Logical().Write("/sys/rotate", nil)
	if err != nil {
		return err
	}
	return nil
}
