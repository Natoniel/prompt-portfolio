package mission

import (
	"errors"
	"sort"
	"strings"
	"time"
)

// Task описывает учебную задачу с оценкой сложности.
// Подсказка: попробуйте добавить поле Category и сделать фильтрацию по темам.
type Task struct {
	Title      string
	Difficulty int // от 1 (легко) до 5 (сложно)
	Completed  bool
}

// DayPlan хранит план на день и текущую статистику.
// Подсказка: можно расширить структуру полем Notes для дневника обучения.
type DayPlan struct {
	Date           time.Time
	Tasks          []Task
	FocusMinutes   int
	CompletedCount int
}

var (
	ErrEmptyTitle      = errors.New("task title cannot be empty")
	ErrInvalidLevel    = errors.New("difficulty must be between 1 and 5")
	ErrTaskOutOfRange  = errors.New("task index out of range")
	ErrNegativeMinutes = errors.New("focus minutes cannot be negative")
)

// NewDayPlan создает стартовый учебный план.
func NewDayPlan(date time.Time, focusMinutes int) (DayPlan, error) {
	if focusMinutes < 0 {
		return DayPlan{}, ErrNegativeMinutes
	}
	return DayPlan{Date: date, FocusMinutes: focusMinutes, Tasks: []Task{}}, nil
}

// AddTask добавляет задачу в учебный план.
func (p *DayPlan) AddTask(title string, difficulty int) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrEmptyTitle
	}
	if difficulty < 1 || difficulty > 5 {
		return ErrInvalidLevel
	}

	p.Tasks = append(p.Tasks, Task{Title: title, Difficulty: difficulty})
	return nil
}

// CompleteTask отмечает задачу выполненной и обновляет счетчик.
func (p *DayPlan) CompleteTask(index int) error {
	if index < 0 || index >= len(p.Tasks) {
		return ErrTaskOutOfRange
	}
	if !p.Tasks[index].Completed {
		p.Tasks[index].Completed = true
		p.CompletedCount++
	}
	return nil
}

// PriorityView возвращает задачи по убыванию сложности.
// Подсказка: можно заменить сортировку на алгоритм, учитывающий дедлайны.
func (p DayPlan) PriorityView() []Task {
	cloned := make([]Task, len(p.Tasks))
	copy(cloned, p.Tasks)
	sort.SliceStable(cloned, func(i, j int) bool {
		return cloned[i].Difficulty > cloned[j].Difficulty
	})
	return cloned
}

// MomentumScore рассчитывает "энергию дня" от 0 до 100.
// Формула: 70% — доля выполненных задач, 30% — фокус-минуты (с насыщением на 180 мин).
func (p DayPlan) MomentumScore() int {
	if len(p.Tasks) == 0 {
		if p.FocusMinutes >= 180 {
			return 30
		}
		return (p.FocusMinutes * 30) / 180
	}

	completion := (float64(p.CompletedCount) / float64(len(p.Tasks))) * 70
	focus := float64(min(p.FocusMinutes, 180)) / 180 * 30
	return int(completion + focus)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
