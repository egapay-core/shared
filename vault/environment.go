package vault

// ClientEnvironment - the environment of the client
// Host - the host of the client (URL of the client, e.g. https://client.vault.com)
// Token - the token of the client (JWT of the vault)
// Vault - the vault of the client (ID of the vault)
type ClientEnvironment struct {
	Host  string `env:"OP_CONNECT_HOST, required"`
	Token string `env:"OP_CONNECT_TOKEN, required"`
	Vault string `env:"OP_VAULT, required"`
}
