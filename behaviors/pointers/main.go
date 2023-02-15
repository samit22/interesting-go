// What gets printed for line 14, 15 and 16
package main

func ptr(num *int) int {
	*num = 4
	return *num
}

func sum(n int) int {
	return n + ptr(&n)
}

func main() {
	println(sum(3))
	println(sum(4))
	println(sum(5))
}
