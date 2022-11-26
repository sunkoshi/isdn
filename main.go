package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/orted-org/isdn/internal/function_executor"
	"github.com/orted-org/isdn/internal/lang_handler"
)

func main() {
	file, err := os.Open("./test_codes/jsm/jsm.zip")
	if err != nil {
		log.Fatal(err)
	}

	lh, err := lang_handler.New()
	if err != nil {
		log.Fatal(err)
	}

	fe, err := function_executor.New(lh, function_executor.FunctionExecutorParams{
		RequestID: "1",
		Code:      file,
		Language:  "Javascript",
		IsZip:     true,
		Input:     function_executor.FunctionInput{},
	})
	if err != nil {
		log.Fatal(err)
	}

	output := fe.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
