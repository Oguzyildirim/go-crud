// Code generated by counterfeiter. DO NOT EDIT.
package resttesting

import (
	"context"
	"sync"

	"github.com/Oguzyildirim/go-crud/internal"
	"github.com/Oguzyildirim/go-crud/internal/rest"
)

type FakeUserService struct {
	CreateStub        func(context.Context, string, string, string, string) (internal.User, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	createReturns struct {
		result1 internal.User
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 internal.User
		result2 error
	}
	DeleteStub        func(context.Context, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindStub        func(context.Context, string) (internal.User, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	findReturns struct {
		result1 internal.User
		result2 error
	}
	findReturnsOnCall map[int]struct {
		result1 internal.User
		result2 error
	}
	FindByCountryStub        func(context.Context, string) ([]internal.User, error)
	findByCountryMutex       sync.RWMutex
	findByCountryArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	findByCountryReturns struct {
		result1 []internal.User
		result2 error
	}
	findByCountryReturnsOnCall map[int]struct {
		result1 []internal.User
		result2 error
	}
	UpdateStub        func(context.Context, string, string, string, string, string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 string
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserService) Create(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string) (internal.User, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUserService) CreateCalls(stub func(context.Context, string, string, string, string) (internal.User, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeUserService) CreateArgsForCall(i int) (context.Context, string, string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeUserService) CreateReturns(result1 internal.User, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) CreateReturnsOnCall(i int, result1 internal.User, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 internal.User
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Delete(arg1 context.Context, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserService) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeUserService) DeleteCalls(stub func(context.Context, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeUserService) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) Find(arg1 context.Context, arg2 string) (internal.User, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.FindStub
	fakeReturns := fake.findReturns
	fake.recordInvocation("Find", []interface{}{arg1, arg2})
	fake.findMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeUserService) FindCalls(stub func(context.Context, string) (internal.User, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeUserService) FindArgsForCall(i int) (context.Context, string) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) FindReturns(result1 internal.User, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindReturnsOnCall(i int, result1 internal.User, result2 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 internal.User
			result2 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindByCountry(arg1 context.Context, arg2 string) ([]internal.User, error) {
	fake.findByCountryMutex.Lock()
	ret, specificReturn := fake.findByCountryReturnsOnCall[len(fake.findByCountryArgsForCall)]
	fake.findByCountryArgsForCall = append(fake.findByCountryArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.FindByCountryStub
	fakeReturns := fake.findByCountryReturns
	fake.recordInvocation("FindByCountry", []interface{}{arg1, arg2})
	fake.findByCountryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) FindByCountryCallCount() int {
	fake.findByCountryMutex.RLock()
	defer fake.findByCountryMutex.RUnlock()
	return len(fake.findByCountryArgsForCall)
}

func (fake *FakeUserService) FindByCountryCalls(stub func(context.Context, string) ([]internal.User, error)) {
	fake.findByCountryMutex.Lock()
	defer fake.findByCountryMutex.Unlock()
	fake.FindByCountryStub = stub
}

func (fake *FakeUserService) FindByCountryArgsForCall(i int) (context.Context, string) {
	fake.findByCountryMutex.RLock()
	defer fake.findByCountryMutex.RUnlock()
	argsForCall := fake.findByCountryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) FindByCountryReturns(result1 []internal.User, result2 error) {
	fake.findByCountryMutex.Lock()
	defer fake.findByCountryMutex.Unlock()
	fake.FindByCountryStub = nil
	fake.findByCountryReturns = struct {
		result1 []internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) FindByCountryReturnsOnCall(i int, result1 []internal.User, result2 error) {
	fake.findByCountryMutex.Lock()
	defer fake.findByCountryMutex.Unlock()
	fake.FindByCountryStub = nil
	if fake.findByCountryReturnsOnCall == nil {
		fake.findByCountryReturnsOnCall = make(map[int]struct {
			result1 []internal.User
			result2 error
		})
	}
	fake.findByCountryReturnsOnCall[i] = struct {
		result1 []internal.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Update(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserService) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeUserService) UpdateCalls(stub func(context.Context, string, string, string, string, string) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeUserService) UpdateArgsForCall(i int) (context.Context, string, string, string, string, string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeUserService) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.findByCountryMutex.RLock()
	defer fake.findByCountryMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rest.UserService = new(FakeUserService)
