package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

func VerifyAddress(address string) bool {
	if common.IsHexAddress(address) {
		return true
	} else {
		return false
	}
}
