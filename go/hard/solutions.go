package hard

/*
============================================================

============================================================
Time Complexity: O()
Space Complexity: O()
*/

/*
============================================================
Count Ways to Build Good Strings
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func assignEdgeWeights(edges [][]int) int {
	const mod = 1_000_000_007
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	depthSlice := make([]int, n)
	seenSlice := make([]bool, n)
	queue := []int{0}
	seenSlice[0] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if !seenSlice[neighbor] {
				seenSlice[neighbor] = true
				depthSlice[neighbor] = depthSlice[node] + 1
				queue = append(queue, neighbor)
			}
		}
	}
	maxDepth := -1
	for _, depth := range depthSlice {
		if depth > maxDepth {
			maxDepth = depth
		}
	}
	output := 1
	for i := 0; i < maxDepth-1; i++ {
		output = (output * 2) % mod
	}
	return output
}

/*
============================================================
Find Coins
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func findCoins(numWays []int) []int {
	n := len(numWays)
	output := []int{}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if dp[i] == numWays[i] {
			continue
		}
		if dp[i] != numWays[i]-1 {
			return []int{}
		}
		dp[i]++
		coinValue := i + 1
		output = append(output, coinValue)
		for j := 0; j < n; j++ {
			if dp[j] != 0 && (j+coinValue) < n {
				dp[j+coinValue] += dp[j]
			}
		}
	}
	return output
}

/*
============================================================
Maximal Square
============================================================
Time Complexity: O(?)
Space Complexity: O(?)
*/
func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	intMatrix := make([][]int, rows)
	for i := range matrix {
		intMatrix[i] = make([]int, cols)
		for j := range matrix[i] {
			intMatrix[i][j] = int(matrix[i][j]) - 48
		}
	}
	output := 0
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minf := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if intMatrix[y][x] > 0 &&
				x > 0 &&
				y > 0 &&
				intMatrix[y-1][x-1] > 0 &&
				intMatrix[y-1][x] > 0 &&
				intMatrix[y][x-1] > 0 {
				intMatrix[y][x] = minf(intMatrix[y-1][x], minf(intMatrix[y-1][x-1], intMatrix[y][x-1])) + 1
			}
			output = maxf(output, intMatrix[y][x])
		}
	}
	return output * output
}
