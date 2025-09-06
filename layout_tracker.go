package main

type Tracker struct {
	Serial string `json:"serial" dynamodbav:"serial"`
	Artist string `json:"artist" dynamodbav:"artist"`

	Created  string `json:"created" dynamodbav:"created"`
	Started  string `json:"started" dynamodbav:"started"`
	Finished string `json:"finished" dynamodbav:"finished"`
	Deleted  string `json:"deleted" dynamodbav:"deleted"`

	Type     string `json:"type" dynamodbav:"type"`
	Shape    string `json:"shape" dynamodbav:"shape"`
	Length   string `json:"length" dynamodbav:"length"`
	Graphics string `json:"graphics" dynamodbav:"graphics"`

	Notes []string `json:"notes" dynamodbav:"notes"`

	Rush string `json:"rush" dynamodbav:"rush"`

	SideCnc    string `json:"side_cnc" dynamodbav:"side_cnc"`
	SidePour   string `json:"side_pour" dynamodbav:"side_pour"`
	WoodCore   string `json:"wood_core" dynamodbav:"wood_core"`
	WoodCart   string `json:"wood_cart" dynamodbav:"wood_cart"`
	SubBase    string `json:"sub_base" dynamodbav:"sub_base"`
	SubTop     string `json:"sub_top" dynamodbav:"sub_top"`
	SubCart    string `json:"sub_cart" dynamodbav:"sub_cart"`
	EdGlue     string `json:"ed_glue" dynamodbav:"ed_glue"`
	EdCart     string `json:"ed_cart" dynamodbav:"ed_cart"`
	LayPress   string `json:"lay_press" dynamodbav:"lay_press"`
	LayInspect string `json:"lay_inspect" dynamodbav:"lay_inspect"`
	FinTune    string `json:"fin_tune" dynamodbav:"fin_tune"`
	FinInspect string `json:"fin_inspect" dynamodbav:"fin_inspect"`
	WaxWax     string `json:"wax_wax" dynamodbav:"wax_wax"`
	WaxInspect string `json:"wax_inspect" dynamodbav:"wax_inspect"`
	WaxTop     string `json:"wax_top" dynamodbav:"wax_top"`
}

// type TrackerJson struct {
// 	Serial string `json:"serial"`
// 	Artist string `json:"artist"`

// 	Created  string `json:"created"`
// 	Started  string `json:"started"`
// 	Finished string `json:"finished"`
// 	Deleted  string `json:"deleted"`

// 	Type     string `json:"type"`
// 	Shape    string `json:"shape"`
// 	Length   string `json:"length"`
// 	Graphics string `json:"graphics"`

// 	Notes []string `json:"notes"`

// 	Rush string `json:"rush"`

// 	SideCnc    string `json:"side_cnc"`
// 	SidePour   string `json:"side_pour"`
// 	WoodCore   string `json:"wood_core"`
// 	WoodCart   string `json:"wood_cart"`
// 	SubBase    string `json:"sub_base"`
// 	SubTop     string `json:"sub_top"`
// 	SubCart    string `json:"sub_cart"`
// 	EdGlue     string `json:"ed_glue"`
// 	EdCart     string `json:"ed_cart"`
// 	LayPress   string `json:"lay_press"`
// 	LayInspect string `json:"lay_inspect"`
// 	FinTune    string `json:"fin_tune"`
// 	FinInspect string `json:"fin_inspect"`
// 	WaxWax     string `json:"wax_wax"`
// 	WaxInspect string `json:"wax_inspect"`
// 	WaxTop     string `json:"wax_top"`
// }

// type TrackerDynamo struct {
// 	Serial string `dynamodbav:"serial"`
// 	Artist string `dynamodbav:"artist"`

// 	Created  string `dynamodbav:"created"`
// 	Started  string `dynamodbav:"started"`
// 	Finished string `dynamodbav:"finished"`
// 	Deleted  string `dynamodbav:"deleted"`

// 	Type     string `dynamodbav:"type"`
// 	Shape    string `dynamodbav:"shape"`
// 	Length   string `dynamodbav:"length"`
// 	Graphics string `dynamodbav:"graphics"`

// 	Notes []string `dynamodbav:"notes"`

// 	Rush string `dynamodbav:"rush"`

// 	SideCnc    string `dynamodbav:"side_cnc"`
// 	SidePour   string `dynamodbav:"side_pour"`
// 	WoodCore   string `dynamodbav:"wood_core"`
// 	WoodCart   string `dynamodbav:"wood_cart"`
// 	SubBase    string `dynamodbav:"sub_base"`
// 	SubTop     string `dynamodbav:"sub_top"`
// 	SubCart    string `dynamodbav:"sub_cart"`
// 	EdGlue     string `dynamodbav:"ed_glue"`
// 	EdCart     string `dynamodbav:"ed_cart"`
// 	LayPress   string `dynamodbav:"lay_press"`
// 	LayInspect string `dynamodbav:"lay_inspect"`
// 	FinTune    string `dynamodbav:"fin_tune"`
// 	FinInspect string `dynamodbav:"fin_inspect"`
// 	WaxWax     string `dynamodbav:"wax_wax"`
// 	WaxInspect string `dynamodbav:"wax_inspect"`
// 	WaxTop     string `dynamodbav:"wax_top"`
// }
