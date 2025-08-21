package main

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return []int{}
	}
	stack := []int{}

	for _, value := range asteroids {
		if len(stack) > 0 && value < 0 {
			for last := len(stack) - 1; last >= 0; last-- {
				if stack[last] > 0 && stack[last] <= (value*-1) {
					stack = stack[:last]
				} else {
					stack = append(stack, value)
				}
			}

		} else {
			stack = append(stack, value)
		}

	}

	return stack
}

// alive is much better than asteroid = 0
func asteroidCollision3(asteroids []int) []int {
	var stack []int

	for _, asteroid := range asteroids {
		alive := true

		for alive && len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
			top := stack[len(stack)-1]

			if top < -asteroid {
				stack = stack[:len(stack)-1]
			} else if top == -asteroid {
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				alive = false
			}
		}

		if alive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
