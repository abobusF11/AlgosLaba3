package main

import (
	"fmt"
	"os"

	"algosLaba3/internal/delivery/cli"
	"algosLaba3/internal/usecase"
)

func main() {
	validator := usecase.NewValidator(
		usecase.WithAngleBrackets(),
		usecase.WithIgnoreOthers(true),
	)

	handler := cli.NewHandler(validator)

	if err := handler.Run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
