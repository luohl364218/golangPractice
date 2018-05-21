package main

import "fmt"

func main()  {
	//测试切片
	//testSlice()
	//测试自定义“slice
	//testSlice2()
	//测试切片的内存中的地址关系
	//testSlice3()
	//测试切片的初始化与append
	//testInitByArray()
	//testInitByMake()
	//测试拷贝
	//testCopy()
	//测试string作为切片
	//testString()
	//测试修改string
	testModifyString()
}

//slice的内存布局  代码模拟
type slice struct {
	ptr *[5]int
	len int
	cap int
}

//slice方法的代码模拟
func make1(s slice, cap int) slice {
	s.ptr = new([5]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice)  {
	s.ptr[1] = 1000
}

func testSlice2()  {
	var s1 slice
	s1 = make1(s1, 10)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func testSlice()  {

	var slice []int
	var arr [5]int = [5]int{1,2,3,4,5}

	// [arr[2], arr[3])
	slice = arr[2 : 5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	slice = arr[1 : ]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = arr[ : 5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println("before modify:", slice)
	modifySlice(slice)
	fmt.Println("after modify:", slice)

	//容量同上， 长度变1
	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func modifySlice(s []int)  {
	//切片是引用类型，可以直接修改
	s[1] = 1000
}

//切片的地址
func testSlice3()  {
	var a = [10]int{1,2,3,4,5}
	//切片其实就是指针，指向第一个元素
	b := a[1:5]
	//%p	十六进制表示，前缀 0x
	fmt.Printf("%p\n", b)
	fmt.Println(&a[1])
	//结果：地址一致
}

//切片的初始化方法一：数组
func testInitByArray()  {
	var arr [5]int = [...]int{1,2,3,4,5}
	var arr2 = []int{6,7,8,9,10}
	s := arr[1 : 5]

	//s:0xc042076038, arr[1]:0xc042076038
	fmt.Printf("s:%p, arr[1]:%p\n", s, &arr[1])

	//扩容测试  4   4
	fmt.Println("before append: len of s:", len(s), " cap of s:", cap(s))
	s = append(s, 10)
	//扩容测试  5   8
	fmt.Println("after append: len of s:", len(s), " cap of s:", cap(s))
	s = append(s, 10)
	//扩容测试  6   8
	fmt.Println("after append: len of s:", len(s), " cap of s:", cap(s))
	s = append(s, 10)
	//扩容测试  7   8
	fmt.Println("after append: len of s:", len(s), " cap of s:", cap(s))
	s = append(s, 10)
	//扩容测试  8   8
	fmt.Println("after append: len of s:", len(s), " cap of s:", cap(s))
	s = append(s, 10)
	//扩容测试  9   16
	fmt.Println("after append: len of s:", len(s), " cap of s:", cap(s))
	//扩容测试结论——翻倍原则扩容

	//切片追加切片 使用...
	s = append(s, arr2...)

	fmt.Println(s)
	fmt.Println(arr)
	//s:0xc042088000, arr[1]:0xc042076038   地址不一样了，说明切片在追加元素后会新开内存
	fmt.Printf("s:%p, arr[1]:%p\n", s, &arr[1])
}

//切片的初始化方法一：make
func testInitByMake()  {
	var slice1 []int = make([]int, 5)

	slice2 := make([]int, 5)

	slice3 := make([]int, 5, 20)

	slice1 = append(slice1, 10)
	fmt.Println("after append: len of slice1:", len(slice1), " cap of slice1:", cap(slice1))
	slice2 = append(slice2, 20)
	fmt.Println("after append: len of slice2:", len(slice2), " cap of slice2:", cap(slice2))
	slice3 = append(slice3, 30)
	fmt.Println("after append: len of slice3:", len(slice3), " cap of slice3:", cap(slice3))
}

//切片拷贝
func testCopy() {

	s1 := []int{1,2,3,4,5}
	s2 := make([]int, 10)
	fmt.Println(s1)
	fmt.Println(s2)
	//copy内置函数
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := []int{1,2,3}
	fmt.Println(s3)
	s3 = append(s3, s2...)
	fmt.Println(s3)
	s3 = append(s3, 4,5,6)
	fmt.Println(s3)
}

//string底层也是一个byte数组，所以也可以作为切片
func testString() {

	s := "hello world"
	s1 := s[:5]
	s2 := s[6:]

	fmt.Println(s1)
	fmt.Println(s2)
}

//string是不可变的，那怎么修改呢？
func testModifyString()  {

	s := "你好吗"
	slice := []rune(s)

	fmt.Printf("before s:%p, slice:%p\n", &s, slice)
	slice[0] = '他'
	fmt.Printf("after s:%p, slice:%p\n", &s, slice)
	fmt.Println(s)
	s = string(slice)
	fmt.Println(s)
	fmt.Printf("finally s:%p, slice:%p\n", &s, slice)
}
