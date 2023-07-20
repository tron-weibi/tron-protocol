package api

import (
	"context"
	"github.com/tien202/tron-protocol/common/base58"
	"github.com/tien202/tron-protocol/core"
	"log"
	"testing"
)

func TestAccountResourceMessage_GetEnergyUsed(t *testing.T) {
	//walletClient := GetNewWalletClient()
	//walletClient
}

func TestGetAccount(t *testing.T) {
	walletClient := GetNewWalletClient()
	acc := new(core.Account)
	addr := "TZ5dPxnxd4rRZb7nudcorifD9zfxi2NSRY"
	acc.Address, _ = base58.Decode(addr)
	result, err := walletClient.GetAccount(context.Background(), acc)
	if err != nil {
		log.Fatalf("can not get result of %s", addr)
	}

	//contract := new(contract.Acc)
	log.Println("result:", result)
}
