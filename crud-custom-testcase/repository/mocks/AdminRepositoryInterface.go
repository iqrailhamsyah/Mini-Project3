// Code generated by mockery v2.29.0. DO NOT EDIT.

package mocks

import (
	entities "crud/entities"

	mock "github.com/stretchr/testify/mock"
)

// AdminRepositoryInterface is an autogenerated mock type for the AdminRepositoryInterface type
type AdminRepositoryInterface struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: customer
func (_m *AdminRepositoryInterface) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	ret := _m.Called(customer)

	var r0 *entities.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.Customer) (*entities.Customer, error)); ok {
		return rf(customer)
	}
	if rf, ok := ret.Get(0).(func(*entities.Customer) *entities.Customer); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.Customer) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAdminById provides a mock function with given fields: id, admin
func (_m *AdminRepositoryInterface) DeleteAdminById(id uint, admin *entities.Actor) error {
	ret := _m.Called(id, admin)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *entities.Actor) error); ok {
		r0 = rf(id, admin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCustomerById provides a mock function with given fields: id, customer
func (_m *AdminRepositoryInterface) DeleteCustomerById(id uint, customer *entities.Customer) error {
	ret := _m.Called(id, customer)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, *entities.Customer) error); ok {
		r0 = rf(id, customer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchCustomersFromAPI provides a mock function with given fields:
func (_m *AdminRepositoryInterface) FetchCustomersFromAPI() ([]*entities.Customer, error) {
	ret := _m.Called()

	var r0 []*entities.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.Customer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdminById provides a mock function with given fields: id
func (_m *AdminRepositoryInterface) GetAdminById(id uint) (*entities.Actor, error) {
	ret := _m.Called(id)

	var r0 *entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entities.Actor, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entities.Actor); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCustomers provides a mock function with given fields: first_name, last_name, email, page, pageSize
func (_m *AdminRepositoryInterface) GetAllCustomers(first_name string, last_name string, email string, page int, pageSize int) ([]*entities.Customer, error) {
	ret := _m.Called(first_name, last_name, email, page, pageSize)

	var r0 []*entities.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) ([]*entities.Customer, error)); ok {
		return rf(first_name, last_name, email, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) []*entities.Customer); ok {
		r0 = rf(first_name, last_name, email, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int, int) error); ok {
		r1 = rf(first_name, last_name, email, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByEmail provides a mock function with given fields: email
func (_m *AdminRepositoryInterface) GetCustomerByEmail(email string) (*entities.Customer, error) {
	ret := _m.Called(email)

	var r0 *entities.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Customer, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Customer); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerById provides a mock function with given fields: id
func (_m *AdminRepositoryInterface) GetCustomerById(id uint) (*entities.Customer, error) {
	ret := _m.Called(id)

	var r0 *entities.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entities.Customer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entities.Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginAdmin provides a mock function with given fields: username
func (_m *AdminRepositoryInterface) LoginAdmin(username string) (*entities.Actor, error) {
	ret := _m.Called(username)

	var r0 *entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.Actor); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterAdmin provides a mock function with given fields: admin
func (_m *AdminRepositoryInterface) RegisterAdmin(admin *entities.Actor) (*entities.Actor, error) {
	ret := _m.Called(admin)

	var r0 *entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.Actor) (*entities.Actor, error)); ok {
		return rf(admin)
	}
	if rf, ok := ret.Get(0).(func(*entities.Actor) *entities.Actor); ok {
		r0 = rf(admin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(*entities.Actor) error); ok {
		r1 = rf(admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCustomersFromAPI provides a mock function with given fields: url
func (_m *AdminRepositoryInterface) SaveCustomersFromAPI(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdminById provides a mock function with given fields: id, admin
func (_m *AdminRepositoryInterface) UpdateAdminById(id uint, admin *entities.Actor) (*entities.Actor, error) {
	ret := _m.Called(id, admin)

	var r0 *entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, *entities.Actor) (*entities.Actor, error)); ok {
		return rf(id, admin)
	}
	if rf, ok := ret.Get(0).(func(uint, *entities.Actor) *entities.Actor); ok {
		r0 = rf(id, admin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, *entities.Actor) error); ok {
		r1 = rf(id, admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdminRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdminRepositoryInterface creates a new instance of AdminRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdminRepositoryInterface(t mockConstructorTestingTNewAdminRepositoryInterface) *AdminRepositoryInterface {
	mock := &AdminRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
