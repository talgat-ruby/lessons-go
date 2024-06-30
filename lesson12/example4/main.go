package main

//mock_data = [
//{"x": 10001, "y": 1},
//{"x": 10004, "y": 2},
//{"x": 10006, "y": 4},
//{"x": 10010, "y": 0},
//{"x": 10017, "y": 3},
//]

func getData() []map[string]int {
	m := []map[string]int{
		{"x": 10001, "y": 1},
		{"x": 10004, "y": 2},
		{"x": 10006, "y": 4},
		{"x": 10010, "y": 0},
		{"x": 10017, "y": 3},
	}
	return m
}

func main() {
	data := getData()

	intervals := make([]int, 0, len(data)-1)

	for i := 1; i < len(data); i++ {

		interval := data[i] - data[i-1]
	}
}
