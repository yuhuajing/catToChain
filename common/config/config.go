package config

import (
	"main/common/ethconn"
)

var (
	// Owner:0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f
	//797391c7bd2e156e52329ceb6471496798e0c125ef35c4c3393329bd2a64f3f5
	EthServer = "https://eth-sepolia.g.alchemy.com/v2/t39LUOfkEMNMz_uxQk2gkwK3kJ1HwDZF"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0x467cFb821B45eE9B30FB684c1C116942a92B9452"
)

// 浏览器
// https://sepolia.etherscan.io/address/0x467cFb821B45eE9B30FB684c1C116942a92B9452
