package main

import (
	"github.com/elliotchance/phpserialize"
	"fmt"
)

func main() {
	out, err := phpserialize.Marshal(3.2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))

	var in float64
	err = phpserialize.Unmarshal(out, &in)

	fmt.Println(in)
}
