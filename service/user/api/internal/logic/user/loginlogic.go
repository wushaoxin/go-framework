package user

import (
    "context"
    "github.com/golang-jwt/jwt/v4"
    "go-framework/common/code/userCode"
    "go-framework/common/response"
    "go-framework/service/user/model"
    "strings"
    "time"
    
    "go-framework/service/user/api/internal/svc"
    "go-framework/service/user/api/internal/types"
    
    "github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
    if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
        return nil, response.InvalidParam
    }
    
    userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
    switch err {
    case nil:
    case model.ErrNotFound:
        return nil, userCode.UserNotFund
    default:
        return nil, err
    }
    
    if userInfo.Password != req.Password {
        return nil, userCode.InvalidPasswd
    }
    
    // ---start---
    now := time.Now().Unix()
    accessExpire := l.svcCtx.Config.Auth.AccessExpire
    jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
    if err != nil {
        return nil, err
    }
    // ---end---
    
    return &types.LoginReply{
        Id:           userInfo.Id,
        Name:         userInfo.Name,
        Gender:       userInfo.Gender,
        AccessToken:  jwtToken,
        AccessExpire: now + accessExpire,
        RefreshAfter: now + accessExpire/2,
    }, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
    claims := make(jwt.MapClaims)
    claims["exp"] = iat + seconds
    claims["iat"] = iat
    claims["userId"] = userId
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims = claims
    return token.SignedString([]byte(secretKey))
}
