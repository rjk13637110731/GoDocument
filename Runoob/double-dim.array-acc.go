package main

import "fmt"

func main() {
	sites := [2][2]string{}

	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	fmt.Println(sites)
}
