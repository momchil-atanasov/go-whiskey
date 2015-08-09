// Generated by 'github.com/momchil-atanasov/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/go-whiskey/graphics"
)

type GraphicsStub struct {
	StubGUID                              int
	CreateFloat2AttributeArrayStub        func(arg1 int) (result1 alias1.Float2AttributeArray, result2 error)
	createFloat2AttributeArrayMutex       sync.RWMutex
	createFloat2AttributeArrayArgsForCall []struct {
		arg1 int
	}
	createFloat2AttributeArrayReturns struct {
		result1 alias1.Float2AttributeArray
		result2 error
	}
	CreateFloat3AttributeArrayStub        func(arg1 int) (result1 alias1.Float3AttributeArray, result2 error)
	createFloat3AttributeArrayMutex       sync.RWMutex
	createFloat3AttributeArrayArgsForCall []struct {
		arg1 int
	}
	createFloat3AttributeArrayReturns struct {
		result1 alias1.Float3AttributeArray
		result2 error
	}
	CreateFloat4UniformStub        func() (result1 alias1.Float4Uniform, result2 error)
	createFloat4UniformMutex       sync.RWMutex
	createFloat4UniformArgsForCall []struct {
	}
	createFloat4UniformReturns struct {
		result1 alias1.Float4Uniform
		result2 error
	}
	CreateFloat4x4UniformStub        func() (result1 alias1.Float4x4Uniform, result2 error)
	createFloat4x4UniformMutex       sync.RWMutex
	createFloat4x4UniformArgsForCall []struct {
	}
	createFloat4x4UniformReturns struct {
		result1 alias1.Float4x4Uniform
		result2 error
	}
	CreateIndexArrayStub        func(arg1 int) (result1 alias1.IndexArray, result2 error)
	createIndexArrayMutex       sync.RWMutex
	createIndexArrayArgsForCall []struct {
		arg1 int
	}
	createIndexArrayReturns struct {
		result1 alias1.IndexArray
		result2 error
	}
	CreateMaterialStub        func() (result1 alias1.Material, result2 error)
	createMaterialMutex       sync.RWMutex
	createMaterialArgsForCall []struct {
	}
	createMaterialReturns struct {
		result1 alias1.Material
		result2 error
	}
	RendererStub        func() (result1 alias1.Renderer)
	rendererMutex       sync.RWMutex
	rendererArgsForCall []struct {
	}
	rendererReturns struct {
		result1 alias1.Renderer
	}
}

var _ alias1.Graphics = new(GraphicsStub)

