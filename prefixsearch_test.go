package prefixsearch_test

import (
	"fmt"
	"sort"

	"github.com/kanocz/go-prefixsearch"
)

// ExampleAutoComplete just creates an object and does simple test
func ExampleAutoComplete() {
	st := prefixsearch.New()
	st.Add("Hello world!", 1)
	st.Add("New impressions", 2)
	st.Add("Hello golang!", 3)
	st.Add("Привет, мир!", 4)
	st.Add("Надо же :)", 5)

	// converting []interface{} to []int to use sort.Ints function :)
	intResult := []int{}
	for _, x := range st.AutoComplete("HE") {
		intResult = append(intResult, x.(int))
	}

	sort.Ints(intResult)
	fmt.Println(intResult)
	// Output: [1 3]
}

// ExampleSearch shows another possible usage of this package
func ExampleSearch() {
	st := prefixsearch.New()
	st.Add("Hello world!", 1)
	st.Add("New impressions", 2)
	st.Add("Hello golang!", 3)
	st.Add("Привет, мир!", 4)
	st.Add("Надо же :)", 5)

	fmt.Println(st.Search("HE"))
	fmt.Println(st.Search("HELLO WORLD"))
	fmt.Println(st.Search("HELLO WORLD!"))
	fmt.Println(st.Search("HELLO WORLD!!"))

	// Output:
	// <nil>
	// <nil>
	// 1
	// <nil>

}
