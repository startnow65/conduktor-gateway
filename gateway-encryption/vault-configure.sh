#run the vault in dev mode
vault server -dev -dev-root-token-id root

# export the vault server address
export VAULT_ADDR="http://127.0.0.1:8200"

# export the vault token
export VAULT_TOKEN=root
# enable the transit engine
vault secrets enable -path=encryption transit

# enable kv engine
vault secrets enable -version=1 kv
#create the gateway keys path
vault write -f encryption/keys/gateway