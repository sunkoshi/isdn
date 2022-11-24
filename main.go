package main

import (
	"context"
	"fmt"
	"log"

	"github.com/orted-org/dambda/pkg/command_executor"
)

func main() {

	o, err := command_executor.ExecuteContext(context.Background(), command_executor.SH_SHELL, "-c ls -la")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(o)
}
