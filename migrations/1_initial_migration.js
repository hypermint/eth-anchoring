
const Migrations = artifacts.require("Migrations");
const Blocks = artifacts.require("./Blocks.sol");

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Migrations);
  deployer.deploy(Blocks, accounts[0]);
};
