package v1alpha1

const (
	ChainConfigTemplate = `
# crypto
privkey = "45c56be699dca666191ad3446897e0f480da234da896270202514a0e1a587c3f"

# db config
data_path = "./devtools/chain/data"

[graphql]
listening_address = "0.0.0.0:8000"
graphql_uri = "/graphql"
graphiql_uri = "/graphiql"

[network]
listening_address = "0.0.0.0:1337"
rpc_timeout = 10

[[network.bootstraps]]
pubkey = "031288a6788678c25952eba8693b2f278f66e2187004b64ac09416d07f83f96d5b"
address = "0.0.0.0:1888"

[mempool]
pool_size = 20000
broadcast_txs_size = 200
broadcast_txs_interval = 200

[consensus]
public_keys = [ "04188ef9488c19458a963cc57b567adde7db8f8b6bec392d5cb7b67b0abc1ed6cd966edc451f6ac2ef38079460eb965e890d1f576e4039a20467820237cda753f07a8b8febae1ec052190973a1bcf00690ea8fc0168b3fbbccd1c4e402eda5ef22" ]

[executor]
light = false

[logger]
filter = "info"
log_to_console = true
console_show_file_and_line = false
log_path = "logs/"
log_to_file = true
metrics = true
# you can specify log level for modules with config below
# modules_level = { "overlord::state::process" = "debug", core_consensus = "error" }

[rocksdb]
max_open_files = 64
	`

	ChainGenesisTemplate = `
timestamp = 0
prevhash = "44915be5b6c20b0678cf05fcddbbaa832e25d7e6ac538784cd5c24de00d47472"

[[services]]
name = "asset"
payload = '{ "id": "f56924db538e77bb5951eb5ff0d02b88983c49c45eea30e8ae3e7234b311436c", "name": "MutaToken", "symbol": "MT", "supply": 320000011, "issuer": "f8389d774afdad8755ef8e629e5a154fddc6325a" }'

[[services]]
name = "metadata"
payload = '{ "chain_id": "b6a4d7da21443f5e816e8700eea87610e6d769657d6b8ec73028457bf2ca4036", "common_ref": "703873635a6b51513451", "timeout_gap": 20, "cycles_limit": 99999999, "cycles_price": 1, "interval": 3000, "verifier_list": [{"address": "f8389d774afdad8755ef8e629e5a154fddc6325a", "propose_weight": 1, "vote_weight": 1}], "propose_ratio": 15, "prevote_ratio": 10, "precommit_ratio": 10}'
	`
)
