/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/9 21:53
 * @version 1.0
 */
package jiguang

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	appKey       = "275fba1e27d25231fdfb8ad0"
	masterScrect = "23f7fc7631ac73de3ce928d8"
)

func TestNewPushObject(t *testing.T) {
	object := NewPushObject(false)
	notification := NewNotification("测试极光推送，哈哈哈，希望一切顺利o啦")
	object.Notification = notification
	var audience Audience
	audience.AddAlias("51_1")
	object.Audience = audience
	push := object.JPush(appKey, masterScrect)
	fmt.Println(push)
}
func TestAudience_AddTag(t *testing.T) {
	var audience Audience
	audience.AddTag("北京", "天津", "bbb")
	audience.AddTag("北京", "保定", "aaa", "bbb", "ccc")
	bytes, e := json.Marshal(audience)
	fmt.Println(e, string(bytes))
}
