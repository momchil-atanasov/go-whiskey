package client

import "github.com/momchil-atanasov/go-whiskey/math"

type UniformLocation interface{}
type AttributeLocation interface{}

//go:generate counterfeiter -o client_fakes/fake_uniform.go ./ UniformClient

type UniformClient interface {
	BindVec4Uniform(UniformLocation, math.Vec4)
	BindMat4x4Uniform(UniformLocation, math.Mat4x4)
}
