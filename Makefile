NETWORK ?= "local"
TRUFFLE ?= npm run truffle --
ABIGEN  ?= abigen

MNEMONIC ?= hazard decade force acid harbor praise wagon dragon they increase build drink
HDW_PATH ?= m/44'/60'/0'/0
GASLIMIT ?= 10000000

.PHONY: build abi migrate

build:
	go build -o build/anchor ./cmd/anchor

abi:
ifdef SOURCE
	$(eval TARGET := $(shell echo ${SOURCE} | tr A-Z a-z))
	@mkdir -p ./build/abi ./pkg/contract
	@mkdir -p ./pkg/contract/$(TARGET)
	@cat ./build/contracts/${SOURCE}.json | jq ".abi" > ./build/abi/${SOURCE}.abi
	$(ABIGEN) --abi ./build/abi/${SOURCE}.abi --pkg $(TARGET) --out ./pkg/contract/$(TARGET)/$(TARGET).go
else
	@echo "'SOURCE={SOURCE}' is required"
endif

migrate:
	$(TRUFFLE) migrate --reset --network=$(NETWORK)

config:
	$(TRUFFLE) exec ./scripts/confgen.js --network=$(NETWORK)

deploy:
	$(MAKE) migrate
	$(MAKE) abi SOURCE=Blocks
	$(MAKE) config

run:
	./build/anchor run --mnemonic="$(MNEMONIC)" --hdw-path="$(HDW_PATH)/0"
