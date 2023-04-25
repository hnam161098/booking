package testing

import (
	"context"
	"errors"
	"grpc/server/customer/models"
	"grpc/server/customer/testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		agrs1 := models.Customer{
			Name:    "Nam",
			Age:     "23",
			Address: "Ha noi",
		}
		customerMock.On("CreateCustomer", context.Background(), &agrs1).Return(&agrs1, nil).Once()
		resultSuccess, err := customerMock.CreateCustomer(context.Background(), &agrs1)
		assert.NoError(t, err)
		assert.NotNil(t, resultSuccess)
		customerMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		agrs2 := models.Customer{
			Name:    "Nam",
			Age:     "23",
			Address: "Ha noi",
		}
		customerMock.On("CreateCustomer", context.Background(), &agrs2).Return(nil, errors.New("Cannot create new customer")).Once()
		got, err := customerMock.CreateCustomer(context.Background(), &agrs2)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("agrs_empty", func(t *testing.T) {
		agrs2 := models.Customer{
			Name:    "",
			Age:     "",
			Address: "",
		}
		customerMock.On("CreateCustomer", context.Background(), &agrs2).Return(nil, errors.New("agrs is empty")).Once()
		resultFailed, err := customerMock.CreateCustomer(context.Background(), &agrs2)
		assert.Error(t, err)
		assert.Nil(t, resultFailed)
		customerMock.AssertExpectations(t)
	})
}

func TestGetCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		agrs1 := models.Customer{
			ID: "123",
		}

		dataFake := models.Customer{
			ID:      "123",
			Name:    "Nam",
			Age:     "23",
			Address: "Ha Noi",
		}
		customerMock.On("GetCustomer", context.Background(), &agrs1).Return(&dataFake, nil).Once()
		got, err := customerMock.GetCustomer(context.Background(), &agrs1)
		assert.NotNil(t, got)
		assert.NoError(t, err)
		customerMock.AssertExpectations(t)
	})

	t.Run("id_empty", func(t *testing.T) {
		agrs2 := models.Customer{
			ID: "",
		}
		customerMock.On("GetCustomer", context.Background(), &agrs2).Return(nil, errors.New("ID is empty")).Once()
		got, err := customerMock.GetCustomer(context.Background(), &agrs2)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("not_exist", func(t *testing.T) {
		agrs3 := models.Customer{
			ID: "111",
		}
		customerMock.On("GetCustomer", context.Background(), &agrs3).Return(nil, nil).Once()
		got, err := customerMock.GetCustomer(context.Background(), &agrs3)
		assert.NoError(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})
}

func TestGetAllCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		customerMock.On("GetAllCustomer", context.Background()).Return([]*models.Customer{
			{
				ID:      "1",
				Name:    "A",
				Age:     "23",
				Address: "Ha Noi",
			},
		}, nil).Once()
		got, err := customerMock.GetAllCustomer(context.Background())
		assert.NotNil(t, got)
		assert.NoError(t, err)
		customerMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		customerMock.On("GetAllCustomer", context.Background()).Return(nil, errors.New("Cannot get all customer"))
		got, err := customerMock.GetAllCustomer(context.Background())
		assert.Empty(t, got)
		assert.NotNil(t, err)
		customerMock.AssertExpectations(t)
	})
}

func TestUpdateCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		agrs := models.Customer{
			ID:      "123",
			Name:    "Nam 123",
			Age:     "24",
			Address: "Yen Bai",
		}

		customerMock.On("UpdateCustomer", context.Background(), &agrs).Return(&agrs, nil)
		got, err := customerMock.UpdateCustomer(context.Background(), &agrs)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("fail_id_empty", func(t *testing.T) {
		agrs := models.Customer{
			ID:      "",
			Name:    "Nam 123",
			Age:     "24",
			Address: "Yen Bai",
		}
		customerMock.On("UpdateCustomer", context.Background(), &agrs).Return(nil, errors.New("id is empty"))
		got, err := customerMock.UpdateCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		agrs := models.Customer{
			ID:      "111",
			Name:    "Nam 123",
			Age:     "24",
			Address: "Yen Bai",
		}
		customerMock.On("UpdateCustomer", context.Background(), &agrs).Return(nil, errors.New("Cannot update customer information"))
		got, err := customerMock.UpdateCustomer(context.Background(), &agrs)
		assert.Empty(t, got)
		assert.NotNil(t, err)
		customerMock.AssertExpectations(t)

	})
}

func TestAddTagsCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		tags := []string{
			"tag 1",
			"tag 2",
		}
		agrs := models.AddTags{
			ID:   "123",
			Tags: tags,
		}

		want := models.Customer{
			ID:      "123",
			Name:    "Nam",
			Age:     "23",
			Address: "Ha noi",
			Tags:    tags,
		}
		customerMock.On("AddTagsCustomer", context.Background(), &agrs).Return(&want, nil)
		got, err := customerMock.AddTagsCustomer(context.Background(), &agrs)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("fail_id_empty", func(t *testing.T) {
		tags := []string{
			"tag 1",
			"tag 2",
		}
		agrs := models.AddTags{
			ID:   "",
			Tags: tags,
		}
		customerMock.On("AddTagsCustomer", context.Background(), &agrs).Return(nil, errors.New("id is empty"))
		got, err := customerMock.AddTagsCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("fail_tags_empty", func(t *testing.T) {
		tags := []string{
			"",
			"",
		}
		agrs := models.AddTags{
			ID:   "",
			Tags: tags,
		}
		customerMock.On("AddTagsCustomer", context.Background(), &agrs).Return(nil, errors.New("tag is empty"))
		got, err := customerMock.AddTagsCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		tags := []string{
			"123",
			"456",
		}
		agrs := models.AddTags{
			ID:   "123",
			Tags: tags,
		}
		customerMock.On("AddTagsCustomer", context.Background(), &agrs).Return(nil, errors.New("Cannot add tags customer"))
		got, err := customerMock.AddTagsCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		customerMock.AssertExpectations(t)
	})
}

func TestDeleteCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}
	t.Run("success", func(t *testing.T) {
		agrs := models.Customer{
			ID: "123",
		}
		customerMock.On("DeleteCustomer", context.Background(), &agrs).Return(1, nil)
		got, err := customerMock.DeleteCustomer(context.Background(), &agrs)
		assert.NoError(t, err)
		assert.NotEmpty(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("fail_id_empty", func(t *testing.T) {
		agrs := models.Customer{
			ID: "",
		}
		customerMock.On("DeleteCustomer", context.Background(), &agrs).Return(0, errors.New("id is empty"))
		got, err := customerMock.DeleteCustomer(context.Background(), &agrs)
		assert.Equal(t, 0, got)
		assert.Error(t, err)
		customerMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		agrs := models.Customer{
			ID: "666",
		}
		customerMock.On("DeleteCustomer", context.Background(), &agrs).Return(0, errors.New("Cannot delete customer"))
		got, err := customerMock.DeleteCustomer(context.Background(), &agrs)
		assert.Equal(t, 0, got)
		assert.Error(t, err)
		customerMock.AssertExpectations(t)
	})
}

func TestDeleteTagsOfCustomer(t *testing.T) {
	customerMock := mocks.CustomerHandler{}

	t.Run("success", func(t *testing.T) {
		agrs := models.DeleteTags{
			ID: "123",
			Tags: []string{
				"123",
				"456",
			},
		}
		want := models.Customer{
			ID:   "123",
			Tags: []string{},
		}
		customerMock.On("DeleteTagsOfCustomer", context.Background(), &agrs).Return(&want, nil)
		got, err := customerMock.DeleteTagsOfCustomer(context.Background(), &agrs)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		customerMock.AssertExpectations(t)
	})

	t.Run("fail_id_empty", func(t *testing.T) {
		agrs := models.DeleteTags{
			ID: "",
			Tags: []string{
				"123",
				"456",
			},
		}
		customerMock.On("DeleteTagsOfCustomer", context.Background(), &agrs).Return(nil, errors.New("id is empty"))
		got, err := customerMock.DeleteTagsOfCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("failed", func(t *testing.T) {
		agrs := models.DeleteTags{
			ID: "666",
			Tags: []string{
				"123",
				"456",
			},
		}
		customerMock.On("DeleteTagsOfCustomer", context.Background(), &agrs).Return(nil, errors.New("cannot delete tags customer"))
		got, err := customerMock.DeleteTagsOfCustomer(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}
