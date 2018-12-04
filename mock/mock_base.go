package mock

import (
	"fmt"
	"gopkg.in/jucardi/go-logger-lib.v1/log"
)

type MockBase struct {
	times map[string]int
	when  map[string]WhenHandler
}

func New() *MockBase {
	return &MockBase{
		times: map[string]int{},
		when:  map[string]WhenHandler{},
	}
}

// When indicates what the expected behavior should be when a function is invoked.
func (m *MockBase) When(funcName string, f WhenHandler) {
	m.when[funcName] = func(args ...interface{}) []interface{} {
		log.Debug(funcName, " func invoked")
		m.times[funcName]++
		return f(args...)
	}
}

// WhenReturn allows to set the return args without the need of a WhenHandler.
func (m *MockBase) WhenReturn(funcName string, retArgs ...interface{}) {
	m.When(funcName, func(args ...interface{}) []interface{} {
		return retArgs
	})
}

// Times asserts that the amount of times a function was invoked matches the provided 'expected'.
func (m *MockBase) Times(funcName string) int {
	return m.times[funcName]
}

// BulkTimes same as 'Times', but verifies multiple function calls in one call.
func (m *MockBase) BulkTimes(names []string, expected []int) error {
	if len(expected) != len(names) {
		panic("The arrays used in BulkTime must be the same size.")
	}

	for i, v := range names {
		if expected[i] != m.Times(v) {
			return fmt.Errorf("field %s times mismatch, expected %d, got %d", v, expected[i], m.Times(v))
		}
	}
	return nil
}

func (m *MockBase) Invoke(name string, args ...interface{}) []interface{} {
	if f, ok := m.when[name]; ok {
		return f(args...)
	}
	m.times[name]++
	return nil
}

func (m *MockBase) ReturnSingleArg(name string, args ...interface{}) interface{} {
	ret := m.Invoke(name, args...)
	if len(ret) > 0 && ret[0] != nil {
		return ret[0]
	}
	return nil
}

func (m *MockBase) ReturnDoubleArg(name string, args ...interface{}) (interface{}, interface{}) {
	ret := m.Invoke(name, args...)

	if len(ret) < 2 {
		log.Warn("Expected 2 returns for '%s', found %d", name, len(ret))
		return nil, nil
	}

	return ret[0], ret[1]
}

func (m *MockBase) ReturnString(name string, args ...interface{}) string {
	if val, ok := m.ReturnSingleArg(name, args...).(string); ok {
		return val
	}
	return ""
}

func (m *MockBase) ReturnBool(name string, args ...interface{}) bool {
	if val, ok := m.ReturnSingleArg(name, args...).(bool); ok {
		return val
	}
	return false
}

func (m *MockBase) ReturnInt(name string, args ...interface{}) int {
	if val, ok := m.ReturnSingleArg(name, args...).(int); ok {
		return val
	}
	return 0
}

func (m *MockBase) ReturnError(name string, args ...interface{}) error {
	if val, ok := m.ReturnSingleArg(name, args...).(error); ok {
		return val
	}
	return nil
}

func (m *MockBase) ReturnSingleArgWithError(name string, args ...interface{}) (interface{}, error) {
	ret, err := m.ReturnDoubleArg(name, args...)

	if err != nil {
		return ret, err.(error)
	}

	return ret, nil
}
