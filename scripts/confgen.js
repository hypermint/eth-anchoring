var Blocks = artifacts.require("./Blocks.sol")
var fs = require("fs");
var ejs = require("ejs");
var targets = [
    // dst, src
    ['./pkg/consts/gen.go', './scripts/template/gen.go.tpl']
];

module.exports = function(callback) {
  targets.forEach(function(item) {
    ejs.renderFile(item[1], {
        BlocksAddress: Blocks.address
      }, null, function(err, str){
        if (err) {
          throw err;
        }
        fs.writeFileSync(item[0], str);
        console.log('generate', item[0]);
      });
  });
  callback();
};
