package main

import (
	"fmt"

	"github.com/CZero/gofuncy/lfs"
)

func main() {
	fmt.Println(lfs.SilentAtoi("5"))
	var Err error
	lfs.PanErr(Err)
	fmt.Println(lfs.SilentAtoi("g"))

}
