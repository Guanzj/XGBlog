/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 14:49:43
 * @LastEditTime: 2020-03-05 14:59:45
 * @FilePath: /XGBlog/app/protocol/protocol.go
 * @Description:定义返回协议
 */

package protocol

// Resp 返回协议结构体
type Resp struct {
	Ret  int32       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// AdminJwtTokenInfo 估计是jwt用户验证返回的token协议结构体
type AdminJwtTokenInfo struct {
	AdminID    int32
	Username   string
	Expiretime int32
}
