package main

import (
	"fmt"

	"github.com/lhboy1984/godatayes"
)

func main() {

	fmt.Println(godatayes.GetData(godatayes.NewArgument("getSecID", map[string]interface{}{
		"ticker": "000001",
	})))

}
