package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ismael@gmail.com
nepu email is nepu@edu.cn
my twitter is twiter@qq.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)
							\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match {
		fmt.Println(m)
	}

}
