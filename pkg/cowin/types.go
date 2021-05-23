package cowin

type CowinResponse struct {
	Centers []Center `json:"centers"`
}

type Center struct {
	Name     string
	Address  string
	District string `json:"district_name"`
	State    string `json:"state_name"`
	Pincode  int
	FeeType  string `json:"fee_type"`
	Sessions []Session
}

type Session struct {
	Date                   string
	AvailableCapacity      int `json:"available_capacity"`
	AvailableCapacityDose1 int `json:"available_capacity_dose1"`
	AvailableCapacityDose2 int `json:"available_capacity_dose2"`
	MinAge                 int `json:"min_age_limit"`
	Vaccine                string
	Slots                  []string
}

type Filters struct {
	Age          int
	Dose         int
	DistrictCode string
}

type StatesResponse struct {
	States []State
}

type State struct {
	Name string `json:"state_name"`
	Code int    `json:"state_id"`
}

type DistrictsResponse struct {
	Districts []District
}

type District struct {
	Name string `json:"district_name"`
	Code int    `json:"district_id"`
}
