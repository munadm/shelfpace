package plan

import (
	"fmt"
	"io"
	"time"
)

// Split divides the total number of pages into a slice of integers representing the number of pages to read each day over a given number of days.
// It distributes the pages as evenly as possible, with any remainder distributed one page at a time to the first few days.
func Split(pages, days int) ([]int, error) {
	if days <= 0 {
		return nil, fmt.Errorf("days must be greater than 0, got %d", days)
	}
	if pages < 0 {
		return nil, fmt.Errorf("pages must be non-negative, got %d", pages)
	}
	base := pages / days
	rem := pages % days
	result := make([]int, days) // starts an slice of length days with all values initialized to 0
	for i := 0; i < days; i++ {
		result[i] = base
		if i < rem {
			result[i]++
		}
	}
	return result, nil
}

// WritePlan is a function that writes the dated page table to an io.Writer
func WritePlan(pages []int, today time.Time, w io.Writer) error {
	if len(pages) == 0 {
		return fmt.Errorf("pages must be non-negative, got %d", pages)
	}
	// print the total number of pages to read
	total := 0

	for i, p := range pages {
		day := today.AddDate(0, 0, i)
		fmt.Fprintf(w, "Day %d (%s): %d pages\n", i+1, day.Format("2006-01-02"), p)
		total += p
	}

	fmt.Fprintf(w, "Total pages: %d\n", total)
	return nil
}
