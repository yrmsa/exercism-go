package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("error")
	}

	// bottom-up approach (?)
	arr := make([]int, 0, n)

	for num := 1; num <= n; num++ {

		if num == 1 || num == 2 {
			arr = append(arr, num-1)
			continue
		}

		counter := 0
		i := num
		for i > 1 {
			// check if index is exists
			if i >= 0 && i-1 < len(arr) {
				counter += arr[i-1]
				i = 1
				continue
			} 

			if i % 2 == 0 {
				i = i / 2
			} else {
				i = (i * 3) + 1
			}
			counter++
		}

		arr = append(arr, counter)
	}

	return arr[n-1], nil
}