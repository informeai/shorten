package entities

//Shorten is struct of entitie for store db.
type Shorten struct {
	Id     string `json:"id" bson:"key"`
	Url    string `json:"url" bson:"url"`
	Visits int    `json:"visits" bson:"visits"`
}