func (stub *GraphicsStub) CreateFloat2AttributeArray(arg1 int) (alias1.Float2AttributeArray, error) {
	stub.createFloat2AttributeArrayMutex.Lock()
	defer stub.createFloat2AttributeArrayMutex.Unlock()
	stub.createFloat2AttributeArrayArgsForCall = append(stub.createFloat2AttributeArrayArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateFloat2AttributeArrayStub != nil {
		return stub.CreateFloat2AttributeArrayStub(arg1)
	} else {
		return stub.createFloat2AttributeArrayReturns.result1, stub.createFloat2AttributeArrayReturns.result2
	}
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayCallCount() int {
	stub.createFloat2AttributeArrayMutex.RLock()
	defer stub.createFloat2AttributeArrayMutex.RUnlock()
	return len(stub.createFloat2AttributeArrayArgsForCall)
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayArgsForCall(index int) int {
	stub.createFloat2AttributeArrayMutex.RLock()
	defer stub.createFloat2AttributeArrayMutex.RUnlock()
	return stub.createFloat2AttributeArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateFloat2AttributeArrayReturns(result1 alias1.Float2AttributeArray, result2 error) {
	stub.createFloat2AttributeArrayMutex.Lock()
	defer stub.createFloat2AttributeArrayMutex.Unlock()
	stub.createFloat2AttributeArrayReturns = struct {
		result1 alias1.Float2AttributeArray
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateFloat3AttributeArray(arg1 int) (alias1.Float3AttributeArray, error) {
	stub.createFloat3AttributeArrayMutex.Lock()
	defer stub.createFloat3AttributeArrayMutex.Unlock()
	stub.createFloat3AttributeArrayArgsForCall = append(stub.createFloat3AttributeArrayArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateFloat3AttributeArrayStub != nil {
		return stub.CreateFloat3AttributeArrayStub(arg1)
	} else {
		return stub.createFloat3AttributeArrayReturns.result1, stub.createFloat3AttributeArrayReturns.result2
	}
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayCallCount() int {
	stub.createFloat3AttributeArrayMutex.RLock()
	defer stub.createFloat3AttributeArrayMutex.RUnlock()
	return len(stub.createFloat3AttributeArrayArgsForCall)
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayArgsForCall(index int) int {
	stub.createFloat3AttributeArrayMutex.RLock()
	defer stub.createFloat3AttributeArrayMutex.RUnlock()
	return stub.createFloat3AttributeArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateFloat3AttributeArrayReturns(result1 alias1.Float3AttributeArray, result2 error) {
	stub.createFloat3AttributeArrayMutex.Lock()
	defer stub.createFloat3AttributeArrayMutex.Unlock()
	stub.createFloat3AttributeArrayReturns = struct {
		result1 alias1.Float3AttributeArray
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateFloat4Uniform() (alias1.Float4Uniform, error) {
	stub.createFloat4UniformMutex.Lock()
	defer stub.createFloat4UniformMutex.Unlock()
	stub.createFloat4UniformArgsForCall = append(stub.createFloat4UniformArgsForCall, struct {
	}{})
	if stub.CreateFloat4UniformStub != nil {
		return stub.CreateFloat4UniformStub()
	} else {
		return stub.createFloat4UniformReturns.result1, stub.createFloat4UniformReturns.result2
	}
}
func (stub *GraphicsStub) CreateFloat4UniformCallCount() int {
	stub.createFloat4UniformMutex.RLock()
	defer stub.createFloat4UniformMutex.RUnlock()
	return len(stub.createFloat4UniformArgsForCall)
}
func (stub *GraphicsStub) CreateFloat4UniformReturns(result1 alias1.Float4Uniform, result2 error) {
	stub.createFloat4UniformMutex.Lock()
	defer stub.createFloat4UniformMutex.Unlock()
	stub.createFloat4UniformReturns = struct {
		result1 alias1.Float4Uniform
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateFloat4x4Uniform() (alias1.Float4x4Uniform, error) {
	stub.createFloat4x4UniformMutex.Lock()
	defer stub.createFloat4x4UniformMutex.Unlock()
	stub.createFloat4x4UniformArgsForCall = append(stub.createFloat4x4UniformArgsForCall, struct {
	}{})
	if stub.CreateFloat4x4UniformStub != nil {
		return stub.CreateFloat4x4UniformStub()
	} else {
		return stub.createFloat4x4UniformReturns.result1, stub.createFloat4x4UniformReturns.result2
	}
}
func (stub *GraphicsStub) CreateFloat4x4UniformCallCount() int {
	stub.createFloat4x4UniformMutex.RLock()
	defer stub.createFloat4x4UniformMutex.RUnlock()
	return len(stub.createFloat4x4UniformArgsForCall)
}
func (stub *GraphicsStub) CreateFloat4x4UniformReturns(result1 alias1.Float4x4Uniform, result2 error) {
	stub.createFloat4x4UniformMutex.Lock()
	defer stub.createFloat4x4UniformMutex.Unlock()
	stub.createFloat4x4UniformReturns = struct {
		result1 alias1.Float4x4Uniform
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateIndexArray(arg1 int) (alias1.IndexArray, error) {
	stub.createIndexArrayMutex.Lock()
	defer stub.createIndexArrayMutex.Unlock()
	stub.createIndexArrayArgsForCall = append(stub.createIndexArrayArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateIndexArrayStub != nil {
		return stub.CreateIndexArrayStub(arg1)
	} else {
		return stub.createIndexArrayReturns.result1, stub.createIndexArrayReturns.result2
	}
}
func (stub *GraphicsStub) CreateIndexArrayCallCount() int {
	stub.createIndexArrayMutex.RLock()
	defer stub.createIndexArrayMutex.RUnlock()
	return len(stub.createIndexArrayArgsForCall)
}
func (stub *GraphicsStub) CreateIndexArrayArgsForCall(index int) int {
	stub.createIndexArrayMutex.RLock()
	defer stub.createIndexArrayMutex.RUnlock()
	return stub.createIndexArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateIndexArrayReturns(result1 alias1.IndexArray, result2 error) {
	stub.createIndexArrayMutex.Lock()
	defer stub.createIndexArrayMutex.Unlock()
	stub.createIndexArrayReturns = struct {
		result1 alias1.IndexArray
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateMaterial() (alias1.Material, error) {
	stub.createMaterialMutex.Lock()
	defer stub.createMaterialMutex.Unlock()
	stub.createMaterialArgsForCall = append(stub.createMaterialArgsForCall, struct {
	}{})
	if stub.CreateMaterialStub != nil {
		return stub.CreateMaterialStub()
	} else {
		return stub.createMaterialReturns.result1, stub.createMaterialReturns.result2
	}
}
func (stub *GraphicsStub) CreateMaterialCallCount() int {
	stub.createMaterialMutex.RLock()
	defer stub.createMaterialMutex.RUnlock()
	return len(stub.createMaterialArgsForCall)
}
func (stub *GraphicsStub) CreateMaterialReturns(result1 alias1.Material, result2 error) {
	stub.createMaterialMutex.Lock()
	defer stub.createMaterialMutex.Unlock()
	stub.createMaterialReturns = struct {
		result1 alias1.Material
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) Renderer() alias1.Renderer {
	stub.rendererMutex.Lock()
	defer stub.rendererMutex.Unlock()
	stub.rendererArgsForCall = append(stub.rendererArgsForCall, struct {
	}{})
	if stub.RendererStub != nil {
		return stub.RendererStub()
	} else {
		return stub.rendererReturns.result1
	}
}
func (stub *GraphicsStub) RendererCallCount() int {
	stub.rendererMutex.RLock()
	defer stub.rendererMutex.RUnlock()
	return len(stub.rendererArgsForCall)
}
func (stub *GraphicsStub) RendererReturns(result1 alias1.Renderer) {
	stub.rendererMutex.Lock()
	defer stub.rendererMutex.Unlock()
	stub.rendererReturns = struct {
		result1 alias1.Renderer
	}{result1}
}
