package main

import (
	"io/ioutil"
	"os"

	"github.com/42LoCo42/z85m"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var out []byte
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		out, err = z85m.Decode(in)
	} else {
		out, err = z85m.Encode(in)
	}
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(out)
}
