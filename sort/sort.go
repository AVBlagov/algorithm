package sort

func Bubble(arg *[]int) {
	m := *arg
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i] < m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
	//fmt.Println(arg)
}
