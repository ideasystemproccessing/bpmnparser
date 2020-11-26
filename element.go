package main

import "strings"

type BPMNFlow struct{

	bpmn *BPMN
	Element interface{}
	elemId string
}


func (self * BPMNFlow) find(){
	if self.GetType()=="Gateway" {
		for _, el := range self.bpmn.Process.ExclusiveGateway {
			if el.ID == self.elemId {
				self.Element = &el
				break
			}
		}
	}else if self.GetType()=="Activity" {
		for _, el := range self.bpmn.Process.Task {
			if el.ID == self.elemId {
				self.Element = &el
				break
			}
		}
	} else 	if self.GetType()=="Flow" {
		for _, el := range self.bpmn.Process.SequenceFlow {
			if el.ID == self.elemId {
				self.Element = &el
				break
			}
		}
	}
}
func (self * BPMNFlow)GetBPMN() *BPMN{
	return self.bpmn
}
func (self * BPMNFlow) LoadObjElement(id string,bpmn *BPMN) {
	self.bpmn = bpmn
	self.elemId= id
	if self.Element==nil {
		self.find()
	}


}

func (self * BPMNFlow) GetType() string{
	if self.elemId!=""{
		return strings.Split(self.elemId,"_")[0]
	}else {
		return ""
	}
}


func ( self * BPMNFlow)GetElement() interface{}{
	return self.Element
}