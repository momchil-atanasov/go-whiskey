// Generated by 'github.com/mokiat/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/mokiat/go-whiskey/graphics"
)

type Float2AttributeArrayStub struct {
	StubGUID        int
	SizeStub        func() (result1 int)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct {
	}
	sizeReturns struct {
		result1 int
	}
	PutFloat2Stub        func(arg1 int, arg2 float32, arg3 float32)
	putFloat2Mutex       sync.RWMutex
	putFloat2ArgsForCall []struct {
		arg1 int
		arg2 float32
		arg3 float32
	}
	Float2Stub        func(arg1 int) (result1 float32, result2 float32)
	float2Mutex       sync.RWMutex
	float2ArgsForCall []struct {
		arg1 int
	}
	float2Returns struct {
		result1 float32
		result2 float32
	}
}

var _ alias1.Float2AttributeArray = new(Float2AttributeArrayStub)

func (stub *Float2AttributeArrayStub) Size() int {
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
func (stub *Float2AttributeArrayStub) SizeCallCount() int {
	stub.sizeMutex.RLock()
	defer stub.sizeMutex.RUnlock()
	return len(stub.sizeArgsForCall)
}
func (stub *Float2AttributeArrayStub) SizeReturns(result1 int) {
	stub.sizeMutex.Lock()
	defer stub.sizeMutex.Unlock()
	stub.sizeReturns = struct {
		result1 int
	}{result1}
}
func (stub *Float2AttributeArrayStub) PutFloat2(arg1 int, arg2 float32, arg3 float32) {
	stub.putFloat2Mutex.Lock()
	defer stub.putFloat2Mutex.Unlock()
	stub.putFloat2ArgsForCall = append(stub.putFloat2ArgsForCall, struct {
		arg1 int
		arg2 float32
		arg3 float32
	}{arg1, arg2, arg3})
	if stub.PutFloat2Stub != nil {
		stub.PutFloat2Stub(arg1, arg2, arg3)
	}
}
func (stub *Float2AttributeArrayStub) PutFloat2CallCount() int {
	stub.putFloat2Mutex.RLock()
	defer stub.putFloat2Mutex.RUnlock()
	return len(stub.putFloat2ArgsForCall)
}
func (stub *Float2AttributeArrayStub) PutFloat2ArgsForCall(index int) (int, float32, float32) {
	stub.putFloat2Mutex.RLock()
	defer stub.putFloat2Mutex.RUnlock()
	return stub.putFloat2ArgsForCall[index].arg1, stub.putFloat2ArgsForCall[index].arg2, stub.putFloat2ArgsForCall[index].arg3
}
func (stub *Float2AttributeArrayStub) Float2(arg1 int) (float32, float32) {
	stub.float2Mutex.Lock()
	defer stub.float2Mutex.Unlock()
	stub.float2ArgsForCall = append(stub.float2ArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.Float2Stub != nil {
		return stub.Float2Stub(arg1)
	} else {
		return stub.float2Returns.result1, stub.float2Returns.result2
	}
}
func (stub *Float2AttributeArrayStub) Float2CallCount() int {
	stub.float2Mutex.RLock()
	defer stub.float2Mutex.RUnlock()
	return len(stub.float2ArgsForCall)
}
func (stub *Float2AttributeArrayStub) Float2ArgsForCall(index int) int {
	stub.float2Mutex.RLock()
	defer stub.float2Mutex.RUnlock()
	return stub.float2ArgsForCall[index].arg1
}
func (stub *Float2AttributeArrayStub) Float2Returns(result1 float32, result2 float32) {
	stub.float2Mutex.Lock()
	defer stub.float2Mutex.Unlock()
	stub.float2Returns = struct {
		result1 float32
		result2 float32
	}{result1, result2}
}
