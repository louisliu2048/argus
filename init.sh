KEY="mykey"
CHAINID="argus_9000-1"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
rm -rf ~/.argusd*
make install

# Set client config
argusd config keyring-backend $KEYRING
argusd config chain-id $CHAINID

# if $KEY exists it should be deleted
argusd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for argus (Moniker can be anything, chain-id must be an integer)
argusd init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to aargus
cat $HOME/.argusd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aargus"' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json
cat $HOME/.argusd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aargus"' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json
cat $HOME/.argusd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aargus"' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json
cat $HOME/.argusd/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="aargus"' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json

# Set gas limit in genesis
cat $HOME/.argusd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.argusd/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.argusd/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.argusd/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.argusd/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.argusd/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.argusd/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
argusd add-genesis-account $KEY 100000000000000000000000000aargus --keyring-backend $KEYRING
argusd add-genesis-account 0x97DA0F6C071C051127B92Da941e259B1104c6a8F 100000000000000000000000000aargus

#total_supply=100000000000000000000000000
#cat $HOME/.argusd/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.argusd/config/tmp_genesis.json && mv $HOME/.argusd/config/tmp_genesis.json $HOME/.argusd/config/genesis.json

# Sign genesis transaction
argusd gentx $KEY 1000000000000000000000aargus --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
argusd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
argusd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
argusd start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001aargus --json-rpc.api eth,txpool,personal,net,debug,web3
