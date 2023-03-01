const { sum } = require("./sum");
const { v4 } = require("uuid");
function handle(input) {
  jsonInput = JSON.parse(input);

  if (jsonInput.doSum) {
    result = sum(jsonInput.a, jsonInput.b);
    return JSON.stringify({ sum: result });
  }

  if (jsonInput.getUUID) {
    return JSON.stringify({ uuid: v4() });
  }
}

module.exports = handle;
