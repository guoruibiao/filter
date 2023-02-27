package main

import (
	"fmt"
)

var handlersPool map[string]Handler = make(map[string]Handler)

func Register(key, operator string, handler Handler) {
	funcName := getHandlerFunc(key, operator)
	handlersPool[funcName] = handler
}

func getHandlerFunc(key, operator string) string {
	return fmt.Sprintf("%s_%s", key, operator)
}

func Do(params map[string]interface{}, presetConditions []Condition) bool {
	var succ = true
	for key, val := range params {
		// 循环处理每一个关键字 及其对应的校验条件
		for _, presetCondition := range presetConditions {
			if key != presetCondition.Keyword {
				continue
			}
			funcName := getHandlerFunc(key, presetCondition.Operator)
			h, ok := handlersPool[funcName]
			if !ok {
				succ = false
				break
			}

			handlePass, err := h(val, presetCondition.Operator, presetCondition.Value)
			if err != nil {
				fmt.Println(err.Error())
			}
			if err != nil || !handlePass {
				succ = false
				break
			}

		}
		if !succ {
			break
		}
	}
	return succ
}
