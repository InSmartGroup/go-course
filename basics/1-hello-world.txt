package main // this is an entry point for every executable program

import "fmt" // module import; "fmt" stands for "format" and is responsible for the standard input/output

func main() { // this function should always be named "main", otherwise the compiler won't be able to compile and run the program
	fmt.Println("Hello World!")
}

/*
GO allows building executable binary files from our code
We can do this by typing "go build [file_name]" in terminal
Next, we can run the program by typing "./[file_name]" in terminal
*/