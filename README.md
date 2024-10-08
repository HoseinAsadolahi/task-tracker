# Task Tracker CLI

## Description
- Complete and simple solution for [roadmap.sh](https://roadmap.sh) backend [task tracker cli project](https://roadmap.sh/projects/task-tracker)
- This is a simple command line interface to manage your daily task, mark them as todo or in-progress or done, list all of the tasks or list them by specific status.
- I used [cobra](https://github.com/spf13/cobra) for cli and [lip gloss](https://github.com/charmbracelet/lipgloss) for styling the text

## How to run
- clone the project
```shell
git clone https://github.com/HoseinAsadolahi/task-tracker.git
cd task-tracker
```
- To run this use commands below
```shell
go build -o task-tracker
or
go build -o task-tracker.exe
```
then
```shell
# Adding a new task
./<file-name> add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./<file-name> update 1 "Buy groceries and cook dinner"
./<file-name> delete 1

# Marking a task as in progress or done
./<file-name> mark-in-progress 1
./<file-name> mark-done 1

# Listing all tasks
./<file-name> list

# Listing tasks by status
./<file-name> list done
./<file-name> list todo
./<file-name> list in-progress
```