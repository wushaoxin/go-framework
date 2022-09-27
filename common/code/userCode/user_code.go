package userCode

import "go-framework/common/response"

var (
    UserNotFund   = response.NewCodeError(100_001, "用户未找到")
    InvalidPasswd = response.NewCodeError(100_003, "用户密码不正确")
)
