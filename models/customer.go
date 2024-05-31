package models

type CustomersBill struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Age       int    `json:"age,omitempty" bson:"age,omitempty"`
	Gender    string `json:"gender,omitempty" bson:"gender,omitempty"`
}

type Customer struct {
	PersonalData PersonalData `json:"personal_data,omitempty" bson:"personal_data,omitempty"`
	Bills        []Bills      `json:"bills,omitempty" bson:"bills,omitempty"`
}

type PersonalData struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Age       int    `json:"age,omitempty" bson:"age,omitempty"`
	Gender    string `json:"gender,omitempty" bson:"gender,omitempty"`
}
