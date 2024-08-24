package main

import (
	"fmt"
	"go-with-tests/mocking"
	"io"
	"net/http"
	"os"
	"time"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	sleeper := mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, &sleeper)
}
