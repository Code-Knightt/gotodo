# GoToDo

```
Create, manage, complete and delete your tasks
right from the command line. Powered by GO and cobra-cli

Usage:
  gotodo [command]

Available Commands:
  add         Add a task
  complete    Complete a task
  delete      Delete a task from the list
  help        Help about any command
  list        List all tasks

Flags:
  -h, --help     help for gotodo

Use "gotodo [command] --help" for more information about a command.
```

### Install

```
go install github.com/Code-knightt/gotodo
```

### Add task

```
Create a new task

Examples:
  gotodo add [Task_Name]
```

### Complete task

```
Complete a task by inputting the ID

Examples:
  gotodo complete [ID]
```

### Delete task

```
Delete a task by inputting the ID

Examples:
  gotodo delete [ID]
```

### List all tasks

```
List all your pending tasks

Usage:
  gotodo list [flags]
```
