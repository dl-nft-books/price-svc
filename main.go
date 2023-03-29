package main

import (
	"os"

	"github.com/dl-nft-books/price-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
