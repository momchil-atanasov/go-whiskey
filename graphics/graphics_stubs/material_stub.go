// Generated by 'github.com/momchil-atanasov/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/go-whiskey/graphics"
)

type MaterialStub struct {
	StubGUID                      int
	EnableDiffuseColorStub        func(arg1 bool)
	enableDiffuseColorMutex       sync.RWMutex
	enableDiffuseColorArgsForCall []struct {
		arg1 bool
	}
	DiffuseColorEnabledStub        func() (result1 bool)
	diffuseColorEnabledMutex       sync.RWMutex
	diffuseColorEnabledArgsForCall []struct {
	}
	diffuseColorEnabledReturns struct {
		result1 bool
	}
}

var _ alias1.Material = new(MaterialStub)

func (stub *MaterialStub) EnableDiffuseColor(arg1 bool) {
	stub.enableDiffuseColorMutex.Lock()
	defer stub.enableDiffuseColorMutex.Unlock()
	stub.enableDiffuseColorArgsForCall = append(stub.enableDiffuseColorArgsForCall, struct {
		arg1 bool
	}{arg1})
	if stub.EnableDiffuseColorStub != nil {
		stub.EnableDiffuseColorStub(arg1)
	}
}
func (stub *MaterialStub) EnableDiffuseColorCallCount() int {
	stub.enableDiffuseColorMutex.RLock()
	defer stub.enableDiffuseColorMutex.RUnlock()
	return len(stub.enableDiffuseColorArgsForCall)
}
func (stub *MaterialStub) EnableDiffuseColorArgsForCall(index int) bool {
	stub.enableDiffuseColorMutex.RLock()
	defer stub.enableDiffuseColorMutex.RUnlock()
	return stub.enableDiffuseColorArgsForCall[index].arg1
}
func (stub *MaterialStub) DiffuseColorEnabled() bool {
	stub.diffuseColorEnabledMutex.Lock()
	defer stub.diffuseColorEnabledMutex.Unlock()
	stub.diffuseColorEnabledArgsForCall = append(stub.diffuseColorEnabledArgsForCall, struct {
	}{})
	if stub.DiffuseColorEnabledStub != nil {
		return stub.DiffuseColorEnabledStub()
	} else {
		return stub.diffuseColorEnabledReturns.result1
	}
}
func (stub *MaterialStub) DiffuseColorEnabledCallCount() int {
	stub.diffuseColorEnabledMutex.RLock()
	defer stub.diffuseColorEnabledMutex.RUnlock()
	return len(stub.diffuseColorEnabledArgsForCall)
}
func (stub *MaterialStub) DiffuseColorEnabledReturns(result1 bool) {
	stub.diffuseColorEnabledMutex.Lock()
	defer stub.diffuseColorEnabledMutex.Unlock()
	stub.diffuseColorEnabledReturns = struct {
		result1 bool
	}{result1}
}
