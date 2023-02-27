# filter

初步设想：设计一套开关配置组件，用于快速上下线某些功能。

## 1 在编辑平台增加配置
此处忽略

## 2 在业务层读取配置，并关联入参进行校验
```go
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
```