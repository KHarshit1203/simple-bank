// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/KHarshit1203/simple-bank/service/db/gen"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, arg
func (_m *Store) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEntry provides a mock function with given fields: ctx, arg
func (_m *Store) CreateEntry(ctx context.Context, arg db.CreateEntryParams) (db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) (db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateEntryParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransfer provides a mock function with given fields: ctx, arg
func (_m *Store) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) (db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Store) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: ctx, arg
func (_m *Store) DeleteAccount(ctx context.Context, arg db.DeleteAccountParams) error {
	ret := _m.Called(ctx, arg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, db.DeleteAccountParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccount provides a mock function with given fields: ctx, id
func (_m *Store) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntry provides a mock function with given fields: ctx, id
func (_m *Store) GetEntry(ctx context.Context, id int64) (db.Entry, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Entry, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Entry); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransfer provides a mock function with given fields: ctx, id
func (_m *Store) GetTransfer(ctx context.Context, id int64) (db.Transfer, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, username
func (_m *Store) GetUser(ctx context.Context, username string) (db.User, error) {
	ret := _m.Called(ctx, username)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAccounts provides a mock function with given fields: ctx, arg
func (_m *Store) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) ([]db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) []db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListAccountsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListEntries provides a mock function with given fields: ctx, arg
func (_m *Store) ListEntries(ctx context.Context, arg db.ListEntriesParams) ([]db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListEntriesParams) ([]db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListEntriesParams) []db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListEntriesParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransfers provides a mock function with given fields: ctx, arg
func (_m *Store) ListTransfers(ctx context.Context, arg db.ListTransfersParams) ([]db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListTransfersParams) ([]db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListTransfersParams) []db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListTransfersParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurgeUserAccounts provides a mock function with given fields: ctx, owner
func (_m *Store) PurgeUserAccounts(ctx context.Context, owner string) error {
	ret := _m.Called(ctx, owner)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, owner)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferTx provides a mock function with given fields: ctx, arg
func (_m *Store) TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.TransferTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) (db.TransferTxResult, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) db.TransferTxResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.TransferTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.TransferTxParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: ctx, arg
func (_m *Store) UpdateBalance(ctx context.Context, arg db.UpdateBalanceParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateBalanceParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateBalanceParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateBalanceParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
