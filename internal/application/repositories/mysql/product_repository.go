package mysqlrepo

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	cuserr "github.com/mproyyan/gin-rest-api/errors"
	"github.com/mproyyan/gin-rest-api/internal/application/domain"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	sql, _, err := sq.Select("id", "name").From("products").ToSql()
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	rows, err := pr.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := new(domain.Product)
		rows.Scan(&product.ID, &product.Name)
		if err = rows.Err(); err != nil {
			return nil, err
		}

		// append scanned result
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) Save(ctx context.Context, product domain.Product) (*domain.Product, error) {
	sql, args, err := sq.Insert("products").Columns("name").Values(product.Name).ToSql()
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	result, err := pr.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	product.ID = int(id)

	return &product, nil
}

func (pr *ProductRepository) Find(ctx context.Context, productId int) (*domain.Product, error) {
	sql, args, err := sq.Select("id", "name").From("products").Where("id = ?", productId).Limit(1).ToSql()
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	rows, err := pr.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		rows.Scan(&product.ID, &product.Name)
		if err = rows.Err(); err != nil {
			return nil, cuserr.NewInternalServerErr().Wrap(err)
		}
	} else {
		notFound := fmt.Errorf("you tried to search for a product with id %d, and no results were found", productId)
		return nil, cuserr.NewProductNotFoundErr().Wrap(notFound)
	}

	return &product, nil
}

func (pr *ProductRepository) Update(ctx context.Context, product domain.Product) (*domain.Product, error) {
	sql, args, err := sq.Update("products").Set("name", product.Name).Where("id = ?", product.ID).ToSql()
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	_, err = pr.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, cuserr.NewInternalServerErr().Wrap(err)
	}

	return &product, nil
}

func (pr *ProductRepository) Delete(ctx context.Context, productId int) error {
	sql, args, err := sq.Delete("products").Where("id = ?", productId).ToSql()
	if err != nil {
		return cuserr.NewInternalServerErr().Wrap(err)
	}

	_, err = pr.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		cuserr.NewInternalServerErr().Wrap(err)
	}

	return nil
}
