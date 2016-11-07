package main

import "fmt"

func main() {

	s := "ttk"

	switch s {
	case "brk":
	case "krts":
	default:
		panic(fmt.Sprintf("panic is trigered %q", s))
	}

}
