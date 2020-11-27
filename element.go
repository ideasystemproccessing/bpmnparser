package main

import (
	"strings"
)

type Element struct{

	bpmn *BPMN
	Element interface{}
	elemId string
	outGoings []string
	PrevState *SequenceFlow
}


func (self * Element) find(){
	if self.GetType()=="Gateway" {
		for _, el := range self.bpmn.Process.ExclusiveGateway {
			if el.ID == self.elemId {
				self.Element = &el
				self.outGoings=el.Outgoing
				break
			}
		}
	}else if self.GetType()=="Activity" {
		for _, el := range self.bpmn.Process.Task {
			if el.ID == self.elemId {
				self.Element = &el
				self.outGoings=el.Outgoing

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
	}else 	if self.GetType()=="Event" {
		for _, el := range self.bpmn.Process.EndEvent {
			if el.ID == self.elemId {

				self.Element = &el
				break
			}
		}
			if self.bpmn.Process.StartEvent.ID == self.elemId {
				self.Element = &self.bpmn.Process.StartEvent
				self.outGoings= self.bpmn.Process.StartEvent.Outgoing
			}


	}
}
func (self * Element)GetBPMN() *BPMN{
	return self.bpmn
}
func (self * Element) LoadObjElement(id string,bpmn *BPMN) {
	self.bpmn = bpmn
	self.elemId= id
	if self.Element==nil {
		self.find()
	}


}

func (self * Element) GetType() string{
	if self.elemId!=""{
		return strings.Split(self.elemId,"_")[0]
	}else {
		return ""
	}
}


func ( self * Element)GetElement() interface{}{
	return self.Element
}

func ( self * Element)GetOutGoings() []string{
	return self.outGoings
}
