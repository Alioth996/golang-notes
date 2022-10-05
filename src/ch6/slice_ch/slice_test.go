package slicech_test

import (
	"testing"
)

func TestSliceBeginByVar(t *testing.T) {
	// slice简单理解为可变长度的数组

	// var关键字声明,不加具体的数组长度则声明一个slice
	var numList []int
	for i := 1; i < 5; i++ {
		numList = append(numList, i)

	}

	t.Log(numList)

}

func TestSliceSimpleDefine(t *testing.T) {
	personList := []string{"张三", "李四", "王五", "王八蛋"}

	for index, v := range personList {
		t.Logf("personlist[%d]:%s", index, v)
	}
}

func TestSlicebyMake(t *testing.T) {

	// make(type,len,/cap/)
	priceList := make([]int, 5) // 创建一个长度为10的切片, 填充其元素值为int类型默认值0

	//创建一个可访问长度为5,容量为10,填充元素为0 的切片,后续长度超出,则容量翻倍.并且容量总是翻倍向后拓展
	pointList := make([]float64, 5, 5)
	t.Log(priceList)
	pointList = append(pointList, 1.5) // append返回一个新切片,把之前的切片拷贝一份,并开辟一份cap大小的存储空间,可能会引起性能问题
	t.Log(pointList, cap(pointList))   // 此时cap=20

}

func TestSliceShareMem(t *testing.T) {
	// 切片共享数据
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// 研究切片的len与cap变化
	spring := months[:6]
	summer := months[5:]

	springLen := len(spring) // 6
	springCap := cap(spring) // 12

	summer[0] = "hello"
	summerLen := len(summer) // 7
	summerCap := cap(summer) // 7

	t.Log(spring, springLen, springCap)
	t.Log(summer, summerLen, summerCap)
	t.Log(months)

}

func TestSliceComparing(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}

	// slice之间不能进行比骄,只能nil比较
	// if s1 == s2 {
	// 	t.Log("same")
	// }
	s1, s2 = s2, s1
}
