package main


type BpmnRuleCondition struct {
	ElemType      string  `json:"elem_type"   bson:"elem_type"`
	ConditionType string  `json:"condition_type"   bson:"condition_type"`
	ConditionID   int64   `json:"condition_id"   bson:"condition_id"`
	Values        []Value `json:"values"   bson:"values"`
}

type Value struct {
	Key   string `json:"key"   bson:"key"`
	Value string `json:"value"   bson:"value"`
}




type ConditionFlowElement struct {
	ConditionType string      `json:"condition_type"   bson:"condition_type"`
	ConditionID   int64       `json:"condition_id"   bson:"condition_id"`
	ElementType   string      `json:"element_type"   bson:"element_type"`
	ConditionParams string `json:"condition_params"   bson:"condition_params"`
	TrueState     *ConditionFlowElement `json:"true_state"   bson:"true_state"`
	FalseState    *ConditionFlowElement `json:"false_state"   bson:"false_state"`
	Next          []*ConditionFlowElement `json:"next"   bson:"next"`

}

