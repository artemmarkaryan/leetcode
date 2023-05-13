package romantointeger

var singles = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var doubles = map[[2]string]int{
	{"I", "V"}: 4,
	{"I", "X"}: 9,
	{"X", "L"}: 40,
	{"X", "C"}: 90,
	{"C", "D"}: 400,
	{"C", "M"}: 900,
}

func romanToInt(s string) (sum int) {
	for i := 0; i < len(s); {
		this := s[i : i+1]
		if i == len(s)-1 {
			sum += singles[this]
			break
		}
		next := s[i+1 : i+2]
		if singles[this] < singles[next] {
			sum += doubles[[2]string{this, next}]
			i += 2
		} else {
			sum += singles[this]
			i += 1
		}
	}
	return
}
