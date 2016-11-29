package main

func add(x, y int) int {
	z := x + y
	return z
}

func main() {
	x := 0x100
	y := 0X100
	go add(x, y)
}
