package userInfo

import (
	"context"
	"encoding/json"
	"fmt"
	"li17server/internal/consts"
	"li17server/internal/model"
	"li17server/internal/service"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
)

type sUserInfo struct {
	cache *gcache.Cache
	url   string
	///
	c *resty.Client
}

type respUserInfo struct {
	Status  int             `json:"status"`
	ErrCode int             `json:"errorCode"`
	Msg     string          `json:"msg"`
	Data    *model.UserInfo `json:"data"`
}

func (s *sUserInfo) GetUserInfo(ctx context.Context, userToken string) (userInfo *model.UserInfo, err error) {

	g.Log().Debug(ctx, "GetUserInfo:", userToken)
	if userToken == "" {
		g.Log().Error(ctx, "GetUserInfo:", userToken)
		return nil, gerror.NewCode(consts.AuthError())
	}
	//todo: check userToekn
	if err != nil {
		g.Log().Error(ctx, "GetUserInfo:", userToken)
		return nil, gerror.NewCode(consts.AuthError())
	}
	///
	// 用户信息示例
	// "id": 10,
	// "appPubKey": "038c90b87d77f2cc3d26132e1ea26e14646d663e3f43f17180345df3d54b8b5c70",
	// "email": "sunwenhao0421@163.com",
	// "loginType": "tkey-auth0-twitter-cyan",
	// "address": "0xe73E35d8Ecc3972481138D01799ED3934cc57853",
	// "keyHash": "U2FsdGVkX1/O6j9czaWzdjjDo/XPjk1hI8pIoaxSuS52zIxVuStK/nS07ucgiM5si8NjN97rAux3aH7Ld2i5oO8UuL6tpNZmLMG9ZpwVTxvGkCa3H14vTxWNz+yBoWG8",
	// "create_time": 1691118876
	info, err := s.getUserInfo(ctx, userToken)
	return info, err
	// return &model.UserInfo{
	// 	Id: 10,
	// 	// AppPubKey:  "038c90b87d77f2cc3d26132e1ea26e14646d663e3f43f17180345df3d54b8b5c70",
	// 	AppPubKey:  utility.GenNewSid(),
	// 	Email:      "sunwenhao0421@163.com",
	// 	LoginType:  "tkey-auth0-twitter-cyan",
	// 	Address:    "0xe73E35d8Ecc3972481138D01799ED3934cc57853",
	// 	KeyHash:    "U2FsdGVkX1/O6j9czaWzdjjDo/XPjk1hI8pIoaxSuS52zIxVuStK/nS07ucgiM5si8NjN97rAux3aH7Ld2i5oO8UuL6tpNZmLMG9ZpwVTxvGkCa3H14vTxWNz+yBoWG8",
	// 	CreateTime: 1691118876,
	// }, nil
}

func (s *sUserInfo) getUserInfo(ctx context.Context, token string) (*model.UserInfo, error) {
	resp, err := s.c.R().
		SetQueryParams(map[string]string{
			"token": token,
		}).
		// EnableTrace().
		Get(s.url)
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}
	userInfo := respUserInfo{}
	err = json.Unmarshal(resp.Body(), &userInfo)
	if err != nil {
		g.Log().Error(ctx, "getUserInfo:", err, token)
		return nil, err
	}
	return userInfo.Data, nil
}

func new() *sUserInfo {
	url, err := gcfg.Instance().Get(context.Background(), "userToken")
	if err != nil {
		panic(err)
	}

	///
	s := &sUserInfo{
		cache: gcache.New(),
		url:   url.String(),
		c:     resty.New(),
	}
	_, err = s.getUserInfo(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiIwMjI1YmI1MmU5NTcyMDUwZmZjMGM4MGRjZDBhYTBmNjQyNDFjMDk5ZDAzZjFlYTFjODEzMmZkMzViY2Q3MDBiMWMiLCJpYXQiOjE2OTQ0Mjk5OTEsImV4cCI6MTcyNTk2NTk5MX0.8YaF5spnD1SjI-NNbBCIBj9H5pspXMMkPJrKk23LdnM")
	if err != nil {
		panic(err)
	}
	return s
}

func init() {
	service.RegisterUserInfo(new())
}
