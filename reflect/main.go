package main

import (
	"fmt"
	"reflect"

	"csdn-code/reflect/entity"
)

func main() {
	reflectTry()
	useReflectKind()
	userReflectStruct()
	reflectUpdateField()
	useReflectGetTag()
	useReflectMethodCall()
}

// use reflect get variable type and value
func reflectTry() {
	num := 100
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.ValueOf(num))

	s := "hello gopher"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
}

func useReflectKind() {
	reflectKind(1)
	reflectKind(0.01)
	reflectKind([]int{1, 2, 3})
	reflectKind(map[string]struct{}{})
}

func reflectKind(v interface{}) {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Int:
		fmt.Printf("int value is: %d\n", value.Int())
	case reflect.Float64:
		fmt.Printf("float64 value is: %f\n", value.Float())
	case reflect.Slice:
		fmt.Printf("slice value is: %v\n", value.Slice(0, 3))
	default:
		fmt.Printf("defaule type is: %v\n", value)
	}
}

func userReflectStruct() {
	user := entity.User{
		Id:   101,
		Name: "user",
		Age:  123,
	}

	reflectStruct(user)
	reflectStruct(&user)
	reflectStruct(1)
}

func reflectStruct(v interface{}) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	fmt.Printf("v type is %v\n", typ)

	// 判断该变量是否是结构体类型
	if typ.Kind() == reflect.Struct {
		// 获取成员变量个数
		fieldCount := typ.NumField()
		for i := 0; i < fieldCount; i++ {
			field := typ.Field(i)
			// CanInterface 未导出的成员变量返回 false，获取不到具体的值，否则会导致 panic
			if value.Field(i).CanInterface() {
				fmt.Printf("field name is: %v, field type is %v, field value is: %v\n", field.Name, field.Type, value.Field(i).Interface())
			} else {
				fmt.Printf("field name is: %v, field type is %v\n", field.Name, field.Type)
			}
		}
	}

	// 注意，对于值接收者和指针接受者来说，方法有不一样的可见性
	// 传入的是结构体的拷贝的话，只能获取到值接收者的方法，传入指针可以获取到所有包外可见的方法
	fmt.Println(typ.NumMethod())

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("v metnod name is: %s, type is %s\n", method.Name, method.Type)
	}
}

func reflectUpdateField() {
	// 修改变量的值
	str := "hello gopher"

	tStr := reflect.TypeOf(&str)
	fmt.Println(tStr.Elem()) // tStr.Elem()  只有 tStr.Kind() 为 Array、Chan、Map、Ptr、Slice时才能调用，否则 panic

	vStr := reflect.ValueOf(&str)
	if vStr.Elem().CanSet() { // vStr.Elem()  只有 vStr.Kind() 为 Ptr、Interface 时才可以调用，否则 panic
		vStr.Elem().SetString("hello world")
	}

	fmt.Printf("now str is: %s\n", str)

	// 修改结构体成员变量的值
	user := entity.User{
		Id:   1,
		Name: "gopher",
		Age:  10,
	}

	vName := reflect.ValueOf(&user.Name)
	vName.Elem().SetString("tom")

	vAge := reflect.ValueOf(&user.Age)
	vAge.Elem().SetInt(100)

	fmt.Printf("now user is: %+v\n", user)
}

func useReflectGetTag() {
	user := entity.User{
		Id:   1,
		Name: "gopher",
		Age:  10,
	}

	reflectGetTag(&user)
}

func reflectGetTag(v interface{}) {
	table := make(map[string]string)

	tUser := reflect.TypeOf(v)
	fmt.Println(tUser.Kind())

	tUser = reflect.TypeOf(v).Elem()
	fmt.Println(tUser.Kind())
	for i := 0; i < tUser.NumField(); i++ {
		table[tUser.Field(i).Name] = tUser.Field(i).Tag.Get("json") // User 结构体中的 json tag，也可以自定义
	}

	fmt.Printf("tag table is: %+v", table)
}

func useReflectMethodCall() {
	methodName := "GetId"
	user := entity.User{
		Name: "gopher",
	}

	reflectMethodCall(&user, methodName)
}

func reflectMethodCall(v interface{}, methodName string) {
	tUser := reflect.TypeOf(v)
	_, ok := tUser.MethodByName(methodName)
	if ok {
		vUser := reflect.ValueOf(v)
		method := vUser.MethodByName(methodName)
		resp := method.Call([]reflect.Value{reflect.ValueOf(int64(10))})
		res := resp[0]
		fmt.Println(res.Int())
	}
}
