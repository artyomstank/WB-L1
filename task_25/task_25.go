package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d) //блокируем операцией чтения из канала
}

func main() {
	fmt.Println("sleep")
	sleep(2 * time.Second)
	fmt.Println("end sleep func!")
}
