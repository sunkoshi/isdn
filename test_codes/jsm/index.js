const { sum } = require("./sum");
const { v4 } = require("uuid");
function handle(input) {
  jsonInput = JSON.parse(input);
  result = sum(jsonInput.a, jsonInput.b);
  return JSON.stringify({ sum: result, uuid: v4() });
}

module.exports = handle;
