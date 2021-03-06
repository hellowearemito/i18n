// Code generated by counterfeiter. DO NOT EDIT.
package i18n

import (
	"io"
	"sync"
)

type fakeService struct {
	AvailableLanguagesStub        func() []Language
	availableLanguagesMutex       sync.RWMutex
	availableLanguagesArgsForCall []struct{}
	availableLanguagesReturns     *struct {
		result1 []Language
	}
	availableLanguagesReturnsOnCall map[int]struct {
		result1 []Language
	}
	LoadMessagesStub        func(fileName string, r io.Reader) error
	loadMessagesMutex       sync.RWMutex
	loadMessagesArgsForCall []struct {
		fileName string
		r        io.Reader
	}
	loadMessagesReturns *struct {
		result1 error
	}
	loadMessagesReturnsOnCall map[int]struct {
		result1 error
	}
	LoadCategoryStub        func(category string) error
	loadCategoryMutex       sync.RWMutex
	loadCategoryArgsForCall []struct {
		category string
	}
	loadCategoryReturns *struct {
		result1 error
	}
	loadCategoryReturnsOnCall map[int]struct {
		result1 error
	}
	TfuncStub        func(language string, additionalLanguages ...string) (TranslateFunc, error)
	tfuncMutex       sync.RWMutex
	tfuncArgsForCall []struct {
		language            string
		additionalLanguages []string
	}
	tfuncReturns *struct {
		result1 TranslateFunc
		result2 error
	}
	tfuncReturnsOnCall map[int]struct {
		result1 TranslateFunc
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeService) AvailableLanguages() []Language {
	fake.availableLanguagesMutex.Lock()
	ret, specificReturn := fake.availableLanguagesReturnsOnCall[len(fake.availableLanguagesArgsForCall)]
	fake.availableLanguagesArgsForCall = append(fake.availableLanguagesArgsForCall, struct{}{})
	fake.recordInvocation("AvailableLanguages", []interface{}{})
	fake.availableLanguagesMutex.Unlock()
	if fake.AvailableLanguagesStub != nil {
		return fake.AvailableLanguagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	if fake.availableLanguagesReturns == nil {
		panic("Unexpected method call: Service.AvailableLanguages()")
	}
	return fake.availableLanguagesReturns.result1
}

func (fake *fakeService) AvailableLanguagesCallCount() int {
	fake.availableLanguagesMutex.RLock()
	defer fake.availableLanguagesMutex.RUnlock()
	return len(fake.availableLanguagesArgsForCall)
}

func (fake *fakeService) AvailableLanguagesReturns(result1 []Language) {
	fake.AvailableLanguagesStub = nil
	fake.availableLanguagesReturns = &struct {
		result1 []Language
	}{result1}
}

func (fake *fakeService) AvailableLanguagesReturnsOnCall(i int, result1 []Language) {
	fake.AvailableLanguagesStub = nil
	if fake.availableLanguagesReturnsOnCall == nil {
		fake.availableLanguagesReturnsOnCall = make(map[int]struct {
			result1 []Language
		})
	}
	fake.availableLanguagesReturnsOnCall[i] = struct {
		result1 []Language
	}{result1}
}

func (fake *fakeService) LoadMessages(fileName string, r io.Reader) error {
	fake.loadMessagesMutex.Lock()
	ret, specificReturn := fake.loadMessagesReturnsOnCall[len(fake.loadMessagesArgsForCall)]
	fake.loadMessagesArgsForCall = append(fake.loadMessagesArgsForCall, struct {
		fileName string
		r        io.Reader
	}{fileName, r})
	fake.recordInvocation("LoadMessages", []interface{}{fileName, r})
	fake.loadMessagesMutex.Unlock()
	if fake.LoadMessagesStub != nil {
		return fake.LoadMessagesStub(fileName, r)
	}
	if specificReturn {
		return ret.result1
	}
	if fake.loadMessagesReturns == nil {
		panic("Unexpected method call: Service.LoadMessages()")
	}
	return fake.loadMessagesReturns.result1
}

func (fake *fakeService) LoadMessagesCallCount() int {
	fake.loadMessagesMutex.RLock()
	defer fake.loadMessagesMutex.RUnlock()
	return len(fake.loadMessagesArgsForCall)
}

func (fake *fakeService) LoadMessagesArgsForCall(i int) (string, io.Reader) {
	fake.loadMessagesMutex.RLock()
	defer fake.loadMessagesMutex.RUnlock()
	return fake.loadMessagesArgsForCall[i].fileName, fake.loadMessagesArgsForCall[i].r
}

func (fake *fakeService) LoadMessagesReturns(result1 error) {
	fake.LoadMessagesStub = nil
	fake.loadMessagesReturns = &struct {
		result1 error
	}{result1}
}

func (fake *fakeService) LoadMessagesReturnsOnCall(i int, result1 error) {
	fake.LoadMessagesStub = nil
	if fake.loadMessagesReturnsOnCall == nil {
		fake.loadMessagesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadMessagesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakeService) LoadCategory(category string) error {
	fake.loadCategoryMutex.Lock()
	ret, specificReturn := fake.loadCategoryReturnsOnCall[len(fake.loadCategoryArgsForCall)]
	fake.loadCategoryArgsForCall = append(fake.loadCategoryArgsForCall, struct {
		category string
	}{category})
	fake.recordInvocation("LoadCategory", []interface{}{category})
	fake.loadCategoryMutex.Unlock()
	if fake.LoadCategoryStub != nil {
		return fake.LoadCategoryStub(category)
	}
	if specificReturn {
		return ret.result1
	}
	if fake.loadCategoryReturns == nil {
		panic("Unexpected method call: Service.LoadCategory()")
	}
	return fake.loadCategoryReturns.result1
}

func (fake *fakeService) LoadCategoryCallCount() int {
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	return len(fake.loadCategoryArgsForCall)
}

func (fake *fakeService) LoadCategoryArgsForCall(i int) string {
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	return fake.loadCategoryArgsForCall[i].category
}

func (fake *fakeService) LoadCategoryReturns(result1 error) {
	fake.LoadCategoryStub = nil
	fake.loadCategoryReturns = &struct {
		result1 error
	}{result1}
}

func (fake *fakeService) LoadCategoryReturnsOnCall(i int, result1 error) {
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

func (fake *fakeService) Tfunc(language string, additionalLanguages ...string) (TranslateFunc, error) {
	fake.tfuncMutex.Lock()
	ret, specificReturn := fake.tfuncReturnsOnCall[len(fake.tfuncArgsForCall)]
	fake.tfuncArgsForCall = append(fake.tfuncArgsForCall, struct {
		language            string
		additionalLanguages []string
	}{language, additionalLanguages})
	fake.recordInvocation("Tfunc", []interface{}{language, additionalLanguages})
	fake.tfuncMutex.Unlock()
	if fake.TfuncStub != nil {
		return fake.TfuncStub(language, additionalLanguages...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	if fake.tfuncReturns == nil {
		panic("Unexpected method call: Service.Tfunc()")
	}
	return fake.tfuncReturns.result1, fake.tfuncReturns.result2
}

func (fake *fakeService) TfuncCallCount() int {
	fake.tfuncMutex.RLock()
	defer fake.tfuncMutex.RUnlock()
	return len(fake.tfuncArgsForCall)
}

func (fake *fakeService) TfuncArgsForCall(i int) (string, []string) {
	fake.tfuncMutex.RLock()
	defer fake.tfuncMutex.RUnlock()
	return fake.tfuncArgsForCall[i].language, fake.tfuncArgsForCall[i].additionalLanguages
}

func (fake *fakeService) TfuncReturns(result1 TranslateFunc, result2 error) {
	fake.TfuncStub = nil
	fake.tfuncReturns = &struct {
		result1 TranslateFunc
		result2 error
	}{result1, result2}
}

func (fake *fakeService) TfuncReturnsOnCall(i int, result1 TranslateFunc, result2 error) {
	fake.TfuncStub = nil
	if fake.tfuncReturnsOnCall == nil {
		fake.tfuncReturnsOnCall = make(map[int]struct {
			result1 TranslateFunc
			result2 error
		})
	}
	fake.tfuncReturnsOnCall[i] = struct {
		result1 TranslateFunc
		result2 error
	}{result1, result2}
}

func (fake *fakeService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.availableLanguagesMutex.RLock()
	defer fake.availableLanguagesMutex.RUnlock()
	fake.loadMessagesMutex.RLock()
	defer fake.loadMessagesMutex.RUnlock()
	fake.loadCategoryMutex.RLock()
	defer fake.loadCategoryMutex.RUnlock()
	fake.tfuncMutex.RLock()
	defer fake.tfuncMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *fakeService) recordInvocation(key string, args []interface{}) {
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

var _ Service = new(fakeService)
