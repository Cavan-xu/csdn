package main

import (
	"fmt"
	"unsafe"

	"csdn-code/unsafe/entity"
)

func main() {
	operateVariable()
	operateStruct()
}

func operateVariable() {
	var a int32 = 8
	var f int64 = 20

	ptr := &a
	fmt.Println(ptr)

	// 先将 *int64 类型转化为 *Arbitrary 类型再转化为 *int32类型
	ptr = (*int32)(unsafe.Pointer(&f))
	fmt.Println(ptr)

	*ptr = 10

	fmt.Println(a)
	fmt.Println(f)
}

func operateStruct() {
	user := new(entity.User)
	// user.name = "jack"
	fmt.Printf("%+v\n", user)

	// 突破第一个私有变量，因为是结构体的第一个字段，所以不需要额外的指针计算
	p := (*string)(unsafe.Pointer(user))
	*p = "张伟"
	fmt.Printf("%+v\n", user)

	// 突破第二个私有变量，因为是第二个成员字段，需要偏移一个字符串占用的长度即 16 个字节
	ptrId := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(user)) + uintptr(16)))
	*ptrId = 1
	fmt.Printf("%+v", user)
}
