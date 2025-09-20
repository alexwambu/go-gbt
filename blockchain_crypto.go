package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey, &privateKey.PublicKey
}

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func SignMessage(priv *ecdsa.PrivateKey, msg []byte) ([]byte, []byte) {
	hash := sha256.Sum256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	return r.Bytes(), s.Bytes()
}

func VerifySignature(pub *ecdsa.PublicKey, msg, rBytes, sBytes []byte) bool {
	hash := sha256.Sum256(msg)
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)
	return ecdsa.Verify(pub, hash[:], r, s)
}

func AddressFromPub(pub *ecdsa.PublicKey) string {
	pubBytes := append(pub.X.Bytes(), pub.Y.Bytes()...)
	hash := Keccak256(pubBytes)
	return "0x" + hex.EncodeToString(hash[12:]) // Ethereum style address
}
