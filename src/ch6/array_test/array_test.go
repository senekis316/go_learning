package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	arr4 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1[1])
	t.Log(arr1, arr3)
	t.Log(arr4)
}

func TestArrayTravel(t *testing.T) {
	arr4 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr4); i++ {
		t.Log(arr4[i])
	}
	for idx, e := range arr4 {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[0:3]
	arr4_sec := arr3[:]
	t.Log(arr3_sec)
	t.Log(arr4_sec)
}

func TestSliceInit(t *testing.T) {

	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))

}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
}

func TestSliceComparing(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	/*if a == b {
		t.Log("equal")
	}*/
}
