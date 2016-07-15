package main

//要修改字符串,可先将其转换成 []rune 或 []byte,完成后再转换为 string。无无论哪种转换,都会重新分配内存,并复制字节数组。
func main() {
	s := "abcd"
	bs := []byte(s)

	bs[1] = 'B'
	println(s)
	println(string(bs))

	us := []rune(s)
	us[1] = 'J'
	println(s)
	println(string(us))

}
