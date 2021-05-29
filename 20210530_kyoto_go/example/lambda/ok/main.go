package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handle)
}

// Handle はアプリケーションのハンドラ
func Handle(ctx context.Context) error {
	// ...
	return nil
}
