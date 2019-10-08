#!/bin/bash
set -e

# init validator
/hmd tendermint init-validator --home="${HMD_HOME}" --mnemonic="${MNEMONIC}" --hdw_path="${HDW_PATH}/${HDW_VALIDATOR_IDX}"

# create a new account
ADDR1=$(/hmcli new --password=password --silent --home="${HMCLI_HOME}" --mnemonic="${MNEMONIC}" --hdw_path="${HDW_PATH}/1")
ADDR2=$(/hmcli new --password=password --silent --home="${HMCLI_HOME}" --mnemonic="${MNEMONIC}" --hdw_path="${HDW_PATH}/2")
/hmd init --address=${ADDR1} --home=${HMD_HOME}
