package day13_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunner(t *testing.T) {
	parseJsonObject("{\"abc\":123}")
	parseJsonObject("{\"abc\":\"123\"}")
	parseJsonObject("{\"animals\":[\"cow\",\"chicken\",\"cicada\"]}")
	parseJsonObject("{\"company\":{\"name\":\"Indeed\"}}",)
}

func TestParseJsonObject(t *testing.T) {
	testcases := []struct {
		name string
		input string
		expectedOutput *JsonObject
		expectingError bool
	} {
		{
			"empty json",
			"{}",
			&JsonObject{},
			false,
		},
		{
			"int value json",
			"{\"age\":10}",
			func() *JsonObject {
				key := "age"
				intValue := 10
				primitiveValue := JsonPrimitive{intValue: &intValue}
				return &JsonObject{key: &key, primitiveValue: &primitiveValue}
			}(),
			false,
		},
		{
			"string value json",
			"{\"name\":\"Ronald\"}",
			func() *JsonObject {
				key := "name"
				stringValue := "Ronald"
				primitiveValue := JsonPrimitive{stringValue: &stringValue}
				return &JsonObject{key: &key, primitiveValue: &primitiveValue}
			}(),
			false,
		},
		{
			"array value json",
			"{\"animals\":[\"cow\",\"chicken\",\"cicada\"]}",
			func() *JsonObject {
				key := "animals"
				arrayValue1, arrayValue2, arrayValue3 := "cow", "chicken", "cicada"
				primitiveValue1 := JsonPrimitive{stringValue: &arrayValue1}
				primitiveValue2 := JsonPrimitive{stringValue: &arrayValue2}
				primitiveValue3 := JsonPrimitive{stringValue: &arrayValue3}
				return &JsonObject{key: &key, arrayValue: []JsonPrimitive{primitiveValue1, primitiveValue2, primitiveValue3}}
			}(),
			false,
		},
		{
			"nested json",
			"{\"company\":{\"name\":\"Indeed\"}}",
			func() *JsonObject {
				outerKey := "company"
				innerKey := "name"
				stringValue := "Indeed"
				primitiveValue := JsonPrimitive{stringValue: &stringValue}
				return &JsonObject{key: &outerKey, objectValue: &JsonObject{key: &innerKey, primitiveValue: &primitiveValue}}
			}(),
			false,
		},
		{
			"invalid json",
			"{",
			nil,
			true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			output, err := parseJsonObject(testcase.input)
			assert.Equal(t, testcase.expectedOutput, output)
			if testcase.expectingError {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestParseJsonPrimitive(t *testing.T) {
	testcases := []struct {
		name string
		input string
		expectedOutput *JsonPrimitive
		expectingError bool
	} {
		{
			"int value",
			"100",
			func() *JsonPrimitive {
				intValue := 100
				return &JsonPrimitive{intValue: &intValue}
			}(),
			false,
		},
		{
			"string value",
			"\"test\"",
			func() *JsonPrimitive {
				stringValue := "test"
				return &JsonPrimitive{stringValue: &stringValue}
			}(),
			false,
		},
		{
			"invalid value",
			"test",
			nil,
			true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			output, err := parseJsonPrimitive(testcase.input)
			assert.Equal(t, testcase.expectedOutput, output)
			if testcase.expectingError {
				assert.NotNil(t, err)
			}
		})
	}
}