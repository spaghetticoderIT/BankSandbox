package create

type CreateBankDto struct {
	Name         string `json:"name"`
	CountryCode  string `json:"countryCode"`
	BIC          string `json:"bic"`
	OwnersNumber uint16 `json:"owners"`
}
