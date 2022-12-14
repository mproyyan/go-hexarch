package tests

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mproyyan/gin-rest-api/internal/application/domain"
	mysqlrepo "github.com/mproyyan/gin-rest-api/internal/application/repositories/mysql"
	mysqlservice "github.com/mproyyan/gin-rest-api/internal/application/services/mysql"
	"github.com/stretchr/testify/assert"
)

var p *domain.Product = &domain.Product{
	ID:   1,
	Name: "Product Test",
}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func newService(db *sql.DB) *mysqlservice.ProductService {
	return mysqlservice.NewProductService(
		db,
		mysqlrepo.NewProductRepository(),
	)
}

func expectationFulfilled(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestProductServiceFindAllSuccess(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findAllQuery := "SELECT id, name FROM products"

	data := sqlmock.NewRows([]string{"ID", "Name"}).
		AddRow(p.ID, p.Name).
		AddRow(p.ID, p.Name)

	mock.ExpectQuery(findAllQuery).WillReturnRows(data)

	service := newService(db)
	products, err := service.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	expectationFulfilled(t, mock)
}

func TestProductServiceFindAllFailed(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findAllQuery := "SELECT id, name FROM products"

	mock.ExpectQuery(findAllQuery).WillReturnError(errors.New("find all product failur"))

	service := newService(db)
	products, err := service.FindAll(context.Background())

	assert.Error(t, err)
	assert.Nil(t, products)
	expectationFulfilled(t, mock)
}

func TestProductServiceCreateSuccess(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectExec("INSERT INTO products").
		WithArgs(p.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	service := newService(db)
	product, err := service.Create(
		context.Background(),
		domain.ProductCreateRequest{Name: "Product Test"},
	)

	assert.NoError(t, err)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Name, product.Name)
	expectationFulfilled(t, mock)
}

func TestProductServiceCreateFailed(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectExec("INSERT INTO products").
		WithArgs(p.Name).
		WillReturnError(errors.New("create product failur"))

	service := newService(db)
	product, err := service.Create(
		context.Background(),
		domain.ProductCreateRequest{Name: "Product Test"},
	)

	assert.Error(t, err)
	assert.Nil(t, product)
	expectationFulfilled(t, mock)
}

func TestProductServiceFindSuccess(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.ID, p.Name)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM products WHERE id = ? LIMIT 1")).
		WithArgs(p.ID).
		WillReturnRows(data)

	service := newService(db)
	product, err := service.Find(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, p.Name, product.Name)
	expectationFulfilled(t, mock)
}

func TestProductServiceFindFailed(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name FROM products WHERE id = ? LIMIT 1")).
		WithArgs(p.ID).
		WillReturnError(errors.New("product not found"))

	service := newService(db)
	product, err := service.Find(context.Background(), 1)

	assert.Error(t, err)
	assert.Nil(t, product)
	expectationFulfilled(t, mock)
}

func TestProductServiceUpdateSuccess(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	updateQuery := regexp.QuoteMeta(`UPDATE products SET name = ? WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.ID, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.ID).WillReturnRows(data)
	mock.ExpectExec(updateQuery).WithArgs("Updated", 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	service := newService(db)
	product, err := service.Update(
		context.Background(),
		domain.ProductUpdateRequest{ID: 1, Name: "Updated"},
	)

	assert.NoError(t, err)
	assert.Equal(t, p.ID, product.ID)
	assert.Equal(t, "Updated", product.Name)
	expectationFulfilled(t, mock)
}

func TestProductServiceUpdateFailed(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	updateQuery := regexp.QuoteMeta(`UPDATE products SET name = ? WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.ID, p.Name)

	mock.ExpectBegin()
	mock.ExpectQuery(findQuery).WithArgs(p.ID).WillReturnRows(data)
	// product update return error then rollback the transactiom
	mock.ExpectExec(updateQuery).WithArgs("Updated", 1).WillReturnError(errors.New("product update failur"))
	mock.ExpectRollback()

	service := newService(db)
	product, err := service.Update(
		context.Background(),
		domain.ProductUpdateRequest{ID: 1, Name: "Updated"},
	)

	assert.Error(t, err)
	assert.Nil(t, product)
	expectationFulfilled(t, mock)
}

func TestProductServiceDeleteSuccess(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)
	deleteQuery := regexp.QuoteMeta(`DELETE FROM products WHERE id = ?`)

	data := sqlmock.NewRows([]string{"Id", "Name"}).AddRow(p.ID, p.Name)

	mock.ExpectQuery(findQuery).WithArgs(p.ID).WillReturnRows(data)
	mock.ExpectExec(deleteQuery).WithArgs(p.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	service := newService(db)
	err := service.Delete(context.Background(), 1)

	assert.NoError(t, err)
	expectationFulfilled(t, mock)
}

func TestProductServiceDeleteFailed(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	findQuery := regexp.QuoteMeta(`SELECT id, name FROM products WHERE id = ? LIMIT 1`)

	mock.ExpectQuery(findQuery).WithArgs(p.ID).WillReturnError(errors.New("product not found"))

	service := newService(db)
	err := service.Delete(context.Background(), 1)

	assert.Error(t, err)
	expectationFulfilled(t, mock)
}
