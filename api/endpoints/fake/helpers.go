package fake

import "github.com/stretchr/testify/mock"

func mockReturn[T any](args mock.Arguments) (*T, error) {
	if obj, ok := args.Get(0).(*T); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func mockReturnSlice[T any](args mock.Arguments) ([]*T, error) {
	if obj, ok := args.Get(0).([]*T); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
