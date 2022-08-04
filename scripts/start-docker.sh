#!/bin/bash

KEY="mykey"
CHAINID="argus_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t argus-datadir.XXXXX)

echo "create and add new keys"
./argusd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Argus with moniker=$MONIKER and chain-id=$CHAINID"
./argusd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./argusd add-genesis-account \
"$(./argusd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aargus,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./argusd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./argusd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./argusd validate-genesis --home $DATA_DIR

echo "starting argus node $i in background ..."
./argusd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started argus node"
tail -f /dev/null
