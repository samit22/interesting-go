// interesting switch
// what gets printed in
package main

func foo() bool {
	return false
}

func main() {
	caseA()
	caseB()
	caseC()
}

func caseA() {
	print("CASE A: ")
	switch foo(); {
	case false:
		println("False")
	case true:
		println("True")
	}
}

func caseB() {
	print("CASE B: ")
	switch foo() {
	case false:
		println("False")
	case true:
		println("True")
	}
}

func caseC() {
	print("CASE C: ")
	switch {
	case false:
		println("False")
	case true:
		println("True")
	}
}
