package models

type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"` //omitempty is used to ignore the empty fields
	//make email unique
	Email    string `json:"email" bson:"email"`  
	Password string `json:"password" bson:"password"`
	Todos    []Todo `json:"todos" bson:"todos"`
}
