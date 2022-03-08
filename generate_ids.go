package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/hydra-booster/idgen"
)

func main() {

	idCount := 1
	if len(os.Args) > 1 {
		idCount, _ = strconv.Atoi(os.Args[1])
	}
	// fmt.Fprintf(os.Stderr, "🐉 Generating %d IDs\n", idCount)

	idGenerator := idgen.NewBalancedIdentityGenerator()

	for i := 0; i < idCount; i++ {
		priv, err := idGenerator.AddBalanced()
		// priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)

		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate Peer ID: %w", err))
		}
		pub := priv.GetPublic()
		b, err := crypto.MarshalPrivateKey(priv)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to extract private key bytes: %w", err))
		}
		privStr := base64.StdEncoding.EncodeToString(b)
		id, _ := peer.IDFromPublicKey(pub)
		idStr := id.Pretty()
		fmt.Println(idStr)
		fmt.Println(privStr)
	}

}
