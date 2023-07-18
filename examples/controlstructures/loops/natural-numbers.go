// nested package, examples/controlstructures/loops
package loops

import (
	"fmt"
	"strconv"
)

func GetNaturalNumbers(initial, last int) string {
	var output string

	for i := initial; i <= last; i++ {
		if i != last {
			output += fmt.Sprintf("%v, ", i)
		} else {
			output += strconv.Itoa(i)
		}
	}

	return output
}
