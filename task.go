package main

import "strings"

type BPMNTask struct{
	bpmn *BPMN
	Element *Task
	elemId string
}
func (self * BPMNTask)GetBPMN() *BPMN{
	return self.bpmn
}

func (self * BPMNTask) find(){
	for _, el:=range self.bpmn.Process.Task{
		if el.ID==self.elemId {
			self.Element=&el
			break

		}
	}
}
func (self * BPMNTask) LoadObjElement(id string,bpmn *BPMN) {
	self.bpmn = bpmn
	self.elemId= id
	if self.Element==nil {
		self.find()
	}

}

func (self * BPMNTask) GetType() string{
	if self.elemId!=""{
		return strings.Split(self.elemId,"_")[0]
	}else {
		return ""
	}
}

func ( self * BPMNTask)GetElement() interface{}{
	return self.Element
}
