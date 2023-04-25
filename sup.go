package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress("CONTRACT-ADDRESS-HERE")
	instance, err := NewRandomNumber(contractAddress, client)
	if err != nil {
		panic(err)
	}

	opts := &bind.TransactOpts{
		From: common.HexToAddress("YOUR-ADDRESS-HERE"),
	}

	randomNumber, err := instance.GetRandNumber(opts)
	if err != nil {
		panic(err)
	}

	fmt.Println(randomNumber)
}
