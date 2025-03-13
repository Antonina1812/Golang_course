package main

import (
	"fmt"
	b "golang/basics"
	p "golang/first_practice"
	"io"
	"os"
)

func FirstPracticePrint(in io.Reader, out io.Writer) error {
	err := p.Scanner(in, out)
	if err != nil {
		return fmt.Errorf("error")
	}
	return nil
}

func BasicsPrint() {
	//b.VariablesPrint()
	//b.StringsPrint()
	//b.ConstantsPrint()
	//b.PointersPrint()
	//b.ArraysPrint()
	//b.SlicesPrint()
	//b.MapsPrint()
	//b.ConditionsPrint()
	//b.LoopsPrint()
	//b.FunctionsPrint()
	//b.DeferPrint()
	//b.PanicPrint()
	//b.SrtructuresPrint()
	//b.MethodsPrint()
	b.InterfacesPrint()
}

func main() {
	//BasicsPrint()
	FirstPracticePrint(os.Stdin, os.Stdout)
}
