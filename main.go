package main

import (
	"fmt"
	alg "golang/algorithms_and_data_structures"
	a "golang/async"
	b "golang/basics"
	p "golang/first_practice"
	j "golang/json"
	l "golang/leetcode"
	"io"
)

func AtgorithmsPrint() {
	alg.StackPrint()
	alg.QueuePrint()
}

func LeetcodePrint() {
	l.LeetcodeArrayPrint()
}

func JsonPrint() {
	//j.JSONPrint()
	//j.TagsPrint()
	//j.DynamicPrint()
	j.ReflectPrint()
}

func AsyncPrint() {
	//a.GoroutinesPrint()
	//a.ChannelsPrint()
	//a.SelectPrint()
	//a.TimePrint()
	//a.ContextPrint()
	//a.AsynqWorkPrint()
	//a.WorkerPoolPrint()
	a.WaitGroupPrint()
}

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
	//FirstPracticePrint(os.Stdin, os.Stdout)
	//AsyncPrint()
	JsonPrint()

	//LeetcodePrint()
	//AtgorithmsPrint()
}
