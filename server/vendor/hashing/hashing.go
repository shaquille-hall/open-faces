//  Author: Shaquille Hall
//  Date: December 6, 2019
//  Title: Blockchain Hashing Logic in Golang
//
//  This package executes all blockchain hash operations:

package OpenFacesHashing

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func GenerateHash(nodeIndex int, data string, prevHash string) (string) {
	// Converting ints to string for conversion to []byte array. 
	newIndex := strconv.Itoa(nodeIndex)
	stringForHash := newIndex + data + prevHash
 	sum32 :=sha256.Sum256([]byte(stringForHash))
  sum := sum32[:]
  hash := hex.EncodeToString(sum)
  return hash
}

func GenerateHashAndNonce(nodeIndex int, data string, prevHash string) (int, string) {
	// Converting ints to string for conversion to []byte array. 
	newIndex := strconv.Itoa(nodeIndex)
	stringForHash := newIndex + data + prevHash
	nonce, hash := getNonceAndHash(stringForHash);
	return nonce, hash
}

func GetGenesisHash() string {
// The genesis hash starts our blockchain as the 'Previous Hash' for our first node.

	return strings.Repeat("0", 32)
}

func generateHashFromStringAndIndex(inputString string, potentialNonce int) string {
	potentialNonceString := strconv.Itoa(potentialNonce)
  newInputString := inputString + potentialNonceString
	sum32 :=sha256.Sum256([]byte(newInputString))
	sum := sum32[:]
	hash := hex.EncodeToString(sum)
	return hash
}

func isValidHash(hash string) bool {
	leadingZeros := strings.Repeat("0", 3)
	return strings.HasPrefix(hash, leadingZeros)
}

func getNonceAndHash(stringToHash string) (int, string) {
// Using proof of work, we generate our hash and nonce

	for i := 0; ; i++ {
		potentialHash := generateHashFromStringAndIndex(stringToHash, i)
		if isValidHash(potentialHash) {
			return i, potentialHash
		}
	}
}


