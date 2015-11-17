go-prefixsearch
===============

Golang string prefix search (for autocomplete)

## Installation

Install and update this go package with `go get -u github.com/kanocz/go-prefixsearch`

## Examples

```go
spackage main

import (
	"fmt"

	"github.com/kanocz/go-prefixsearch"
)

type someInfo struct {
	ID      int
	Name    string
	Comment string
}

func main() {

	data := []someInfo{
		{1, "Hello world!", "First example"},
		{2, "New impressions", "Second example"},
		{3, "Hello golang!", "Some other important info"},
		{4, "Привет, мир!", "Unicode symbols also work"},
		{5, "こんにちは世界", "Even this one may work"},
	}

	st := prefixsearch.New()
	for _, x := range data {
		st.Add(x.Name, x)
	}

	fmt.Println(st.AutoComplete("HE"))
  // Output: [{1 Hello world! First example} {3 Hello golang! Some other important info}]
	fmt.Println(st.AutoComplete("こん"))
  // Output: [{5 こんにちは世界 Even this one may work}]

	fmt.Println(st.Search("HE"))
  // Output: <nil>
	fmt.Println(st.Search("HELLO WORLD!"))
  // Output: {1 Hello world! First example}
	fmt.Println(st.Search("HELLO WORLD!!"))
  // Output: <nil>
}
```

[godoc]: http://godoc.org/github.com/tatsushid/go-fastping
