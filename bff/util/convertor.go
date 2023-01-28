package util

import (
	"fmt"
	"strconv"
)

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func ToBool(v interface{}) bool {
	if value, ok := v.(bool); ok {
		return value
	}
	return false
}

func ToInt(v interface{}) int {
	if value, err := strconv.Atoi(ToString(v)); err != nil {
		return 0
	} else {
		return value
	}
}
