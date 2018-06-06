package main

import (
	"fmt"
	"strings"
)

func main() {
	q := "carp lake flooding index:test"

	qp, i := getIndex(q)

	fmt.Println(qp)
	fmt.Println(i)

}

func getIndex(q string) (string, string) {
	// token string on spaces
	qs := strings.Split(q, " ")
	index := ""
	// check each string for a token
	for i := range qs {
		if strings.Contains(qs[i], "index:") {
			is := strings.Split(qs[i], ":")
			index = is[1] // get the second index item which would be the index:value value we are after
			// remove from qs index
			qs = removeIndex(qs, i)
		}
	}

	return strings.Join(qs, " "), index
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
