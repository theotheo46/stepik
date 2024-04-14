package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

type TaskArray []Task

func (a TaskArray) Len() int           { return len(a) }
func (a TaskArray) Less(i, j int) bool { return a[i].Dur > a[j].Dur }
func (a TaskArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	//return tasks, errors.New("not implemented")
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	t, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, errors.New("not implemented")
	} else {
		return t, nil
	}
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	re := regexp.MustCompile(`(^\d{1,2}:\d\d) - (\d{1,2}:\d\d) (.+)`)
	var tasks []Task
	tasks_map := map[string]time.Duration{}
	for _, line := range lines {
		groups := re.FindStringSubmatch(line) //len(goups) == 0 - error
		if len(groups) != 4 {
			return nil, errors.New("Строка журнала с задачей не соответствует формату")
		}
		t1 := groups[1]
		t1_time, err := time.Parse("15:04", t1)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("время %s парсится неправильно", t1))
		}
		t2 := groups[2]
		t2_time, err := time.Parse("15:04", t2)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("время %s парсится неправильно", t2))
		}
		task_name := groups[3]
		duration := t2_time.Sub(t1_time)
		if duration <= 0 {
			return nil, errors.New(fmt.Sprintf("время t2 %s <= t1 %s", t2, t1))
		}
		tasks_map[task_name] = tasks_map[task_name] + duration
		//tasks = append(tasks, Task{date, t2_time.Sub(t1_time), task_name})
	}
	for k, v := range tasks_map {
		tasks = append(tasks, Task{date, v, k})
	}
	return tasks, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Sort(TaskArray(tasks))
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 7:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`
	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
