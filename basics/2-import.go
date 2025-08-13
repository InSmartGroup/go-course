package main

// we can import multiple modules from the standard library
import (
	foo "fmt"  // we can also assign a custom name to a module
	"net/http"
)

func main() {
	foo.Println("Hello, GO Standard Library!") // now we specify our custom module name instead of a standard name

	response, error := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if error != nil {
		foo.Println("Error:", error)
		return
	}

	defer response.Body.Close()

	foo.Println("HTTP Response Status:", response.Status)
}
