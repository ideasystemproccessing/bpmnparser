package main

import "strings"

type BPMNGateWay struct{

	bpmn *BPMN
	Element *ExclusiveGateway
	elemId string
}
func (self * BPMNGateWay)GetBPMN() *BPMN{
	return self.bpmn
}

func (self * BPMNGateWay) find(){

	for _, el:=range self.bpmn.Process.ExclusiveGateway{
		if el.ID==self.elemId {
			self.Element=&el
			break
		}
	}
}
func (self * BPMNGateWay) LoadObjElement(id string,bpmn *BPMN) {
	self.bpmn = bpmn
	self.elemId= id
	if self.Element==nil {
		self.find()
	}

}

func (self * BPMNGateWay) GetType() string{
	if self.elemId!=""{
		return strings.Split(self.elemId,"_")[0]
	}else {
		return ""
	}

}

func ( self * BPMNGateWay)GetElement() interface{}{
	return self.Element
}


