Simple To Do List (Go)

A lightweight command-line To Do List application written in Go that lets you add, list, and mark to do items as done. Items are stored persistently in a JSON file (`tasks.json`).

--------------------------------------------------------
## Features

- Add to-do items with optional notes
- List all to-do items with their status (done/not done)
- Mark to-do items as done by their ID
- Persist items between runs in a JSON file

--------------------------------------------------------

## Clone
git clone https://github.com/your-username/toDoList-Go.git
cd toDoList-Go

## Examples
go build -o todo
./todo --add "Task Name" --note "Note for the task"
./todo --list
./todo --done 1 //1 = the tast number/index

or

go run main.go --add "Task Name" --note "Note for the task"
go run main.go --list
go run main.go --done 1 //1 = the tast number/index
