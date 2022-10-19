package main

import (
	"os"

	"gitlab.com/tokend/nft-books/price-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
