package handler

import (
	"context"
	"errors"
	"github.com/Jack-ZL/go_rookie/config"
	"goRookie_project/common"
	adminProto "goRookie_project/server/proto/admin"
	"goRookie_project/server/store"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AdminServer struct {
}

/*
后台管理登录
*/
func (a *AdminServer) AdminLogin(ctx context.Context, req *adminProto.AdminLoginReq) (*adminProto.AdminLoginRsp, error) {
	if strings.TrimSpace(req.UserName) == "" {
		return nil, errors.New("登录用户名为空")
	}
	if strings.TrimSpace(req.Password) == "" {
		return nil, errors.New("登录密码为空")
	}
	if strings.TrimSpace(req.Captcha) == "" {
		return nil, errors.New("验证码为空")
	}

	admin := &store.Admin{}
	adminUser, err := admin.GetAdminByName(strings.TrimSpace(req.UserName))
	if err != nil {
		return nil, errors.New("获取登录用户失败")
	}

	if adminUser.Salt == "" {
		return nil, errors.New("登录用户信息错误")
	}

	merge := common.DoubleStrMerge(strings.TrimSpace(req.Password), adminUser.Salt)
	if common.Md5Encrypt(merge) != adminUser.Password {
		return nil, errors.New("密码错误")
	}

	token := common.MakeSignature(config.Conf.App["token_key"].(string), strconv.FormatInt(time.Now().Unix(), 10))
	// 用户token缓存
	common.RedisSet(common.ADMIN_TOKEN_KEY, token, common.TOKEN_VALID_TIME)
	return &adminProto.AdminLoginRsp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: &adminProto.Admin{
			Id:     adminUser.Id,
			Name:   adminUser.Name,
			Mobile: adminUser.Mobile,
			Email:  adminUser.Email,
			Token:  token,
		},
	}, nil
}
