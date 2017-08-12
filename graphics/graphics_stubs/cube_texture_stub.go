// Generated by 'github.com/mokiat/gostub'

package graphics_stubs

import (
	sync "sync"

	alias1 "github.com/mokiat/go-whiskey/graphics"
)

type CubeTextureStub struct {
	StubGUID            int
	SetTexelStub        func(arg1 alias1.CubeSide, arg2 int, arg3 int, arg4 alias1.RGBAColor)
	setTexelMutex       sync.RWMutex
	setTexelArgsForCall []struct {
		arg1 alias1.CubeSide
		arg2 int
		arg3 int
		arg4 alias1.RGBAColor
	}
	TexelStub        func(arg1 alias1.CubeSide, arg2 int, arg3 int) (result1 alias1.RGBAColor)
	texelMutex       sync.RWMutex
	texelArgsForCall []struct {
		arg1 alias1.CubeSide
		arg2 int
		arg3 int
	}
	texelReturns struct {
		result1 alias1.RGBAColor
	}
}

var _ alias1.CubeTexture = new(CubeTextureStub)

func (stub *CubeTextureStub) SetTexel(arg1 alias1.CubeSide, arg2 int, arg3 int, arg4 alias1.RGBAColor) {
	stub.setTexelMutex.Lock()
	defer stub.setTexelMutex.Unlock()
	stub.setTexelArgsForCall = append(stub.setTexelArgsForCall, struct {
		arg1 alias1.CubeSide
		arg2 int
		arg3 int
		arg4 alias1.RGBAColor
	}{arg1, arg2, arg3, arg4})
	if stub.SetTexelStub != nil {
		stub.SetTexelStub(arg1, arg2, arg3, arg4)
	}
}
func (stub *CubeTextureStub) SetTexelCallCount() int {
	stub.setTexelMutex.RLock()
	defer stub.setTexelMutex.RUnlock()
	return len(stub.setTexelArgsForCall)
}
func (stub *CubeTextureStub) SetTexelArgsForCall(index int) (alias1.CubeSide, int, int, alias1.RGBAColor) {
	stub.setTexelMutex.RLock()
	defer stub.setTexelMutex.RUnlock()
	return stub.setTexelArgsForCall[index].arg1, stub.setTexelArgsForCall[index].arg2, stub.setTexelArgsForCall[index].arg3, stub.setTexelArgsForCall[index].arg4
}
func (stub *CubeTextureStub) Texel(arg1 alias1.CubeSide, arg2 int, arg3 int) alias1.RGBAColor {
	stub.texelMutex.Lock()
	defer stub.texelMutex.Unlock()
	stub.texelArgsForCall = append(stub.texelArgsForCall, struct {
		arg1 alias1.CubeSide
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	if stub.TexelStub != nil {
		return stub.TexelStub(arg1, arg2, arg3)
	} else {
		return stub.texelReturns.result1
	}
}
func (stub *CubeTextureStub) TexelCallCount() int {
	stub.texelMutex.RLock()
	defer stub.texelMutex.RUnlock()
	return len(stub.texelArgsForCall)
}
func (stub *CubeTextureStub) TexelArgsForCall(index int) (alias1.CubeSide, int, int) {
	stub.texelMutex.RLock()
	defer stub.texelMutex.RUnlock()
	return stub.texelArgsForCall[index].arg1, stub.texelArgsForCall[index].arg2, stub.texelArgsForCall[index].arg3
}
func (stub *CubeTextureStub) TexelReturns(result1 alias1.RGBAColor) {
	stub.texelMutex.Lock()
	defer stub.texelMutex.Unlock()
	stub.texelReturns = struct {
		result1 alias1.RGBAColor
	}{result1}
}
