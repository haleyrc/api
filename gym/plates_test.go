package gym

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/haleyrc/api/log"
)

func TestBuildTree(t *testing.T) {
	log.SetDebug(true)
	tree, err := buildTree(
		context.Background(),
		NewPlateSet([]Plate{
			{Weight: 45},
			{Weight: 35},
			{Weight: 25},
			{Weight: 10},
			{Weight: 5},
			{Weight: 5},
			{Weight: 2.5},
		}),
	)
	t.Log(len(tree))
	if err != nil {
		t.Error("unexpected error:", err)
	}
	saveResult(t, "tree.json", struct {
		Tree map[float64]PlateSet
	}{tree})
}

func TestPermutePlates(t *testing.T) {
	log.SetDebug(true)
	plates := []Plate{
		{Weight: 45},
		{Weight: 35},
		{Weight: 25},
		{Weight: 10},
		{Weight: 5},
		{Weight: 5},
		{Weight: 2.5},
	}
	result := permutePlates(
		context.Background(),
		[][]Plate{{}},
		plates,
	)
	t.Log(len(result))
	saveResult(t, "permutations.json", result)
	want := int(math.Pow(2, float64(len(plates))))
	got := len(result)
	if got != want {
		t.Errorf("should have %d results, but got %d", want, got)
	}
}

func TestDedupePermutations(t *testing.T) {
	log.SetDebug(true)
	ctx := context.Background()
	perms := permutePlates(
		context.Background(),
		[][]Plate{{}},
		[]Plate{
			{Weight: 45},
			{Weight: 35},
			{Weight: 25},
			{Weight: 10},
			{Weight: 5},
			{Weight: 5},
			{Weight: 2.5},
		},
	)
	deduped := dedupePermutations(ctx, perms)
	saveResult(t, "deduped.json", deduped)
	t.Log(len(deduped))
}

func saveResult(t *testing.T, fn string, i interface{}) {
	f, err := os.Create(fn)
	if err != nil {
		t.Logf("error saving %s: %v", fn, err)
		return
	}
	defer f.Close()

	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		t.Logf("error saving %s: %v", fn, err)
		return
	}

	fmt.Fprintln(f, string(b))
}
