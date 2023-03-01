package main

import "encoding/json"

func JSONParse(body []byte, v interface{}) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func JSONStringify(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}
