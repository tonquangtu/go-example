package main

func main()  {
	println(magicDeferFuckingGo())
}

func magicDeferFuckingGo() int {
	x := 1
	defer func() { x = 2 }()
	return x
}
