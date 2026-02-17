package mission

import (
	"testing"
	"time"
)

func TestAddTaskValidation(t *testing.T) {
	plan, err := NewDayPlan(time.Now(), 30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := plan.AddTask("", 3); err != ErrEmptyTitle {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}

	if err := plan.AddTask("Графы", 6); err != ErrInvalidLevel {
		t.Fatalf("expected ErrInvalidLevel, got %v", err)
	}
}

func TestCompleteTaskAndMomentum(t *testing.T) {
	plan, err := NewDayPlan(time.Now(), 120)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_ = plan.AddTask("Task A", 3)
	_ = plan.AddTask("Task B", 4)

	if err := plan.CompleteTask(0); err != nil {
		t.Fatalf("unexpected error on complete: %v", err)
	}

	score := plan.MomentumScore()
	if score != 55 { // 35 from completion + 20 from focus
		t.Fatalf("expected score 55, got %d", score)
	}
}

func TestPriorityView(t *testing.T) {
	plan, _ := NewDayPlan(time.Now(), 10)
	_ = plan.AddTask("easy", 1)
	_ = plan.AddTask("hard", 5)
	_ = plan.AddTask("medium", 3)

	ordered := plan.PriorityView()
	if ordered[0].Title != "hard" || ordered[1].Title != "medium" || ordered[2].Title != "easy" {
		t.Fatalf("priority order is invalid: %#v", ordered)
	}
}
