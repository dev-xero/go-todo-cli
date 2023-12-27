# `GO TODO CLI`
A light-weight and fast in-memory todo-list command-line application, written in GO.

## Usage

```bash
AVAILABLE COMMANDS:

todo > new [to-do]
todo > list
todo > check [index]
todo > remove [index]
todo > clear
```

### 1. Adding an item
add an item to the todo-list, use `new` or `:n`:
```bash
todo > new [item]
todo > :n [item]
```

### 2. Listing all items
to list every item in the todo-list, use `list` or `:l`:
```bash
todo > list
todo > :l
```

### 3. Completing to-dos
to complete a to-do, use `done` or `:d`:
```bash
todo > done [index]
todo > :d [index]
```

### 4. Removing to-dos
to remove a to-do, use `remove` or `:r`:
```bash
todo > remove [index]
todo > :r [index]
```

### 5. Clearing all items
to clear the todo-list, use `clear` or `:c`:
```bash
todo > clear
todo > :c
```
