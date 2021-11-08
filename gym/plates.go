package gym

import (
	"context"

	"github.com/haleyrc/api/log"
)

type Plate struct {
	Weight float64
	Bumper bool
}

func NewPlateSet(plates []Plate) PlateSet {
	return PlateSet{Plates: plates}
}

type PlateSet struct {
	Plates []Plate
}

func (p PlateSet) Len() int {
	return len(p.Plates)
}

func (p PlateSet) Less(i, j int) bool {
	return p.Plates[i].Weight < p.Plates[j].Weight
}

func (p *PlateSet) Swap(i, j int) {
	p.Plates[i], p.Plates[j] = p.Plates[j], p.Plates[i]
}

func (p PlateSet) TotalWeight() float64 {
	return totalWeight(p.Plates)
}

func buildTree(ctx context.Context, plateSet PlateSet) (map[float64]PlateSet, error) {
	plates := plateSet.Plates[:]

	// First we get all of the possible permutations of the plates that we were provided
	permutations := permutePlates(ctx, [][]Plate{{}}, plates)

	// Then, since we're just greedily permuting, we'll end up with duplicates, so we
	// filter those out.
	uniquePermutations := dedupePermutations(ctx, permutations)

	// Next, we build a lookup table matching each of the permutations to their total
	// weight (with weight as the key). This allows us to find the set of plates that
	// matches a target weight in constant time. Since we can be fairly sure that the
	// total number of plates in a given weight set is fairly low, we aren't overly
	// concerned with the edge case of having to permute millions of combinations and
	// this is much simpler than doing it in place every time.
	tree := make(map[float64]PlateSet)
	for _, p := range uniquePermutations {
		log.Debug("Plates:", p)
		weight := totalWeight(p)
		log.Debug("Weight:", weight)
		tree[weight] = PlateSet{Plates: p}
	}

	return tree, nil
}

func totalWeight(plates []Plate) float64 {
	total := float64(0)
	for _, p := range plates {
		total += p.Weight
	}
	return total
}

func dedupePermutations(ctx context.Context, permutations [][]Plate) [][]Plate {
	set := map[float64][]Plate{}
	for _, perm := range permutations {
		wt := totalWeight(perm)
		if existing, exists := set[wt]; exists {
			if len(perm) < len(existing) {
				set[wt] = perm
			}
		}
		set[wt] = perm
	}
	deduped := make([][]Plate, 0, len(set))
	for _, v := range set {
		deduped = append(deduped, v)
	}
	return deduped
}

func permutePlates(ctx context.Context, initial [][]Plate, plates []Plate) [][]Plate {
	log.Debug("Plates:", plates)

	// Our resulting set of solutions is twice as long as the set we already have,
	// since all we're doing is adding the next available weight to each entry and
	// taking the union of those and the initial set.
	union := make([][]Plate, 0, 2*len(initial))

	// First we add all the solutions we already have
	for _, ps := range initial {
		curr := make([]Plate, 0, len(ps))
		for _, p := range ps {
			curr = append(curr, Plate{Weight: p.Weight})
		}
		union = append(union, curr)
	}

	// Now we go through and add the next plate to each of the existing solutions.
	// This could be done in one pass with the loop above, but for clarity's sake
	// at this point, we do it in two steps. Considering the maximum number of plates
	// we can reasonably expect, this isn't likely to cause any issues.
	next := plates[0]
	log.Debug("Current plate:", next)
	for _, ps := range initial {
		log.Debug("Extending set:", ps)
		newSet := append(ps, Plate{Weight: next.Weight})
		log.Debug("Created new set:", newSet)
		union = append(union, newSet)
	}

	// If this was the last plate, we're done
	if len(plates) == 1 {
		return union
	}

	// Otherwise, we do the same process with the new set of solutions as the initial
	// set and the remaining plates as our set.
	return permutePlates(ctx, union, plates[1:])
}

// TODO: Sort plate set before calculating
// TODO: Remove bumpers from regular plate set before calculating
// TODO: Sort results before returning
// TODO: Put bumpers at beginning of sorted results if used
func CalculatePlates(
	ctx context.Context,
	targetWeight float64,
	barWeight float64,
	plateSet []Plate,
	requireBumpers bool,
	preferBumpers bool,
) ([]Plate, error) {
	log.Debug("Target weight: ", targetWeight)
	log.Debug("Bar weight: ", barWeight)
	log.Debug("Plate set: ", plateSet)
	log.Debug("Require bumpers: ", requireBumpers)
	log.Debug("Prefer bumpers: ", preferBumpers)

	// Subtract bar from total
	log.Debug("Subtracting bar weight...")
	total := targetWeight - barWeight
	log.Debug("New target: ", total)

	// Divide remainder by 2
	log.Debug("Calculating weight for one side...")
	total = total / 2
	log.Debug("New target: ", total)

	// Subtract bumpers if required
	if requireBumpers {
		// TODO
	}

	// TODO: Handler prefer bumpers by calculating with bumper weight removed if possible

	// Calculate plates for remainder
	// TODO: Sort plates first
	result := []Plate{}
	for _, plate := range plateSet {
		if plate.Weight <= total {
			result = append(result, plate)
			log.Debug("Added plate to result: ", plate)
			total -= plate.Weight
			log.Debug("New target: ", total)
		}
	}

	resultWeight := (totalWeight(result) * 2) + barWeight
	log.Debug("Achieved total weight: ", resultWeight)

	delta := targetWeight - resultWeight
	deltaPerc := (delta / targetWeight) * 100
	log.Debugf("Delta: %f (%f%%)", delta, deltaPerc)

	return result, nil
}
