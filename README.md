# `GO TODO CLI`
A light-weight, fast, in-memory todo-list command-line application, written in GO.

## Usage

```bash
AVAILABLE COMMANDS:

todo > new: [item]
todo > list
todo > done: [index]
todo > remove: [index]
todo > clear
todo > quit
```

### 1. Adding an item
add an item to the todo-list, use `new:` or `n:`:
```bash
todo > new: [item]
todo > n: [item]
```

### 2. Listing all items
to list every item in the todo-list, use `list:` or `l:`:
```bash
todo > list
todo > l
```

### 3. Completing to-dos
To complete a to-do, use `done:` or `d:`:
```bash
todo > done: [number] or all
todo > d: [number] or all
```

### 4. Removing to-dos
to remove a to-do, use `remove:` or `r:`:
```bash
todo > remove: [number]
todo > r: [number]
```

### 5. Clearing all items
To clear the todo-list, use `clear` or `c`:
```bash
todo > clear
todo > c
```

### 6. Quitting the program
To quit the program, use `quit` or `q`:
```bash
todo > quit
todo > q
```
