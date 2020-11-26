package main

func main() {
	// Open our xmlFile
err,bpmn:=NewBPMN("export.bpmn")
	if err!=nil {
		panic(err)
	}
	getFirstStep:=bpmn.GetStartElement()
	for _,el:=range getFirstStep{
		switch el.GetType() {
		case "Gateway":
			println(el.GetElement().(*ExclusiveGateway).ID)
		case "Activity":
			println(el.GetElement().(*Task).ID)

		}

	}
	//for _,v:=range bpmn.Process.Task{
	//	println(v.ID)
	//}
	//fmt.Println(bpmn.Process)
	//for _,v:=range

	//f:=getFlowId(bpmn.refrence.Process.StartEvent.Outgoing[0],bpmn.refrence)
	//next:= strings.Split(f.TargetRef,"_")[0]
	//if next == "Gateway" {
	//	flow := getGateWayId(f.TargetRef, bpmn.refrence)
	//	//Goto True and false Calc Operation
	//	//ElementCalc(flow,bpmn)
	//	println(flow.ID)
	//}else if next == "Activity" {
	//	flow :=getTaskId(f.TargetRef, bpmn.refrence)
	//	println(flow.ID)
	//	//GoToActivity Calc
	//}


		//for _,target:=range flow.Outgoing{
		//	switch strings.Split(target.TargetRef,"_")[0] {
		//	case "Gateway":
		//		gateway:=getGateWayId(flow.TargetRef,bpmn)
		//		println(gateway.RuleCondition)
		//	case "Task":
		//		task:=getTaskId(flow.TargetRef,bpmn)
		//		println(task.RuleCondition)
		//	case "Event":
		//
		//
		//	}
		//}

	//}


}


//func ElementCalc(){
//
//}
//func getFlowId(id string,diag *BPMN) *SequenceFlow{
//	for _, el:=range diag.Process.SequenceFlow{
//		if el.ID==id {
//			return &el
//		}
//	}
//	return nil
//}
//func getTaskId(id string,diag *BPMN) *Task{
//	for _, el:=range diag.Process.Task{
//		if el.ID==id {
//			return &el
//		}
//	}
//	return nil
//}
//func getGateWayId(id string,diag *BPMN) *ExclusiveGateway{
//	for _, el:=range diag.Process.ExclusiveGateway{
//		if el.ID==id {
//			return &el
//		}
//	}
//	return nil
//}
//package main
//
//import (
//	"encoding/xml"
//	"fmt"
//	"log"
//	"strings"
//)
//
//type Animal int
//
//const (
//	Unknown Animal = iota
//	Gopher
//	Zebra
//)
//
//func (a *Animal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	var s string
//	if err := d.DecodeElement(&s, &start); err != nil {
//		return err
//	}
//	switch strings.ToLower(s) {
//	default:
//		*a = Unknown
//	case "gopher":
//		*a = Gopher
//	case "zebra":
//		*a = Zebra
//	}
//
//	return nil
//}
//
//func (a Animal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	var s string
//	switch a {
//	default:
//		s = "unknown"
//	case Gopher:
//		s = "gopher"
//	case Zebra:
//		s = "zebra"
//	}
//	return e.EncodeElement(s, start)
//}
//
//func main() {
//	blob := `
//	<animals>
//		<animal>gopher</animal>
//		<animal>armadillo</animal>
//		<animal>zebra</animal>
//		<animal>unknown</animal>
//		<animal>gopher</animal>
//		<animal>bee</animal>
//		<animal>gopher</animal>
//		<animal>zebra</animal>
//	</animals>`
//	var zoo struct {
//		Animals []Animal `xml:"animal"`
//	}
//	if err := xml.Unmarshal([]byte(blob), &zoo); err != nil {
//		log.Fatal(err)
//	}
//
//	census := make(map[Animal]int)
//	for _, animal := range zoo.Animals {
//		census[animal] += 1
//	}
//
//	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
//		census[Gopher], census[Zebra], census[Unknown])
//
//}
