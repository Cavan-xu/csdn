package main

import (
	"fmt"
	"unsafe"

	"csdn-code/unsafe/entity"
)

func main() {
	operateVariable()
	operateStruct()

	data := StringToBytes("hello gopher")
	str := BytesToString(data)
	fmt.Println(str)
}

func operateVariable() {
	var a int32 = 8
	var f int64 = 20

	// int32 的指针
	ptr := &a

	// 先将 *int64 类型转化为 *Arbitrary 类型再转化为 *int32类型
	ptr = (*int32)(unsafe.Pointer(&f))
	*ptr = 10

	fmt.Println(a)
	fmt.Println(f)
}

func operateStruct() {
	user := new(entity.User)
	// user.name = "jack"
	fmt.Printf("%+v\n", user)

	// 突破第一个私有变量，因为是结构体的第一个字段，所以不需要额外的指针计算
	*(*string)(unsafe.Pointer(user)) = "张伟"
	fmt.Printf("%+v\n", user)

	// 突破第二个私有变量，因为是第二个成员字段，需要偏移一个字符串占用的长度即 16 个字节
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(user)) + uintptr(16))) = 1
	fmt.Printf("%+v\n", user)
}

func StringToBytes(str string) []byte {
	var b []byte
	// 切片的底层数组、len字段，指向字符串的底层数组，len字段
	*(*string)(unsafe.Pointer(&b)) = str

	// 切片的 cap 字段赋值为 len(str)的长度，切片的指针、len 字段各占八个字节，直接偏移16个字节
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*uintptr(8))) = len(str)

	return b
}

func BytesToString(data []byte) string {
	// 直接转换
	return *(*string)(unsafe.Pointer(&data))
}