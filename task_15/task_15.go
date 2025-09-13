package main


var justString string

func someFunc(){
	v := createHugeString(1 &lt;&lt; 10)
	b:=make([]byte,100)//создаем новый массив длинны 100, чтобы избежать удержания памяти большим массивом,
	//  длинна которого после 100-го элемента не используется, но остается в памяти
	copy(b,v[:100])
	justString=string(b)
}
func main(){
	someFunc()
}