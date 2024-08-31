package service

import "testing"

type mocks struct { //nolint:unused
	accountRepository     *MockAccountRepository
	transactionRepository *MockTransactionRepository
	generalRepository     *MockGeneralRepository
	userRepository        *MockUserRepository
	permissionsService    *MockAccountPermissionsService
	accountService        *MockAccountService
}

func newMocks(t *testing.T) mocks { //nolint:unused
	return mocks{
		accountRepository:     NewMockAccountRepository(t),
		transactionRepository: NewMockTransactionRepository(t),
		generalRepository:     NewMockGeneralRepository(t),
		userRepository:        NewMockUserRepository(t),
		permissionsService:    NewMockAccountPermissionsService(t),
		accountService:        NewMockAccountService(t),
	}
}
