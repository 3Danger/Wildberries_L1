package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме
	способна определить тип переменной:
		int, string, bool, channel из переменной типа interface{}.
*/

type TestStruct struct {
	someString string
}

func Tester(detector func(interface{})) {
	detector(int8(1))
	detector(int64(1))
	detector(struct{ name string }{"AAA"})
	detector(struct{}{})
	detector(TestStruct{"someValue"})
	detector('r')
	detector(new(int))
	detector(new(chan rune))
	detector(func() {})
	detector(detector)
	fmt.Println()
}

func main() {
	fmt.Println("\treflect.TypeOf(in)")
	Tester(DetectorTypeOne)

	fmt.Println("\n\treflect.ValueOf(in).Kind()")
	Tester(DetectorTypeTwo)

	/*
		!!Wrong solution
		fmt.Println("\n\tfmt.Sprintf()")
		Tester(DetectorTypeThree)
	*/
}

func DetectorTypeOne(in interface{}) {
	fmt.Println(reflect.TypeOf(in))
}
func DetectorTypeTwo(in interface{}) {
	fmt.Println(reflect.ValueOf(in).Kind())
}
func DetectorTypeThree(in interface{}) {
	fmt.Println(fmt.Sprintf("%T", in))
}
