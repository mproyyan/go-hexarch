// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	databases "github.com/mproyyan/gin-rest-api/internal/adapters/databases"
	domain "github.com/mproyyan/gin-rest-api/internal/application/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, dbtx, productId
func (_m *ProductRepository) Delete(ctx context.Context, dbtx databases.DBTX, productId int) error {
	ret := _m.Called(ctx, dbtx, productId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, databases.DBTX, int) error); ok {
		r0 = rf(ctx, dbtx, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, dbtx, productId
func (_m *ProductRepository) Find(ctx context.Context, dbtx databases.DBTX, productId int) (*domain.Product, error) {
	ret := _m.Called(ctx, dbtx, productId)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, databases.DBTX, int) *domain.Product); ok {
		r0 = rf(ctx, dbtx, productId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, databases.DBTX, int) error); ok {
		r1 = rf(ctx, dbtx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: ctx, dbtx
func (_m *ProductRepository) FindAll(ctx context.Context, dbtx databases.DBTX) ([]*domain.Product, error) {
	ret := _m.Called(ctx, dbtx)

	var r0 []*domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, databases.DBTX) []*domain.Product); ok {
		r0 = rf(ctx, dbtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, databases.DBTX) error); ok {
		r1 = rf(ctx, dbtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, dbtx, product
func (_m *ProductRepository) Save(ctx context.Context, dbtx databases.DBTX, product domain.Product) (*domain.Product, error) {
	ret := _m.Called(ctx, dbtx, product)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, databases.DBTX, domain.Product) *domain.Product); ok {
		r0 = rf(ctx, dbtx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, databases.DBTX, domain.Product) error); ok {
		r1 = rf(ctx, dbtx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, dbtx, product
func (_m *ProductRepository) Update(ctx context.Context, dbtx databases.DBTX, product domain.Product) (*domain.Product, error) {
	ret := _m.Called(ctx, dbtx, product)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, databases.DBTX, domain.Product) *domain.Product); ok {
		r0 = rf(ctx, dbtx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, databases.DBTX, domain.Product) error); ok {
		r1 = rf(ctx, dbtx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
