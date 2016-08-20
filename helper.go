package vaultpolicymanager

import (
	"fmt"
	. "github.com/gobins/vaultutil"
	vaultapi "github.com/hashicorp/vault/api"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitializeAdminClient(auth_type string) (*vaultapi.Client, error) {
	config := vaultapi.DefaultConfig()
	config.ReadEnvironment()
	c, _ := vaultapi.NewClient(config)

	authclient, err := GetVaultClient(auth_type)

	if err != nil {
		fmt.Println("Error Initializing client:", err.Error())
		return c, err
	}

	_, err = authclient.Authenticate()

	if err != nil {
		fmt.Println("Error Authenticating client:", err.Error())
		return c, err
	}
	token := authclient.GetToken()

	c.SetToken(token)
	return c, nil
}

func InitializeUserPassClient(username, password string) (*vaultapi.Client, error) {
	config := vaultapi.DefaultConfig()
	config.ReadEnvironment()
	c, _ := vaultapi.NewClient(config)
	user_token, err := GetUserToken("userpass", "test_user", "password")

	if err != nil {
		fmt.Println("Error Initializing Userpass client:", err.Error())
		return c, err
	}
	c.SetToken(user_token)
	return c, nil
}

func CreateUser(c *vaultapi.Client, username, password string) error {
	data := map[string]interface{}{"password": password}
	_, err := c.Logical().Write("auth/userpass/users/"+username, data)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(c *vaultapi.Client, username string) error {
	_, err := c.Logical().Delete("auth/userpass/users/" + username)

	if err != nil {
		return err
	}
	return nil
}

func GetUserToken(auth_type, username, password string) (string, error) {
	var token string
	os.Setenv("VAULT_USER", username)
	os.Setenv("VAULT_PASSWORD", password)
	client, err := GetVaultClient(auth_type)
	if err != nil {
		return token, err
	}
	_, err = client.Authenticate()
	if err != nil {
		return token, err
	}
	token = client.GetToken()
	return token, err
}
