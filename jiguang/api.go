/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/7 16:51
 * @version 1.0
 */
package jiguang

import (
	"encoding/base64"
	"fmt"
)

const (
	J_PUSH_API        = "https://api.jpush.cn/v3/push"
	J_GROUP_PUSH_API  = "https://api.jpush.cn/v3/grouppush"     //该 API 用于为开发者在 portal 端创建的应用分组创建推送。groupkey 可以在创建的分组信息中获取，使用起来同 appkey 类似，但在使用的时候前面要加上 “group-” 前缀。
	J_VALIDATE_API    = "https://api.jpush.cn/v3/push/validate" //该 API 只用于验证推送调用是否能够成功，与推送 API 的区别在于：不向用户发送任何消息。 其他字段说明：同推送 API。
	PLATFORM_IOS      = "ios"
	PLATFORM_ANDROID  = "android"
	PLATFORM_WINPHONE = "winphone"
)

/**
 * <p>Description: (生成极光authorization) </p>
 * @author lizhi_duan
 * @date 2018/9/9 21:07
 */
func NewAuthrization(appKey string, masterSecret string) string {
	key := appKey + ":" + masterSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(key))
	fmt.Printf("source: %s\n To\n base64encode: %s\n", key, "Basic "+encoded)
	return "Basic " + encoded
}
