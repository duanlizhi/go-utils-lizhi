/**
 * <p>Description: (极光推送api) </>
 * @author lizhi_duan
 * @date 2018/9/7 09:39
 * @version 1.0
 */
package jiguang

import (
	"encoding/json"
	"fmt"
	"go-utils-lizhi/httpClient"
	"log"
)

/**
 * <p>Description: (极光推送消息) </p>
 * @author lizhi_duan
 * @date 2018/9/9 21:25
 */
func (pushObj *PushObject) JPush(appKey string, masterSecret string) string {
	authrization := NewAuthrization(appKey, masterSecret)
	client := httpClient.NewMyHttpClient()
	client.Headers = map[string]string{
		"Authorization": authrization,
	}
	bytes, err := json.Marshal(pushObj)
	fmt.Printf("push err: %v\n push json str: %s\n", err, string(bytes))
	response, err := client.PostJson(J_PUSH_API, pushObj)
	if err != nil {
		log.Fatalf("Jpush err,appKey:%s\n,masterSecret:%s\n,pushObj:%v,\n err:%v", appKey, masterSecret, pushObj, err)
		return err.Error()
	}
	s, e := response.ToString()
	if e != nil {
		log.Printf("response body to string err: %v\n", e)
	}
	return s
}
