// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.ibm.com/almaden-containers/ubiquity/model"
)

type FakeStorageClient struct {
	ActivateStub        func() error
	activateMutex       sync.RWMutex
	activateArgsForCall []struct{}
	activateReturns     struct {
		result1 error
	}
	CreateVolumeStub        func(name string, opts map[string]interface{}) error
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		name string
		opts map[string]interface{}
	}
	createVolumeReturns struct {
		result1 error
	}
	RemoveVolumeStub        func(name string, forceDelete bool) error
	removeVolumeMutex       sync.RWMutex
	removeVolumeArgsForCall []struct {
		name        string
		forceDelete bool
	}
	removeVolumeReturns struct {
		result1 error
	}
	ListVolumesStub        func() ([]model.VolumeMetadata, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []model.VolumeMetadata
		result2 error
	}
	GetVolumeStub        func(name string) (volumeMetadata model.VolumeMetadata, volumeConfigDetails model.SpectrumConfig, err error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		name string
	}
	getVolumeReturns struct {
		result1 model.VolumeMetadata
		result2 model.SpectrumConfig
		result3 error
	}
	AttachStub        func(name string) (string, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		name string
	}
	attachReturns struct {
		result1 string
		result2 error
	}
	DetachStub        func(name string) error
	detachMutex       sync.RWMutex
	detachArgsForCall []struct {
		name string
	}
	detachReturns struct {
		result1 error
	}
	GetPluginNameStub        func() string
	getPluginNameMutex       sync.RWMutex
	getPluginNameArgsForCall []struct{}
	getPluginNameReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorageClient) Activate() error {
	fake.activateMutex.Lock()
	fake.activateArgsForCall = append(fake.activateArgsForCall, struct{}{})
	fake.recordInvocation("Activate", []interface{}{})
	fake.activateMutex.Unlock()
	if fake.ActivateStub != nil {
		return fake.ActivateStub()
	} else {
		return fake.activateReturns.result1
	}
}

func (fake *FakeStorageClient) ActivateCallCount() int {
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	return len(fake.activateArgsForCall)
}

func (fake *FakeStorageClient) ActivateReturns(result1 error) {
	fake.ActivateStub = nil
	fake.activateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) CreateVolume(name string, opts map[string]interface{}) error {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		name string
		opts map[string]interface{}
	}{name, opts})
	fake.recordInvocation("CreateVolume", []interface{}{name, opts})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(name, opts)
	} else {
		return fake.createVolumeReturns.result1
	}
}

func (fake *FakeStorageClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeStorageClient) CreateVolumeArgsForCall(i int) (string, map[string]interface{}) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].name, fake.createVolumeArgsForCall[i].opts
}

func (fake *FakeStorageClient) CreateVolumeReturns(result1 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) RemoveVolume(name string, forceDelete bool) error {
	fake.removeVolumeMutex.Lock()
	fake.removeVolumeArgsForCall = append(fake.removeVolumeArgsForCall, struct {
		name        string
		forceDelete bool
	}{name, forceDelete})
	fake.recordInvocation("RemoveVolume", []interface{}{name, forceDelete})
	fake.removeVolumeMutex.Unlock()
	if fake.RemoveVolumeStub != nil {
		return fake.RemoveVolumeStub(name, forceDelete)
	} else {
		return fake.removeVolumeReturns.result1
	}
}

func (fake *FakeStorageClient) RemoveVolumeCallCount() int {
	fake.removeVolumeMutex.RLock()
	defer fake.removeVolumeMutex.RUnlock()
	return len(fake.removeVolumeArgsForCall)
}

func (fake *FakeStorageClient) RemoveVolumeArgsForCall(i int) (string, bool) {
	fake.removeVolumeMutex.RLock()
	defer fake.removeVolumeMutex.RUnlock()
	return fake.removeVolumeArgsForCall[i].name, fake.removeVolumeArgsForCall[i].forceDelete
}

func (fake *FakeStorageClient) RemoveVolumeReturns(result1 error) {
	fake.RemoveVolumeStub = nil
	fake.removeVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) ListVolumes() ([]model.VolumeMetadata, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeStorageClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeStorageClient) ListVolumesReturns(result1 []model.VolumeMetadata, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []model.VolumeMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageClient) GetVolume(name string) (volumeMetadata model.VolumeMetadata, volumeConfigDetails model.SpectrumConfig, err error) {
	fake.getVolumeMutex.Lock()
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVolume", []interface{}{name})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(name)
	} else {
		return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
	}
}

func (fake *FakeStorageClient) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeStorageClient) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].name
}

func (fake *FakeStorageClient) GetVolumeReturns(result1 model.VolumeMetadata, result2 model.SpectrumConfig, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 model.VolumeMetadata
		result2 model.SpectrumConfig
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStorageClient) Attach(name string) (string, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Attach", []interface{}{name})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(name)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2
	}
}

func (fake *FakeStorageClient) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeStorageClient) AttachArgsForCall(i int) string {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].name
}

func (fake *FakeStorageClient) AttachReturns(result1 string, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorageClient) Detach(name string) error {
	fake.detachMutex.Lock()
	fake.detachArgsForCall = append(fake.detachArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Detach", []interface{}{name})
	fake.detachMutex.Unlock()
	if fake.DetachStub != nil {
		return fake.DetachStub(name)
	} else {
		return fake.detachReturns.result1
	}
}

func (fake *FakeStorageClient) DetachCallCount() int {
	fake.detachMutex.RLock()
	defer fake.detachMutex.RUnlock()
	return len(fake.detachArgsForCall)
}

func (fake *FakeStorageClient) DetachArgsForCall(i int) string {
	fake.detachMutex.RLock()
	defer fake.detachMutex.RUnlock()
	return fake.detachArgsForCall[i].name
}

func (fake *FakeStorageClient) DetachReturns(result1 error) {
	fake.DetachStub = nil
	fake.detachReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorageClient) GetPluginName() string {
	fake.getPluginNameMutex.Lock()
	fake.getPluginNameArgsForCall = append(fake.getPluginNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPluginName", []interface{}{})
	fake.getPluginNameMutex.Unlock()
	if fake.GetPluginNameStub != nil {
		return fake.GetPluginNameStub()
	} else {
		return fake.getPluginNameReturns.result1
	}
}

func (fake *FakeStorageClient) GetPluginNameCallCount() int {
	fake.getPluginNameMutex.RLock()
	defer fake.getPluginNameMutex.RUnlock()
	return len(fake.getPluginNameArgsForCall)
}

func (fake *FakeStorageClient) GetPluginNameReturns(result1 string) {
	fake.GetPluginNameStub = nil
	fake.getPluginNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStorageClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.activateMutex.RLock()
	defer fake.activateMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.removeVolumeMutex.RLock()
	defer fake.removeVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.detachMutex.RLock()
	defer fake.detachMutex.RUnlock()
	fake.getPluginNameMutex.RLock()
	defer fake.getPluginNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStorageClient) recordInvocation(key string, args []interface{}) {
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

var _ model.StorageClient = new(FakeStorageClient)
