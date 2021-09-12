package main

import "time"

func main() {

	uni := NewUniverse()
	uni.Seed()
	for i := 0; i < 9999; i++ {
		uni.Print()
		uni.Next()
		time.Sleep(time.Second * 1)
	}
}
