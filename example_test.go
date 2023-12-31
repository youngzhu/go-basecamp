package basecamp_test

import (
	"github.com/youngzhu/go-basecamp"
	"time"
)

func ExampleAddScheduleEntry() {

	basecamp.AddSchedule("MeTime", "Schedule",
		basecamp.ScheduleEntry{
			Summary:  "!!! Important Meeting",
			StartsAt: time.Date(2023, 11, 5, 8, 0, 0, 0, time.Local),
			EndsAt:   time.Date(2023, 11, 5, 10, 0, 0, 0, time.Local),
		})

	// Output:
	//
}

func ExampleCreateCard() {
	basecamp.AddCard("Profession", "Card Table", "In progress",
		basecamp.Card{
			Title: "Launch a product",
		})

	// Output:
	//
}

func ExampleCreateTodo() {
	basecamp.AddTodo("MeTime", "To-dos", "券",
		basecamp.Todo{
			Content: "Buy mask",
		})

	// Output:
	//
}
