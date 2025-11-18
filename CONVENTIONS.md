# Coding Conventions & Organization

This document outlines the conventions and organizational rules for this LeetCode practice repository.

## 📁 File Organization

### Java
- **Practice Files**: Solutions are organized in `Practice1.java` through `Practice6.java`
- **New Practice File**: Create a new Practice file when the current one reaches ~4,000 lines (IDE performance threshold)
- **References File**: Common algorithms, utilities, and reference implementations go in `References.java`

### Go
- **By Difficulty**: Solutions are organized by difficulty level
  - `go/easy/solutions.go` - Easy difficulty problems
  - `go/medium/solutions.go` - Medium difficulty problems
  - `go/hard/solutions.go` - Hard difficulty problems
- **Draft File**: Use `draft.go` in the root for work-in-progress solutions (gitignored)
- **Reference File**: Use `go/golangReference.go` for quick testing and Go syntax experiments

### SQL
- All SQL problems in `sql/SQLPractice.sql`

## 📝 Solution Format

### Comment Block Template

Each solution should include a comment block with:

**Java:**
```java
/*
============================================================
Problem Name
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
public ReturnType functionName(parameters) {
    // solution code
}
```

**Go:**
```go
/*
============================================================
Problem Name
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func functionName(parameters) returnType {
    // solution code
}
```

### Guidelines
- **Problem Name**: Use the full problem title (no problem number needed)
- **Complexity**: Leave as `O(?)` if not analyzed yet
- **Function Names**: Use descriptive names based on the problem (camelCase for Go/Java)

## 🔄 Workflow

### Adding New Solutions

1. **Work in Progress**
   - Java: Add to the current Practice file or start testing in `References.java`
   - Go: Use `draft.go` for experimental solutions

2. **Completed Solutions**
   - Java: Keep in the current Practice file
   - Go: Move from `draft.go` to the appropriate difficulty file (`easy/`, `medium/`, `hard/`)
   - Add proper comment block with problem name

3. **Helper Functions**
   - Keep helper functions immediately after the main solution function
   - Use clear naming to indicate they're helpers (e.g., `helperName`, `dfsHelper`)

### Git Workflow
- Commit when adding significant number of solutions or making organizational changes
- Use descriptive commit messages
- Progress screenshots go in `resources/progress/`

## 🎯 Quality Standards

### Code Quality
- Write clean, readable code
- Add comments for complex logic
- Include complexity analysis when possible
- Test solutions before committing

### Organization
- One problem per function (main solution)
- Helper functions stay with their main function
- No orphaned or unused code
- Keep files focused on LeetCode solutions only

## 📊 Progress Tracking

- Take screenshots of LeetCode progress periodically
- Save in `resources/progress/` with date format: `MM-YY.png`
- Update main README when reaching milestones

## 🚫 What NOT to Include

- API testing code
- Non-LeetCode practice projects
- Compiled files (`.class`, binaries)
- IDE-specific files beyond `.idea/` directory
- Personal optimization calculators or unrelated code

## 🔧 IDE & Tools

- **Java**: IntelliJ IDEA configuration in `java/.idea/`
- **Go**: Go modules configured (`go.mod`)
- **Git**: `.gitignore` handles common files
- Performance: Create new Practice file at ~4,000 lines to maintain IDE performance
