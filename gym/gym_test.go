package gym_test

import (
	"context"
	"testing"

	"github.com/haleyrc/api/gym"
	"github.com/haleyrc/api/log"
)

func TestCalculatePlates(t *testing.T) {
	log.SetDebug(true)

	ctx := context.Background()
	barWeight := float64(45.0)
	targetWeight := float64(225)
	plateSet := []gym.Plate{
		{Weight: 45},
		{Weight: 35},
		{Weight: 25},
		{Weight: 10, Bumper: true},
		{Weight: 10},
		{Weight: 5},
		{Weight: 5},
	}
	requireBumpers := false
	preferBumpers := false

	plates, err := gym.CalculatePlates(ctx, targetWeight, barWeight, plateSet, requireBumpers, preferBumpers)
	if err != nil {
		t.Errorf("CalculatePlates failed: %v", err)
	}
	t.Log(plates)
}
