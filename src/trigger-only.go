package main

import (
	"context"

	"github.com/kafkaesque-io/pulsar-go-function-template/src/pf"
)

// HandleExclamation handles exlamation
func HandleExclamation(ctx context.Context, in []byte) ([]byte, error) {
	return []byte(string(in) + "!"), nil
}

func main() {
	pf.Start(HandleExclamation)
}