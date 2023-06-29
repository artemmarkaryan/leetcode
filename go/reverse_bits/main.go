package main

func reverseBits(num uint32) (r uint32) {
	for i := 0; i < 31; i++ {
		r |= num & 1
		r <<= 1
		num >>= 1
	}
	r |= num & 1

	return r
}
