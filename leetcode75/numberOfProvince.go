package leetcode75

func dfsProvinceCount(isConnected [][]int, visited []bool, i int) {
	visited[i] = true

	for j := range len(isConnected) {
		if isConnected[i][j] == 1 && !visited[j] {
			dfsProvinceCount(isConnected, visited, j)
		}
	}

}
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	provinceCount := 0

	for i := range n {
		if !visited[i] {

			provinceCount++

			dfsProvinceCount(isConnected, visited, i)
		}
	}
	return provinceCount
}
