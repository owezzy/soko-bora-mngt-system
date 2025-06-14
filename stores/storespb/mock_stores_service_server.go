// Code generated by mockery v2.53.4. DO NOT EDIT.

package storespb

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockStoresServiceServer is an autogenerated mock type for the StoresServiceServer type
type MockStoresServiceServer struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) AddProduct(_a0 context.Context, _a1 *AddProductRequest) (*AddProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddProduct")
	}

	var r0 *AddProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *AddProductRequest) (*AddProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *AddProductRequest) *AddProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AddProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *AddProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStore provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) CreateStore(_a0 context.Context, _a1 *CreateStoreRequest) (*CreateStoreResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateStore")
	}

	var r0 *CreateStoreResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateStoreRequest) *CreateStoreResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateStoreResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateStoreRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecreaseProductPrice provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) DecreaseProductPrice(_a0 context.Context, _a1 *DecreaseProductPriceRequest) (*DecreaseProductPriceResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DecreaseProductPrice")
	}

	var r0 *DecreaseProductPriceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *DecreaseProductPriceRequest) (*DecreaseProductPriceResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *DecreaseProductPriceRequest) *DecreaseProductPriceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DecreaseProductPriceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *DecreaseProductPriceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableParticipation provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) DisableParticipation(_a0 context.Context, _a1 *DisableParticipationRequest) (*DisableParticipationResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DisableParticipation")
	}

	var r0 *DisableParticipationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *DisableParticipationRequest) (*DisableParticipationResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *DisableParticipationRequest) *DisableParticipationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DisableParticipationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *DisableParticipationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableParticipation provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) EnableParticipation(_a0 context.Context, _a1 *EnableParticipationRequest) (*EnableParticipationResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for EnableParticipation")
	}

	var r0 *EnableParticipationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *EnableParticipationRequest) (*EnableParticipationResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *EnableParticipationRequest) *EnableParticipationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EnableParticipationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *EnableParticipationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatalog provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) GetCatalog(_a0 context.Context, _a1 *GetCatalogRequest) (*GetCatalogResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCatalog")
	}

	var r0 *GetCatalogResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetCatalogRequest) (*GetCatalogResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetCatalogRequest) *GetCatalogResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetCatalogResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetCatalogRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParticipatingStores provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) GetParticipatingStores(_a0 context.Context, _a1 *GetParticipatingStoresRequest) (*GetParticipatingStoresResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetParticipatingStores")
	}

	var r0 *GetParticipatingStoresResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetParticipatingStoresRequest) (*GetParticipatingStoresResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetParticipatingStoresRequest) *GetParticipatingStoresResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetParticipatingStoresResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetParticipatingStoresRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) GetProduct(_a0 context.Context, _a1 *GetProductRequest) (*GetProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProduct")
	}

	var r0 *GetProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductRequest) (*GetProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductRequest) *GetProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStore provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) GetStore(_a0 context.Context, _a1 *GetStoreRequest) (*GetStoreResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetStore")
	}

	var r0 *GetStoreResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetStoreRequest) (*GetStoreResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetStoreRequest) *GetStoreResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetStoreResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetStoreRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStores provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) GetStores(_a0 context.Context, _a1 *GetStoresRequest) (*GetStoresResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetStores")
	}

	var r0 *GetStoresResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetStoresRequest) (*GetStoresResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetStoresRequest) *GetStoresResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetStoresResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetStoresRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncreaseProductPrice provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) IncreaseProductPrice(_a0 context.Context, _a1 *IncreaseProductPriceRequest) (*IncreaseProductPriceResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IncreaseProductPrice")
	}

	var r0 *IncreaseProductPriceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *IncreaseProductPriceRequest) (*IncreaseProductPriceResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *IncreaseProductPriceRequest) *IncreaseProductPriceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*IncreaseProductPriceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *IncreaseProductPriceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RebrandProduct provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) RebrandProduct(_a0 context.Context, _a1 *RebrandProductRequest) (*RebrandProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RebrandProduct")
	}

	var r0 *RebrandProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *RebrandProductRequest) (*RebrandProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *RebrandProductRequest) *RebrandProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RebrandProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *RebrandProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RebrandStore provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) RebrandStore(_a0 context.Context, _a1 *RebrandStoreRequest) (*RebrandStoreResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RebrandStore")
	}

	var r0 *RebrandStoreResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *RebrandStoreRequest) (*RebrandStoreResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *RebrandStoreRequest) *RebrandStoreResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RebrandStoreResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *RebrandStoreRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveProduct provides a mock function with given fields: _a0, _a1
func (_m *MockStoresServiceServer) RemoveProduct(_a0 context.Context, _a1 *RemoveProductRequest) (*RemoveProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RemoveProduct")
	}

	var r0 *RemoveProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *RemoveProductRequest) (*RemoveProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *RemoveProductRequest) *RemoveProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RemoveProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *RemoveProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedStoresServiceServer provides a mock function with no fields
func (_m *MockStoresServiceServer) mustEmbedUnimplementedStoresServiceServer() {
	_m.Called()
}

// NewMockStoresServiceServer creates a new instance of MockStoresServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStoresServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStoresServiceServer {
	mock := &MockStoresServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
