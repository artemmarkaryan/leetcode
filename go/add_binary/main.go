package main

var m = map[byte]int{
	'1': 1,
	'0': 0,
}

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	inc := 0
	r := make([]byte, 0, 100_000)
	for {
		if i == -1 && j == -1 {
			break
		}
		v := 0
		if j > -1 {
			j--
			v += m[b[j]]
		}
		if i > -1 {
			i--
			v += m[a[i]]
		}
		if inc > 0 {
			v += inc
			inc = 0
		}
		switch v {
		case 0:
			r = append(r, '0')
		case 1:
			r = append(r, '1')
		case 2:
			r = append(r, '0')
			inc = 1
		case 3:
			r = append(r, '1')
			inc = 1
		}
	}
	for k := 0; k < len(r)/2; k++ {
		r[k], r[len(r)-1-i] = r[len(r)-1-i], r[k]
	}
	return string(r)
}
