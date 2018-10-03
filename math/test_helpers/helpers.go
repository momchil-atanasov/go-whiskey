package test_helpers

import (
	"fmt"

	gomath "math"

	"github.com/mokiat/go-whiskey/math"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

const FloatMargin = 0.0001

func EqualMat4x4(expectedValue math.Mat4x4) types.GomegaMatcher {
	return QuickMatcher(func(actual interface{}) (MatchStatus, error) {
		matrix, ok := actual.(math.Mat4x4)
		if !ok {
			return MatchStatus{}, fmt.Errorf("EqualMat4x4 matcher expects a math.Mat4x4")
		}
		areEqual := areEqualFloat32(matrix.M11, expectedValue.M11, FloatMargin) &&
			areEqualFloat32(matrix.M21, expectedValue.M21, FloatMargin) &&
			areEqualFloat32(matrix.M31, expectedValue.M31, FloatMargin) &&
			areEqualFloat32(matrix.M41, expectedValue.M41, FloatMargin) &&
			areEqualFloat32(matrix.M12, expectedValue.M12, FloatMargin) &&
			areEqualFloat32(matrix.M22, expectedValue.M22, FloatMargin) &&
			areEqualFloat32(matrix.M32, expectedValue.M32, FloatMargin) &&
			areEqualFloat32(matrix.M42, expectedValue.M42, FloatMargin) &&
			areEqualFloat32(matrix.M13, expectedValue.M13, FloatMargin) &&
			areEqualFloat32(matrix.M23, expectedValue.M23, FloatMargin) &&
			areEqualFloat32(matrix.M33, expectedValue.M33, FloatMargin) &&
			areEqualFloat32(matrix.M43, expectedValue.M43, FloatMargin) &&
			areEqualFloat32(matrix.M14, expectedValue.M14, FloatMargin) &&
			areEqualFloat32(matrix.M24, expectedValue.M24, FloatMargin) &&
			areEqualFloat32(matrix.M34, expectedValue.M34, FloatMargin) &&
			areEqualFloat32(matrix.M44, expectedValue.M44, FloatMargin)

		if !areEqual {
			return FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto equal\n\t%#v", matrix, expectedValue),
				fmt.Sprintf("Expected\n\t%#v\nnot to equal\n\t%#v", matrix, expectedValue),
			), nil
		}
		return SuccessMatchStatus(), nil
	})
}

func AssertFloatEquals(actualValue, expectedValue float32) {
	Ω(actualValue).Should(BeNumerically("~", expectedValue, FloatMargin))
}

func AssertVec2Equals(vector math.Vec2, expectedX, expectedY float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
}

func AssertVec3Equals(vector math.Vec3, expectedX, expectedY, expectedZ float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
	AssertFloatEquals(vector.Z, expectedZ)
}

func AssertVec4Equals(vector math.Vec4, expectedX, expectedY, expectedZ, expectedW float32) {
	AssertFloatEquals(vector.X, expectedX)
	AssertFloatEquals(vector.Y, expectedY)
	AssertFloatEquals(vector.Z, expectedZ)
	AssertFloatEquals(vector.W, expectedW)
}

func AssertMat4x4Equals(matrix math.Mat4x4,
	m11, m12, m13, m14,
	m21, m22, m23, m24,
	m31, m32, m33, m34,
	m41, m42, m43, m44 float32) {

	AssertFloatEquals(matrix.M11, m11)
	AssertFloatEquals(matrix.M12, m12)
	AssertFloatEquals(matrix.M13, m13)
	AssertFloatEquals(matrix.M14, m14)

	AssertFloatEquals(matrix.M21, m21)
	AssertFloatEquals(matrix.M22, m22)
	AssertFloatEquals(matrix.M23, m23)
	AssertFloatEquals(matrix.M24, m24)

	AssertFloatEquals(matrix.M31, m31)
	AssertFloatEquals(matrix.M32, m32)
	AssertFloatEquals(matrix.M33, m33)
	AssertFloatEquals(matrix.M34, m34)

	AssertFloatEquals(matrix.M41, m41)
	AssertFloatEquals(matrix.M42, m42)
	AssertFloatEquals(matrix.M43, m43)
	AssertFloatEquals(matrix.M44, m44)
}

type MatchStatus struct {
	Success                bool
	FailureMessage         string
	NegativeFailureMessage string
}

func SuccessMatchStatus() MatchStatus {
	return MatchStatus{
		Success: true,
	}
}

func FailureMatchStatus(failureMessage, negativeFailureMessage string) MatchStatus {
	return MatchStatus{
		Success:                false,
		FailureMessage:         failureMessage,
		NegativeFailureMessage: negativeFailureMessage,
	}
}

func areEqualFloat32(a, b float32, margin float64) bool {
	return gomath.Abs(float64(a)-float64(b)) < margin
}

type MatchHandler func(value interface{}) (MatchStatus, error)

func QuickMatcher(handler MatchHandler) types.GomegaMatcher {
	return &quickMatcher{
		handler: handler,
	}
}

type quickMatcher struct {
	handler MatchHandler
	status  MatchStatus
}

func (m *quickMatcher) Match(actual interface{}) (success bool, err error) {
	m.status, err = m.handler(actual)
	if err != nil {
		return false, err
	}
	return m.status.Success, nil
}

func (m *quickMatcher) FailureMessage(actual interface{}) (message string) {
	return m.status.FailureMessage
}

func (m *quickMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return m.status.NegativeFailureMessage
}
