package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"xmlParser/bpmn_parser"
)
func CheckProperty(els []*bpmn_parser.Element,bpmn *bpmn_parser.Bpmn,proc *ConditionFlowElement){
	var node string
	println(len(els))

	for _,el:=range els {
		new_proc:=new(ConditionFlowElement)
		switch el.GetType() {
		case "Gateway":
			node=el.GetElement().(*bpmn_parser.ExclusiveGateway).ID
			new_proc.ConditionParams=el.Element.(*bpmn_parser.ExclusiveGateway).RuleCondition
		case "Activity":
			node= el.GetElement().(*bpmn_parser.Task).ID
			new_proc.ConditionParams=el.Element.(*bpmn_parser.Task).RuleCondition

		case "Event":
			node = el.GetElement().(*bpmn_parser.EndEvent).ID
			new_proc.ConditionParams=el.Element.(*bpmn_parser.EndEvent).RuleCondition

		}
		println(node)
		println(el.PrevState.TestStatus)

		new_proc.ConditionType=el.ElemId
		new_proc.ElementType=el.GetElemType()
		if el.PrevState.TestStatus=="true"{
			proc.TrueState = new_proc
		}else if el.PrevState.TestStatus =="false" {
			proc.FalseState = new_proc

		}else {
			proc.Next=append(proc.Next,new_proc)
		}

		getFirstStep:=bpmn.ForwardElement(node)
		CheckProperty(getFirstStep,bpmn,new_proc)
	}
}
func main() {
	// Open our xmlFile
	err, bpmn := bpmn_parser.NewBPMN("export.bpmn")
	if err != nil {
		panic(err)
	}
	//getFirstStep := bpmn.Start()
	getFirstStep:=bpmn.ForwardElement(bpmn.GetStartElement().ID)
	proc:=new(ConditionFlowElement)
	proc.ConditionType=bpmn.GetStartElement().ID
	proc.ElementType= bpmn_parser.START_EVENT
	CheckProperty(getFirstStep, bpmn,proc)


	v,err:=jsoniter.Marshal(proc)
	if err!=nil{
		panic(err)
	}

	fmt.Println(string(v))

}