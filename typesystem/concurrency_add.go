package typesystem

func add(in, out chan int){
	a := <- in
	b := <- in
	out <- a + b
}

func main(){
	in := make(chan int, 0)
	out := make(chan int, 0)

	go add(in, out)

	in <- 1
	in <- 2

	print(<-out)
}