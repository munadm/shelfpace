package plan_test

import (
	"slices"
	"testing"

	"github.com/munadm/shelfpace/internal/plan"
)

func assertSplit(t *testing.T, pages, days int, expected []int) {
	t.Helper()
	got, err := plan.Split(pages, days)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(got))
	}
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
func TestSplit(t *testing.T) {
	var splitTests = []struct {
		name    string
		pages   int
		days    int
		want    []int
		wantErr bool
	}{
		{name: "remainder front-loaded", pages: 10, days: 3, want: []int{4, 3, 3}, wantErr: false},
		{name: "equal distribution", pages: 9, days: 3, want: []int{3, 3, 3}, wantErr: false},
		{name: "unequal distribution", pages: 2, days: 3, want: []int{1, 1, 0}, wantErr: false},
		{name: "zero days", pages: 10, days: 0, want: nil, wantErr: true},
		{name: "negative pages", pages: -1, days: 3, want: nil, wantErr: true},
	}
	t.Run("Split", func(t *testing.T) {
		for _, tt := range splitTests {
			t.Run(tt.name, func(t *testing.T) {
				if tt.wantErr {
					_, err := plan.Split(tt.pages, tt.days)
					if err == nil {
						t.Errorf("expected error for pages=%d, days=%d, got nil", tt.pages, tt.days)
					}
				} else {
					assertSplit(t, tt.pages, tt.days, tt.want)
				}
			})
		}
	})
}
