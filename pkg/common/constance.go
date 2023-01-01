package common

import "time"

const (
	// ExpireTime 2个星期不登录就过期
	ExpireTime = time.Hour * 24 * 14
	// UpdateTokenTime 离过期时间还有 30 min的时候才回去更新token
	UpdateTokenTime = ExpireTime / 2
)
