// Generated by 'github.com/mokiat/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/mokiat/go-whiskey/graphics"
)

type AttributeArrayStub struct {
	StubGUID        int
	SizeStub        func() (result1 int)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct {
	}
	sizeReturns struct {
		result1 int
	}
}

var _ alias1.AttributeArray = new(AttributeArrayStub)

func (stub *AttributeArrayStub) Size() int {
	stub.sizeMutex.Lock()
	defer stub.sizeMutex.Unlock()
	stub.sizeArgsForCall = append(stub.sizeArgsForCall, struct {
	}{})
	if stub.SizeStub != nil {
		return stub.SizeStub()
	} else {
		return stub.sizeReturns.result1
	}
}
func (stub *AttributeArrayStub) SizeCallCount() int {
	stub.sizeMutex.RLock()
	defer stub.sizeMutex.RUnlock()
	return len(stub.sizeArgsForCall)
}
func (stub *AttributeArrayStub) SizeReturns(result1 int) {
	stub.sizeMutex.Lock()
	defer stub.sizeMutex.Unlock()
	stub.sizeReturns = struct {
		result1 int
	}{result1}
}
