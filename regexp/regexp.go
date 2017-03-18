package main

import (
	"regexp"
	"fmt"
	"bytes"
)

// https://gobyexample.com/regular-expressions
func main() {

	// 簡単に含まれるか確認
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	// パターン作成
	r, _ := regexp.Compile("p([a-z]+)ch")

	// パターンが含まれるか
	fmt.Println(r.MatchString("peach")) // true

	// パターンを検索
	fmt.Println(r.FindString("peach punch")) // peach

	// パターンを検索して、その文字のインデックスを取得
	fmt.Println(r.FindStringIndex("peach punch")) // [0 5]

	// 最初にヒットしたものと、中のグループを取り出す
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// Similarly this will return information about the
	// indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex( // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(r.Match([]byte("peach"))) // true

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
