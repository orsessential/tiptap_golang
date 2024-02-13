package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	wrd, _ := buf.ReadString('\n')
	wrd = strings.TrimSpace(wrd)
	str_wrd := strings.Split(wrd, "")
	var new_map = map[string]int{}
	for _, val := range str_wrd {
		fmt.Println(val)
		value, exist := new_map[val]
		if exist {
			new_map[val] = value + 1
		} else {
			new_map[val] = 1
		}
	}
	fmt.Println(new_map)
}
