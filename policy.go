package vaultpolicymanager

import (
	vaultapi "github.com/hashicorp/vault/api"
	"io/ioutil"
	"path/filepath"
)

func ApplyGroupPolicy(c *vaultapi.Client, group_name, policy_name string) error {
	data := map[string]interface{}{"policies": policy_name}
	_, err := c.Logical().Write("auth/ldap/groups/"+group_name, data)
	if err != nil {
		return err
	}
	return nil
}

func ApplyUserPolicy(c *vaultapi.Client, user_name, policy_name string) error {
	data := map[string]interface{}{"policies": policy_name}
	_, err := c.Logical().Write("auth/userpass/users/"+user_name, data)
	if err != nil {
		return err
	}
	return nil
}

func RevokeGroupPolicy(c *vaultapi.Client, group_name string) error {
	data := map[string]interface{}{"policies": ""}
	_, err := c.Logical().Write("auth/ldap/groups/"+group_name, data)
	if err != nil {
		return err
	}
	return nil
}

func RevokeUserPolicy(c *vaultapi.Client, user_name string) error {
	data := map[string]interface{}{"policies": ""}
	_, err := c.Logical().Write("auth/userpass/users/"+user_name, data)
	if err != nil {
		return err
	}
	return nil
}

func CreateUpdatePolicy(c *vaultapi.Client, policy_name, policy_file string) error {
	absPath, _ := filepath.Abs(policy_file)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}
	err = c.Sys().PutPolicy(policy_name, string(data))
	if err != nil {
		return err
	}
	return nil
}

func DeletePolicy(c *vaultapi.Client, policy_name string) error {
	err := c.Sys().DeletePolicy(policy_name)
	if err != nil {
		return err
	}
	return nil
}

func ReadPolicy(c *vaultapi.Client, policy_name string) error {
	_, err := c.Sys().GetPolicy(policy_name)
	if err != nil {
		return err
	}
	return nil
}
