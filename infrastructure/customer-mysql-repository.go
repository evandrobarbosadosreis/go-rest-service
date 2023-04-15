package infrastructure

import (
	"database/sql"
	"time"

	"github.com/evandrobarbosadosreis/go-rest-development/domain"
	"github.com/evandrobarbosadosreis/go-rest-development/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type customerMySqlRepository struct {
	client *sqlx.DB
}

func (repository *customerMySqlRepository) FindAll(status domain.Status) ([]domain.Customer, *domain.AppError) {

	var err error

	result := make([]domain.Customer, 0)

	if status.IsValid() {
		err = repository.client.Select(&result, "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?", status.GetValue())
	} else {
		err = repository.client.Select(&result, "select customer_id, name, city, zipcode, date_of_birth, status from customers")
	}

	if err != nil {
		logger.Error("Error getting customers: " + err.Error())
		return nil, domain.NewUnexpectedError("unexpected database error")
	}

	return result, nil
}

func (repository *customerMySqlRepository) FindById(id string) (*domain.Customer, *domain.AppError) {

	var result domain.Customer

	err := repository.client.Get(&result, "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?", id)

	if err == nil {
		return &result, nil
	}

	if err == sql.ErrNoRows {
		return nil, domain.NewNotFoundError("customer not found")
	}

	logger.Error("Error reading row: " + err.Error())

	return nil, domain.NewUnexpectedError("unexpected database error")
}

func NewMySqlRepository() domain.CustomerRepository {

	db, err := sqlx.Open("mysql", "root:mysql@tcp(localhost)/banking")

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return &customerMySqlRepository{client: db}
}
