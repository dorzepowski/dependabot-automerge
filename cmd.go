package main

import (
	"log"

	"github.com/BuxOrg/go-buxclient"
	"github.com/bitcoinschema/go-bitcoin/v2"
)

func main() {
	// Example xPub
	masterKey, _ := bitcoin.GenerateHDKey(bitcoin.SecureSeedLength)

	// Create a client
	buxClient, err := buxclient.New(
		buxclient.WithXPriv(masterKey.String()),
		buxclient.WithHTTP("localhost:3001"),
		buxclient.WithDebugging(true),
		buxclient.WithSignRequest(true),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("client loaded - bux debug: %v", buxClient.IsDebug())
}
