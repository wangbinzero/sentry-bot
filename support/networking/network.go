package networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
)

func JSONRequestDynamicHeaders(httpClient *http.Client,
	method string,
	reqURL string,
	data string,
	headers map[string]HeaderFn,
	responseData interface{},
	errorKey string) error {
	headersMap := map[string]string{}
	for header, fn := range headers {
		headersMap[header] = fn(method, reqURL, data)
	}

	return JSONRequest(httpClient, method, reqURL, data, headersMap, responseData, errorKey)
}

// 构建http 请求客户端
func JSONRequest(httpClient *http.Client,
	method string,
	reqURL string,
	data string,
	headers map[string]string,
	responseData interface{},
	errorKey string) error {

	req, err := http.NewRequest(method, reqURL, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("无法创建http请求: %s", err)
	}

	//添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("无法执行http请求:%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("无法读取http返回值:%s", err)
	}

	bodyString := string(body)

	//读取媒体类型
	contentType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))

	if err != nil {
		return fmt.Errorf("无法读取媒体类型 Content-Type值:%s", err)
	}
	if contentType != "application/json" && contentType != "application/hal+json" {
		return fmt.Errorf("无效的媒体类型: %s", err)
	}

	if errorKey != "" {
		var errorResponse interface{}
		err = json.Unmarshal(body, &errorResponse)

		if err != nil {
			return fmt.Errorf("无法反序列化 response body: %s", err)
		}

		switch err := errorResponse.(type) {
		case map[string]interface{}:
			if _, ok := err[errorKey]; ok {
				return fmt.Errorf("error in response, bodyString: %s", bodyString)
			}
		}
	}

	if responseData != nil {
		// parse response, the passed in responseData should be a pointer
		err = json.Unmarshal(body, responseData)
		if e != nil {
			return fmt.Errorf("could not unmarshall response body into json: %s | response body: %s", err, bodyString)
		}
	}
	return nil

}
