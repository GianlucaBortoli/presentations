package main

import "fmt"

func main() {
	const c = "日語ab"
	for i := 0; i < len(c); i++ {
		fmt.Printf("%x starts at byte position %d\n", c[i], i)
    }
}
