// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
)

/******** 包没有找到。后面学习到包的知识后再回来修复 ***********/
func main() {
	result, err := github.SearchIssues("repo:golang/go is:open json decoder")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
