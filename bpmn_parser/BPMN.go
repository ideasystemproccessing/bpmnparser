package bpmn_parser

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

type ExclusiveGateway struct {
	Text              string   `xml:",chardata"`
	ID                string   `xml:"id,attr"`
	RuleCondition     string   `xml:"ruleCondition,attr"`
	TestStatus        string   `xml:"testStatus,attr"`
	ExtensionElements string   `xml:"extensionElements"`
	Incoming          []string `xml:"incoming"`
	Outgoing          []string `xml:"outgoing"`
	Name              string   `xml:"name,attr"`
}

type SequenceFlow struct {
	Text              string `xml:",chardata"`
	ID                string `xml:"id,attr"`
	RuleCondition     string `xml:"ruleCondition,attr"`
	TestStatus        string `xml:"testStatus,attr"`
	ExtensionElements string `xml:"extensionElements"`
	Name              string `xml:"name,attr"`
	SourceRef         string `xml:"sourceRef,attr"`
	TargetRef         string `xml:"targetRef,attr"`
}
type TerminateEndEvent struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}

type EndEvent struct {
	Text                     string             `xml:",chardata"`
	ID                       string             `xml:"id,attr"`
	RuleCondition            string             `xml:"ruleCondition,attr"`
	TestStatus               string             `xml:"testStatus,attr"`
	ExtensionElements        string             `xml:"extensionElements"`
	Incoming                 []string           `xml:"incoming"`
	Name                     string             `xml:"name,attr"`
	TerminateEventDefinition *TerminateEndEvent `xml:"terminateEventDefinition"`
}
type Task struct {
	Text              string   `xml:",chardata"`
	ID                string   `xml:"id,attr"`
	RuleCondition     string   `xml:"ruleCondition,attr"`
	TestStatus        string   `xml:"testStatus,attr"`
	ExtensionElements string   `xml:"extensionElements"`
	Incoming          []string `xml:"incoming"`
	Name              string   `xml:"name,attr"`
	Outgoing          []string `xml:"outgoing"`
}
type StartEvent struct {
	RuleCondition     string   `xml:"ruleCondition,attr"`
	TestStatus        string   `xml:"testStatus,attr"`
	ExtensionElements string   `xml:"extensionElements"`
	Text              string   `xml:",chardata"`
	ID                string   `xml:"id,attr"`
	Name              string   `xml:"name,attr"`
	Outgoing          []string `xml:"outgoing"`
}
type IntermediateEvent struct {
	Text                string   `xml:",chardata"`
	ID                  string   `xml:"id,attr"`
	RuleCondition       string   `xml:"ruleCondition,attr"`
	TestStatus          string   `xml:"testStatus,attr"`
	ExtensionElements   string   `xml:"extensionElements"`
	Incoming            []string `xml:"incoming"`
	Name                string   `xml:"name,attr"`
	Outgoing            []string `xml:"outgoing"`
	LinkEventDefinition struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"linkEventDefinition"`
}
type BPMN struct {
	XMLName         xml.Name `xml:"definitions"`
	Text            string   `xml:",chardata"`
	Xmlns           string   `xml:"xmlns,attr"`
	Bpmndi          string   `xml:"bpmndi,attr"`
	Omgdc           string   `xml:"omgdc,attr"`
	Omgdi           string   `xml:"omgdi,attr"`
	Xsi             string   `xml:"xsi,attr"`
	Jrules          string   `xml:"jrules,attr"`
	TargetNamespace string   `xml:"targetNamespace,attr"`
	SchemaLocation  string   `xml:"schemaLocation,attr"`
	Process         struct {
		Text                   string              `xml:",chardata"`
		ID                     string              `xml:"id,attr"`
		StartEvent             StartEvent          `xml:"startEvent"`
		ExclusiveGateway       []ExclusiveGateway  `xml:"exclusiveGateway"`
		SequenceFlow           []SequenceFlow      `xml:"sequenceFlow"`
		EndEvent               []EndEvent          `xml:"endEvent"`
		Task                   []Task              `xml:"task"`
		IntermediateThrowEvent []IntermediateEvent `xml:"intermediateThrowEvent"`
	} `xml:"process"`
	BPMNDiagram struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		BPMNPlane struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			BpmnElement string `xml:"bpmnElement,attr"`
			BPMNEdge    []struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"id,attr"`
				BpmnElement string `xml:"bpmnElement,attr"`
				Waypoint    []struct {
					Text string `xml:",chardata"`
					X    string `xml:"x,attr"`
					Y    string `xml:"y,attr"`
				} `xml:"waypoint"`
			} `xml:"BPMNEdge"`
			BPMNShape []struct {
				Text            string `xml:",chardata"`
				ID              string `xml:"id,attr"`
				BpmnElement     string `xml:"bpmnElement,attr"`
				IsMarkerVisible string `xml:"isMarkerVisible,attr"`
				Bounds          struct {
					Text   string `xml:",chardata"`
					X      string `xml:"x,attr"`
					Y      string `xml:"y,attr"`
					Width  string `xml:"width,attr"`
					Height string `xml:"height,attr"`
				} `xml:"Bounds"`
			} `xml:"BPMNShape"`
		} `xml:"BPMNPlane"`
		BPMNLabelStyle []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Font struct {
				Text            string `xml:",chardata"`
				Name            string `xml:"name,attr"`
				Size            string `xml:"size,attr"`
				IsBold          string `xml:"isBold,attr"`
				IsItalic        string `xml:"isItalic,attr"`
				IsUnderline     string `xml:"isUnderline,attr"`
				IsStrikeThrough string `xml:"isStrikeThrough,attr"`
			} `xml:"Font"`
		} `xml:"BPMNLabelStyle"`
	} `xml:"BPMNDiagram"`
}

