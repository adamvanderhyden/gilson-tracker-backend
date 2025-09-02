package main

type TrackerJson struct {
	Serial string `json:"serial"`

	TimeCreate string `json:"time_create"`
	TimeStart  string `json:"time_start"`
	TimeFinish string `json:"time_finish"`

	Notes []string `json:"notes"`

	IsRush    bool `json:"is_rush"`
	IsDeleted bool `json:"is_deleted"`

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

	WhoSideCnc    string `json:"who_side_cnc"`
	WhoSidePour   string `json:"who_side_pour"`
	WhoWoodCore   string `json:"who_wood_core"`
	WhoWoodCart   string `json:"who_wood_cart"`
	WhoSubBase    string `json:"who_sub_base"`
	WhoSubTop     string `json:"who_sub_top"`
	WhoSubCart    string `json:"who_sub_cart"`
	WhoEdGlue     string `json:"who_ed_glue"`
	WhoEdCart     string `json:"who_ed_cart"`
	WhoLayPress   string `json:"who_lay_press"`
	WhoLayInspect string `json:"who_lay_inspect"`
	WhoFinTune    string `json:"who_fin_tune"`
	WhoFinInspect string `json:"who_fin_inspect"`
	WhoWaxWax     string `json:"who_wax_wax"`
	WhoWaxInspect string `json:"who_wax_inspect"`
	WhoWaxTop     string `json:"who_wax_top"`
}

type TrackerDynamo struct {
	Serial string `dynamodbav:"serial"`

	TimeCreate string `dynamodbav:"time_create"`
	TimeStart  string `dynamodbav:"time_start"`
	TimeFinish string `dynamodbav:"time_finish"`

	Notes []string `dynamodbav:"notes"`

	IsRush    bool `dynamodbav:"is_rush"`
	IsDeleted bool `dynamodbav:"is_deleted"`

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

	WhoSideCnc    string `dynamodbav:"who_side_cnc"`
	WhoSidePour   string `dynamodbav:"who_side_pour"`
	WhoWoodCore   string `dynamodbav:"who_wood_core"`
	WhoWoodCart   string `dynamodbav:"who_wood_cart"`
	WhoSubBase    string `dynamodbav:"who_sub_base"`
	WhoSubTop     string `dynamodbav:"who_sub_top"`
	WhoSubCart    string `dynamodbav:"who_sub_cart"`
	WhoEdGlue     string `dynamodbav:"who_ed_glue"`
	WhoEdCart     string `dynamodbav:"who_ed_cart"`
	WhoLayPress   string `dynamodbav:"who_lay_press"`
	WhoLayInspect string `dynamodbav:"who_lay_inspect"`
	WhoFinTune    string `dynamodbav:"who_fin_tune"`
	WhoFinInspect string `dynamodbav:"who_fin_inspect"`
	WhoWaxWax     string `dynamodbav:"who_wax_wax"`
	WhoWaxInspect string `dynamodbav:"who_wax_inspect"`
	WhoWaxTop     string `dynamodbav:"who_wax_top"`
}
