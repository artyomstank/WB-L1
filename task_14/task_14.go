package main

import "fmt"

func whatType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("переменная типа: int")
	case string:
		fmt.Println("переменная типа: string")
	case bool:
		fmt.Println("переменная типа: bool")
	case chan int:
		fmt.Println("переменная типа: chan int")
	case chan string:
		fmt.Println("переменная типа: chan string")
	case chan bool:
		fmt.Println("переменная типа: chan bool")
	default:
		fmt.Printf("у переменной другой тип: %T\n", v)
	}
}

func main() {
	var val_int int = 1
	var val_float float64 = 1
	var val_string string = "string"
	var val_bool bool = true
	var ch_int chan int = make(chan int)
	var ch_string chan string = make(chan string)
	var ch_bool chan bool = make(chan bool)

	whatType(val_int)
	whatType(val_string)
	whatType(val_bool)
	whatType(ch_int)
	whatType(ch_string)
	whatType(ch_bool)
	whatType(val_float)
}
