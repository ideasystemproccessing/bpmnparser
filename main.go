package main

import (
	"fmt"
	"git.ispfarm.com/mehdi/bpmnparser/bpmn_parser"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

var Errors []string
var Paths []string

func ParseToJson(els []*bpmn_parser.Element, bpmn *bpmn_parser.Bpmn, proc *ConditionFlowElement) {
	//println(len(els))

	for _, el := range els {
		new_proc := new(ConditionFlowElement)
		param := new(BpmnRuleCondition)
		switch el.GetType() {
		case "Gateway":
			if el.Element.(*bpmn_parser.ExclusiveGateway).RuleCondition != "" {
				err := jsoniter.Unmarshal([]byte(el.Element.(*bpmn_parser.ExclusiveGateway).RuleCondition), param)
				if err != nil {
					panic(err)
				}
			} else {
				param = nil
			}
		case "Activity":
			if el.Element.(*bpmn_parser.Task).RuleCondition != "" {

				err := jsoniter.Unmarshal([]byte(el.Element.(*bpmn_parser.Task).RuleCondition), param)
				if err != nil {
					panic(err)
				}
			} else {
				param = nil
			}
		case "Event":
			if endEvent, ok := el.GetElement().(*bpmn_parser.EndEvent); ok {
				if el.Element.(*bpmn_parser.EndEvent).RuleCondition != "" {
					err := jsoniter.Unmarshal([]byte(endEvent.RuleCondition), param)
					if err != nil {
						panic(err)
					}
				} else {
					param = nil
				}
			} else if interEvent, ok := el.GetElement().(*bpmn_parser.IntermediateEvent); ok {
				if el.Element.(*bpmn_parser.IntermediateEvent).RuleCondition != "" {
					err := jsoniter.Unmarshal([]byte(interEvent.RuleCondition), param)
					if err != nil {
						panic(err)
					}
				} else {
					param = nil
				}
			}
		}
		//println(node)
		new_proc.ConditionParams = param
		new_proc.ConditionType = el.ElemId
		new_proc.ElementType = el.GetElemType()
		if el.PrevState.Name == "Y" {
			proc.TrueState = new_proc
		} else if el.PrevState.Name == "N" {
			proc.FalseState = new_proc

		} else {
			proc.Next = append(proc.Next, new_proc)
		}

		getFirstStep := bpmn.ForwardElement(el.ElemId)
		ParseToJson(getFirstStep, bpmn, new_proc)
	}
}

func main() {
	// BPMN Load File
	err, bpmn := bpmn_parser.NewBPMN("export.bpmn")
	if err != nil {
		panic(err)
	}

	//====================================================
	// Parse To Json Struct
	getFirstStep := bpmn.ForwardElement(bpmn.GetStartElement().ID)
	proc := new(ConditionFlowElement)
	proc.ConditionType = bpmn.GetStartElement().ID
	proc.ElementType = bpmn_parser.START_EVENT
	ParseToJson(getFirstStep, bpmn, proc)

	v, err := jsoniter.Marshal(proc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(v))

	//====================================================
	// Element Validation
	if len(bpmn.GetStartElement().Outgoing) == 0 {
		Errors = append(Errors, "StartPoint Connection is required")

	}
	if len(bpmn.GetStartElement().Outgoing) > 1 {
		Errors = append(Errors, "The starting point can only have one connection ")

	}
	DiagValidate(getFirstStep, bpmn, proc)
	for _, err := range Errors {
		println(err)
	}

	//====================================================
	//Path Validations
	var path string
	path = bpmn.GetStartElement().ID + " , "
	SuccessEndValidation(getFirstStep, bpmn, &path)

	for _, s := range Paths {
		if strings.Contains(s, bpmn_parser.END_EVENT) || strings.Contains(s, bpmn_parser.TERMINATE_END_EVENT) {
			fmt.Println(s)
		}
	}
}

func DiagValidate(els []*bpmn_parser.Element, bpmn *bpmn_parser.Bpmn, proc *ConditionFlowElement) {

	for _, el := range els {
		if el.PrevState != nil {
			if strings.Split(el.PrevState.SourceRef, "_")[0] == "Gateway" {
				if el.PrevState.Name != "Y" && el.PrevState.Name != "N" {
					Errors = append(Errors, "The names of the Arrows must be `N` or `Y` : "+el.PrevState.Name)
				}
			}
		}
		new_proc := new(ConditionFlowElement)
		param := new(BpmnRuleCondition)
		switch el.GetType() {
		case "Gateway":
			if len(el.GetOutGoings()) > 2 {
				Errors = append(Errors, "This element should not have more than two connections : "+el.ElemId)
			}
			//for _,selection_target:=range el.GetOutGoings(){
			//
			//	x:=0
			//	for _,target:=range el.GetOutGoings(){
			//		if selection_target==target && strings.Split(target,"_")[0]!="Event"{
			//			x++
			//		}
			//	}
			//	if x>1{
			//		Errors=append(Errors,"Two connections to one destination are not possible : "+node)
			//
			//	}
			//}
		case "Activity":
			//if len(el.GetOutGoings())==1{
			//	typeOfNext:=strings.Split(el.GetOutGoings()[0],"_")[0]
			//	if typeOfNext=="Activity" {
			//		Errors = append(Errors, "This element can only have Gateway or `Terminate Event` connections : "+node)
			//	}
			//}else
			if len(el.GetOutGoings()) > 1 {
				Errors = append(Errors, "The minimum number of Task  is one Connection : "+el.ElemId)

			} else if len(el.GetOutGoings()) == 0 {
				Errors = append(Errors, "In task Design , A connection is required : "+el.ElemId)
			}

		case "Event":

		}

		new_proc.ConditionParams = param
		new_proc.ConditionType = el.ElemId
		new_proc.ElementType = el.GetElemType()
		if el.PrevState.Name == "Y" {
			proc.TrueState = new_proc
		} else if el.PrevState.Name == "N" {
			proc.FalseState = new_proc

		} else {
			proc.Next = append(proc.Next, new_proc)
		}

		getFirstStep := bpmn.ForwardElement(el.ElemId)
		DiagValidate(getFirstStep, bpmn, new_proc)
	}
}

func SuccessEndValidation(els []*bpmn_parser.Element, bpmn *bpmn_parser.Bpmn, path *string) {

	Paths = append(Paths, *path)

	var newPath string
	for _, el := range els {
		if el.GetType() == "Activity" {
			newPath = *path + el.ElemId + " Act " + " , "
		} else {
			newPath = *path + el.ElemId + " , "

		}

		if el.GetElemType() == bpmn_parser.END_EVENT ||  el.GetElemType() == bpmn_parser.TERMINATE_END_EVENT{
			newPath += el.GetElemType()
		}
		getFirstStep := bpmn.ForwardElement(el.ElemId)
		SuccessEndValidation(getFirstStep, bpmn, &newPath)
	}
}
