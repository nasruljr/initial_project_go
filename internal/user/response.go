package user

type PersonalInfoResponse struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

type ResponseGetUsers struct {
	Records []Users `json:"records"`
	Total   int64   `json:"total"`
}
