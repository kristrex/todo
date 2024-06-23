package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	lst := *t
	if index <= 0 || index > len(lst) {
		return errors.New("invalid index")
	}
	lst[index-1].CompletedAt = time.Now()
	lst[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	lst := *t
	if index <= 0 || index > len(lst) {
		return errors.New("invalid index")
	}
	*t = append(lst[:index-1], lst[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File not Exist!!!")
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for index, item := range *t {
		index++
		task := makeBlue(item.Task)
		done := makeRed("no")
		timeCreated := makeGray(item.CreatedAt.Format(time.RFC822))
		timeCompleted := makeGray(item.CompletedAt.Format(time.RFC822))

		if item.Done {
			task = makeGreen(fmt.Sprintf("\u2705 %s", item.Task))
			done = makeGreen("yes")
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: done},
			{Text: timeCreated},
			{Text: timeCompleted},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	compl, notcompl := t.CountTask()
	compltask := makeGreen(strconv.Itoa(compl))
	notcompltask := makeRed(strconv.Itoa(notcompl))

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("Task completed %s\nUnfulfilled task %s", compltask, notcompltask)},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountTask() (int, int) {
	totalcompl := 0
	total := 0
	for _, item := range *t {
		if item.Done {
			totalcompl++
		} else {
			total++
		}
	}
	return total, totalcompl
}
