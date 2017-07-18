package main

import (
	"fmt"

	"github.com/robyerts/goTest/stringutil"

	"github.com/robyerts/goTest/myPackage"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(myPackage.Mee())
}
