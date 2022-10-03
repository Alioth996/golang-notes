package arraych_test

import "testing"

func TestArrayBegin(t *testing.T) {
	// var arr [3]int 声明长度为3,元素类型为int的数组

	// arr := [3]int{}  数组声明简写

	arr := [...]int{1, 2, 30} // 声明数组并初始化数组元素

	// 遍历数组 range
	for _, v := range arr {
		t.Log(v)
	}
	t.Log(len(arr))

	// 多维数组
	// var arr1 [3][4]int

	arr1 := [...][2]int{{1, 2}, {2, 3}, {2}}

	t.Log(arr1)
	t.Logf("arr1 len %d", len(arr1))
}

func TestArraySlice(t *testing.T) {
	arr := [...]int{1, 2, 33, 22, 56}

	arr1 := arr[:2] // (] 左闭右开区间,防止数组越界截取, arr[1:len(arr)]

	t.Log(arr1)
}
