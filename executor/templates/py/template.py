from index import handle


def main():
    inputFile = open("input.in", "r")
    input = inputFile.read()
    inputFile.close()
    result = handle(input)

    outputFile = open("output.out", "w")
    outputFile.write(result)
    outputFile.close()


main()