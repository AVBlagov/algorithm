package sort

func Sort(arr *[]int) {
	m := *arr
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i] < m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
}

func Select(arr *[]int) {
	m := *arr
	for i := 0; i < len(m)-1; i++ {
		min := i
		for j := i + 1; j < len(m); j++ {
			if m[min] > m[j] {
				min = j
			}
		}
		if min != i {
			m[i], m[min] = m[min], m[i]
		}
	}
}

func Bubble(arr *[]int) {
	m := *arr
	for i := 0; i < len(m); i++ {
		for j := len(m) - 1; j > i; j-- {
			if m[j-1] > m[j] {
				m[j], m[j-1] = m[j-1], m[j]
			}
		}
	}
}
