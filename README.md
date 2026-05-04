# Task CLI

A project for https://roadmap.sh/projects/task-tracker

A simple command-line task tracker built in Go.  
It allows you to manage tasks directly from your terminal, storing data in a local JSON file.

---

## Features

- Add, update, and delete tasks  
- Mark tasks as **in progress** or **done**  
- List all tasks or filter by status  
- Persistent storage using `tasks.json`  
- Simple and minimal CLI interface  

---

## Installation

Clone the repository and build the binary:

```bash
git clone <your-repo-url>
cd task-cli
go build -o task-cli
```

---

## Usage

```bash
task-cli <command> [arguments]
```

If no command is provided, a usage message is shown.

---

## Commands

### Add a task
```bash
task-cli add "Buy groceries"
```

### Update a task
```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task
```bash
task-cli delete 1
```

### Mark task as in progress
```bash
task-cli mark-in-progress 1
```

### Mark task as done
```bash
task-cli mark-done 1
```

### List tasks
```bash
task-cli list
```

### List tasks by status
```bash
task-cli list done
task-cli list in-progress
task-cli list todo
```

---

## Data Storage

Tasks are stored in a local `tasks.json` file.  
Each task contains:

- ID  
- Description  
- Status (`todo`, `in-progress`, `done`)  
- Created and updated timestamps  

---

## Project Structure

- `main.go` – CLI entry point and command routing  
- `commands.go` – command definitions and registry  
- `tasks.go` – data model and JSON persistence  
- `command_*.go` – individual command implementations  

---

## Example Output

```bash
[ID: 1] Buy groceries (todo)
[ID: 2] Finish project (in-progress)
[ID: 3] Read book (done)
```

---

## Future Improvements

- Add due dates and priorities  
- Improve filtering and sorting  
- Add tests  
- Support config or custom storage location  
