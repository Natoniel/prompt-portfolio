package main

import (
	"fmt"
	"log"
	"time"

	"github.com/example/learnvoyage/internal/mission"
)

func main() {
	plan, err := mission.NewDayPlan(time.Now(), 90)
	if err != nil {
		log.Fatal(err)
	}

	_ = plan.AddTask("Прочитать главу по структурам данных", 4)
	_ = plan.AddTask("Решить 2 задачи на указатели", 5)
	_ = plan.AddTask("Написать мини-конспект", 2)
	_ = plan.CompleteTask(0)

	fmt.Println("=== LearnVoyage: учебная экспедиция дня ===")
	fmt.Printf("Дата: %s\n", plan.Date.Format("2006-01-02"))
	fmt.Printf("Фокус-время: %d мин\n", plan.FocusMinutes)
	fmt.Printf("Прогресс: %d/%d задач\n", plan.CompletedCount, len(plan.Tasks))
	fmt.Printf("Momentum Score: %d/100\n\n", plan.MomentumScore())

	fmt.Println("Приоритетный список:")
	for i, task := range plan.PriorityView() {
		status := "⏳"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s [%d/5] %s\n", i+1, task.Title, task.Difficulty, status)
	}
}
