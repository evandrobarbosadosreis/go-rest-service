package domain

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Birth   string `db:"date_of_birth"`
	Zipcode string
	Status  string
}
