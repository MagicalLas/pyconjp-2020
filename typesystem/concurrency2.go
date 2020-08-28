package typesystem

func hello(){
	println("Hello Las")
}

func main2(){
	go hello()
	println("Hello Pycon")
}