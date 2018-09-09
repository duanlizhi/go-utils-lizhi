/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/9 20:26
 * @version 1.0
 */
package httpClient

import "github.com/ddliu/go-httpclient"

/**
 * <p>Description: (创建http客户端) </p>
 * @author lizhi_duan
 * @date 2018/9/9 20:38
 */
func NewMyHttpClient() *httpclient.HttpClient {
	client := &httpclient.HttpClient{}
	return client
}
