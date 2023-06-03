package day13_5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// original implementation considered key to be needed, but actually we can just have a value for example: [1,2,3]
// each integer is also considered a json object
// we can consider instead to use json primitive to abstract these values
type JsonObject struct {
	key         *string
	primitiveValue *JsonPrimitive
	//intValue    *int
	//stringValue *string
	//arrayValue  []JsonObject
	arrayValue []JsonPrimitive
	objectValue *JsonObject
}

type JsonPrimitive struct {
	intValue *int
	stringValue *string
}

//var jsonRegexp = regexp.MustCompile("{\"(.*)\":(.*),?\"(.*)\":(.*)}")
var jsonRegexp = regexp.MustCompile("{\"([^:]+)\":(.*)}")

// think about breaking down a json into smaller piece
// - need to use a queue to keep track of sub-json object, and recurse
// - first implement simple case
func parseJsonObject(json string) (*JsonObject, error) {
	if json == "{}" {
		return &JsonObject{}, nil
	}

	if jsonRegexp.MatchString(json) {
		matches := jsonRegexp.FindStringSubmatch(json)
		key := matches[1]
		value := matches[2]

		switch {
		case isIntValue(value) || isStringValue(value):
			primitive, err := parseJsonPrimitive(value)
			if err != nil {
				return nil, err
			}
			return &JsonObject{key: &key, primitiveValue: primitive}, nil
		case isArrayValue(value):
			arrayValue, err := getArrayValue(value)
			if err != nil {
				return nil, err
			}
			return &JsonObject{key: &key, arrayValue: arrayValue}, nil
		case isObjectValue(value):
			objectValue, err := parseJsonObject(value)
			if err != nil {
				return nil, err
			}
			return &JsonObject{key: &key, objectValue: objectValue}, nil
		default:
			return nil, fmt.Errorf("unable to parse json value: %s", value)
		}
	}

	return nil, fmt.Errorf("not implemented yet")
}

func parseJsonPrimitive(value string) (*JsonPrimitive, error) {
	var primitive JsonPrimitive
	switch {
	case isIntValue(value):
		primitive = getIntPrimitive(value)
	case isStringValue(value):
		primitive = getStringPrimitive(value)
	default:
		return nil, fmt.Errorf("unable to parse json primitive value: %s", value)
	}

	return &primitive, nil
}

func isIntValue(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func getIntPrimitive(value string) JsonPrimitive {
	intValue, _ := strconv.Atoi(value)
	return JsonPrimitive{intValue: &intValue}
}

func isStringValue(value string) bool {
	const quotationChar string = "\""

	return getChar(value, 0) == quotationChar && getChar(value, len(value) - 1) == quotationChar
}

func getStringPrimitive(value string) JsonPrimitive {
	stringContent := value[1:len(value) - 1]
	return JsonPrimitive{stringValue: &stringContent}
}

func isArrayValue(value string) bool {
	const openBracket string = "["
	const closedBracket string = "]"

	return getChar(value, 0) == openBracket && getChar(value, len(value) - 1) == closedBracket
}

func getArrayValue(value string) ([]JsonPrimitive, error) {
	const delimiter string = ","

	arrayContent := value[1:len(value) - 1]
	tokens := strings.Split(arrayContent, delimiter)

	primitives := make([]JsonPrimitive, 0)
	for _, token := range tokens {
		primitive, err := parseJsonPrimitive(token)
		if err != nil {
			return nil, err
		}

		primitives = append(primitives, *primitive)
	}

	return primitives, nil
}

func isObjectValue(value string) bool {
	const openBrace string = "{"
	const closedBrace string = "}"

	return getChar(value, 0) == openBrace && getChar(value, len(value) - 1) == closedBrace
}

func getChar(s string, index int) string {
	return s[index:index + 1]
}