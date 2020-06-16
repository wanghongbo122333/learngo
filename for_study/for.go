package forstudy

const N = 1000

func ForSlice(s []string) {
	for i := 0; i < len(s); i++ {
		_, _ = i, s[i]
	}
}
func RangeForSlice(s []string) {
	for i, v := range s {
		_, _ = i, v
	}
}
func ForMap(m map[int]string) {
	for i := 0; i < len(m); i++ {
		_, _ = i, m[i]
	}
}
func RangeforMap(m map[int]string) {
	for i, j := range m {
		_, _ = i, j
	}
}
