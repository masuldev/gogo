package main

import (
	"fmt"
	"strconv"
)

func main() {
	scan()
}

func convert() {
	var i int
	var k int64
	var f float64
	var s string

	i, _ = strconv.Atoi("350")
	fmt.Println(i)
	k, _ = strconv.ParseInt("cc7ffdd", 16, 32)
	k, _ = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)
	f, _ = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	s = strconv.Itoa(340)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)
}

func scan() {
	var num int
	fmt.Sscanf("57", "%d", &num)

	var s string
	s = fmt.Sprint(3.14)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)
}
