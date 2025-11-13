# Go-ToDo
*A simple CLI based todo list written in Go using the Cobra CLI library*

## Usage
- `$ todo` initialize a todo list in the current project directory (if none already exists)\n
- `$ todo help` show commands and options
- `$ todo list` list all todos for current project in a tabular format, responive to the current terminal size. 
    - Options flags:
        - Show only completed items `-c`
        - Show only incomplete items `-i`
        - Show items created today `-t`
        - Show items created this week `-w`
        - Show items created this month `-m`
- `$ todo add <item>` add a todo
- `$ todo update <id> <item description>` update the item description
- `$ todo check <id>` mark an item as complete, or use `-i` to mark it as incomplete
- `$ todo delete <id>` delete a todo item

ToDo items are stored in `<directory name>_ToDoList.csv` in the directory todo is called, allowing for project-scoped todos

## ToDo List Format

| Field | Type |
|----------|----------|
| Id | int64 |
| Item | string |
| IsComplete | boolean |
| CreatedDate | date |



