package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type ToDoItem struct {
	Id          int
	Item        string
	IsComplete  bool
	CreatedDate time.Time
}

type ToDoList struct {
	Filename string
	ToDos    []ToDoItem
	NextId   int
	Count    int
}

func InitToDoList(filename string) *ToDoList {
	tdl := &ToDoList{Filename: filename}

	if err := tdl.ReadFromCSV(); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
	}

	return tdl
}

// Convert ToDoList into slice of strings slices for writing to CSV
func (tdl *ToDoList) unpackToDoList() (*[][]string, error) {
	output := make([][]string, tdl.Count)

	for i, v := range tdl.ToDos {
		output[i] = make([]string, 4)
		output[i][0] = strconv.Itoa(v.Id)
		output[i][1] = v.Item
		output[i][2] = strconv.FormatBool(v.IsComplete)
		output[i][3] = v.CreatedDate.Format(time.DateOnly)
	}
	return &output, nil
}

// Write ToDoList to CSV file
func (tdl *ToDoList) WriteToCSV() error {

	// TODO Lock file for wrtiting

	file, err := os.Create(tdl.Filename)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	stringTDL, err := tdl.unpackToDoList()
	if err != nil {
		return err
	}

	writer.Write([]string{"Id", "Item", "IsComplete", "CreatedDate"})

	writer.WriteAll(*stringTDL)
	return nil
}

