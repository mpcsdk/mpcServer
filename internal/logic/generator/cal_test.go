package generator

import (
	"strings"
	"testing"
)

func TestHashMsg(t *testing.T) {
	msg := "abhdhd"
	s := &sGenerator{}
	// b, _ := hex.DecodeString(msg)
	hash := s.hashMessage(nil, strings.Trim(s.hashMsg(nil, msg), "0x"))
	if hash != "0xc561ef1578395650761751a0c411b69354627e29a9503ee663f843db8e16d2c7" {
		t.Error(hash)
	}
	///
}

func TestHashHexMsg(t *testing.T) {
	msg := "9a67af39e089e377bb94f236fd7deeffae283f615e039dbaa22b1084e1e2f008"
	s := &sGenerator{}
	// b, _ := hex.DecodeString(msg)
	hash := s.hashMessage(nil, msg)
	if hash != "0x2183ec69418545ba73baf5352043c52ef5ce8b09e668776cc6386cdb60feaccb" {
		t.Error(hash)
	}
	///
}

func TestHashDomain(t *testing.T) {
	msg := "0xfa6f4f7d2bf9d0892b232102908d97315c9684d9551b572bc977484b3a3cc913"
	s := &sGenerator{}
	hash := s.hashMessage(nil, msg)
	if hash != "0x37a10e3f652ae9ed2b58829345304a2b8bdbeb7cae4c192e1fe5f8668b552548" {
		t.Error(hash)
	}
	///
}
