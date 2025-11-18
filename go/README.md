# Go LeetCode Solutions

This directory contains Go solutions to LeetCode problems, organized by difficulty.

## Structure

```
.
├── easy/
│   └── solutions.go      # Easy difficulty problems
├── medium/
│   └── solutions.go      # Medium difficulty problems
├── hard/
│   └── solutions.go      # Hard difficulty problems
├── golangReference.go    # Go scratchpad for quick testing
└── draft.go              # Working file (gitignored)
```

## How to Use

- Solutions are organized by difficulty level
- `draft.go` is for working on new solutions before moving them to the appropriate difficulty file
- `golangReference.go` can be used for quick Go syntax testing and experimentation

## Running Code

```bash
# Run the reference file
go run golangReference.go

# Run tests for a specific package
cd easy
go test
```
