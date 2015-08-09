// Generated by 'github.com/momchil-atanasov/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/go-whiskey/graphics"
)

type RendererStub struct {
	StubGUID              int
	InitializeStub        func() (result1 error)
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
	}
	initializeReturns struct {
		result1 error
	}
	UseMaterialStub        func(arg1 alias1.Material)
	useMaterialMutex       sync.RWMutex
	useMaterialArgsForCall []struct {
		arg1 alias1.Material
	}
	BindAttributeStub        func(arg1 alias1.AttributeName, arg2 alias1.AttributeArray)
	bindAttributeMutex       sync.RWMutex
	bindAttributeArgsForCall []struct {
		arg1 alias1.AttributeName
		arg2 alias1.AttributeArray
	}
	BindUniformStub        func(arg1 alias1.UniformName, arg2 alias1.Uniform)
	bindUniformMutex       sync.RWMutex
	bindUniformArgsForCall []struct {
		arg1 alias1.UniformName
		arg2 alias1.Uniform
	}
	RenderStub        func(arg1 alias1.IndexArray, arg2 alias1.SequenceType)
	renderMutex       sync.RWMutex
	renderArgsForCall []struct {
		arg1 alias1.IndexArray
		arg2 alias1.SequenceType
	}
	FlushStub        func() (result1 error)
	flushMutex       sync.RWMutex
	flushArgsForCall []struct {
	}
	flushReturns struct {
		result1 error
	}
}

var _ alias1.Renderer = new(RendererStub)

func (stub *RendererStub) Initialize() error {
	stub.initializeMutex.Lock()
	defer stub.initializeMutex.Unlock()
	stub.initializeArgsForCall = append(stub.initializeArgsForCall, struct {
	}{})
	if stub.InitializeStub != nil {
		return stub.InitializeStub()
	} else {
		return stub.initializeReturns.result1
	}
}
func (stub *RendererStub) InitializeCallCount() int {
	stub.initializeMutex.RLock()
	defer stub.initializeMutex.RUnlock()
	return len(stub.initializeArgsForCall)
}
func (stub *RendererStub) InitializeReturns(result1 error) {
	stub.initializeMutex.Lock()
	defer stub.initializeMutex.Unlock()
	stub.initializeReturns = struct {
		result1 error
	}{result1}
}
func (stub *RendererStub) UseMaterial(arg1 alias1.Material) {
	stub.useMaterialMutex.Lock()
	defer stub.useMaterialMutex.Unlock()
	stub.useMaterialArgsForCall = append(stub.useMaterialArgsForCall, struct {
		arg1 alias1.Material
	}{arg1})
	if stub.UseMaterialStub != nil {
		stub.UseMaterialStub(arg1)
	}
}
func (stub *RendererStub) UseMaterialCallCount() int {
	stub.useMaterialMutex.RLock()
	defer stub.useMaterialMutex.RUnlock()
	return len(stub.useMaterialArgsForCall)
}
func (stub *RendererStub) UseMaterialArgsForCall(index int) alias1.Material {
	stub.useMaterialMutex.RLock()
	defer stub.useMaterialMutex.RUnlock()
	return stub.useMaterialArgsForCall[index].arg1
}
func (stub *RendererStub) BindAttribute(arg1 alias1.AttributeName, arg2 alias1.AttributeArray) {
	stub.bindAttributeMutex.Lock()
	defer stub.bindAttributeMutex.Unlock()
	stub.bindAttributeArgsForCall = append(stub.bindAttributeArgsForCall, struct {
		arg1 alias1.AttributeName
		arg2 alias1.AttributeArray
	}{arg1, arg2})
	if stub.BindAttributeStub != nil {
		stub.BindAttributeStub(arg1, arg2)
	}
}
func (stub *RendererStub) BindAttributeCallCount() int {
	stub.bindAttributeMutex.RLock()
	defer stub.bindAttributeMutex.RUnlock()
	return len(stub.bindAttributeArgsForCall)
}
func (stub *RendererStub) BindAttributeArgsForCall(index int) (alias1.AttributeName, alias1.AttributeArray) {
	stub.bindAttributeMutex.RLock()
	defer stub.bindAttributeMutex.RUnlock()
	return stub.bindAttributeArgsForCall[index].arg1, stub.bindAttributeArgsForCall[index].arg2
}
func (stub *RendererStub) BindUniform(arg1 alias1.UniformName, arg2 alias1.Uniform) {
	stub.bindUniformMutex.Lock()
	defer stub.bindUniformMutex.Unlock()
	stub.bindUniformArgsForCall = append(stub.bindUniformArgsForCall, struct {
		arg1 alias1.UniformName
		arg2 alias1.Uniform
	}{arg1, arg2})
	if stub.BindUniformStub != nil {
		stub.BindUniformStub(arg1, arg2)
	}
}
func (stub *RendererStub) BindUniformCallCount() int {
	stub.bindUniformMutex.RLock()
	defer stub.bindUniformMutex.RUnlock()
	return len(stub.bindUniformArgsForCall)
}
func (stub *RendererStub) BindUniformArgsForCall(index int) (alias1.UniformName, alias1.Uniform) {
	stub.bindUniformMutex.RLock()
	defer stub.bindUniformMutex.RUnlock()
	return stub.bindUniformArgsForCall[index].arg1, stub.bindUniformArgsForCall[index].arg2
}
func (stub *RendererStub) Render(arg1 alias1.IndexArray, arg2 alias1.SequenceType) {
	stub.renderMutex.Lock()
	defer stub.renderMutex.Unlock()
	stub.renderArgsForCall = append(stub.renderArgsForCall, struct {
		arg1 alias1.IndexArray
		arg2 alias1.SequenceType
	}{arg1, arg2})
	if stub.RenderStub != nil {
		stub.RenderStub(arg1, arg2)
	}
}
func (stub *RendererStub) RenderCallCount() int {
	stub.renderMutex.RLock()
	defer stub.renderMutex.RUnlock()
	return len(stub.renderArgsForCall)
}
func (stub *RendererStub) RenderArgsForCall(index int) (alias1.IndexArray, alias1.SequenceType) {
	stub.renderMutex.RLock()
	defer stub.renderMutex.RUnlock()
	return stub.renderArgsForCall[index].arg1, stub.renderArgsForCall[index].arg2
}
func (stub *RendererStub) Flush() error {
	stub.flushMutex.Lock()
	defer stub.flushMutex.Unlock()
	stub.flushArgsForCall = append(stub.flushArgsForCall, struct {
	}{})
	if stub.FlushStub != nil {
		return stub.FlushStub()
	} else {
		return stub.flushReturns.result1
	}
}
func (stub *RendererStub) FlushCallCount() int {
	stub.flushMutex.RLock()
	defer stub.flushMutex.RUnlock()
	return len(stub.flushArgsForCall)
}
func (stub *RendererStub) FlushReturns(result1 error) {
	stub.flushMutex.Lock()
	defer stub.flushMutex.Unlock()
	stub.flushReturns = struct {
		result1 error
	}{result1}
}
