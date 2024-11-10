package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
	"surge/cmdutil/common"
	"surge/cmdutil/verbose"
)

type SurgeApiHeader struct {
	Server string
	Port   string
	XKey   string
}

type Response common.Response
type requestBody common.RequestBody
type responseError common.ResponseError

//func init() {
//	var surge_api_config cmd.SurgeApiHeader
//	surge_api_config.Server = cmd.GetConfigKey("server", "127.0.0.1")
//	surge_api_config.Port = cmd.GetConfigKey("port", "6170")
//	surge_api_config.XKey = cmd.GetConfigKey("x-key", "")
//}

// Get 请求
func Get(url string) (Response, error) {
	var requestBody requestBody
	requestBody.Method = "GET"
	requestBody.Url = url
	return requestServices(requestBody)
}

// Post 请求
func Post(url string, payload string) (Response, error) {
	var requestbody requestBody
	requestbody.Method = "POST"
	requestbody.Url = url
	requestbody.Payload = payload
	return requestServices(requestbody)
}

// requestServices 执行 HTTP 请求
func requestServices(requestBody requestBody) (Response, error) {
	var surgeApiConfig SurgeApiHeader

	var responseBody Response
	surgeApiConfig.Server = GetConfigKey("server", "127.0.0.1")
	surgeApiConfig.Port = GetConfigKey("port", "6171")
	surgeApiConfig.XKey = GetConfigKey("x-key", "123123")

	var regUrl = fmt.Sprintf("http://%s:%s%s", surgeApiConfig.Server, surgeApiConfig.Port, requestBody.Url)
	var result responseError
	client := &http.Client{}
	payload := strings.NewReader(requestBody.Payload)
	req, err := http.NewRequest(requestBody.Method, regUrl, payload)

	if err != nil {
		verbose.Println(" err", err)
		return Response{}, errors.New(err.Error())
	}
	req.Header.Add("X-Key", surgeApiConfig.XKey)
	req.Header.Add("Accept", " */*")

	res, errDo := client.Do(req)
	if errDo != nil {
		verbose.Println(" errDo", errDo)
		return Response{}, errors.New(errDo.Error())
	}
	defer res.Body.Close()

	requestResponse, err := ioutil.ReadAll(res.Body)

	responseBody.Request = common.RequestBody(requestBody)
	responseBody.Status = res.Status
	responseBody.StatusCode = res.StatusCode

	switch res.StatusCode {
	case 200: // 正常 200 请求
		responseBody.Context = string(requestResponse)
		verbose.LogHttpRequest(common.Response(responseBody))
		return responseBody, nil
	case 500: // API 服务端内部 500 错误
		return Response{}, errors.New("surge 客户端内部错误(API 请求返回 500)")
	default: // 非正常 200 请求（业务失败）
		unmarshalError := json.Unmarshal([]byte(requestResponse), &result)
		if unmarshalError != nil {
			// JSON 解析失败
			responseBody.Context = "JSON 解析失败"
			verbose.LogHttpRequest(common.Response(responseBody))
			return responseBody, errors.New("JSON 解析错误" + unmarshalError.Error())
		} else {
			// JSON 解析成功，返回具体错误信息
			responseBody.Context = string(requestResponse)
			verbose.LogHttpRequest(common.Response(responseBody))
			return responseBody, errors.New(result.ErrorStr)
		}
	}
}

func GetConfigKey(key string, defaultValue string) string {
	if err := viper.Get(key); err != nil {
		// 配置中包含指定键值
		return viper.Get(key).(string)
	} else {
		// 配置中不包含指定键值
		return defaultValue
	}
}

func main() {

}
