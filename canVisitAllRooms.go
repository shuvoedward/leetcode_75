package main

func canVisitAllRoomsBFS(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true
	count := 1

	for len(queue) > 0 {
		currentRoom := queue[0]
		queue = queue[1:]
		for _, key := range rooms[currentRoom] {
			if !visited[key] {
				visited[key] = true
				queue = append(queue, key)
				count++
			}
		}
	}
	return count == n
}

func canVisitAllRoomsDFS(rooms [][]int) bool {
	n := len(rooms)
	stack := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	count := 1

	for len(stack) > 0 {
		currentRoom := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, key := range rooms[currentRoom] {
			if !visited[key] {
				visited[key] = true
				stack = append(stack, key)
				count++
			}
		}
	}

	return count == n
}

// Recursive Depth first Search
func canVisitAllRoomsDFS2(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	count := 0

	var dfs func(int)
	dfs = func(roomID int) {
		if visited[roomID] {
			return
		}

		visited[roomID] = true
		count++

		for _, key := range rooms[roomID] {
			dfs(key)
		}

	}
	dfs(0)
	return count == n
}
