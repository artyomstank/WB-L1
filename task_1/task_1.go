package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Println("hello, I can speak")
}

func (h *Human) Work() {
	fmt.Println("human working")
}

type Action struct {
	Human
	ActionName string
}

func (a *Action) DoWbTask() {
	fmt.Println("action do wb")
}

func (a *Action) Work() {
	fmt.Println("action work")
}

func main() {
	act := Action{
		Human:      Human{Name: "Ethan", Age: 29},
		ActionName: "Task_1",
	}
	act.Speak()      //вызываем родительский метод у дочерней структуры
	act.Work()       //вызываем переопределённый метод
	act.Human.Work() //вызываем родительский(первоначальный метод)
	act.DoWbTask()   //вызов метода у Action
}
