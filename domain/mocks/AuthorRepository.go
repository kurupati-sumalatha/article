package mocks

import (
	"art/domain"
	"context"

	mock "github.com/stretchr/testify/mock"
)

type AuthorRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AuthorRepository) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Author
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Author); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Author)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
