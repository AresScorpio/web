// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/web/getresource"
)

type FakeResourcesDB struct {
	GetPipelineNameStub        func() string
	getPipelineNameMutex       sync.RWMutex
	getPipelineNameArgsForCall []struct{}
	getPipelineNameReturns     struct {
		result1 string
	}
	GetConfigStub        func() (atc.Config, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct{}
	getConfigReturns     struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}
	GetResourceStub        func(string) (db.SavedResource, error)
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		arg1 string
	}
	getResourceReturns struct {
		result1 db.SavedResource
		result2 error
	}
	GetResourceHistoryCursorStub        func(string, int, bool, int) ([]*db.VersionHistory, bool, error)
	getResourceHistoryCursorMutex       sync.RWMutex
	getResourceHistoryCursorArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 bool
		arg4 int
	}
	getResourceHistoryCursorReturns struct {
		result1 []*db.VersionHistory
		result2 bool
		result3 error
	}
	GetResourceHistoryMaxIDStub        func(int) (int, error)
	getResourceHistoryMaxIDMutex       sync.RWMutex
	getResourceHistoryMaxIDArgsForCall []struct {
		arg1 int
	}
	getResourceHistoryMaxIDReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeResourcesDB) GetPipelineName() string {
	fake.getPipelineNameMutex.Lock()
	fake.getPipelineNameArgsForCall = append(fake.getPipelineNameArgsForCall, struct{}{})
	fake.getPipelineNameMutex.Unlock()
	if fake.GetPipelineNameStub != nil {
		return fake.GetPipelineNameStub()
	} else {
		return fake.getPipelineNameReturns.result1
	}
}

func (fake *FakeResourcesDB) GetPipelineNameCallCount() int {
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	return len(fake.getPipelineNameArgsForCall)
}

func (fake *FakeResourcesDB) GetPipelineNameReturns(result1 string) {
	fake.GetPipelineNameStub = nil
	fake.getPipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResourcesDB) GetConfig() (atc.Config, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct{}{})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub()
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3
	}
}

func (fake *FakeResourcesDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeResourcesDB) GetConfigReturns(result1 atc.Config, result2 db.ConfigVersion, result3 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourcesDB) GetResource(arg1 string) (db.SavedResource, error) {
	fake.getResourceMutex.Lock()
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getResourceMutex.Unlock()
	if fake.GetResourceStub != nil {
		return fake.GetResourceStub(arg1)
	} else {
		return fake.getResourceReturns.result1, fake.getResourceReturns.result2
	}
}

func (fake *FakeResourcesDB) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeResourcesDB) GetResourceArgsForCall(i int) string {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return fake.getResourceArgsForCall[i].arg1
}

func (fake *FakeResourcesDB) GetResourceReturns(result1 db.SavedResource, result2 error) {
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 db.SavedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourcesDB) GetResourceHistoryCursor(arg1 string, arg2 int, arg3 bool, arg4 int) ([]*db.VersionHistory, bool, error) {
	fake.getResourceHistoryCursorMutex.Lock()
	fake.getResourceHistoryCursorArgsForCall = append(fake.getResourceHistoryCursorArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 bool
		arg4 int
	}{arg1, arg2, arg3, arg4})
	fake.getResourceHistoryCursorMutex.Unlock()
	if fake.GetResourceHistoryCursorStub != nil {
		return fake.GetResourceHistoryCursorStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.getResourceHistoryCursorReturns.result1, fake.getResourceHistoryCursorReturns.result2, fake.getResourceHistoryCursorReturns.result3
	}
}

func (fake *FakeResourcesDB) GetResourceHistoryCursorCallCount() int {
	fake.getResourceHistoryCursorMutex.RLock()
	defer fake.getResourceHistoryCursorMutex.RUnlock()
	return len(fake.getResourceHistoryCursorArgsForCall)
}

func (fake *FakeResourcesDB) GetResourceHistoryCursorArgsForCall(i int) (string, int, bool, int) {
	fake.getResourceHistoryCursorMutex.RLock()
	defer fake.getResourceHistoryCursorMutex.RUnlock()
	return fake.getResourceHistoryCursorArgsForCall[i].arg1, fake.getResourceHistoryCursorArgsForCall[i].arg2, fake.getResourceHistoryCursorArgsForCall[i].arg3, fake.getResourceHistoryCursorArgsForCall[i].arg4
}

func (fake *FakeResourcesDB) GetResourceHistoryCursorReturns(result1 []*db.VersionHistory, result2 bool, result3 error) {
	fake.GetResourceHistoryCursorStub = nil
	fake.getResourceHistoryCursorReturns = struct {
		result1 []*db.VersionHistory
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourcesDB) GetResourceHistoryMaxID(arg1 int) (int, error) {
	fake.getResourceHistoryMaxIDMutex.Lock()
	fake.getResourceHistoryMaxIDArgsForCall = append(fake.getResourceHistoryMaxIDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.getResourceHistoryMaxIDMutex.Unlock()
	if fake.GetResourceHistoryMaxIDStub != nil {
		return fake.GetResourceHistoryMaxIDStub(arg1)
	} else {
		return fake.getResourceHistoryMaxIDReturns.result1, fake.getResourceHistoryMaxIDReturns.result2
	}
}

func (fake *FakeResourcesDB) GetResourceHistoryMaxIDCallCount() int {
	fake.getResourceHistoryMaxIDMutex.RLock()
	defer fake.getResourceHistoryMaxIDMutex.RUnlock()
	return len(fake.getResourceHistoryMaxIDArgsForCall)
}

func (fake *FakeResourcesDB) GetResourceHistoryMaxIDArgsForCall(i int) int {
	fake.getResourceHistoryMaxIDMutex.RLock()
	defer fake.getResourceHistoryMaxIDMutex.RUnlock()
	return fake.getResourceHistoryMaxIDArgsForCall[i].arg1
}

func (fake *FakeResourcesDB) GetResourceHistoryMaxIDReturns(result1 int, result2 error) {
	fake.GetResourceHistoryMaxIDStub = nil
	fake.getResourceHistoryMaxIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ getresource.ResourcesDB = new(FakeResourcesDB)
