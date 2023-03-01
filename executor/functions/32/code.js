const fs = require("fs");
const path = require("path");
const handle = require("./index");

const inputFile = path.join(__dirname, "input.in");
const outputFile = path.join(__dirname, "output.out");

fs.readFile(inputFile, "utf8", (err, data) => {
  if (err) {
    console.log(err);
  } else {
    const result = handle(data);
    if (!result || result == "") result = {};
    fs.writeFile(outputFile, result, (err) => {
      if (err) {
        console.log(err);
      }
    });
  }
});
