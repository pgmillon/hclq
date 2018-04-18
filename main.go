package main

import (
	"io/ioutil"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"strings"
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 3 {
		inputFile := os.Args[1]
		paramName := os.Args[2]

		bytes, e := ioutil.ReadFile(inputFile)
		check(e)

		file, e := hcl.ParseBytes(bytes)
		check(e)

		items := file.Node.(*ast.ObjectList).Filter(paramName).Items

		if len(items) > 0 {
			literal, ok := items[0].Val.(*ast.LiteralType)
			if !ok {
				panic("Not a literal field")
			}

			print(strings.Replace(literal.Token.Text, "\"", "", -1))
		} else {
			panic("No such key")
		}
	} else {
		fmt.Printf("Usage: %s terraform.tfvars SOME_PROPERTY", os.Args[0])
	}

}
