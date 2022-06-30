package filter

import (
	"fmt"
)

type Condition struct {
	Key      string
	Value    interface{}
	Operator string
}

type Handler func(input interface{}, operator string, value interface{}) (bool, error)

var handlersPool map[string]Handler = make(map[string]Handler)

func Register(key, operator string, handler Handler) {
	funcName := fmt.Sprintf("%s%s", key, operator)
	handlersPool[funcName] = handler
}

func Do(params map[string]interface{}, presetConditions []Condition) bool {
	fmt.Println("-----------------")
	fmt.Printf("pool=%#v\n", handlersPool)
	var succ = true
	for key, val := range params {
		//
		for _, presetCondition := range presetConditions {
			funcName := key + presetCondition.Operator
			fmt.Println("-----------------")
			fmt.Printf("funcName=%#v\n", funcName)
			h, ok := handlersPool[funcName]
			if !ok {
				succ = false
				break
			}

			handlePass, err := h(val, presetCondition.Operator, presetCondition.Value)
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
