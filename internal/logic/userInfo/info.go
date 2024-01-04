package userInfo

import (
	"context"
	"encoding/json"
	"mpcServer/internal/model"
	"mpcServer/internal/service"

	"github.com/mpcsdk/mpcCommon/mpccode"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/errors/gerror"
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
	if userToken == "" {
		return nil, mpccode.CodeParamInvalid()
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
	if info == nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userToken", userToken),
		))
		return nil, err
	}
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
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("token", token),
		))
		return nil, err
	}
	userInfo := respUserInfo{}
	err = json.Unmarshal(resp.Body(), &userInfo)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("token", token),
			mpccode.ErrDetail("resp", resp),
		))
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
