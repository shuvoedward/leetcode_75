package backtracking

/*
The algorithm uses DFS (Depth-First Search) with two choices at each index:

Include the current element
Exclude the current element

						dfs(0) - subset=[]
                               |
                    ┌──────────┴──────────┐
                    │                     │
               Include 'a'           Exclude 'a'
                    │                     │
              dfs(1) - subset=[a]    dfs(1) - subset=[]
                    |                     |
         ┌──────────┴────────┐   ┌────────┴────────┐
         │                   │   │                 │
    Include 'b'        Exclude 'b'  Include 'b'  Exclude 'b'
         │                   │      │                │
   dfs(2)-[a,b]        dfs(2)-[a]  dfs(2)-[b]    dfs(2)-[]
         |                   |      |                |
    ┌────┴────┐         ┌────┴────┐ ┌───┴────┐   ┌───┴────┐
    │         │         │         │ │        │   │        │
  Inc 'c'  Exc 'c'   Inc 'c'  Exc 'c' Inc 'c' Exc 'c' Inc 'c' Exc 'c'
    │         │         │         │   │        │   │        │
dfs(3)    dfs(3)    dfs(3)    dfs(3) dfs(3)  dfs(3) dfs(3) dfs(3)
[a,b,c]   [a,b]     [a,c]     [a]    [b,c]   [b]    [c]    []
   ↓         ↓         ↓        ↓      ↓       ↓      ↓      ↓
  ADD       ADD       ADD      ADD    ADD     ADD    ADD    ADD
*/

func Subset(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			result = append(result, temp)
			return
		}
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return result
}
