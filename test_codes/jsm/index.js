const { sum } = require("./sum");
function handle(input) {
  jsonInput = JSON.parse(input);
  result = sum(jsonInput.a, jsonInput.b);
  return JSON.stringify({ sum: result });
}

module.exports = handle;
