// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/worker"
)

type FakeWorker struct {
	CreateContainerStub        func(worker.Identifier, worker.ContainerSpec) (worker.Container, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 worker.Identifier
		arg2 worker.ContainerSpec
	}
	createContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindContainerForIdentifierStub        func(worker.Identifier) (worker.Container, error)
	findContainerForIdentifierMutex       sync.RWMutex
	findContainerForIdentifierArgsForCall []struct {
		arg1 worker.Identifier
	}
	findContainerForIdentifierReturns struct {
		result1 worker.Container
		result2 error
	}
	FindContainersForIdentifierStub        func(worker.Identifier) ([]worker.Container, error)
	findContainersForIdentifierMutex       sync.RWMutex
	findContainersForIdentifierArgsForCall []struct {
		arg1 worker.Identifier
	}
	findContainersForIdentifierReturns struct {
		result1 []worker.Container
		result2 error
	}
	LookupContainerStub        func(string) (worker.Container, error)
	lookupContainerMutex       sync.RWMutex
	lookupContainerArgsForCall []struct {
		arg1 string
	}
	lookupContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns struct {
		result1 string
	}
	ActiveContainersStub        func() int
	activeContainersMutex       sync.RWMutex
	activeContainersArgsForCall []struct{}
	activeContainersReturns struct {
		result1 int
	}
	SatisfiesStub        func(worker.ContainerSpec) bool
	satisfiesMutex       sync.RWMutex
	satisfiesArgsForCall []struct {
		arg1 worker.ContainerSpec
	}
	satisfiesReturns struct {
		result1 bool
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns struct {
		result1 string
	}
}

func (fake *FakeWorker) CreateContainer(arg1 worker.Identifier, arg2 worker.ContainerSpec) (worker.Container, error) {
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 worker.Identifier
		arg2 worker.ContainerSpec
	}{arg1, arg2})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(arg1, arg2)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeWorker) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeWorker) CreateContainerArgsForCall(i int) (worker.Identifier, worker.ContainerSpec) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].arg1, fake.createContainerArgsForCall[i].arg2
}

func (fake *FakeWorker) CreateContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindContainerForIdentifier(arg1 worker.Identifier) (worker.Container, error) {
	fake.findContainerForIdentifierMutex.Lock()
	fake.findContainerForIdentifierArgsForCall = append(fake.findContainerForIdentifierArgsForCall, struct {
		arg1 worker.Identifier
	}{arg1})
	fake.findContainerForIdentifierMutex.Unlock()
	if fake.FindContainerForIdentifierStub != nil {
		return fake.FindContainerForIdentifierStub(arg1)
	} else {
		return fake.findContainerForIdentifierReturns.result1, fake.findContainerForIdentifierReturns.result2
	}
}

func (fake *FakeWorker) FindContainerForIdentifierCallCount() int {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return len(fake.findContainerForIdentifierArgsForCall)
}

func (fake *FakeWorker) FindContainerForIdentifierArgsForCall(i int) worker.Identifier {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return fake.findContainerForIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorker) FindContainerForIdentifierReturns(result1 worker.Container, result2 error) {
	fake.FindContainerForIdentifierStub = nil
	fake.findContainerForIdentifierReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindContainersForIdentifier(arg1 worker.Identifier) ([]worker.Container, error) {
	fake.findContainersForIdentifierMutex.Lock()
	fake.findContainersForIdentifierArgsForCall = append(fake.findContainersForIdentifierArgsForCall, struct {
		arg1 worker.Identifier
	}{arg1})
	fake.findContainersForIdentifierMutex.Unlock()
	if fake.FindContainersForIdentifierStub != nil {
		return fake.FindContainersForIdentifierStub(arg1)
	} else {
		return fake.findContainersForIdentifierReturns.result1, fake.findContainersForIdentifierReturns.result2
	}
}

func (fake *FakeWorker) FindContainersForIdentifierCallCount() int {
	fake.findContainersForIdentifierMutex.RLock()
	defer fake.findContainersForIdentifierMutex.RUnlock()
	return len(fake.findContainersForIdentifierArgsForCall)
}

func (fake *FakeWorker) FindContainersForIdentifierArgsForCall(i int) worker.Identifier {
	fake.findContainersForIdentifierMutex.RLock()
	defer fake.findContainersForIdentifierMutex.RUnlock()
	return fake.findContainersForIdentifierArgsForCall[i].arg1
}

func (fake *FakeWorker) FindContainersForIdentifierReturns(result1 []worker.Container, result2 error) {
	fake.FindContainersForIdentifierStub = nil
	fake.findContainersForIdentifierReturns = struct {
		result1 []worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) LookupContainer(arg1 string) (worker.Container, error) {
	fake.lookupContainerMutex.Lock()
	fake.lookupContainerArgsForCall = append(fake.lookupContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.lookupContainerMutex.Unlock()
	if fake.LookupContainerStub != nil {
		return fake.LookupContainerStub(arg1)
	} else {
		return fake.lookupContainerReturns.result1, fake.lookupContainerReturns.result2
	}
}

func (fake *FakeWorker) LookupContainerCallCount() int {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return len(fake.lookupContainerArgsForCall)
}

func (fake *FakeWorker) LookupContainerArgsForCall(i int) string {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return fake.lookupContainerArgsForCall[i].arg1
}

func (fake *FakeWorker) LookupContainerReturns(result1 worker.Container, result2 error) {
	fake.LookupContainerStub = nil
	fake.lookupContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeWorker) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorker) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) ActiveContainers() int {
	fake.activeContainersMutex.Lock()
	fake.activeContainersArgsForCall = append(fake.activeContainersArgsForCall, struct{}{})
	fake.activeContainersMutex.Unlock()
	if fake.ActiveContainersStub != nil {
		return fake.ActiveContainersStub()
	} else {
		return fake.activeContainersReturns.result1
	}
}

func (fake *FakeWorker) ActiveContainersCallCount() int {
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	return len(fake.activeContainersArgsForCall)
}

func (fake *FakeWorker) ActiveContainersReturns(result1 int) {
	fake.ActiveContainersStub = nil
	fake.activeContainersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) Satisfies(arg1 worker.ContainerSpec) bool {
	fake.satisfiesMutex.Lock()
	fake.satisfiesArgsForCall = append(fake.satisfiesArgsForCall, struct {
		arg1 worker.ContainerSpec
	}{arg1})
	fake.satisfiesMutex.Unlock()
	if fake.SatisfiesStub != nil {
		return fake.SatisfiesStub(arg1)
	} else {
		return fake.satisfiesReturns.result1
	}
}

func (fake *FakeWorker) SatisfiesCallCount() int {
	fake.satisfiesMutex.RLock()
	defer fake.satisfiesMutex.RUnlock()
	return len(fake.satisfiesArgsForCall)
}

func (fake *FakeWorker) SatisfiesArgsForCall(i int) worker.ContainerSpec {
	fake.satisfiesMutex.RLock()
	defer fake.satisfiesMutex.RUnlock()
	return fake.satisfiesArgsForCall[i].arg1
}

func (fake *FakeWorker) SatisfiesReturns(result1 bool) {
	fake.SatisfiesStub = nil
	fake.satisfiesReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) Description() string {
	fake.descriptionMutex.Lock()
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	} else {
		return fake.descriptionReturns.result1
	}
}

func (fake *FakeWorker) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeWorker) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

var _ worker.Worker = new(FakeWorker)
