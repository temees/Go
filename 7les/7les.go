package main

import (
	"fmt"
)

type Struct struct {
	s string
}

func (str Struct) pointerMethod() {
	fmt.Println("pointerMethod", &str.s)
}
func (str Struct) valueMethod() {
	fmt.Println("valueMethod", str.s)

}

func main() {

	var valueStruct = Struct{s: "valueStruct"}
	var pointerStruct = &Struct{s: "pointerStruct"}

	var emptyInterfaceInstancePointerStruct interface{} = pointerStruct
	pMI := emptyInterfaceInstancePointerStruct.(pointerMethodInterface)

	var emptyInterfaceInstanceValueStruct interface{} = valueStruct
	vMI := emptyInterfaceInstanceValueStruct.(valueMethodInterface)

	valueStruct.s = "VALUE"

	valueStruct.pointerMethod()
	valueStruct.valueMethod()

	pointerStruct.pointerMethod()
	pointerStruct.valueMethod()

	pMI.pointerMethod()
	vMI.valueMethod()

}

type pointerMethodInterface interface {
	pointerMethod()
}
type valueMethodInterface interface {
	valueMethod()
}
