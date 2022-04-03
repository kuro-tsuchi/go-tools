package xstrings

import (
	"fmt"

	"github.com/huandu/xstrings"
)

func MainExec() {
	str := "helloworld"
	count := xstrings.Count(str, "w")
	fmt.Println(count)
}
