package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"xmlParser/bpmn_parser"
)
func ParseToJson(els []*bpmn_parser.Element,bpmn *bpmn_parser.Bpmn,proc *ConditionFlowElement){
	var node string
	//println(len(els))

	for _,el:=range els {
		new_proc:=new(ConditionFlowElement)
		param:=new(BpmnRuleCondition)
		switch el.GetType() {
		case "Gateway":
			node=el.GetElement().(*bpmn_parser.ExclusiveGateway).ID
			if el.Element.(*bpmn_parser.ExclusiveGateway).RuleCondition!="" {
				err := jsoniter.Unmarshal([]byte(el.Element.(*bpmn_parser.ExclusiveGateway).RuleCondition), param)
				if err != nil {
					panic(err)
				}
			}else {
				param=nil
			}
		case "Activity":
			node= el.GetElement().(*bpmn_parser.Task).ID
			if el.Element.(*bpmn_parser.Task).RuleCondition!="" {

				err := jsoniter.Unmarshal([]byte(el.Element.(*bpmn_parser.Task).RuleCondition), param)
				if err != nil {
					panic(err)
				}
			}else {
				param=nil
			}
		case "Event":
			if endEvent,ok:=el.GetElement().(*bpmn_parser.EndEvent);ok {
				node = endEvent.ID
				if el.Element.(*bpmn_parser.EndEvent).RuleCondition != "" {
					err := jsoniter.Unmarshal([]byte(endEvent.RuleCondition), param)
					if err != nil {
						panic(err)
					}
				} else {
					param = nil
				}
			}else if interEvent,ok:=el.GetElement().(*bpmn_parser.IntermediateEvent);ok{
				node = interEvent.ID
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
		//println(el.PrevState.TestStatus)
		new_proc.ConditionParams=param
		new_proc.ConditionType=el.ElemId
		new_proc.ElementType=el.GetElemType()
		if el.PrevState.Name=="Y"{
			proc.TrueState = new_proc
		}else if el.PrevState.Name =="N" {
			proc.FalseState = new_proc

		}else {
			proc.Next=append(proc.Next,new_proc)
		}

		getFirstStep:=bpmn.ForwardElement(node)
		ParseToJson(getFirstStep,bpmn,new_proc)
	}
}
var Errors []string
var Paths []string
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
	ParseToJson(getFirstStep, bpmn,proc)


	v,err:=jsoniter.Marshal(proc)
	if err!=nil{
		panic(err)
	}

	fmt.Println(string(v))
	if len(bpmn.GetStartElement().Outgoing)==0{
		Errors=append(Errors,"StartPoint Connection is required")

	}
	if len(bpmn.GetStartElement().Outgoing)>1{
		Errors=append(Errors,"The starting point can only have one connection ")

	}
	DiagValidate(getFirstStep, bpmn,proc)
	for _,err:=range Errors{
		println(err)
	}

	var path  string
	path=bpmn.GetStartElement().ID+" , "
	SuccessEndValidation(getFirstStep,bpmn,&path)

	for _,s:=range Paths{
		if strings.Contains(s,"End") || strings.Contains(s,"Out_Of_Com") {
			fmt.Println(s)
		}
	}
}


func DiagValidate(els []*bpmn_parser.Element,bpmn *bpmn_parser.Bpmn,proc *ConditionFlowElement){
	var node string

	for _,el:=range els {
		if el.PrevState!=nil {
			if strings.Split(el.PrevState.SourceRef,"_")[0]=="Gateway"{
				if el.PrevState.Name!="Y" && el.PrevState.Name!="N"{
					Errors=append(Errors,"The names of the Arrows must be `N` or `Y` : "+el.PrevState.Name)
				}
			}
		}
		new_proc:=new(ConditionFlowElement)
		param:=new(BpmnRuleCondition)
		switch el.GetType() {
		case "Gateway":
			node=el.GetElement().(*bpmn_parser.ExclusiveGateway).ID
			if len(el.GetOutGoings())>2{
				Errors=append(Errors,"This element should not have more than two connections : "+node)
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
			node= el.GetElement().(*bpmn_parser.Task).ID
			//if len(el.GetOutGoings())==1{
			//	typeOfNext:=strings.Split(el.GetOutGoings()[0],"_")[0]
			//	if typeOfNext=="Activity" {
			//		Errors = append(Errors, "This element can only have Gateway or `Terminate Event` connections : "+node)
			//	}
			//}else
			if len(el.GetOutGoings())>1{
				Errors=append(Errors,"The minimum number of Task  is one Connection : "+node)

			}else if len(el.GetOutGoings())==0{
				Errors = append(Errors, "In task Design , A connection is required : "+node)
			}


		case "Event":
			if endEvent,ok:=el.GetElement().(*bpmn_parser.EndEvent);ok {

				node = endEvent.ID
			}else if interEvent,ok:=el.GetElement().(*bpmn_parser.IntermediateEvent);ok {
				node = interEvent.ID

			}



		}

		//println(el.PrevState.TestStatus)
		new_proc.ConditionParams=param
		new_proc.ConditionType=el.ElemId
		new_proc.ElementType=el.GetElemType()
		if el.PrevState.Name=="Y"{
			proc.TrueState = new_proc
		}else if el.PrevState.Name =="N" {
			proc.FalseState = new_proc

		}else {
			proc.Next=append(proc.Next,new_proc)
		}

		getFirstStep:=bpmn.ForwardElement(node)
		DiagValidate(getFirstStep,bpmn,new_proc)
	}
}

func SuccessEndValidation(els []*bpmn_parser.Element,bpmn *bpmn_parser.Bpmn,path * string){
	var node string
	//println(len(els))
	Paths=append(Paths,*path)

var newPath string
	for _,el:=range els {

		switch el.GetType() {
		case "Gateway":
			node=el.GetElement().(*bpmn_parser.ExclusiveGateway).ID
		case "Activity":
			node=el.GetElement().(*bpmn_parser.Task).ID
		case "Event":
			if endEvent,ok:=el.GetElement().(*bpmn_parser.EndEvent);ok {
				node = endEvent.ID
			}else if interEvent,ok:=el.GetElement().(*bpmn_parser.IntermediateEvent);ok {
				node = interEvent.ID
			}
		}


		newPath = *path + node + " , "

		if el.GetElemType()=="End" || el.GetElemType()=="Out_Of_Commitment" {
			 newPath+= el.GetElemType()
		}

		getFirstStep:=bpmn.ForwardElement(node)
		SuccessEndValidation(getFirstStep,bpmn,&newPath)
	}
}