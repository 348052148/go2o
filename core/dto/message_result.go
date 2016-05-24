/**
 * Copyright 2015 @ z3q.net.
 * name : message_result
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package dto

//操作消息结果
type MessageResult struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Tag     int    `json:"tag"`
}
