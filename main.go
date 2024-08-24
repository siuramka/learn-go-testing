package main

import (
	"fmt"
	"go-with-tests/di"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	repo := di.NewSumRepo()
	repo.Add(10)

	di.Calculate(repo)

	test := repo.GetAll()
	_ = append(test, 3)

	fmt.Printf("%v", test)
}