// Read ToDoList from CSV file
func (tdl *ToDoList) ReadFromCSV() error {

	if _, err := os.Stat(tdl.Filename); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(tdl.Filename)
	if err != nil {
		fmt.Printf("Unable to open file '%v', error: %v\n", tdl.Filename, err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to read file. Error: %v\n", err)
		return err
	}

	// fmt.Printf("Records: %v\n", records)
	// fmt.Printf("Rows in CSV: %v\n", len(records))

	start := 1 // Exclude header

	todos := make([]ToDoItem, 0, len(records)-start)

	for _, v := range records[start:] {
		if len(v) == 0 || (len(v) > 0 && strings.TrimSpace(v[0]) == "") {
			continue
		}

		id, _ := strconv.Atoi(v[0])

		var ic bool = false

		if strings.ToLower(v[2]) == "true" {
			ic = true
		}

		var cd time.Time

		cd, err := time.Parse(time.DateOnly, v[3])
		if err != nil {
			fmt.Printf("Unable to parse date: %v, error: %v\n", v[3], err)
			cd = time.Now()
		}

		todos = append(todos, ToDoItem{
			Id:          id,
			Item:        v[1],
			IsComplete:  ic,
			CreatedDate: cd,
		})

	}
	// fmt.Printf("Rows: %v\n", rows)

	tdl.ToDos = todos
	tdl.Count = len(todos)
	tdl.NextId = len(todos) + 1
	return nil
}

type FilterStatus string

const (
	FilterAll        FilterStatus = "All"
	FilterComplete   FilterStatus = "Complete"
	FilterIncomplete FilterStatus = "Incomplete"
)

type FilterPeriod string

const (
	PeriodAll       FilterPeriod = "All"
	PeriodToday     FilterPeriod = "Today"
	PeriodThisWeek  FilterPeriod = "ThisWeek"
	PeriodThisMonth FilterPeriod = "ThisMonth"
)

type ToDoFilter struct {
	Complete        FilterStatus
	CreatedInPeriod FilterPeriod
}

func bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Filter todo items based on filter criteria
func (tdl *ToDoList) filterToDos(filter *ToDoFilter) *[]ToDoItem {

	outputToDoList := make([]ToDoItem, 0, tdl.Count)

	var periodThreshold time.Time
	if filter.CreatedInPeriod != PeriodAll {
		var now time.Time = bod(time.Now())
		switch filter.CreatedInPeriod {
		case PeriodToday:
			periodThreshold = now.AddDate(0, 0, -1)
		case PeriodThisWeek:
			periodThreshold = now.AddDate(0, 0, -7)
		case PeriodThisMonth:
			periodThreshold = now.AddDate(0, -1, 0)
		}
	}

	for _, v := range tdl.ToDos {
		matchesComplete := filter.Complete == FilterAll ||
			(filter.Complete == FilterComplete && v.IsComplete) ||
			(filter.Complete == FilterIncomplete && !v.IsComplete)

		periodMatches := filter.CreatedInPeriod == PeriodAll ||
			v.CreatedDate.After(periodThreshold)

		if matchesComplete && periodMatches {
			outputToDoList = append(outputToDoList, v)
		}
	}

	return &outputToDoList
}

// Print ToDoList
func (tdl *ToDoList) DisplayTodoList(filter *ToDoFilter) {

	filteredToDos := tdl.filterToDos(filter)

	// const CheckSymbol string = "\u2713"
	const BoldCheckSymbol string = "\u2714"

	fmt.Printf("%-4s %-6s %-45s %-15s\n", "Id", "Done", "Item", "Created")
	fmt.Println(strings.Repeat("-", 90))

	for _, v := range *filteredToDos {
		cb := "[   ]"
		if v.IsComplete {
			cb = fmt.Sprintf("[ %s ]", BoldCheckSymbol)
		}

		fmt.Printf("%-4d %-6s %-45s %-15s\n",
			v.Id,
			cb,
			v.Item,
			v.CreatedDate.Format(time.DateOnly))
	}
}

func (tdl *ToDoList) AddToDo(item string) {

	tdl.ToDos = append(tdl.ToDos, ToDoItem{
		Id:          tdl.NextId,
		Item:        item,
		IsComplete:  false,
		CreatedDate: time.Now(),
	})
	tdl.Count++
	tdl.NextId++
	tdl.WriteToCSV()
	fmt.Printf("Added item '%s' to ToDo list, id: %d\n", item, max(tdl.NextId-1, 0))

}

func (tdl *ToDoList) UpdateToDoItem(Id int, Item string) {

	if Id >= tdl.NextId {
		fmt.Printf("Id not in ToDo list, max Id: %v\n", max(tdl.NextId-1, 0))
		return
	}

	if Item == "" {
		fmt.Println("Item must have value")
	}

	var prevItemName string

	for i, v := range tdl.ToDos {
		if v.Id == Id {
			prevItemName = tdl.ToDos[i].Item
			tdl.ToDos[i].Item = Item
			break
		}
	}
	tdl.WriteToCSV()
	fmt.Printf("Updated '%s' %d to '%s'\n", prevItemName, Id, Item)
}

func (tdl *ToDoList) CheckToDo(Id int, IsComplete bool) {

	if Id >= tdl.NextId {
		fmt.Printf("Id not in ToDo list, max Id: %v\n", max(tdl.NextId-1, 0))
		return
	}
	var updateItem string
	for i, v := range tdl.ToDos {
		if v.Id == Id {
			tdl.ToDos[i].IsComplete = IsComplete
			updateItem = tdl.ToDos[i].Item
			break
		}
	}
	tdl.WriteToCSV()

	if IsComplete {
		fmt.Printf("Checked off '%s' %d\n", updateItem, Id)
	} else {
		fmt.Printf("Unchecked '%s' %d\n", updateItem, Id)
	}
}

func (tdl *ToDoList) DeleteToDo(Id int) {

	if Id >= tdl.NextId {
		fmt.Printf("Id not in ToDo list, max Id: %v\n", max(tdl.NextId-1, 0))
		return
	}

	newToDoList := make([]ToDoItem, 0, tdl.Count)

	var nextId int = 1
	for _, v := range tdl.ToDos {
		if v.Id != Id {
			v.Id = nextId
			newToDoList = append(newToDoList, v)
			nextId++
		}
	}

	tdl.ToDos = newToDoList
	tdl.Count = len(newToDoList)
	tdl.NextId = nextId
	tdl.WriteToCSV()
	fmt.Printf("Deleted Item %d\n", Id)
}
