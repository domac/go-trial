package main

func main() {
	s := "abc"

	println(s[0] == '\x61', s[1] == 'b', s[2] == '\x63')
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
}
