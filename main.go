package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// 模拟获取的上游传参信息
	requestParams := make(map[string]interface{})
	requestParams["uid"] = "123456789"
	requestParams["version"] = "13.25.0.10"

	// 加载本地已配置的校验条件
	bytes, err := os.ReadFile("./data/demo1.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type tmp struct {
		PresetConditions []Condition `json:"conditions"`
	}
	t := tmp{}
	if err := json.Unmarshal(bytes, &t); err != nil {
		fmt.Println(err.Error())
		return
	}

	// 调用 filter 处理，看
	pass := Do(requestParams, t.PresetConditions)
	fmt.Println("本次校验结果：", pass)
}
