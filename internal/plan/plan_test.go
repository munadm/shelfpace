package plan_test

import (
	"bytes"
	"flag"
	"os"
	"slices"
	"testing"
	"time"

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

var update = flag.Bool("update", false, "rewrite golden files")

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

func TestWritePlan(t *testing.T) {
	var writePlanTests = []struct {
		name   string
		pages  []int
		golden string
	}{
		{name: "simple case", pages: []int{4, 3, 3}, golden: "testdata/writeplan_simple.golden"},
		{name: "zero pages", pages: []int{0, 0, 0}, golden: "testdata/writeplan_zero.golden"},
		{name: "single day", pages: []int{10}, golden: "testdata/writeplan_single.golden"},
	}
	t.Run("WritePlan", func(t *testing.T) {
		for _, tt := range writePlanTests {
			t.Run(tt.name, func(t *testing.T) {
				var buf bytes.Buffer
				today := time.Date(2024, 6, 29, 0, 0, 0, 0, time.UTC)
				err := plan.WritePlan(tt.pages, today, &buf)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				got := buf.String()
				if *update {
					if err := os.WriteFile(tt.golden, []byte(got), 0o644); err != nil {
						t.Fatalf("failed to write golden file: %v", err)
					}
				}
				wantBytes, err := os.ReadFile(tt.golden)
				if err != nil {
					t.Fatalf("failed to read golden file: %v", err)
				}
				want := string(wantBytes)
				if got != want {
					t.Errorf("WritePlan() = %q, want %q", got, want)
				}
			})
		}
	})
}
