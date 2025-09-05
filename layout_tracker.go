package main

type TrackerJson struct {
	Serial string `json:"serial"`
	Artist string `json:"artist"`

	Created  string `json:"created"`
	Started  string `json:"started"`
	Finished string `json:"finished"`
	Deleted  string `json:"deleted"`

	Type     string `json:"type"`
	Shape    string `json:"shape"`
	Length   string `json:"length"`
	Graphics string `json:"graphics"`

	Notes []string `json:"notes"`

	Rush string `json:"rush"`

	SideCnc    string `json:"side_cnc"`
	SidePour   string `json:"side_pour"`
	WoodCore   string `json:"wood_core"`
	WoodCart   string `json:"wood_cart"`
	SubBase    string `json:"sub_base"`
	SubTop     string `json:"sub_top"`
	SubCart    string `json:"sub_cart"`
	EdGlue     string `json:"ed_glue"`
	EdCart     string `json:"ed_cart"`
	LayPress   string `json:"lay_press"`
	LayInspect string `json:"lay_inspect"`
	FinTune    string `json:"fin_tune"`
	FinInspect string `json:"fin_inspect"`
	WaxWax     string `json:"wax_wax"`
	WaxInspect string `json:"wax_inspect"`
	WaxTop     string `json:"wax_top"`
}

type TrackerDynamo struct {
	Serial string `dynamodbav:"serial"`
	Artist string `dynamodbav:"artist"`

	Created  string `dynamodbav:"created"`
	Started  string `dynamodbav:"started"`
	Finished string `dynamodbav:"finished"`
	Deleted  string `dynamodbav:"deleted"`

	Type     string `dynamodbav:"type"`
	Shape    string `dynamodbav:"shape"`
	Length   string `dynamodbav:"length"`
	Graphics string `dynamodbav:"graphics"`

	Notes []string `dynamodbav:"notes"`

	Rush string `dynamodbav:"rush"`

	SideCnc    string `dynamodbav:"side_cnc"`
	SidePour   string `dynamodbav:"side_pour"`
	WoodCore   string `dynamodbav:"wood_core"`
	WoodCart   string `dynamodbav:"wood_cart"`
	SubBase    string `dynamodbav:"sub_base"`
	SubTop     string `dynamodbav:"sub_top"`
	SubCart    string `dynamodbav:"sub_cart"`
	EdGlue     string `dynamodbav:"ed_glue"`
	EdCart     string `dynamodbav:"ed_cart"`
	LayPress   string `dynamodbav:"lay_press"`
	LayInspect string `dynamodbav:"lay_inspect"`
	FinTune    string `dynamodbav:"fin_tune"`
	FinInspect string `dynamodbav:"fin_inspect"`
	WaxWax     string `dynamodbav:"wax_wax"`
	WaxInspect string `dynamodbav:"wax_inspect"`
	WaxTop     string `dynamodbav:"wax_top"`
}
