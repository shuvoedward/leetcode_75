package main

import "math"

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

func asteroidCollision2(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid < 0 {
			topAsteroid := stack[len(stack)-1]

			if math.Abs(float64(asteroid)) > float64(topAsteroid) {
				stack = stack[:len(stack)-1] // top asteroid explodes
			} else if math.Abs(float64(asteroid)) == float64(topAsteroid) {
				stack = stack[:len(stack)-1] // both explodes
				asteroid = 0
				break

			} else {
				// current asteroid is smaller than on top of the stack
				asteroid = 0 // asteroid explodes
				break
			}
		}

		if asteroid != 0 {
			stack = append(stack, asteroid)
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
