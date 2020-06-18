package main

import (
	"fmt"
)

var a []string

func main() {
	s := "shonphand"
	var ok bool
	for _,v:=range s {

		ok = found(string(v))
		if !ok {
			fmt.Printf("%s", string(v))
			a = append(a, string(v))
		}

	}
}

func found(s string) bool {
	flag := 0
	for _, v := range a {
		if s == v {
			flag = 1
			break
		}

	}
	if flag == 1 {
		return true
	} else {
		return false
	}
}
