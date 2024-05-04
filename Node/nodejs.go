package Node

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// Response 结构体用于解析响应JSON数据
type Response struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// QueryAPI 函数发送请求到本地API，并解析响应
func QueryNodeQuery(question string, targetID int, history string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"bind":    targetID,
		"query":   question,
		"history": history,
	})
	if err != nil {
		return "调用Nodejs后端失败", err
	}
	resp, err := http.Post("http://"+viper.GetString("node_config.node_prefix")+"/admin_api/query", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "调用Nodejs后端失败", err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "调用Nodejs后端失败", err
	}
	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "调用Nodejs后端失败", err
	}
	return response.Data, nil
}

func QueryNodeChatGPT(question string) (string, error) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"query": question,
	})

	resp, err := http.Post("http://"+viper.GetString("node_config.node_prefix")+"/admin_api/query/normal", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "调用Nodejs后端失败", err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "调用Nodejs后端失败", err
	}
	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "调用Nodejs后端失败", err
	}
	return response.Data, nil
}
