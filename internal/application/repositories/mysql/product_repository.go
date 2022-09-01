package mysqlrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/mproyyan/gin-rest-api/internal/adapters/databases"
	"github.com/mproyyan/gin-rest-api/internal/application/domain"
)

type ProductRepository struct {
	DB databases.DBTX
}

func (pr *ProductRepository) FindAll(ctx context.Context) ([]*domain.Product, error) {
	sql, _, err := sq.Select("id", "name").From("products").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pr.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	result, err := pr.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = int(id)

	return &product, nil
}

func (pr *ProductRepository) Find(ctx context.Context, productId int) (*domain.Product, error) {
	sql, args, err := sq.Select("id", "name").From("products").Where("id = ?", productId).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pr.DB.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var product *domain.Product
	if rows.Next() {
		rows.Scan(&product.ID, &product.Name)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (pr *ProductRepository) Update(ctx context.Context, product domain.Product) (*domain.Product, error) {
	sql, args, err := sq.Update("products").Set("name", product.Name).Where("id = ?", product.ID).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = pr.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr *ProductRepository) Delete(ctx context.Context, productId int) error {
	sql, args, err := sq.Delete("products").Where("id = ?", productId).ToSql()
	if err != nil {
		return err
	}

	_, err = pr.DB.ExecContext(ctx, sql, args...)

	return err
}