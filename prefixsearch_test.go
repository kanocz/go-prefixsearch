package prefixsearch_test

import (
	"fmt"
	"sort"

	"github.com/kanocz/go-prefixsearch"
)

// ExampleAutoComplete just creates an object and does simple test
func ExampleSearchTree_AutoComplete() {
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

// Support of unicode symbols and using struct as value
func ExampleSearchTree_AutoComplete_unicode() {
	data := []struct {
		ID      int
		Name    string
		Comment string
	}{
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

	fmt.Println(st.AutoComplete("こん"))
	// Output: [{5 こんにちは世界 Even this one may work}]
}

// ExampleSearch shows another possible usage of this package
func ExampleSearchTree_Search() {
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
