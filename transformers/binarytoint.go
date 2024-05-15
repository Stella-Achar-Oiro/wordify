package transformers

import (
	"fmt"
	"strconv"
)

func BinaryToInt(bin string) string {
	number, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Printf("Error parsing binary: %s\n", err)
		return bin // Return original string on error
	}
	return fmt.Sprint(number)
}
