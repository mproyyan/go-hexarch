package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/mproyyan/gin-rest-api/internal/application/domain"
	"github.com/mproyyan/gin-rest-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductRepositoryFindAllSuccess(t *testing.T) {
	expected := []*domain.Product{
		{
			ID:   1,
			Name: "Test",
		},
	}

	repo := mocks.NewProductRepository(t)
	repo.On("FindAll", mock.Anything).
		Return(expected, nil)

	result, err := repo.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, expected[0].ID, result[0].ID)
	assert.Equal(t, expected[0].Name, result[0].Name)
}

func TestProductRepositoryFindAllFailed(t *testing.T) {
	repo := mocks.NewProductRepository(t)
	repo.On("FindAll", mock.Anything).
		Return(nil, errors.New("product not found"))

	result, err := repo.FindAll(context.Background())

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestProductRepositorySaveSuccess(t *testing.T) {
	prod := domain.Product{
		Name: "New",
	}

	repo := mocks.NewProductRepository(t)
	repo.On("Save", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, prod domain.Product) *domain.Product {
			prod.ID = 77

			return &prod
		}, nil)

	result, err := repo.Save(context.Background(), prod)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, prod.Name, result.Name)
	assert.Equal(t, 77, result.ID)
}

func TestProductRepositorySaveFailed(t *testing.T) {
	prod := domain.Product{
		Name: "New",
	}

	repo := mocks.NewProductRepository(t)
	repo.On("Save", mock.Anything, mock.Anything).
		Return(nil, errors.New("fail"))

	result, err := repo.Save(context.Background(), prod)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProductRepositoryFindSuccess(t *testing.T) {
	prod := domain.Product{
		ID:   77,
		Name: "Found",
	}

	repo := mocks.NewProductRepository(t)
	repo.On("Find", mock.Anything, mock.AnythingOfType("int")).
		Return(&prod, nil)

	result, err := repo.Find(context.Background(), 77)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, prod.Name, result.Name)
	assert.Equal(t, prod.ID, result.ID)
}

func TestProductRepositoryFindFailed(t *testing.T) {
	repo := mocks.NewProductRepository(t)
	repo.On("Find", mock.Anything, mock.AnythingOfType("int")).
		Return(nil, errors.New("nt found"))

	result, err := repo.Find(context.Background(), 77)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProductRepositoryUpdateSuccess(t *testing.T) {
	prod := domain.Product{
		ID:   99,
		Name: "Edit",
	}

	repo := mocks.NewProductRepository(t)
	repo.On("Update", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, prod domain.Product) *domain.Product {
			prod.Name = "Changed"
			return &prod
		}, nil)

	result, err := repo.Update(context.Background(), prod)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, "Changed", result.Name)
	assert.Equal(t, prod.ID, result.ID)
}

func TestProductRepositoryUpdateFailed(t *testing.T) {
	prod := domain.Product{
		ID:   99,
		Name: "Edit",
	}

	repo := mocks.NewProductRepository(t)
	repo.On("Update", mock.Anything, mock.Anything).
		Return(nil, errors.New("fail"))

	result, err := repo.Update(context.Background(), prod)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestProductRepositoryDeleteSuccess(t *testing.T) {
	repo := mocks.NewProductRepository(t)
	repo.On("Delete", mock.Anything, mock.Anything).
		Return(nil)

	err := repo.Delete(context.Background(), 33)

	assert.NoError(t, err)
	assert.Nil(t, err)
}

func TestProductRepositoryDeleteFailed(t *testing.T) {
	repo := mocks.NewProductRepository(t)
	repo.On("Delete", mock.Anything, mock.Anything).
		Return(errors.New("fail"))

	err := repo.Delete(context.Background(), 33)

	assert.Error(t, err)
}
