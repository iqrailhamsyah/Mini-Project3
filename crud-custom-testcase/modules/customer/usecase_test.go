package customer

import (
	"crud/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	// Create a mock customer repository
	customerRepo := &MockCustomerRepository{}

	// Create an instance of the use case with the mock repository
	uc := UsecaseCustomer{customerRepo}

	// Create a sample customer parameter
	customerParam := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Avatar:    "https://example.com/avatar.jpg",
	}

	// Expected result
	expectedCustomer := entities.Customer{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "johndoe@example.com",
		Avatar:     "https://example.com/avatar.jpg",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Mock the CreateCustomer method of the customer repository
	customerRepo.On("CreateCustomer", &expectedCustomer).Return(nil)

	// Call the CreateCustomer method
	createdCustomer, err := uc.CreateCustomer(customerParam)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the created customer matches the expected customer
	assert.Equal(t, expectedCustomer, createdCustomer)

	// Assert that the CreateCustomer method of the customer repository was called with the expected customer
	customerRepo.AssertCalled(t, "CreateCustomer", &expectedCustomer)
}

func TestGetCustomerById(t *testing.T) {
	// Create a mock customer repository
	mockRepo := &mocks.CustomerRepository{}

	// Create a sample customer
	customer := entities.Customer{
		ID:   1,
		Name: "John Doe",
		// Add other relevant fields
	}

	// Set up expectations
	mockRepo.On("GetCustomerById", uint(1)).Return(&customer, nil)

	// Create an instance of the use case with the mock repository
	useCase := UsecaseCustomer{
		customerRepo: mockRepo,
	}

	// Call the function being tested
	result, err := useCase.GetCustomerById(uint(1))

	// Assert the expected customer is returned
	assert.NoError(t, err)
	assert.Equal(t, customer, result)

	// Assert that the mock repository's GetCustomerById method was called with the correct ID
	mockRepo.AssertCalled(t, "GetCustomerById", uint(1))
}

func TestUpdateCustomerById(t *testing.T) {
	// Mock the customer repository
	mockRepo := &mocks.CustomerRepository{}

	// Create the customer use case with the mocked repository
	uc := UsecaseCustomer{customerRepo: mockRepo}

	// Define the test input
	id := uint(1)
	customer := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
	}

	// Define the expected existing customer data and repository behavior
	existingCustomer := entities.Customer{
		ID:         id,
		First_name: "Old",
		Last_name:  "Data",
		Email:      "old.email@example.com",
		Avatar:     "old_avatar.png",
		UpdatedAt:  time.Now(),
	}
	mockRepo.On("GetCustomerById", id).Return(&existingCustomer, nil)

	// Define the expected updated customer data and repository behavior
	updatedCustomer := existingCustomer
	updatedCustomer.First_name = customer.FirstName
	updatedCustomer.Last_name = customer.LastName
	updatedCustomer.Email = customer.Email
	updatedCustomer.Avatar = customer.Avatar
	updatedCustomer.UpdatedAt = time.Now()
	mockRepo.On("UpdateCustomerById", id, &existingCustomer).Return(&updatedCustomer, nil)

	// Call the function
	result, err := uc.UpdateCustomerById(id, customer)

	// Assert the result
	assert.Nil(t, err)
	assert.Equal(t, updatedCustomer, result)

	// Verify the interactions with the mock repository
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomerById(t *testing.T) {
	// Create a mock repository
	mockRepo := new(MockCustomerRepository)

	// Create a sample existing customer data
	existingData := CustomerData{ID: 1, Name: "John Doe"}

	// Create a sample ID
	id := uint(1)

	// Set up the mock repository to return the sample existing data
	mockRepo.On("GetCustomerById", id).Return(existingData, nil)

	// Create the UsecaseCustomer instance with the mock repository
	uc := UsecaseCustomer{customerRepo: mockRepo}

	// Call the DeleteCustomerById function
	err := uc.DeleteCustomerById(id)

	// Assert that the GetCustomerById method was called with the correct ID
	mockRepo.AssertCalled(t, "GetCustomerById", id)

	// Assert that the DeleteCustomerById method was called with the correct ID and existing data
	mockRepo.AssertCalled(t, "DeleteCustomerById", id, existingData)

	// Assert that no error was returned
	assert.Nil(t, err)
}
