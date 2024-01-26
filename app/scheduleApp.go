package app

import (
	"fmt"
	"math/rand"
	"sync"
)

type Schedule struct {
	Day    string
	Person string
}

func GenerateSchedule(days []string, person []string, ch chan<- Schedule, wg *sync.WaitGroup) {
	defer wg.Done()

	personMap := make(map[string]bool)

	for _, day := range days {
		personIndex := -1

		for personIndex == -1 || personMap[person[personIndex]] {
			personIndex = rand.Intn(len(person))
		}

		personMap[person[personIndex]] = true

		jadwal := Schedule{
			Day:    day,
			Person: person[personIndex],
		}

		ch <- jadwal
	}

	close(ch)
}

func PrintSchedule(ch <-chan Schedule, wg *sync.WaitGroup) {
	defer wg.Done()

	for jadwal := range ch {
		fmt.Printf("Day %s: Piket oleh %s\n", jadwal.Day, jadwal.Person)
	}
}
