var HDWalletProvider = require("truffle-hdwallet-provider");

var mnemonic = "hazard decade force acid harbor praise wagon dragon they increase build drink";

module.exports = {
    // See <http://truffleframework.com/docs/advanced/configuration>
    // to customize your Truffle configuration!
    // contracts_build_directory: "./output",
    networks: {
        local: {
            host: "localhost",
            port: 8545,
            network_id: "*", // Match any network id
            gas: 6721975,
            provider: () =>
                new HDWalletProvider(mnemonic, "http://localhost:8545", 0, 10)
        }
    },

    compilers: {
        solc: {
            settings: {          // See the solidity docs for advice about optimization and evmVersion
                optimizer: {
                    enabled: true,
                    runs: 200
                }
            }
        }
    }
};
