package main

type Condition struct {
	Keyword  string      `json:"keyword"`  // keyword   关键字
	Operator string      `json:"operator"` // operator  操作符
	Value    interface{} `json:"value"`    // value     值
}
