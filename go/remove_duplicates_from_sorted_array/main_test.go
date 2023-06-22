package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums  []int
		want  int
		want2 []int
	}{
		{
			nums:  []int{1, 1, 1, 2, 2, 2, 3, 3},
			want:  3,
			want2: []int{1, 2, 3},
		},
		{
			nums:  []int{1, 2, 3, 3},
			want:  3,
			want2: []int{1, 2, 3},
		},
		{
			nums:  []int{1, 2, 2, 2, 3, 3},
			want:  3,
			want2: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := removeDuplicates(tt.nums)

			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			for i := 0; i < got; i++ {
				if tt.nums[i] != tt.want2[i] {
					t.Errorf("%v: nums (%v) != want2 (%v)", i, tt.nums[i], tt.want2[i])
				}
			}
		})
	}
}
