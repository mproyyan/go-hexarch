package domain

import "context"

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCreateRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=1,max=200,alpha"`
}

type ProductUpdateRequest struct {
	ID   int    `json:"id" form:"id" binding:"required,numeric"`
	Name string `json:"name" form:"name" binding:"required,min=1,max=200,alpha"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*Product, error)
	Save(ctx context.Context, product Product) (*Product, error)
	Find(ctx context.Context, productId int) (*Product, error)
	Update(ctx context.Context, product Product) (*Product, error)
	Delete(ctx context.Context, productId int) error
}

type ProductService interface {
	FindAll(ctx context.Context) []*Product
	Create(ctx context.Context, request ProductCreateRequest) *Product
	Find(ctx context.Context, productId int) *Product
	Update(ctx context.Context, request ProductUpdateRequest) *Product
	Delete(ctx context.Context, productId int)
}
