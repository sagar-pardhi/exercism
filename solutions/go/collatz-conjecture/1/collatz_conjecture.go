package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	steps := 0
	result := n

	for {
		// fmt.Println("result: ", result)
		if result == 1 {
			break
		}

		if result == 0 {
			return 0, errors.New("zero is an error")
		}

		if result < 0 {
			return 0, errors.New("negative number is an error")
		}

		if result%2 == 0 {
			result = result / 2
			steps += 1
		} else {
			result = result*3 + 1
			steps += 1
		}

	}

	// fmt.Println("steps: ", steps)
	return steps, nil
}
