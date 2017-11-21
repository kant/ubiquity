// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/local/scbe"
)

type FakeScbeDataModelWrapper struct {
	GetVolumeStub        func(name string, mustExist bool) (scbe.ScbeVolume, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		name      string
		mustExist bool
	}
	getVolumeReturns struct {
		result1 scbe.ScbeVolume
		result2 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 scbe.ScbeVolume
		result2 error
	}
	DeleteVolumeStub        func(name string) error
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		name string
	}
	deleteVolumeReturns struct {
		result1 error
	}
	deleteVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	InsertVolumeStub        func(volumeName string, wwn string, fstype string) error
	insertVolumeMutex       sync.RWMutex
	insertVolumeArgsForCall []struct {
		volumeName string
		wwn        string
		fstype     string
	}
	insertVolumeReturns struct {
		result1 error
	}
	insertVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	ListVolumesStub        func() ([]scbe.ScbeVolume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []scbe.ScbeVolume
		result2 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 []scbe.ScbeVolume
		result2 error
	}
	UpdateDatabaseVolumeStub        func(newVolume *scbe.ScbeVolume)
	updateDatabaseVolumeMutex       sync.RWMutex
	updateDatabaseVolumeArgsForCall []struct {
		newVolume *scbe.ScbeVolume
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScbeDataModelWrapper) GetVolume(name string, mustExist bool) (scbe.ScbeVolume, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		name      string
		mustExist bool
	}{name, mustExist})
	fake.recordInvocation("GetVolume", []interface{}{name, mustExist})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(name, mustExist)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2
}

func (fake *FakeScbeDataModelWrapper) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeScbeDataModelWrapper) GetVolumeArgsForCall(i int) (string, bool) {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].name, fake.getVolumeArgsForCall[i].mustExist
}

func (fake *FakeScbeDataModelWrapper) GetVolumeReturns(result1 scbe.ScbeVolume, result2 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 scbe.ScbeVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeScbeDataModelWrapper) GetVolumeReturnsOnCall(i int, result1 scbe.ScbeVolume, result2 error) {
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 scbe.ScbeVolume
			result2 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 scbe.ScbeVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeScbeDataModelWrapper) DeleteVolume(name string) error {
	fake.deleteVolumeMutex.Lock()
	ret, specificReturn := fake.deleteVolumeReturnsOnCall[len(fake.deleteVolumeArgsForCall)]
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteVolume", []interface{}{name})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteVolumeReturns.result1
}

func (fake *FakeScbeDataModelWrapper) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *FakeScbeDataModelWrapper) DeleteVolumeArgsForCall(i int) string {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return fake.deleteVolumeArgsForCall[i].name
}

func (fake *FakeScbeDataModelWrapper) DeleteVolumeReturns(result1 error) {
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModelWrapper) DeleteVolumeReturnsOnCall(i int, result1 error) {
	fake.DeleteVolumeStub = nil
	if fake.deleteVolumeReturnsOnCall == nil {
		fake.deleteVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModelWrapper) InsertVolume(volumeName string, wwn string, fstype string) error {
	fake.insertVolumeMutex.Lock()
	ret, specificReturn := fake.insertVolumeReturnsOnCall[len(fake.insertVolumeArgsForCall)]
	fake.insertVolumeArgsForCall = append(fake.insertVolumeArgsForCall, struct {
		volumeName string
		wwn        string
		fstype     string
	}{volumeName, wwn, fstype})
	fake.recordInvocation("InsertVolume", []interface{}{volumeName, wwn, fstype})
	fake.insertVolumeMutex.Unlock()
	if fake.InsertVolumeStub != nil {
		return fake.InsertVolumeStub(volumeName, wwn, fstype)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertVolumeReturns.result1
}

func (fake *FakeScbeDataModelWrapper) InsertVolumeCallCount() int {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return len(fake.insertVolumeArgsForCall)
}

func (fake *FakeScbeDataModelWrapper) InsertVolumeArgsForCall(i int) (string, string, string) {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return fake.insertVolumeArgsForCall[i].volumeName, fake.insertVolumeArgsForCall[i].wwn, fake.insertVolumeArgsForCall[i].fstype
}

func (fake *FakeScbeDataModelWrapper) InsertVolumeReturns(result1 error) {
	fake.InsertVolumeStub = nil
	fake.insertVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModelWrapper) InsertVolumeReturnsOnCall(i int, result1 error) {
	fake.InsertVolumeStub = nil
	if fake.insertVolumeReturnsOnCall == nil {
		fake.insertVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModelWrapper) ListVolumes() ([]scbe.ScbeVolume, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
}

func (fake *FakeScbeDataModelWrapper) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeScbeDataModelWrapper) ListVolumesReturns(result1 []scbe.ScbeVolume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []scbe.ScbeVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeScbeDataModelWrapper) ListVolumesReturnsOnCall(i int, result1 []scbe.ScbeVolume, result2 error) {
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 []scbe.ScbeVolume
			result2 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 []scbe.ScbeVolume
		result2 error
	}{result1, result2}
}


func (fake *FakeScbeDataModelWrapper) UpdateDatabaseVolume(newVolume *scbe.ScbeVolume) {
	fake.updateDatabaseVolumeMutex.Lock()
	fake.updateDatabaseVolumeArgsForCall = append(fake.updateDatabaseVolumeArgsForCall, struct {
		newVolume *scbe.ScbeVolume
	}{newVolume})
	fake.recordInvocation("UpdateDatabaseVolume", []interface{}{newVolume})
	fake.updateDatabaseVolumeMutex.Unlock()
	if fake.UpdateDatabaseVolumeStub != nil {
		fake.UpdateDatabaseVolumeStub(newVolume)
	}
}

func (fake *FakeScbeDataModelWrapper) UpdateDatabaseVolumeCallCount() int {
	fake.updateDatabaseVolumeMutex.RLock()
	defer fake.updateDatabaseVolumeMutex.RUnlock()
	return len(fake.updateDatabaseVolumeArgsForCall)
}

func (fake *FakeScbeDataModelWrapper) UpdateDatabaseVolumeArgsForCall(i int) *scbe.ScbeVolume {
	fake.updateDatabaseVolumeMutex.RLock()
	defer fake.updateDatabaseVolumeMutex.RUnlock()
	return fake.updateDatabaseVolumeArgsForCall[i].newVolume
}

func (fake *FakeScbeDataModelWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.updateDatabaseVolumeMutex.RLock()
	defer fake.updateDatabaseVolumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScbeDataModelWrapper) recordInvocation(key string, args []interface{}) {
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

var _ scbe.ScbeDataModelWrapper = new(FakeScbeDataModelWrapper)