package modals

type ZipCodeEntry struct {
	Pincode    string `json:"pincode"`
	Lang       string `json:"lang"`
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}
