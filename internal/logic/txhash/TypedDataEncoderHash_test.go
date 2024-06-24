package txhash

import "testing"

func TestTypedDataEncoderHash(t *testing.T) {
	signStr := `
		{
		    "types": {
		        "EIP712Domain": [
		            { "name": "name", "type": "string" },
		            { "name": "version", "type": "string" },
		            { "name": "chainId", "type": "uint256" },
		            { "name": "verifyingContract", "type": "address" }
		        ],
		        "Person": [
		            { "name": "name", "type": "string" },
		            { "name": "wallet", "type": "address" }
		        ],
		        "Mail": [
		            { "name": "from", "type": "Person" },
		            { "name": "to", "type": "Person" },
		            { "name": "contents", "type": "string" }
		        ]
		    },
		    "primaryType": "Mail",
		    "domain": {
		        "name": "Ether Mail",
		        "version": "1",
		        "chainId": 1,
		        "verifyingContract": "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC"
		    },
		    "message": {
		        "from": {
		            "name": "Cow",
		            "wallet": "0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826"
		        },
		        "to": {
		            "name": "Bob",
		            "wallet": "0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB"
		        },
		        "contents": "Hello, Bob!"
		    }
		}
	`
	//////
	dataHash, _ := TypedDataEncoderHash(signStr)
	if dataHash != "0xbbbe8d81936cca96cd0cb08dae75780921b901ecc8193d83672c4f919ef36660" {
		t.Fatal("TypedDataEncoderHash is not equal 0xbbbe8d81936cca96cd0cb08dae75780921b901ecc8193d83672c4f919ef36660")
	}

}
