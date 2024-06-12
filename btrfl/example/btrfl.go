package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iamsoloma/butterfly/btrfl"
)

func main() {
	file, err:= os.Create("bloom.btrfl")
	check(err)
	defer file.Close()

	example := make(map[string]string)
	example["Ping"] = "Pong"
	example["Hi!"] = "Bye..."
	Keys, Values := []string{}, []string{}

	for k, v := range example {
		Keys = append(Keys, k)
		Values = append(Values, v)
	}

	Wfile, err := os.OpenFile("bloom.btrfl", os.O_WRONLY, 0666)
	check(err)
	defer Wfile.Close()


	btrfl.WriteKeySpace(Wfile, Keys)
	check(err)

	AWfile, err := os.OpenFile("bloom.btrfl", os.O_APPEND|os.O_WRONLY, 0666)
	check(err)
	defer AWfile.Close()
	Rfile, err := os.OpenFile("bloom.btrfl", os.O_RDONLY, 0666)
	check(err)
	defer Rfile.Close()

	last, err := btrfl.AppendValues(AWfile, Rfile, Values)
	check(err)
	fmt.Println("Last Appended value: " + strconv.Itoa(last))
}

func check(err error) {
	if err!=nil{
		panic(err)
	}
}