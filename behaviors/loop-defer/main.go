// How does loop and defer work
package main

func main() {
	for i := range [3]int{} {
		defer func() { print(i) }()
	}

	for j := 0; j < 3; j++ {
		defer func() { print(j) }()
	}

}
