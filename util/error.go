package util

import "errors"

func MergeErrors(errs ...error) error {
	str := ""
	for i, v := range errs {
		str += v.Error()
		if i != len(errs)-1 {
			str += "\n"
		}
	}
	return errors.New(str)
}
