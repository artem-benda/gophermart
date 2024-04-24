package dto

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

func LuhnStringValidator(fl validator.FieldLevel) bool {
	switch v := fl.Field(); v.Kind() {
	case reflect.String:
		luhnNumberString := v.String()
		parsedNum, err := strconv.ParseInt(luhnNumberString, 10, 64)
		if err != nil {
			return false
		}
		return (parsedNum%10+checksum(parsedNum/10))%10 == 0
	default:
		return false
	}
}

func checksum(n int64) int64 {
	var luhn int64

	for i := 0; n > 0; i++ {
		cur := n % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		n = n / 10
	}
	return luhn % 10
}
