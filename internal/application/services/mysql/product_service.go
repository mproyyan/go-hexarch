package mysqlservice

import (
	"context"
	"database/sql"

	cuserr "github.com/mproyyan/gin-rest-api/errors"
	"github.com/mproyyan/gin-rest-api/internal/application/domain"
)

type ProductService struct {
	DB                *sql.DB
	ProductRepository domain.ProductRepository
}

func NewProductService(db *sql.DB, productRepository domain.ProductRepository) *ProductService {
	return &ProductService{
		DB:                db,
		ProductRepository: productRepository,
	}
}

func (ps *ProductService) transaction(ctx context.Context, cb func(*sql.Tx) error) error {
	// start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return cuserr.NewInternalServerErr().Wrap(err)
	}

	// run callback
	// if callback return error then rollback the transaction
	err = cb(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			// log.Fatalf("Transaction rollback failur, cause of rollback : %s", err.Error())
			return cuserr.NewInternalServerErr().Wrap(rbErr)
		}

		return err
	}

	// if callback successfully executed and no error return
	// then commit the transaction and save all changes
	err = tx.Commit()
	if err != nil {
		return cuserr.NewInternalServerErr().Wrap(err)
	}

	return nil
}

func (ps *ProductService) FindAll(ctx context.Context) ([]*domain.Product, error) {
	products, err := ps.ProductRepository.FindAll(ctx, ps.DB)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) Create(ctx context.Context, request domain.ProductCreateRequest) (*domain.Product, error) {
	newProduct := domain.Product{
		Name: request.Name,
	}

	product, err := ps.ProductRepository.Save(ctx, ps.DB, newProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) Find(ctx context.Context, productId int) (*domain.Product, error) {
	product, err := ps.ProductRepository.Find(ctx, ps.DB, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) Update(ctx context.Context, request domain.ProductUpdateRequest) (*domain.Product, error) {
	var updatedProduct *domain.Product

	err := ps.transaction(ctx, func(tx *sql.Tx) error {
		var err error

		// first find product by id, if not found error returned
		existsProduct, err := ps.ProductRepository.Find(ctx, tx, request.ID)
		if err != nil {
			return err
		}

		// change name to new
		existsProduct.Name = request.Name

		// update product with new data
		updatedProduct, err = ps.ProductRepository.Update(ctx, tx, *existsProduct)
		if err != nil {
			return err
		}

		return err
	})

	// check the transaction is valid or not
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (ps *ProductService) Delete(ctx context.Context, productId int) error {
	// first find product by id, if not found error returned
	existsProduct, err := ps.ProductRepository.Find(ctx, ps.DB, productId)
	if err != nil {
		return err
	}

	err = ps.ProductRepository.Delete(ctx, ps.DB, existsProduct.ID)
	if err != nil {
		return err
	}

	return err
}
