// Code generated by counterfeiter. DO NOT EDIT.
package i18n

import (
	"sync"
)

type fakeLoader struct {
	LoadCategoryStub        func(category string, ml messageLoadable) error
	loadCategoryMutex       sync.RWMutex
	loadCategoryArgsForCall []struct {
		category string
		ml       messageLoadable
	}
	loadCategoryReturns *struct {
		result1 error
	}
	loadCategoryReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeLoader) LoadCategory(category string, ml messageLoadable) error {
	fake.loadCategoryMutex.Lock()
	ret, specificReturn := fake.loadCategoryReturnsOnCall[len(fake.loadCategoryArgsForCall)]
	fake.loadCategoryArgsForCall = append(fake.loadCategoryArgsForCall, struct {
		category string
		ml       messageLoadable
	}{category, ml})
	fake.recordInvocation("LoadCategory", []interface{}{category, ml})
	fake.loadCategoryMutex.Unlock()
	if fake.LoadCategoryStub != nil {
		return fake.LoadCategoryStub(category, ml)
	}
	if specificReturn {
		return ret.result1
	}
	if fake.loadCategoryReturns == nil {
		panic("Unexpected method call: MessageLoader.LoadCategory()")
	}
	return fake.loadCategoryReturns.result1
}

func (fake *fakeLoader) LoadCategoryCallCount() int {
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	return len(fake.loadCategoryArgsForCall)
}

func (fake *fakeLoader) LoadCategoryArgsForCall(i int) (string, messageLoadable) {
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	return fake.loadCategoryArgsForCall[i].category, fake.loadCategoryArgsForCall[i].ml
}

func (fake *fakeLoader) LoadCategoryReturns(result1 error) {
	fake.LoadCategoryStub = nil
	fake.loadCategoryReturns = &struct {
		result1 error
	}{result1}
}

func (fake *fakeLoader) LoadCategoryReturnsOnCall(i int, result1 error) {
	fake.LoadCategoryStub = nil
	if fake.loadCategoryReturnsOnCall == nil {
		fake.loadCategoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadCategoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakeLoader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeLoader) recordInvocation(key string, args []interface{}) {
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

var _ MessageLoader = new(fakeLoader)
