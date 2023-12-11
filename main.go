package main

import (
	"context"
	"fmt"

	"github.com/Apsodius1/orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start application: %s", err)
	}
}