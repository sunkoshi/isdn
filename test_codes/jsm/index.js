const { sum } = require("./sum");
function handle(input) {
  jsonInput = JSON.parse(input);

  switch (jsonInput.operation) {
    case "sum":
      return JSON.stringify({ result: jsonInput.a + jsonInput.b });
    case "mul":
      return JSON.stringify({ result: jsonInput.a * jsonInput.b });

    case "sub":
      return JSON.stringify({ result: jsonInput.a - jsonInput.b });
    case "div":
      return JSON.stringify({ result: jsonInput.a / jsonInput.b });

    default:
      break;
  }
}

module.exports = handle;
