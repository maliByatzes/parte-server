package util

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
)

type Token struct {
	Secret string
	Hash   string
}

func GenerateOTP() (*Token, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return nil, err
	}

	sDNum := bigInt.Int64() + 100000

	sDstr := fmt.Sprintf("%06d", sDNum)

	token := Token{
		Secret: sDstr,
	}

	hash := sha256.Sum256([]byte(token.Secret))

	token.Hash = fmt.Sprintf("%x\n", hash)

	return &token, nil
}

func FormatOTP(otp string) string {
	length := len(otp)
	half := length / 2
	fHalf := otp[:half]
	sHalf := otp[half:]
	words := []string{fHalf, sHalf}
	return strings.Join(words, " ")
}
