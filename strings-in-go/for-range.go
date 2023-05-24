package main

import "fmt"

func main() {
	const c = "日本ab"
	for i, r := range c {
        fmt.Printf("%x (%U) starts at byte position %d\n", r, r, i)
    }
}