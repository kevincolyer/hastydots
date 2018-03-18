package main

import "fmt"
import "strings"
import "strconv"

func warn(s string) {
	fmt.Println("WARNING:" + s)
}

func debug(s string, args ...interface{}) {
	if dbg == true {
		fmt.Printf("Debug:"+s, args...)
	}
	return
}

func output(s string, args ...interface{}) {
	fmt.Printf(s, args...)
	return
}

// splits strings in two of form [.\d*]. Second string is coerced to int value or 1 if missing/err
func a1splitter(s string) (l string, i int) {
	arr := strings.SplitN(s, "", 2) // split into two parts
	l = arr[0]
	i = 1
	if len(arr) == 1 {
		return
	}
	if j, err := strconv.Atoi(arr[1]); err == nil {
		i = j
	}
	return
}
