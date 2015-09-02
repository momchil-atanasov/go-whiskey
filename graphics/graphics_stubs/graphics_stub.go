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
	DeleteAttributeArrayStub        func(arg1 alias1.AttributeArray) (result1 error)
	deleteAttributeArrayMutex       sync.RWMutex
	deleteAttributeArrayArgsForCall []struct {
		arg1 alias1.AttributeArray
	}
	deleteAttributeArrayReturns struct {
		result1 error
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
	DeleteIndexArrayStub        func(arg1 alias1.IndexArray) (result1 error)
	deleteIndexArrayMutex       sync.RWMutex
	deleteIndexArrayArgsForCall []struct {
		arg1 alias1.IndexArray
	}
	deleteIndexArrayReturns struct {
		result1 error
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
	DeleteUniformStub        func(arg1 alias1.Uniform) (result1 error)
	deleteUniformMutex       sync.RWMutex
	deleteUniformArgsForCall []struct {
		arg1 alias1.Uniform
	}
	deleteUniformReturns struct {
		result1 error
	}
	CreateFlatTextureStub        func(arg1 int, arg2 int) (result1 alias1.FlatTexture, result2 error)
	createFlatTextureMutex       sync.RWMutex
	createFlatTextureArgsForCall []struct {
		arg1 int
		arg2 int
	}
	createFlatTextureReturns struct {
		result1 alias1.FlatTexture
		result2 error
	}
	CreateCubeTextureStub        func(arg1 int) (result1 alias1.CubeTexture, result2 error)
	createCubeTextureMutex       sync.RWMutex
	createCubeTextureArgsForCall []struct {
		arg1 int
	}
	createCubeTextureReturns struct {
		result1 alias1.CubeTexture
		result2 error
	}
	DeleteTextureStub        func(arg1 alias1.Texture) (result1 error)
	deleteTextureMutex       sync.RWMutex
	deleteTextureArgsForCall []struct {
		arg1 alias1.Texture
	}
	deleteTextureReturns struct {
		result1 error
	}
	CreateMaterialStub        func(arg1 alias1.MaterialDefinition) (result1 alias1.Material, result2 error)
	createMaterialMutex       sync.RWMutex
	createMaterialArgsForCall []struct {
		arg1 alias1.MaterialDefinition
	}
	createMaterialReturns struct {
		result1 alias1.Material
		result2 error
	}
	DeleteMaterialStub        func(arg1 alias1.Material) (result1 error)
	deleteMaterialMutex       sync.RWMutex
	deleteMaterialArgsForCall []struct {
		arg1 alias1.Material
	}
	deleteMaterialReturns struct {
		result1 error
	}
	InvalidateStub        func() (result1 error)
	invalidateMutex       sync.RWMutex
	invalidateArgsForCall []struct {
	}
	invalidateReturns struct {
		result1 error
	}
	InitializeStub        func() (result1 error)
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
	}
	initializeReturns struct {
		result1 error
	}
	DestroyStub        func() (result1 error)
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
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
	BindTextureStub        func(arg1 alias1.TextureName, arg2 alias1.Texture)
	bindTextureMutex       sync.RWMutex
	bindTextureArgsForCall []struct {
		arg1 alias1.TextureName
		arg2 alias1.Texture
	}
	RenderStub        func(arg1 alias1.SequenceType, arg2 alias1.IndexArray, arg3 int, arg4 int)
	renderMutex       sync.RWMutex
	renderArgsForCall []struct {
		arg1 alias1.SequenceType
		arg2 alias1.IndexArray
		arg3 int
		arg4 int
	}
	FlushStub        func() (result1 error)
	flushMutex       sync.RWMutex
	flushArgsForCall []struct {
	}
	flushReturns struct {
		result1 error
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
func (stub *GraphicsStub) DeleteAttributeArray(arg1 alias1.AttributeArray) error {
	stub.deleteAttributeArrayMutex.Lock()
	defer stub.deleteAttributeArrayMutex.Unlock()
	stub.deleteAttributeArrayArgsForCall = append(stub.deleteAttributeArrayArgsForCall, struct {
		arg1 alias1.AttributeArray
	}{arg1})
	if stub.DeleteAttributeArrayStub != nil {
		return stub.DeleteAttributeArrayStub(arg1)
	} else {
		return stub.deleteAttributeArrayReturns.result1
	}
}
func (stub *GraphicsStub) DeleteAttributeArrayCallCount() int {
	stub.deleteAttributeArrayMutex.RLock()
	defer stub.deleteAttributeArrayMutex.RUnlock()
	return len(stub.deleteAttributeArrayArgsForCall)
}
func (stub *GraphicsStub) DeleteAttributeArrayArgsForCall(index int) alias1.AttributeArray {
	stub.deleteAttributeArrayMutex.RLock()
	defer stub.deleteAttributeArrayMutex.RUnlock()
	return stub.deleteAttributeArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) DeleteAttributeArrayReturns(result1 error) {
	stub.deleteAttributeArrayMutex.Lock()
	defer stub.deleteAttributeArrayMutex.Unlock()
	stub.deleteAttributeArrayReturns = struct {
		result1 error
	}{result1}
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
func (stub *GraphicsStub) DeleteIndexArray(arg1 alias1.IndexArray) error {
	stub.deleteIndexArrayMutex.Lock()
	defer stub.deleteIndexArrayMutex.Unlock()
	stub.deleteIndexArrayArgsForCall = append(stub.deleteIndexArrayArgsForCall, struct {
		arg1 alias1.IndexArray
	}{arg1})
	if stub.DeleteIndexArrayStub != nil {
		return stub.DeleteIndexArrayStub(arg1)
	} else {
		return stub.deleteIndexArrayReturns.result1
	}
}
func (stub *GraphicsStub) DeleteIndexArrayCallCount() int {
	stub.deleteIndexArrayMutex.RLock()
	defer stub.deleteIndexArrayMutex.RUnlock()
	return len(stub.deleteIndexArrayArgsForCall)
}
func (stub *GraphicsStub) DeleteIndexArrayArgsForCall(index int) alias1.IndexArray {
	stub.deleteIndexArrayMutex.RLock()
	defer stub.deleteIndexArrayMutex.RUnlock()
	return stub.deleteIndexArrayArgsForCall[index].arg1
}
func (stub *GraphicsStub) DeleteIndexArrayReturns(result1 error) {
	stub.deleteIndexArrayMutex.Lock()
	defer stub.deleteIndexArrayMutex.Unlock()
	stub.deleteIndexArrayReturns = struct {
		result1 error
	}{result1}
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
func (stub *GraphicsStub) DeleteUniform(arg1 alias1.Uniform) error {
	stub.deleteUniformMutex.Lock()
	defer stub.deleteUniformMutex.Unlock()
	stub.deleteUniformArgsForCall = append(stub.deleteUniformArgsForCall, struct {
		arg1 alias1.Uniform
	}{arg1})
	if stub.DeleteUniformStub != nil {
		return stub.DeleteUniformStub(arg1)
	} else {
		return stub.deleteUniformReturns.result1
	}
}
func (stub *GraphicsStub) DeleteUniformCallCount() int {
	stub.deleteUniformMutex.RLock()
	defer stub.deleteUniformMutex.RUnlock()
	return len(stub.deleteUniformArgsForCall)
}
func (stub *GraphicsStub) DeleteUniformArgsForCall(index int) alias1.Uniform {
	stub.deleteUniformMutex.RLock()
	defer stub.deleteUniformMutex.RUnlock()
	return stub.deleteUniformArgsForCall[index].arg1
}
func (stub *GraphicsStub) DeleteUniformReturns(result1 error) {
	stub.deleteUniformMutex.Lock()
	defer stub.deleteUniformMutex.Unlock()
	stub.deleteUniformReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) CreateFlatTexture(arg1 int, arg2 int) (alias1.FlatTexture, error) {
	stub.createFlatTextureMutex.Lock()
	defer stub.createFlatTextureMutex.Unlock()
	stub.createFlatTextureArgsForCall = append(stub.createFlatTextureArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	if stub.CreateFlatTextureStub != nil {
		return stub.CreateFlatTextureStub(arg1, arg2)
	} else {
		return stub.createFlatTextureReturns.result1, stub.createFlatTextureReturns.result2
	}
}
func (stub *GraphicsStub) CreateFlatTextureCallCount() int {
	stub.createFlatTextureMutex.RLock()
	defer stub.createFlatTextureMutex.RUnlock()
	return len(stub.createFlatTextureArgsForCall)
}
func (stub *GraphicsStub) CreateFlatTextureArgsForCall(index int) (int, int) {
	stub.createFlatTextureMutex.RLock()
	defer stub.createFlatTextureMutex.RUnlock()
	return stub.createFlatTextureArgsForCall[index].arg1, stub.createFlatTextureArgsForCall[index].arg2
}
func (stub *GraphicsStub) CreateFlatTextureReturns(result1 alias1.FlatTexture, result2 error) {
	stub.createFlatTextureMutex.Lock()
	defer stub.createFlatTextureMutex.Unlock()
	stub.createFlatTextureReturns = struct {
		result1 alias1.FlatTexture
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) CreateCubeTexture(arg1 int) (alias1.CubeTexture, error) {
	stub.createCubeTextureMutex.Lock()
	defer stub.createCubeTextureMutex.Unlock()
	stub.createCubeTextureArgsForCall = append(stub.createCubeTextureArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.CreateCubeTextureStub != nil {
		return stub.CreateCubeTextureStub(arg1)
	} else {
		return stub.createCubeTextureReturns.result1, stub.createCubeTextureReturns.result2
	}
}
func (stub *GraphicsStub) CreateCubeTextureCallCount() int {
	stub.createCubeTextureMutex.RLock()
	defer stub.createCubeTextureMutex.RUnlock()
	return len(stub.createCubeTextureArgsForCall)
}
func (stub *GraphicsStub) CreateCubeTextureArgsForCall(index int) int {
	stub.createCubeTextureMutex.RLock()
	defer stub.createCubeTextureMutex.RUnlock()
	return stub.createCubeTextureArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateCubeTextureReturns(result1 alias1.CubeTexture, result2 error) {
	stub.createCubeTextureMutex.Lock()
	defer stub.createCubeTextureMutex.Unlock()
	stub.createCubeTextureReturns = struct {
		result1 alias1.CubeTexture
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) DeleteTexture(arg1 alias1.Texture) error {
	stub.deleteTextureMutex.Lock()
	defer stub.deleteTextureMutex.Unlock()
	stub.deleteTextureArgsForCall = append(stub.deleteTextureArgsForCall, struct {
		arg1 alias1.Texture
	}{arg1})
	if stub.DeleteTextureStub != nil {
		return stub.DeleteTextureStub(arg1)
	} else {
		return stub.deleteTextureReturns.result1
	}
}
func (stub *GraphicsStub) DeleteTextureCallCount() int {
	stub.deleteTextureMutex.RLock()
	defer stub.deleteTextureMutex.RUnlock()
	return len(stub.deleteTextureArgsForCall)
}
func (stub *GraphicsStub) DeleteTextureArgsForCall(index int) alias1.Texture {
	stub.deleteTextureMutex.RLock()
	defer stub.deleteTextureMutex.RUnlock()
	return stub.deleteTextureArgsForCall[index].arg1
}
func (stub *GraphicsStub) DeleteTextureReturns(result1 error) {
	stub.deleteTextureMutex.Lock()
	defer stub.deleteTextureMutex.Unlock()
	stub.deleteTextureReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) CreateMaterial(arg1 alias1.MaterialDefinition) (alias1.Material, error) {
	stub.createMaterialMutex.Lock()
	defer stub.createMaterialMutex.Unlock()
	stub.createMaterialArgsForCall = append(stub.createMaterialArgsForCall, struct {
		arg1 alias1.MaterialDefinition
	}{arg1})
	if stub.CreateMaterialStub != nil {
		return stub.CreateMaterialStub(arg1)
	} else {
		return stub.createMaterialReturns.result1, stub.createMaterialReturns.result2
	}
}
func (stub *GraphicsStub) CreateMaterialCallCount() int {
	stub.createMaterialMutex.RLock()
	defer stub.createMaterialMutex.RUnlock()
	return len(stub.createMaterialArgsForCall)
}
func (stub *GraphicsStub) CreateMaterialArgsForCall(index int) alias1.MaterialDefinition {
	stub.createMaterialMutex.RLock()
	defer stub.createMaterialMutex.RUnlock()
	return stub.createMaterialArgsForCall[index].arg1
}
func (stub *GraphicsStub) CreateMaterialReturns(result1 alias1.Material, result2 error) {
	stub.createMaterialMutex.Lock()
	defer stub.createMaterialMutex.Unlock()
	stub.createMaterialReturns = struct {
		result1 alias1.Material
		result2 error
	}{result1, result2}
}
func (stub *GraphicsStub) DeleteMaterial(arg1 alias1.Material) error {
	stub.deleteMaterialMutex.Lock()
	defer stub.deleteMaterialMutex.Unlock()
	stub.deleteMaterialArgsForCall = append(stub.deleteMaterialArgsForCall, struct {
		arg1 alias1.Material
	}{arg1})
	if stub.DeleteMaterialStub != nil {
		return stub.DeleteMaterialStub(arg1)
	} else {
		return stub.deleteMaterialReturns.result1
	}
}
func (stub *GraphicsStub) DeleteMaterialCallCount() int {
	stub.deleteMaterialMutex.RLock()
	defer stub.deleteMaterialMutex.RUnlock()
	return len(stub.deleteMaterialArgsForCall)
}
func (stub *GraphicsStub) DeleteMaterialArgsForCall(index int) alias1.Material {
	stub.deleteMaterialMutex.RLock()
	defer stub.deleteMaterialMutex.RUnlock()
	return stub.deleteMaterialArgsForCall[index].arg1
}
func (stub *GraphicsStub) DeleteMaterialReturns(result1 error) {
	stub.deleteMaterialMutex.Lock()
	defer stub.deleteMaterialMutex.Unlock()
	stub.deleteMaterialReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) Invalidate() error {
	stub.invalidateMutex.Lock()
	defer stub.invalidateMutex.Unlock()
	stub.invalidateArgsForCall = append(stub.invalidateArgsForCall, struct {
	}{})
	if stub.InvalidateStub != nil {
		return stub.InvalidateStub()
	} else {
		return stub.invalidateReturns.result1
	}
}
func (stub *GraphicsStub) InvalidateCallCount() int {
	stub.invalidateMutex.RLock()
	defer stub.invalidateMutex.RUnlock()
	return len(stub.invalidateArgsForCall)
}
func (stub *GraphicsStub) InvalidateReturns(result1 error) {
	stub.invalidateMutex.Lock()
	defer stub.invalidateMutex.Unlock()
	stub.invalidateReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) Initialize() error {
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
func (stub *GraphicsStub) InitializeCallCount() int {
	stub.initializeMutex.RLock()
	defer stub.initializeMutex.RUnlock()
	return len(stub.initializeArgsForCall)
}
func (stub *GraphicsStub) InitializeReturns(result1 error) {
	stub.initializeMutex.Lock()
	defer stub.initializeMutex.Unlock()
	stub.initializeReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) Destroy() error {
	stub.destroyMutex.Lock()
	defer stub.destroyMutex.Unlock()
	stub.destroyArgsForCall = append(stub.destroyArgsForCall, struct {
	}{})
	if stub.DestroyStub != nil {
		return stub.DestroyStub()
	} else {
		return stub.destroyReturns.result1
	}
}
func (stub *GraphicsStub) DestroyCallCount() int {
	stub.destroyMutex.RLock()
	defer stub.destroyMutex.RUnlock()
	return len(stub.destroyArgsForCall)
}
func (stub *GraphicsStub) DestroyReturns(result1 error) {
	stub.destroyMutex.Lock()
	defer stub.destroyMutex.Unlock()
	stub.destroyReturns = struct {
		result1 error
	}{result1}
}
func (stub *GraphicsStub) UseMaterial(arg1 alias1.Material) {
	stub.useMaterialMutex.Lock()
	defer stub.useMaterialMutex.Unlock()
	stub.useMaterialArgsForCall = append(stub.useMaterialArgsForCall, struct {
		arg1 alias1.Material
	}{arg1})
	if stub.UseMaterialStub != nil {
		stub.UseMaterialStub(arg1)
	}
}
func (stub *GraphicsStub) UseMaterialCallCount() int {
	stub.useMaterialMutex.RLock()
	defer stub.useMaterialMutex.RUnlock()
	return len(stub.useMaterialArgsForCall)
}
func (stub *GraphicsStub) UseMaterialArgsForCall(index int) alias1.Material {
	stub.useMaterialMutex.RLock()
	defer stub.useMaterialMutex.RUnlock()
	return stub.useMaterialArgsForCall[index].arg1
}
func (stub *GraphicsStub) BindAttribute(arg1 alias1.AttributeName, arg2 alias1.AttributeArray) {
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
func (stub *GraphicsStub) BindAttributeCallCount() int {
	stub.bindAttributeMutex.RLock()
	defer stub.bindAttributeMutex.RUnlock()
	return len(stub.bindAttributeArgsForCall)
}
func (stub *GraphicsStub) BindAttributeArgsForCall(index int) (alias1.AttributeName, alias1.AttributeArray) {
	stub.bindAttributeMutex.RLock()
	defer stub.bindAttributeMutex.RUnlock()
	return stub.bindAttributeArgsForCall[index].arg1, stub.bindAttributeArgsForCall[index].arg2
}
func (stub *GraphicsStub) BindUniform(arg1 alias1.UniformName, arg2 alias1.Uniform) {
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
func (stub *GraphicsStub) BindUniformCallCount() int {
	stub.bindUniformMutex.RLock()
	defer stub.bindUniformMutex.RUnlock()
	return len(stub.bindUniformArgsForCall)
}
func (stub *GraphicsStub) BindUniformArgsForCall(index int) (alias1.UniformName, alias1.Uniform) {
	stub.bindUniformMutex.RLock()
	defer stub.bindUniformMutex.RUnlock()
	return stub.bindUniformArgsForCall[index].arg1, stub.bindUniformArgsForCall[index].arg2
}
func (stub *GraphicsStub) BindTexture(arg1 alias1.TextureName, arg2 alias1.Texture) {
	stub.bindTextureMutex.Lock()
	defer stub.bindTextureMutex.Unlock()
	stub.bindTextureArgsForCall = append(stub.bindTextureArgsForCall, struct {
		arg1 alias1.TextureName
		arg2 alias1.Texture
	}{arg1, arg2})
	if stub.BindTextureStub != nil {
		stub.BindTextureStub(arg1, arg2)
	}
}
func (stub *GraphicsStub) BindTextureCallCount() int {
	stub.bindTextureMutex.RLock()
	defer stub.bindTextureMutex.RUnlock()
	return len(stub.bindTextureArgsForCall)
}
func (stub *GraphicsStub) BindTextureArgsForCall(index int) (alias1.TextureName, alias1.Texture) {
	stub.bindTextureMutex.RLock()
	defer stub.bindTextureMutex.RUnlock()
	return stub.bindTextureArgsForCall[index].arg1, stub.bindTextureArgsForCall[index].arg2
}
func (stub *GraphicsStub) Render(arg1 alias1.SequenceType, arg2 alias1.IndexArray, arg3 int, arg4 int) {
	stub.renderMutex.Lock()
	defer stub.renderMutex.Unlock()
	stub.renderArgsForCall = append(stub.renderArgsForCall, struct {
		arg1 alias1.SequenceType
		arg2 alias1.IndexArray
		arg3 int
		arg4 int
	}{arg1, arg2, arg3, arg4})
	if stub.RenderStub != nil {
		stub.RenderStub(arg1, arg2, arg3, arg4)
	}
}
func (stub *GraphicsStub) RenderCallCount() int {
	stub.renderMutex.RLock()
	defer stub.renderMutex.RUnlock()
	return len(stub.renderArgsForCall)
}
func (stub *GraphicsStub) RenderArgsForCall(index int) (alias1.SequenceType, alias1.IndexArray, int, int) {
	stub.renderMutex.RLock()
	defer stub.renderMutex.RUnlock()
	return stub.renderArgsForCall[index].arg1, stub.renderArgsForCall[index].arg2, stub.renderArgsForCall[index].arg3, stub.renderArgsForCall[index].arg4
}
func (stub *GraphicsStub) Flush() error {
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
func (stub *GraphicsStub) FlushCallCount() int {
	stub.flushMutex.RLock()
	defer stub.flushMutex.RUnlock()
	return len(stub.flushArgsForCall)
}
func (stub *GraphicsStub) FlushReturns(result1 error) {
	stub.flushMutex.Lock()
	defer stub.flushMutex.Unlock()
	stub.flushReturns = struct {
		result1 error
	}{result1}
}
