package entities

//Shorten is struct of entitie for store db.
type Shorten struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Visits int    `json:"visits"`
}
