package mysqlservice

import (
	"context"
	"database/sql"
	"log"

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

func (ps *ProductService) FindAll(ctx context.Context) []*domain.Product {
	products, err := ps.ProductRepository.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return products
}

func (ps *ProductService) Create(ctx context.Context, request domain.ProductCreateRequest) *domain.Product {
	newProduct := domain.Product{
		Name: request.Name,
	}

	product, err := ps.ProductRepository.Save(ctx, newProduct)
	if err != nil {
		log.Fatal(err)
	}

	return product
}

func (ps *ProductService) Find(ctx context.Context, productId int) *domain.Product {
	product, err := ps.ProductRepository.Find(ctx, productId)
	if err != nil {
		log.Fatal(err)
	}

	return product
}

func (ps *ProductService) Update(ctx context.Context, request domain.ProductUpdateRequest) *domain.Product {
	// first find product by id, if not found error returned
	existsProduct, err := ps.ProductRepository.Find(ctx, request.ID)
	if err != nil {
		log.Fatal(err)
	}

	// change name to new
	existsProduct.Name = request.Name

	updatedProduct, err := ps.ProductRepository.Update(ctx, *existsProduct)
	if err != nil {
		log.Fatal(err)
	}

	return updatedProduct
}

func (ps *ProductService) Delete(ctx context.Context, productId int) {
	// first find product by id, if not found error returned
	existsProduct, err := ps.ProductRepository.Find(ctx, productId)
	if err != nil {
		log.Fatal(err)
	}

	err = ps.ProductRepository.Delete(ctx, existsProduct.ID)
	if err != nil {
		log.Fatal(err)
	}
}
