BATS ?= ./tools/bats/bin/bats
ANCHOR ?= ../../build/anchor
MNEMONIC ?= hazard decade force acid harbor praise wagon dragon they increase build drink
HDW_PATH ?= m/44'/60'/0'/0

.PHONY: setup tools test

setup:
	docker-compose up -d
	$(MAKE) -C ../../ deploy build

tools:
	@mkdir -p tools
	cd tools && git clone https://github.com/sstephenson/bats.git

test:
	ANCHOR=$(ANCHOR) MNEMONIC="$(MNEMONIC)" HDW_PATH="$(HDW_PATH)/0" $(BATS) run.bats
