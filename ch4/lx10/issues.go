// 将符合搜索条件的 issue 输出为一个表格
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	lessThanOneMonthOld := make([]*github.Issue, 0)
	lessThanOneYearOld := make([]*github.Issue, 0)
	moreThanOneYearOld := make([]*github.Issue, 0)
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreateAt.After(oneMonthAgo) {
			lessThanOneMonthOld = append(lessThanOneMonthOld, item)
		} else if item.CreateAt.After(oneYearAgo) {
			lessThanOneYearOld = append(lessThanOneYearOld, item)
		} else {
			moreThanOneYearOld = append(moreThanOneYearOld, item)
		}
	}

	fmt.Printf("\nLess than an month old:\n")
	printIssues(lessThanOneMonthOld)
	fmt.Printf("\nLess than a year old:\n")
	printIssues(lessThanOneYearOld)
	fmt.Printf("\nMore than a year old:\n")
	printIssues(moreThanOneYearOld)

}

func printIssues(issues, []*github.Issue) {
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
