package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = `~!@#$%^&*()_-+={[}]|\:;"'<,>.?/`
const alphaNumeric = numbers + alphabets
const alphaNumericSymbol = numbers + alphabets + symbols

func RandomInt(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func RandomString(size int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphabets[rand.Intn(k)])
	}
	return sb.String()
}

func RandomBool() bool {
	num := rand.Int63n(2)
	if num == 0 {
		return false
	} else {
		return true
	}
}

func RandomAlphaNumeric(size int) string {
	var sb strings.Builder
	k := len(alphaNumeric)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphaNumeric[rand.Intn(k)])
	}
	return sb.String()
}

func RandomAlphaNumericSymbolString(size int) string {
	var sb strings.Builder
	k := len(alphaNumericSymbol)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphaNumericSymbol[rand.Intn(k)])
	}
	return sb.String()
}

func RandomOTP(size int) int {
	// using all the numbers except 0
	nos := "123456789"
	var sb strings.Builder
	k := len(nos)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphaNumeric[rand.Intn(k)])
	}
	v, _ := strconv.Atoi(sb.String())
	return v
}

func RandomUid() int64 {
	return rand.Int63n(10000000001 - 100000001 + 1)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomAlphaNumeric(10), RandomString(10))
}

func RandomPhone() string {
	nos := "123456789"
	var sb strings.Builder
	k := len(nos)
	for i := 0; i < 10; i++ {
		sb.WriteByte(alphaNumeric[rand.Intn(k)])
	}
	return "+91" + sb.String()
}

func RandomUserType() string {
	var userTypes []string = []string{"vendor", "customer"}
	return userTypes[RandomInt(0, int64(len(userTypes)-1))]
}
