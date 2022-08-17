package main

import (
	"fmt"
	"os"

	"github.com/ardanlabs/service/foundation/logger"
)

func main() {
	// Construct the application logger.
	log, err := logger.New("SHOP-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer log.Sync()
}
