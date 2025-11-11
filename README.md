# ToDo List CLI
*A CLI based todo list written in Go*

## Functionality 
- Using CLI commands, view, add, update, mark as complete, and delete ToDos
- View in tabular, colourful format within the CLI
- Persinstant data via .csv or SQLlite DB

## Format

| Field | Type |
|----------|----------|
| Id | int64 |
| Name | string |
| IsComplete | boolean |
| CreatedDate | date (relative) |

## Order of Ops
1. List
2. Lock / Read CSV
3. Filter
4. Convert to Printable formate
5. Print to terminal


