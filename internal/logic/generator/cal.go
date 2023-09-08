package generator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/model"
	"li17server/internal/service"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GenContextP2
func (s *sGenerator) GenContextP2(ctx context.Context, sid string, private_key2, public_key string, submit bool) error {
	if submit {
		s.pool.Submit(func() {
			s.genContextP2(s.ctx, sid, private_key2, public_key)
		})
	} else {
		s.genContextP2(s.ctx, sid, private_key2, public_key)
	}
	return nil
}

// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and context_p2
// func (s *sGenerator) CalZKProof2(ctx context.Context, token string, zk_proof1 string) (err error) {

// 	s.pool.Submit(func() {
// 		s.calZKProof2(s.ctx, token, zk_proof1)
// 	})

// 	return
// }

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
func (s *sGenerator) CalZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	s.pool.Submit(func() {
		s.calZKProofP2(s.ctx, sid, p1_hash_proof)
	})

	return nil
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sGenerator) CalPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {

	s.pool.Submit(func() {
		s.calPublicKey2(s.ctx, sid, p1_zk_proof)
	})

	return nil
}

// 8.calculate request, recal context_p2
func (s *sGenerator) CalRequest(ctx context.Context, sid string, request string) error {
	s.pool.Submit(func() {
		s.pool.Submit(func() {
			s.calRequest(s.ctx, sid, request)
		})
	})

	return nil
}

// 9.signature
func (s *sGenerator) CalSign(ctx context.Context, req *v1.SignMsgReq, checkRule bool) error {
	analzytx := &model.AnalzyTx{}
	var err error
	if len(req.Msg) > 10 {
		// , sid string, msg string, request string, signData string, checkRule bool) error {
		signtx := &model.SignTx{}
		json.Unmarshal([]byte(req.SignData), signtx)
		///
		if checkRule {
			//todo: exec txs rules
			rst, err := service.Rule().Exec(signtx.Address, signtx.Txs)
			g.Log().Info(ctx, "Rule().Exec:", rst, signtx.Address, signtx.Txs)
			///
			if err != nil || rst != nil && rst.Result == false {
				//todo:
				fmt.Println("rules not passed send smscode:", err)
				//cache req
				val, err := json.Marshal(req)
				if err != nil {
					return gerror.NewCode(consts.CodeInternalError)
				}
				service.Generator().RecordSid(ctx, req.SessionId, consts.KEY_txs, string(val))
				///
				return gerror.NewCode(consts.NeedSmsCodeError(""))
			}
		}
		///analzy tx
		analzytx, err = service.EthTx().AnalzyTxs(ctx, signtx)
		if err != nil {
			g.Log().Error(ctx, "analzyTx:", err, signtx)
			return gerror.NewCode(consts.CodeInternalError)
		}
	}
	// /////sign
	s.pool.Submit(func() {
		s.CalSignTask(s.ctx, req.SessionId, req.Msg, req.Request)
		// recordtx
		service.DB().RecordTxs(ctx, analzytx)
	})

	return nil
}

func detx(rawtx string) (*types.Transaction, error) {

	rawBytes, err := hex.DecodeString(rawtx)
	if err != nil {
		return nil, err
	}
	tx := new(types.Transaction)
	err = tx.UnmarshalBinary(rawBytes)
	if err != nil {
		return nil, err
	}
	fmt.Println(tx)
	fmt.Println(tx.To())
	contractABI := `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_upgradedAddress","type":"address"}],"name":"deprecate","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"deprecated","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_evilUser","type":"address"}],"name":"addBlackList","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"upgradedAddress","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balances","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"maximumFee","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"_totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"unpause","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_maker","type":"address"}],"name":"getBlackListStatus","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowed","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"paused","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"who","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"pause","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getOwner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"newBasisPoints","type":"uint256"},{"name":"newMaxFee","type":"uint256"}],"name":"setParams","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"issue","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"redeem","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"basisPointsRate","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"isBlackListed","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_clearedUser","type":"address"}],"name":"removeBlackList","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"MAX_UINT","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_blackListedUser","type":"address"}],"name":"destroyBlackFunds","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_initialSupply","type":"uint256"},{"name":"_name","type":"string"},{"name":"_symbol","type":"string"},{"name":"_decimals","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"amount","type":"uint256"}],"name":"Issue","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"amount","type":"uint256"}],"name":"Redeem","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"newAddress","type":"address"}],"name":"Deprecate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"feeBasisPoints","type":"uint256"},{"indexed":false,"name":"maxFee","type":"uint256"}],"name":"Params","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_blackListedUser","type":"address"},{"indexed":false,"name":"_balance","type":"uint256"}],"name":"DestroyedBlackFunds","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_user","type":"address"}],"name":"AddedBlackList","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_user","type":"address"}],"name":"RemovedBlackList","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[],"name":"Pause","type":"event"},{"anonymous":false,"inputs":[],"name":"Unpause","type":"event"}]`
	c, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}
	fmt.Println(c)
	// callMsg := ethereum.CallMsg{
	// 	To:   tx.To(),
	// 	Data: tx.Data(), // 交易数据
	// }

	data := tx.Data()
	method, err := c.MethodById(data[:4])
	if err != nil {
		return nil, err
	}
	fmt.Println(method)
	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, data[4:])
	if err != nil {
		return nil, err
	}
	fmt.Println(args)
	return tx, nil
}
