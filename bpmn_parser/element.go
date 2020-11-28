package bpmn_parser

import (
	"strings"
)

const (
	GATEWAY     string = "Bool_Condition"
	FLOW        string = "Arrow"
	TASK        string = "Operation"
	END_EVENT   string = "End"
	START_EVENT string = "Start"
	INTER_EVENT string = "Intermediate"
)

type Element struct {
	bpmn      *BPMN
	Element   interface{}
	ElemId    string
	elemType  string
	outGoings []string
	inComes   []string
	PrevState *SequenceFlow
}

func (self *Element) find() {
	if self.GetType() == "Gateway" {
		for _, el := range self.bpmn.Process.ExclusiveGateway {
			if el.ID == self.ElemId {
				self.Element = &el
				self.elemType = GATEWAY

				self.outGoings = el.Outgoing
				self.inComes = el.Incoming
				break
			}
		}
	} else if self.GetType() == "Activity" {
		for _, el := range self.bpmn.Process.Task {
			if el.ID == self.ElemId {
				self.elemType = TASK
				self.Element = &el
				self.outGoings = el.Outgoing
				self.inComes = el.Incoming

				break
			}
		}
	} else if self.GetType() == "Flow" {
		for _, el := range self.bpmn.Process.SequenceFlow {
			if el.ID == self.ElemId {
				self.elemType = FLOW
				self.Element = &el

				break
			}
		}
	} else if self.GetType() == "Event" {
		for _, el := range self.bpmn.Process.EndEvent {
			if el.ID == self.ElemId {
				self.elemType = END_EVENT
				self.inComes = el.Incoming

				self.Element = &el
				break
			}
		}
		for _, el := range self.bpmn.Process.IntermediateThrowEvent {
			if el.ID == self.ElemId {
				self.elemType = INTER_EVENT
				self.inComes = el.Incoming

				self.Element = &el
				break
			}
		}


	} else if self.bpmn.Process.StartEvent.ID == self.ElemId {
		self.elemType = START_EVENT
		self.Element = &self.bpmn.Process.StartEvent
		self.outGoings = self.bpmn.Process.StartEvent.Outgoing

	}
}
func (self *Element) GetBPMN() *BPMN {
	return self.bpmn
}
func (self *Element) LoadObjElement(id string, bpmn *BPMN) {
	self.bpmn = bpmn
	self.ElemId = id
	if self.Element == nil {
		self.find()
	}

}

func (self *Element) GetType() string {
	if self.ElemId != "" {
		return strings.Split(self.ElemId, "_")[0]
	} else {
		return ""
	}
}

func (self *Element) GetElemType() string {
	return self.elemType
}
func (self *Element) GetElement() interface{} {
	return self.Element
}

func (self *Element) GetOutGoings() []string {
	return self.outGoings
}
func (self *Element) GetIncomes() []string {
	return self.inComes
}
