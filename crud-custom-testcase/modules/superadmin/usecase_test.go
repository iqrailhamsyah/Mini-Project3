package superadmin

import (
	"crud/entities"
	repositories "crud/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) LoginAdmin(username string) (*entities.Actor, error) {
	args := m.Called(username)
	if actor, ok := args.Get(0).(*entities.Actor); ok {
		return actor, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestLoginAdmin_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{}
	usecaseAdmin := UsecaseAdmin{adminRepo: mockRepo}
	adminID := uint(1)
	username := "test_admin"
	password := "password123"
	expectedToken := "jwt_token"

	adminUser := &entities.Actor{
		ID:       adminID,
		Username: username,
		Password: "hashed_password",
	}

	mockRepo.On("LoginAdmin", username).Return(adminUser, nil)

	// Act
	admin, token, err := usecaseAdmin.LoginAdmin(adminID, username, password)

	// Assert
	assert.NotNil(t, admin)
	assert.Equal(t, adminUser, admin)
	assert.Equal(t, expectedToken, token)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{}
	usecaseAdmin := UsecaseAdmin{adminRepo: mockRepo}
	username := "test_admin"
	expectedError := errors.New("repository error")

	mockRepo.On("LoginAdmin", username).Return(nil, expectedError)

	// Act
	admin, token, err := usecaseAdmin.LoginAdmin(0, username, "password")

	// Assert
	assert.Nil(t, admin)
	assert.Empty(t, token)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestLoginAdmin_IncorrectPassword(t *testing.T) {
	// Arrange
	mockRepo := &MockRepository{}
	usecase

	func TestRegisterAdmin(t *testing.T) {
		// Step 2: Create a mock repository implementation
		mockRepo := &MockAdminRepository{} // Replace MockAdminRepository with your own mock implementation

		// Step 3: Set up test environment
		uc := UsecaseAdmin{
			adminRepo: mockRepo,
		}

		// Step 4: Define test inputs
		admin := LoginAdminParam{
			Username: "testadmin",
			Password: "password123",
		}

		// Step 5: Invoke the function
		result, err := uc.RegisterAdmin(admin)

		// Step 6: Verify the result and error
		assert.NoError(t, err) // Assert that no error occurred
		assert.NotNil(t, result) // Assert that the result is not nil

		// Additional assertions specific to the returned result, e.g., checking the created admin's properties
		assert.Equal(t, admin.Username, result.Username)
		assert.Equal(t, entities.False, result.IsVerified)
		assert.Equal(t, entities.False, result.IsActived)

		// Clean up if necessary
		// ...

	}

	func TestGetAdminById(t *testing.T) {
		// Create a mock implementation of the admin repository
		mockAdminRepo := &MockAdminRepository{}

		// Create an instance of the use case with the mock repository
		usecase := UsecaseAdmin{adminRepo: mockAdminRepo}

		// Define the test case input
		adminID := uint(1)

		// Define the expected output
		expectedAdmin := entities.Actor{
			// Define the expected properties of the admin entity
			// ...
		}

		// Set up the mock repository to return the expected admin entity
		mockAdminRepo.On("GetAdminById", adminID).Return(&expectedAdmin, nil)

		// Call the function under test
		admin, err := usecase.GetAdminById(adminID)

		// Assert that the function call was successful
		assert.NoError(t, err)

		// Assert that the returned admin entity matches the expected admin entity
		assert.Equal(t, expectedAdmin, admin)

		// Assert that the mock repository's GetAdminById method was called with the correct input
		mockAdminRepo.AssertCalled(t, "GetAdminById", adminID)
	}

	// MockAdminRepository is a mock implementation of the admin repository
	type MockAdminRepository struct {
		mock.Mock
	}

	// GetAdminById is a mocked method to get an admin entity by ID
	func (m *MockAdminRepository) GetAdminById(id uint) (*entities.Actor, error) {
		args := m.Called(id)
		admin := args.Get(0)
		err := args.Error(1)
		if admin != nil {
			return admin.(*entities.Actor), err
		}
		return nil, err
	}

	func TestUpdateCustomerById(t *testing.T) {
		// Create a mock admin parameter
		admin := AdminParam{
			Username: "newadmin",
			Password: "newpassword",
		}

		// Create a mock existing admin
		existingAdmin := entities.Actor{
			ID:       1,
			Username: "oldadmin",
			Password: "oldpassword",
			UpdatedAt: time.Now().Add(-time.Hour), // Set a past update time
		}

		// Create a mock updated admin
		updatedAdmin := entities.Actor{
			ID:        1,
			Username:  "newadmin",
			Password:  "newhashedpassword",
			UpdatedAt: time.Now(), // Set the current update time
		}

		// Create a mock admin repository
		mockAdminRepo := &MockAdminRepository{
			GetAdminByIdFunc: func(id uint) (*entities.Actor, error) {
				return &existingAdmin, nil
			},
			UpdateAdminByIdFunc: func(id uint, admin *entities.Actor) (*entities.Actor, error) {
				return &updatedAdmin, nil
			},
		}

		// Create an instance of the UsecaseAdmin with the mock repository
		usecaseAdmin := UsecaseAdmin{adminRepo: mockAdminRepo}

		// Call the UpdateCustomerById function
		result, err := usecaseAdmin.UpdateCustomerById(1, admin)

		// Check the returned result
		assert.NoError(t, err)
		assert.Equal(t, updatedAdmin, result)

		// Check the changes in the existing admin
		assert.Equal(t, admin.Username, existingAdmin.Username)
		assert.Equal(t, "newhashedpassword", existingAdmin.Password) // Ensure the password is hashed
		assert.WithinDuration(t, time.Now(), existingAdmin.UpdatedAt, time.Second) // Ensure the update time is set to the current time
	}

	// MockAdminRepository is a mock implementation of the admin repository interface
	type MockAdminRepository struct {
		GetAdminByIdFunc    func(id uint) (*entities.Actor, error)
		UpdateAdminByIdFunc func(id uint, admin *entities.Actor) (*entities.Actor, error)
	}

	func (m *MockAdminRepository) GetAdminById(id uint) (*entities.Actor, error) {
		return m.GetAdminByIdFunc(id)
	}

	func (m *MockAdminRepository) UpdateAdminById(id uint, admin *entities.Actor) (*entities.Actor, error) {
		return m.UpdateAdminByIdFunc(id, admin)
	}

	type MockAdminRepo struct {
		GetAdminByIdFn      func(id uint) (Admin, error)
		DeleteAdminByIdFn   func(id uint, admin Admin) error
	}

	// GetAdminById calls the mock implementation GetAdminByIdFn.
	func (m *MockAdminRepo) GetAdminById(id uint) (Admin, error) {
		return m.GetAdminByIdFn(id)
	}

	// DeleteAdminById calls the mock implementation DeleteAdminByIdFn.
	func (m *MockAdminRepo) DeleteAdminById(id uint, admin Admin) error {
		return m.DeleteAdminByIdFn(id, admin)
	}

	func TestDeleteAdminById(t *testing.T) {
		// Define a sample admin ID
		adminID := uint(1)

		// Create a mock admin data
		mockAdmin := Admin{
			ID:        adminID,
			FirstName: "John",
			LastName:  "Doe",
		}

		// Create a mock admin repository
		mockRepo := &MockAdminRepo{
			GetAdminByIdFn: func(id uint) (Admin, error) {
				if id == adminID {
					return mockAdmin, nil
				}
				return Admin{}, errors.New("Admin not found")
			},
			DeleteAdminByIdFn: func(id uint, admin Admin) error {
				if id == adminID && admin == mockAdmin {
					return nil
				}
				return errors.New("Failed to delete admin")
			},
		}

		// Create an instance of the UsecaseAdmin with the mock repository
		uc := UsecaseAdmin{
			adminRepo: mockRepo,
		}

		// Call the DeleteAdminById function
		err := uc.DeleteAdminById(adminID)

		// Assert that no error occurred
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		// Assert that the DeleteAdminById method of the mock repository was called with the correct arguments
		if mockRepo.DeleteAdminByIdFnCalledWithID != adminID || mockRepo.DeleteAdminByIdFnCalledWithAdmin != mockAdmin {
			t.Errorf("DeleteAdminById was not called with the correct arguments")
		}
	}

	func TestCreateCustomer(t *testing.T) {
		// Mock dependencies
		adminRepo := &MockAdminRepository{}
		uc := UsecaseAdmin{adminRepo}

		// Input data
		customer := CustomerParam{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Avatar:    "https://example.com/avatar.png",
		}

		// Expected output
		expectedCustomer := entities.Customer{
			First_name: "John",
			Last_name:  "Doe",
			Email:      "john.doe@example.com",
			Avatar:     "https://example.com/avatar.png",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Mock repository behavior
		adminRepo.On("CreateCustomer", &expectedCustomer).Return(nil)

		// Call the function
		createdCustomer, err := uc.CreateCustomer(customer)

		// Assert the result
		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer, createdCustomer)

		// Assert that the repository method was called with the expected arguments
		adminRepo.AssertCalled(t, "CreateCustomer", &expectedCustomer)
	}

	func TestDeleteCustomerById(t *testing.T) {
		// Create a mock repository
		mockRepo := &MockAdminRepository{}

		// Create the use case instance
		uc := UsecaseAdmin{
			adminRepo: mockRepo,
		}

		// Set up the test data
		customerID := uint(1)
		existingData := Customer{
			ID:   customerID,
			Name: "John Doe",
			// Add other relevant fields
		}

		// Mock the GetCustomerById function to return the existing data
		mockRepo.On("GetCustomerById", customerID).Return(existingData, nil)

		// Mock the DeleteCustomerById function to return nil error
		mockRepo.On("DeleteCustomerById", customerID, existingData).Return(nil)

		// Call the DeleteCustomerById function
		err := uc.DeleteCustomerById(customerID)

		// Assert that the GetCustomerById function was called with the correct customer ID
		mockRepo.AssertCalled(t, "GetCustomerById", customerID)

		// Assert that the DeleteCustomerById function was called with the correct customer ID and existing data
		mockRepo.AssertCalled(t, "DeleteCustomerById", customerID, existingData)

		// Assert that no error was returned
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	}

	func TestDeleteCustomerById_ErrorGettingCustomerData(t *testing.T) {
		// Create a mock repository
		mockRepo := &MockAdminRepository{}

		// Create the use case instance
		uc := UsecaseAdmin{
			adminRepo: mockRepo,
		}

		// Set up the test data
		customerID := uint(1)

		// Mock the GetCustomerById function to return an error
		expectedErr := errors.New("error getting customer data")
		mockRepo.On("GetCustomerById", customerID).Return(Customer{}, expectedErr)

		// Call the DeleteCustomerById function
		err := uc.DeleteCustomerById(customerID)

		// Assert that the GetCustomerById function was called with the correct customer ID
		mockRepo.AssertCalled(t, "GetCustomerById", customerID)

		// Assert that the DeleteCustomerById function was not called
		mockRepo.AssertNotCalled(t, "DeleteCustomerById", customerID, Customer{})

		// Assert that the expected error was returned
		if err != expectedErr {
			t.Errorf("expected error '%v', got '%v'", expectedErr, err)
		}
	}

	func TestDeleteCustomerById_ErrorDeletingCustomer(t *testing.T) {
		// Create a mock repository
		mockRepo := &MockAdminRepository{}

		// Create the use case instance
		uc := UsecaseAdmin{
			adminRepo: mockRepo,
		}

		// Set up the test data
		customerID := uint(1)
		existingData := Customer{
			ID:   customerID,
			Name: "John Doe",
			// Add other relevant fields
		}

		// Mock the GetCustomerById function to return the existing data
		mockRepo.On("GetCustomerById", customerID).Return(existingData, nil)

		// Mock the DeleteCustomerById function to return an error
		expectedErr := errors.New("error deleting customer")
		mockRepo.On("DeleteCustomerById", customerID, existingData).Return(expectedErr)

		// Call the DeleteCustomerById function
		err := uc.DeleteCustomerById(customerID)

		// Assert that the GetCustomerById function was called with the correct customer ID
		mockRepo.AssertCalled(t, "GetCustomerById", customerID)

		// Assert that the DeleteCustomerById function was called with the correct

		func TestGetAllCustomers(t *testing.T) {
			// Step 1: Set up dependencies and mocks

			// Create a mock repository
			mockRepo := &mocks.AdminRepository{}

			// Initialize the use case with the mock repository
			uc := UsecaseAdmin{adminRepo: mockRepo}

			// Define input parameters
			firstName := "John"
			lastName := "Doe"
			email := "john.doe@example.com"
			page := 1
			pageSize := 10

			// Step 2: Define expected output

			// Create an expected list of customers
			expectedCustomers := []*entities.Customer{
				{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
				{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com"},
			}

			// Step 3: Execute the function

			// Set up the mock repository to return the expected customers
			mockRepo.On("GetAllCustomers", firstName, lastName, email, page, pageSize).Return(expectedCustomers, nil)

			// Call the function being tested
			customers, err := uc.GetAllCustomers(firstName, lastName, email, page, pageSize)

			// Step 4: Compare the actual output with the expected output

			// Assert that the returned customers match the expected customers
			assert.Equal(t, expectedCustomers, customers)

			// Assert that the returned error is nil
			assert.Nil(t, err)

			// Optionally, assert any other conditions or edge cases
		}

		// MockAdminRepo is a mock implementation of the AdminRepository interface
		type MockAdminRepo struct{}

		func (m *MockAdminRepo) SaveCustomersFromAPI(url string) error {
			// Mock implementation that always returns nil
			return nil
		}

		func TestSaveCustomersFromAPI_Success(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: &MockAdminRepo{},
			}

			err := uc.SaveCustomersFromAPI()
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		}

		func TestSaveCustomersFromAPI_Error(t *testing.T) {
			uc := UsecaseAdmin{
				adminRepo: &MockAdminRepo{},
			}

			mockErr := errors.New("mock error")
			uc.adminRepo = &MockAdminRepo{
				SaveCustomersFromAPI: func(url string) error {
					return mockErr
				},
			}

			err := uc.SaveCustomersFromAPI()
			if err != mockErr {
				t.Errorf("Expected error %v, got %v", mockErr, err)
			}
		}

		// MockSuperadminRepo is a mock implementation of the SuperadminRepository interface
		type MockSuperadminRepo struct {
			ApprovedAdminRegisterFn func(id uint) error
		}

		// ApprovedAdminRegister is a mock implementation of the ApprovedAdminRegister method
		func (m *MockSuperadminRepo) ApprovedAdminRegister(id uint) error {
			return m.ApprovedAdminRegisterFn(id)
		}

		func TestUsecaseSuperadmin_ApprovedAdminRegister(t *testing.T) {
			// Create a mock repository
			mockRepo := &MockSuperadminRepo{}

			// Create an instance of the UsecaseSuperadmin with the mock repository
			uc := UsecaseSuperadmin{
				superadminRepo: mockRepo,
			}

			// Test case 1: Successful registration
			mockRepo.ApprovedAdminRegisterFn = func(id uint) error {
				// Simulate a successful registration
				return nil
			}

			err := uc.ApprovedAdminRegister(1)

			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}

			// Test case 2: Error during registration
			mockRepo.ApprovedAdminRegisterFn = func(id uint) error {
				// Simulate an error during registration
				return errors.New("registration error")
			}

			err = uc.ApprovedAdminRegister(1)

			if err == nil {
				t.Error("Expected an error, got nil")
			} else if err.Error() != "registration error" {
				t.Errorf("Expected 'registration error', got '%v'", err.Error())
			}
		}

		func TestRejectedAdminRegister(t *testing.T) {
			// Create a mock superadmin repository
			mockRepo := &mockSuperadminRepository{}

			// Create an instance of the UsecaseSuperadmin
			uc := UsecaseSuperadmin{
				superadminRepo: mockRepo,
			}

			// Set up the test data
			adminID := uint(123)

			// Define the expected error
			expectedError := errors.New("failed to register rejected admin")

			// Set up the mock behavior of the superadmin repository
			mockRepo.On("RejectedAdminRegister", adminID).Return(expectedError)

			// Call the RejectedAdminRegister function
			err := uc.RejectedAdminRegister(adminID)

			// Assert that the returned error matches the expected error
			if err != expectedError {
				t.Errorf("Expected error: %v, but got: %v", expectedError, err)
			}

			// Assert that the RejectedAdminRegister method of the mock repository was called with the correct arguments
			mockRepo.AssertCalled(t, "RejectedAdminRegister", adminID)
		}

		// Define a mock implementation of the superadmin repository
		type mockSuperadminRepository struct {
			mock.Mock
		}

		// Implement the RejectedAdminRegister method of the mock repository
		func (m *mockSuperadminRepository) RejectedAdminRegister(id uint) error {
			args := m.Called(id)
			return args.Error(0)
		}

		type MockSuperadminRepo struct {
			UpdateActivedAdminFunc func(id uint) error
		}

		func (m *MockSuperadminRepo) UpdateActivedAdmin(id uint) error {
			return m.UpdateActivedAdminFunc(id)
		}

		func TestUpdateActivedAdmin_Success(t *testing.T) {
			mockRepo := &MockSuperadminRepo{
				UpdateActivedAdminFunc: func(id uint) error {
					// Simulate a successful update
					return nil
				},
			}

			uc := UsecaseSuperadmin{
				superadminRepo: mockRepo,
			}

			err := uc.UpdateActivedAdmin(1)
			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		}

		func TestUpdateActivedAdmin_Failure(t *testing.T) {
			mockRepo := &MockSuperadminRepo{
				UpdateActivedAdminFunc: func(id uint) error {
					// Simulate an error during update
					return errors.New("failed to update")
				},
			}

			uc := UsecaseSuperadmin{
				superadminRepo: mockRepo,
			}

			err := uc.UpdateActivedAdmin(1)
			if err == nil {
				t.Error("Expected an error, but got nil")
			}
		}


		type mockSuperadminRepo struct {
			updateCalled bool
			updateError  error
		}

		func (m *mockSuperadminRepo) UpdateDeadactivedAdmin(id uint) error {
			m.updateCalled = true
			return m.updateError
		}

		func TestUpdateDeadactivedAdmin_Success(t *testing.T) {
			repo := &mockSuperadminRepo{}
			uc := UsecaseSuperadmin{superadminRepo: repo}
			id := uint(1)

			err := uc.UpdateDeadactivedAdmin(id)

			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}

			if !repo.updateCalled {
				t.Error("Expected UpdateDeadactivedAdmin to be called")
			}
		}

		func TestUpdateDeadactivedAdmin_Error(t *testing.T) {
			repo := &mockSuperadminRepo{updateError: errors.New("database error")}
			uc := UsecaseSuperadmin{superadminRepo: repo}
			id := uint(1)

			err := uc.UpdateDeadactivedAdmin(id)

			if err == nil {
				t.Error("Expected an error, got nil")
			}

			if !repo.updateCalled {
				t.Error("Expected UpdateDeadactivedAdmin to be called")
			}
		}

		// Create a mock repository implementing the superadminRepo interface
		type mockSuperadminRepo struct{}

		// Implement the GetApprovalRequest method for the mock repository
		func (m *mockSuperadminRepo) GetApprovalRequest() ([]*entities.Actor, error) {
			// Prepare and return sample data for testing
			actors := []*entities.Actor{
				&entities.Actor{ID: 1, Name: "Actor 1"},
				&entities.Actor{ID: 2, Name: "Actor 2"},
			}
			return actors, nil
		}

		// Define a test case function for GetApprovalRequest
		func TestGetApprovalRequest(t *testing.T) {
			// Create a UsecaseSuperadmin instance with the mock repository
			uc := UsecaseSuperadmin{
				superadminRepo: &mockSuperadminRepo{},
			}

			// Call the GetApprovalRequest method
			actors, err := uc.GetApprovalRequest()

			// Check the returned values
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			// Define the expected result
			expectedActors := []*entities.Actor{
				&entities.Actor{ID: 1, Name: "Actor 1"},
				&entities.Actor{ID: 2, Name: "Actor 2"},
			}

			// Compare the returned actors with the expected result
			for i := range actors {
				if actors[i].ID != expectedActors[i].ID || actors[i].Name != expectedActors[i].Name {
					t.Errorf("unexpected actor at index %d", i)
				}
			}
		}

		func TestGetAllAdmins_Success(t *testing.T) {
			// 1. Set up the necessary dependencies and mocks
			superadminRepoMock := &YourSuperadminRepoMock{}
			uc := UsecaseSuperadmin{superadminRepo: superadminRepoMock}

			// 2. Mock the GetAllAdmins method of the superadminRepo to return a sample list of admins
			expectedAdmins := []*entities.Actor{
				{ID: 1, Username: "admin1"},
				{ID: 2, Username: "admin2"},
				// Add more sample admins as needed
			}
			superadminRepoMock.On("GetAllAdmins", "username", 1, 10).Return(expectedAdmins, nil)

			// 3. Invoke the GetAllAdmins method
			admins, err := uc.GetAllAdmins("username", 1, 10)

			// 4. Assert the results
			assert.NoError(t, err)
			assert.Equal(t, expectedAdmins, admins)

			// Optionally, you can also assert specific properties of the returned admins
			// assert.Equal(t, 1, admins[0].ID)
			// assert.Equal(t, "admin1", admins[0].Username)

			// 5. Optionally, you can also assert that the GetAllAdmins method of the superadminRepo was called with the correct arguments
			superadminRepoMock.AssertCalled(t, "GetAllAdmins", "username", 1, 10)
		}











