// What gets printed in final data
package main

import "fmt"

func main() {
	var data = make(map[string]interface{})
	data["name"] = "samit"

	fmt.Printf("Before function call %#v\n", data)

	myFunction(data)

	fmt.Printf("After function call 'myFunction': %#v \n", data)

	myFunction2(data)

	fmt.Printf("After function call 'my Function2': %#v \n", data)

	copyVar := data
	copyVar["watch"] = "Madan Bahadur Hari Bahadur"

	fmt.Printf("Final data: %#v \n", data)
}

func myFunction(data map[string]interface{}) {
	data["work"] = "lazy"
}
func myFunction2(data map[string]interface{}) {
	newData := data
	newData["do"] = "sleep"
}
