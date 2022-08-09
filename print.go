package print

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

// JSON - prints JSON output for given interface
func JSON(object interface{}) {
	s, _ := prettyjson.Marshal(object)
	fmt.Println(string(s))
}
