package wrappingerror

import (
	"regexp"
	"strings"
)

func indent(input string, indent string, countOfTimes int) string {
	// Pattern used for breaking into separate lines.
	re := regexp.MustCompile("(?m)^(.+)$")

	// Result pattern used to substitute into.
	indentation := strings.Repeat(indent, countOfTimes)
	resultPattern := []byte(indentation + "$1")

	indentedBytes := re.ReplaceAll([]byte(input), resultPattern)
	return string(indentedBytes)
}
