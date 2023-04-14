package infrastructure

import (
	"database/sql"
	"log"
	"time"

	"github.com/evandrobarbosadosreis/go-rest-development/domain"
	_ "github.com/go-sql-driver/mysql"
)

type customerMySqlRepository struct {
	client *sql.DB
}

func (repository *customerMySqlRepository) FindAll(status domain.Status) ([]domain.Customer, *domain.AppError) {
	result := make([]domain.Customer, 0)

	rows, err := repository.Filter(status)

	if err != nil {
		log.Println("Error getting customers: " + err.Error())
		return result, domain.NewUnexpectedError("unexpected database error")
	}

	for rows.Next() {
		var customer domain.Customer
		err = rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zip, &customer.Birth, &customer.Status)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return result, domain.NewUnexpectedError("unexpected database error")
		}
		result = append(result, customer)
	}

	return result, nil
}

func (repository *customerMySqlRepository) Filter(filter domain.Status) (*sql.Rows, error) {

	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if filter.IsValid() {
		query += " where status = ?"
		return repository.client.Query(query, filter.GetValue())
	}
	return repository.client.Query(query)
}

func (repository *customerMySqlRepository) FindById(id string) (*domain.Customer, *domain.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := repository.client.QueryRow(query, id)
	var customer domain.Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zip, &customer.Birth, &customer.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("customer not found")
		}
		log.Println("Error reading row: " + err.Error())
		return nil, domain.NewUnexpectedError("unexpected database error")
	}

	return &customer, nil
}

func NewMySqlRepository() domain.CustomerRepository {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost)/banking")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return &customerMySqlRepository{client: db}
}
