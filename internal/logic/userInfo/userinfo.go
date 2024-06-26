package userInfo

import (
	"context"
	"mpcServer/internal/config"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/userInfoGeter"
)

type sUserInfo struct {
	url string
	///
	userGeter *userInfoGeter.UserTokenInfoGeter

	dur int
}

func (s *sUserInfo) GetUserInfo(ctx context.Context, userToken string) (userInfo *userInfoGeter.UserInfo, err error) {
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
	// if v, ok := s..Get(ctx, userToken); ok == nil && !v.IsEmpty() {
	// 	info := &userInfoGeter.UserInfo{}
	// 	err = v.Struct(info)
	// 	if err != nil {
	// 		g.Log().Error(ctx, "GetUserInfo:", "toekn:", userToken, "err:", err)
	// 		return nil, mpccode.CodeInternalError()
	// 	}
	// 	return info, nil
	// }
	///
	info, err := s.userGeter.GetUserInfo(ctx, userToken)
	if err != nil {
		g.Log().Error(ctx, "GetUserInfo:", "toekn:", userToken, "err:", err)
		return info, mpccode.CodeInternalError()
	}
	// s.c.Set(ctx, userToken, info, s.dur)
	return info, err
}

func new() *sUserInfo {
	///
	// url, err := gcfg.Instance().Get(context.Background(), "userTokenUrl")
	// if err != nil {
	// 	panic(err)
	// }
	url := config.Config.UserTokenUrl
	//
	s := &sUserInfo{
		// userGeter: userGeter,

		dur: config.Config.Cache.SessionDuration,
	}
	///
	r := g.Redis("cache")
	_, err := r.Conn(gctx.GetInitCtx())
	if err != nil {
		panic(err)
	}
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(r))
	///
	userGeter := userInfoGeter.NewUserInfoGeter(url, s.dur)
	_, err = userGeter.GetUserInfo(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiIwMjI1YmI1MmU5NTcyMDUwZmZjMGM4MGRjZDBhYTBmNjQyNDFjMDk5ZDAzZjFlYTFjODEzMmZkMzViY2Q3MDBiMWMiLCJpYXQiOjE2OTQ0Mjk5OTEsImV4cCI6MTcyNTk2NTk5MX0.8YaF5spnD1SjI-NNbBCIBj9H5pspXMMkPJrKk23LdnM")
	if err != nil && !gerror.HasError(err, mpccode.CodeTokenInvalid()) {
		panic(err)
	}
	//
	s.userGeter = userGeter

	return s
}

func init() {
	service.RegisterUserInfo(new())
}
