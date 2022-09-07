package mysqlservice

import (
	"context"
	"database/sql"

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

func (ps *ProductService) FindAll(ctx context.Context) ([]*domain.Product, error) {
	products, err := ps.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) Create(ctx context.Context, request domain.ProductCreateRequest) (*domain.Product, error) {
	newProduct := domain.Product{
		Name: request.Name,
	}

	product, err := ps.ProductRepository.Save(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) Find(ctx context.Context, productId int) (*domain.Product, error) {
	product, err := ps.ProductRepository.Find(ctx, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) Update(ctx context.Context, request domain.ProductUpdateRequest) (*domain.Product, error) {
	// first find product by id, if not found error returned
	existsProduct, err := ps.ProductRepository.Find(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	// change name to new
	existsProduct.Name = request.Name

	updatedProduct, err := ps.ProductRepository.Update(ctx, *existsProduct)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (ps *ProductService) Delete(ctx context.Context, productId int) error {
	// first find product by id, if not found error returned
	existsProduct, err := ps.ProductRepository.Find(ctx, productId)
	if err != nil {
		return err
	}

	err = ps.ProductRepository.Delete(ctx, existsProduct.ID)
	if err != nil {
		return err
	}

	return nil
}
