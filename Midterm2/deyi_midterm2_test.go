package main

import "testing"

func TestIsBSTOrdered(t *testing.T) {
	var N0, N1, N2, N3, N4, N5, N6, N7, N8, N9 TreeNode
	N0.value = 0
	N1.value = 1
	N2.value = 2
	N3.value = 3
	N4.value = 4
	N5.value = 5
	N6.value = 6
	N7.value = 7
	N8.value = 8
	N9.value = 9
	N5.left = &N3
	N5.right = &N7
	N3.left = &N1
	N3.right = &N4
	N1.left = &N0
	N1.right = &N2
	N7.left = &N6
	N7.right = &N8
	N8.left = nil
	N8.right = &N9
	cases := []struct {
		in   *TreeNode
		want bool
	}{
		{&N5, true},
		{&N3, true},
		{&N7, true},
		{&N1, true},
	}

	for _, test := range cases {
		output := IsBSTOrdered(test.in)
		if output != test.want {
			t.Error("Test failed, {} was inputted, {} was outputted, {} was expected", test.in, output, test.want)
		}
	}
}
