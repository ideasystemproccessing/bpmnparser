package main

import "strings"

type ElementBase struct{
	bpmn *BPMN
	Element interface{}
	elemId string
}



func (self * ElementBase)GetBPMN() *BPMN{
	return self.bpmn
}


func (self * ElementBase) GetType() string{
	if self.elemId!=""{
		return strings.Split(self.elemId,"_")[0]
	}else {
		return ""
	}
}


func ( self * ElementBase)GetElement() interface{}{
	return self.Element
}