package transformers

import (
	"fmt"
	"strconv"
)

func HexadecimalToInt(hex string) string {
	number, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Printf("Error parsing hexadecimal: %s\n", err)
		return hex // Return original string on error
	}
	return fmt.Sprint(number)
}
