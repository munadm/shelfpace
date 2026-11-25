package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/munadm/shelfpace/internal/plan"
)

func main() {
	// define usage if args are not correct
	if len(os.Args) < 2 || os.Args[1] != "pace" {
		fmt.Fprintln(os.Stderr, "usage: shelfpace pace --pages N --at N --by YYYY-MM-DD")
		os.Exit(2)
	}
	// define flags for pace command
	paceCmd := flag.NewFlagSet("pace", flag.ExitOnError)
	pacePages := paceCmd.Int("pages", 0, "Number of pages to read")
	paceAt := paceCmd.Int("at", 0, "How many pages have already been read")
	paceBy := paceCmd.String("by", "", "Deadline date")
	err := paceCmd.Parse(os.Args[2:])

	// set conditions for flags and exit if not met
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing flags:", err)
		os.Exit(2)
	}
	if *paceBy == "" {
		fmt.Fprintln(os.Stderr, "Error: --by flag is required")
		os.Exit(2)
	}
	if *pacePages <= 0 {
		fmt.Fprintln(os.Stderr, "Error: --pages must be greater than 0")
		os.Exit(2)
	}
	if *paceAt < 0 {
		fmt.Fprintln(os.Stderr, "Error: --at must be a non-negative integer")
		os.Exit(2)
	}
	if *pacePages < *paceAt {
		fmt.Fprintln(os.Stderr, "Error: --pages must be greater than or equal to --at")
		os.Exit(1)
	}
	// establish the date to pace
	by, err := time.Parse("2006-01-02", *paceBy)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing date:", err)
		os.Exit(1)
	}
	// calculate the number of days between today and the deadline
	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	days := int(by.Sub(today).Hours() / 24)
	if days <= 0 {
		fmt.Fprintln(os.Stderr, "Error: --by date must be in the future")
		os.Exit(1)
	}
	// split the pages into a plan based on the number of days
	perDay, err := plan.Split(*pacePages-*paceAt, days)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	// print the plan
	total := 0
	for i, pages := range perDay {
		day := today.AddDate(0, 0, i)
		fmt.Printf("%s: %d pages\n", day.Format("Mon Jan 2"), pages)
		total += pages
	}

	fmt.Printf("Total pages: %d\n", total)

}
