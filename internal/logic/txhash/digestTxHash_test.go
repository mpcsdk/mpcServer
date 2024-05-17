package txhash

import "testing"

func TestTxHash(t *testing.T) {
	signStr := `
	{
        "chainId": 9527,
        "address": "0x4878c8FA20664D36F45F6f66Ea19F131276b8F2a",
        "number": 0,
        "txs": [
            {
                "_isUnipassWalletTransaction": true,
                "callType": 0,
                "revertOnError": true,
                "gasLimit": {
                    "type": "BigNumber",
                    "hex": "0x00"
                },
                "target": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
                "value": {
                    "type": "BigNumber",
                    "hex": "0x01"
                },
                "data": "0xa9059cbb000000000000000000000000aa5c1d42f766c98089a233ce1496bce18cfac5840000000000000000000000000000000000000000000000000000000000989680"
            }
        ],
        "txHash": "0x58816202f59de5a3e249898fe413a58cc3d0ecead33f82d288d3e7a28676eff3"
    }
	`
	//////
	txHash := DigestTxHash(signStr)
	if txHash != "0x6cc6f23f55e83864d24013a4e5b802eec2950013c8fa12130ffa6355f15b31e0" {
		t.Fatal("txHash is not equal 0x6cc6f23f55e83864d24013a4e5b802eec2950013c8fa12130ffa6355f15b31e0")
	}

}
