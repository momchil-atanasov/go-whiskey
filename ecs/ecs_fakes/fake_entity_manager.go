// This file was generated by counterfeiter
package ecs_fakes

import (
	"sync"

	"github.com/momchil-atanasov/go-whiskey/ecs"
)

type FakeEntityManager struct {
	CreateEntityStub        func() ecs.Entity
	createEntityMutex       sync.RWMutex
	createEntityArgsForCall []struct{}
	createEntityReturns     struct {
		result1 ecs.Entity
	}
	HasEntityStub        func(ecs.Entity) bool
	hasEntityMutex       sync.RWMutex
	hasEntityArgsForCall []struct {
		arg1 ecs.Entity
	}
	hasEntityReturns struct {
		result1 bool
	}
	DeleteEntityStub        func(ecs.Entity)
	deleteEntityMutex       sync.RWMutex
	deleteEntityArgsForCall []struct {
		arg1 ecs.Entity
	}
	DeleteAllEntitiesStub        func()
	deleteAllEntitiesMutex       sync.RWMutex
	deleteAllEntitiesArgsForCall []struct{}
}

func (fake *FakeEntityManager) CreateEntity() ecs.Entity {
	fake.createEntityMutex.Lock()
	fake.createEntityArgsForCall = append(fake.createEntityArgsForCall, struct{}{})
	fake.createEntityMutex.Unlock()
	if fake.CreateEntityStub != nil {
		return fake.CreateEntityStub()
	} else {
		return fake.createEntityReturns.result1
	}
}

func (fake *FakeEntityManager) CreateEntityCallCount() int {
	fake.createEntityMutex.RLock()
	defer fake.createEntityMutex.RUnlock()
	return len(fake.createEntityArgsForCall)
}

func (fake *FakeEntityManager) CreateEntityReturns(result1 ecs.Entity) {
	fake.CreateEntityStub = nil
	fake.createEntityReturns = struct {
		result1 ecs.Entity
	}{result1}
}

func (fake *FakeEntityManager) HasEntity(arg1 ecs.Entity) bool {
	fake.hasEntityMutex.Lock()
	fake.hasEntityArgsForCall = append(fake.hasEntityArgsForCall, struct {
		arg1 ecs.Entity
	}{arg1})
	fake.hasEntityMutex.Unlock()
	if fake.HasEntityStub != nil {
		return fake.HasEntityStub(arg1)
	} else {
		return fake.hasEntityReturns.result1
	}
}

func (fake *FakeEntityManager) HasEntityCallCount() int {
	fake.hasEntityMutex.RLock()
	defer fake.hasEntityMutex.RUnlock()
	return len(fake.hasEntityArgsForCall)
}

func (fake *FakeEntityManager) HasEntityArgsForCall(i int) ecs.Entity {
	fake.hasEntityMutex.RLock()
	defer fake.hasEntityMutex.RUnlock()
	return fake.hasEntityArgsForCall[i].arg1
}

func (fake *FakeEntityManager) HasEntityReturns(result1 bool) {
	fake.HasEntityStub = nil
	fake.hasEntityReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEntityManager) DeleteEntity(arg1 ecs.Entity) {
	fake.deleteEntityMutex.Lock()
	fake.deleteEntityArgsForCall = append(fake.deleteEntityArgsForCall, struct {
		arg1 ecs.Entity
	}{arg1})
	fake.deleteEntityMutex.Unlock()
	if fake.DeleteEntityStub != nil {
		fake.DeleteEntityStub(arg1)
	}
}

func (fake *FakeEntityManager) DeleteEntityCallCount() int {
	fake.deleteEntityMutex.RLock()
	defer fake.deleteEntityMutex.RUnlock()
	return len(fake.deleteEntityArgsForCall)
}

func (fake *FakeEntityManager) DeleteEntityArgsForCall(i int) ecs.Entity {
	fake.deleteEntityMutex.RLock()
	defer fake.deleteEntityMutex.RUnlock()
	return fake.deleteEntityArgsForCall[i].arg1
}

func (fake *FakeEntityManager) DeleteAllEntities() {
	fake.deleteAllEntitiesMutex.Lock()
	fake.deleteAllEntitiesArgsForCall = append(fake.deleteAllEntitiesArgsForCall, struct{}{})
	fake.deleteAllEntitiesMutex.Unlock()
	if fake.DeleteAllEntitiesStub != nil {
		fake.DeleteAllEntitiesStub()
	}
}

func (fake *FakeEntityManager) DeleteAllEntitiesCallCount() int {
	fake.deleteAllEntitiesMutex.RLock()
	defer fake.deleteAllEntitiesMutex.RUnlock()
	return len(fake.deleteAllEntitiesArgsForCall)
}

var _ ecs.EntityManager = new(FakeEntityManager)
