package filter

import "testing"

func TestDo(t *testing.T) {
	params := map[string]interface{}{
		"version": "13.13.0.0",
	}

	presetConditions := []Condition{
		{
			Key:      "version",
			Operator: "==",
			Value:    "13.13.0.0",
		},
	}

	// test equal
	ret := Do(params, presetConditions)
	t.Log("equal case: ", ret)

	// test not equal
	params["version"] = "13.12"
	ret = Do(params, presetConditions)
	t.Log("not equal case: ", ret)
}
