package main

import (
	"fmt"

	"github.com/dashevo/dashd-go/btcjson"
)

func main() {
	a := btcjson.GetBlockVerboseTxResult{
		Chainlock: true,
	}
	fmt.Println(a)
}
