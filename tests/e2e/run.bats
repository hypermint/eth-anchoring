#!/usr/bin/env bats

@test "current state is uninitialized" {
  run ${ANCHOR} l1 latest --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]
  [ "$output" -eq 0 ]
}

@test "on-shot submission is success" {
  run ${ANCHOR} submit --height=1 --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]
}

@test "submission includes duplicated height is reverted" {
  run ${ANCHOR} submit --height=1 --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 1 ]
}

@test "height=2 submission is success" {
  run ${ANCHOR} submit --height=2 --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]
}

@test "submitted block hash is equal to expected value" {
  run ${ANCHOR} l1 verify --height=1 --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]

  run ${ANCHOR} l1 verify --height=2 --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]
}

@test "current state is valid" {
  run ${ANCHOR} l1 latest --mnemonic="${MNEMONIC}" --hdw-path="${HDW_PATH}"
  [ "$status" -eq 0 ]
  [ "$output" -eq 2 ]
}
