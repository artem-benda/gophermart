package dto

import "strconv"

type LuhnNumber string

// Valid check number is valid or not based on Luhn algorithm
func (n LuhnNumber) valid() bool {
	parsedNum, err := strconv.ParseInt(string(n), 10, 64)
	if err != nil {
		return false
	}
	return (parsedNum%10+checksum(parsedNum/10))%10 == 0
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

type UploadOrderRequest struct {
	OrderID LuhnNumber ``
}

func (r UploadOrderRequest) Valid() bool {
	return r.OrderID.valid()
}

func (r UploadOrderRequest) OrderNumberString() string {
	return string(r.OrderID)
}
