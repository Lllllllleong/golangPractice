# Go LeetCode Solutions

This directory contains Go solutions to LeetCode problems, organized by difficulty.

## 📊 Statistics

- **Total Solutions**: 34 problems
  - Easy: 13 solutions
  - Medium: 18 solutions
  - Hard: 3 solutions
- **Total Lines**: ~882 lines of Go code

## 📁 Structure

```
.
├── easy/
│   └── solutions.go      # 13 easy difficulty problems
├── medium/
│   └── solutions.go      # 18 medium difficulty problems
├── hard/
│   └── solutions.go      # 3 hard difficulty problems
├── golangReference.go    # Go scratchpad for quick testing
├── draft.go              # Working file (gitignored, in root)
└── go.mod                # Module definition
```

## 🔍 Finding Solutions

### Search Tips
1. **By Problem Name**: Open the appropriate difficulty file and use Ctrl+F
2. **By Pattern**: Search for keywords like `Binary Search`, `DFS`, `DP`
3. **All Solutions**: Check the function names starting with lowercase letters

### Solution Format
Each solution includes:
- Problem name as comment header
- Time and space complexity (O(?) if not analyzed)
- Clean, idiomatic Go code

## 🛠️ How to Use

### Working on New Solutions
1. Write and test in `../draft.go` (root directory, gitignored)
2. Once complete, move to appropriate difficulty file
3. Keep helper functions near their main function

### Running Code
```bash
# Run the reference file for quick testing
go run golangReference.go

# Run all tests (if you have test files)
go test ./...

# Run a specific package
cd easy
go test
```

### Adding New Solutions
1. Start in `draft.go` for work-in-progress
2. Use proper comment block format:
   ```go
   /*
   ============================================================
   Problem Name Here
   ============================================================
   Time Complexity: O(?)
   Space Complexity: O(?)
   */
   func solutionName(params) returnType {
       // implementation
   }
   ```
3. Move to `easy/`, `medium/`, or `hard/` when complete
4. Keep helper functions immediately after main function

## 📦 Dependencies

Minimal dependencies - mostly standard library:
- `fmt` - For formatted I/O
- `sort` - For sorting operations
- Standard data structures (maps, slices)

Module: `github.com/Lllllllleong/self-learning-programming/go`