type Bpmn struct {
	reference *BPMN
	filePath  string
}

func (self *Bpmn) loadBpmnFile() error {
	xmlFile, err := os.Open(self.filePath)
	if err != nil {
		return err
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	bpmn := new(BPMN)
	err = xml.Unmarshal(byteValue, bpmn)
	if err != nil {
		return err
	}
	self.reference = bpmn
	return nil
}
func (self *Bpmn) loadXML(data []byte) error{
	bpmn := new(BPMN)
	err := xml.Unmarshal(data, bpmn)
	if err != nil {
		return err
	}
	self.reference = bpmn

	return nil
}

func (self *Bpmn) GetBPMN(path string) (*BPMN, error) {
	self.filePath = path
	if self.reference == nil {
		err := self.loadBpmnFile()
		if err != nil {
			return nil, err
		}
	}
	return self.reference, nil

}
func (self *Bpmn) GetStartElement() StartEvent {
	return self.reference.Process.StartEvent
}
func (self *Bpmn) ForwardElement(elemId string) []*Element {
	els := make([]*Element, 0)
	el := new(Element)
	el.LoadObjElement(elemId, self.reference)

	for _, target := range el.GetOutGoings() {
		f := new(Element)
		f.LoadObjElement(target, self.reference)

		prevState := f.Element.(*SequenceFlow)

		elem := new(Element)
		elem.PrevState = prevState
		elem.LoadObjElement(f.Element.(*SequenceFlow).TargetRef, self.reference)
		els = append(els, elem)

	}
	return els
}

func NewBPMN(data interface{}) (error, *Bpmn) {
	b := new(Bpmn)
	if param,ok:=data.(string);ok {

		_, err := b.GetBPMN(param)
		if err != nil {
			return err, nil
		} else {
			return nil, b
		}
	}else if param,ok:=data.([]byte);ok{
		err := b.loadXML(param)
		if err != nil {
			return err, nil
		} else {
			return nil, b
		}
	}
	return errors.New("Invalid Param type") , nil
}
