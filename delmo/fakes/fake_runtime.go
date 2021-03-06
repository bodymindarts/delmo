// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/bodymindarts/delmo/delmo"
)

type FakeRuntime struct {
	StartAllStub        func(delmo.TestOutput) error
	startAllMutex       sync.RWMutex
	startAllArgsForCall []struct {
		arg1 delmo.TestOutput
	}
	startAllReturns struct {
		result1 error
	}
	StopAllStub        func(delmo.TestOutput) error
	stopAllMutex       sync.RWMutex
	stopAllArgsForCall []struct {
		arg1 delmo.TestOutput
	}
	stopAllReturns struct {
		result1 error
	}
	StopServicesStub        func(delmo.TestOutput, ...string) error
	stopServicesMutex       sync.RWMutex
	stopServicesArgsForCall []struct {
		arg1 delmo.TestOutput
		arg2 []string
	}
	stopServicesReturns struct {
		result1 error
	}
	StartServicesStub        func(delmo.TestOutput, ...string) error
	startServicesMutex       sync.RWMutex
	startServicesArgsForCall []struct {
		arg1 delmo.TestOutput
		arg2 []string
	}
	startServicesReturns struct {
		result1 error
	}
	DestroyServicesStub        func(delmo.TestOutput, ...string) error
	destroyServicesMutex       sync.RWMutex
	destroyServicesArgsForCall []struct {
		arg1 delmo.TestOutput
		arg2 []string
	}
	destroyServicesReturns struct {
		result1 error
	}
	SystemOutputStub        func() ([]byte, error)
	systemOutputMutex       sync.RWMutex
	systemOutputArgsForCall []struct{}
	systemOutputReturns     struct {
		result1 []byte
		result2 error
	}
	ExecuteTaskStub        func(string, delmo.TaskConfig, delmo.TaskEnvironment, delmo.TestOutput) error
	executeTaskMutex       sync.RWMutex
	executeTaskArgsForCall []struct {
		arg1 string
		arg2 delmo.TaskConfig
		arg3 delmo.TaskEnvironment
		arg4 delmo.TestOutput
	}
	executeTaskReturns struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuntime) StartAll(arg1 delmo.TestOutput) error {
	fake.startAllMutex.Lock()
	fake.startAllArgsForCall = append(fake.startAllArgsForCall, struct {
		arg1 delmo.TestOutput
	}{arg1})
	fake.recordInvocation("StartAll", []interface{}{arg1})
	fake.startAllMutex.Unlock()
	if fake.StartAllStub != nil {
		return fake.StartAllStub(arg1)
	} else {
		return fake.startAllReturns.result1
	}
}

func (fake *FakeRuntime) StartAllCallCount() int {
	fake.startAllMutex.RLock()
	defer fake.startAllMutex.RUnlock()
	return len(fake.startAllArgsForCall)
}

func (fake *FakeRuntime) StartAllArgsForCall(i int) delmo.TestOutput {
	fake.startAllMutex.RLock()
	defer fake.startAllMutex.RUnlock()
	return fake.startAllArgsForCall[i].arg1
}

func (fake *FakeRuntime) StartAllReturns(result1 error) {
	fake.StartAllStub = nil
	fake.startAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) StopAll(arg1 delmo.TestOutput) error {
	fake.stopAllMutex.Lock()
	fake.stopAllArgsForCall = append(fake.stopAllArgsForCall, struct {
		arg1 delmo.TestOutput
	}{arg1})
	fake.recordInvocation("StopAll", []interface{}{arg1})
	fake.stopAllMutex.Unlock()
	if fake.StopAllStub != nil {
		return fake.StopAllStub(arg1)
	} else {
		return fake.stopAllReturns.result1
	}
}

func (fake *FakeRuntime) StopAllCallCount() int {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return len(fake.stopAllArgsForCall)
}

func (fake *FakeRuntime) StopAllArgsForCall(i int) delmo.TestOutput {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return fake.stopAllArgsForCall[i].arg1
}

func (fake *FakeRuntime) StopAllReturns(result1 error) {
	fake.StopAllStub = nil
	fake.stopAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) StopServices(arg1 delmo.TestOutput, arg2 ...string) error {
	fake.stopServicesMutex.Lock()
	fake.stopServicesArgsForCall = append(fake.stopServicesArgsForCall, struct {
		arg1 delmo.TestOutput
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("StopServices", []interface{}{arg1, arg2})
	fake.stopServicesMutex.Unlock()
	if fake.StopServicesStub != nil {
		return fake.StopServicesStub(arg1, arg2...)
	} else {
		return fake.stopServicesReturns.result1
	}
}

func (fake *FakeRuntime) StopServicesCallCount() int {
	fake.stopServicesMutex.RLock()
	defer fake.stopServicesMutex.RUnlock()
	return len(fake.stopServicesArgsForCall)
}

func (fake *FakeRuntime) StopServicesArgsForCall(i int) (delmo.TestOutput, []string) {
	fake.stopServicesMutex.RLock()
	defer fake.stopServicesMutex.RUnlock()
	return fake.stopServicesArgsForCall[i].arg1, fake.stopServicesArgsForCall[i].arg2
}

func (fake *FakeRuntime) StopServicesReturns(result1 error) {
	fake.StopServicesStub = nil
	fake.stopServicesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) StartServices(arg1 delmo.TestOutput, arg2 ...string) error {
	fake.startServicesMutex.Lock()
	fake.startServicesArgsForCall = append(fake.startServicesArgsForCall, struct {
		arg1 delmo.TestOutput
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("StartServices", []interface{}{arg1, arg2})
	fake.startServicesMutex.Unlock()
	if fake.StartServicesStub != nil {
		return fake.StartServicesStub(arg1, arg2...)
	} else {
		return fake.startServicesReturns.result1
	}
}

func (fake *FakeRuntime) StartServicesCallCount() int {
	fake.startServicesMutex.RLock()
	defer fake.startServicesMutex.RUnlock()
	return len(fake.startServicesArgsForCall)
}

func (fake *FakeRuntime) StartServicesArgsForCall(i int) (delmo.TestOutput, []string) {
	fake.startServicesMutex.RLock()
	defer fake.startServicesMutex.RUnlock()
	return fake.startServicesArgsForCall[i].arg1, fake.startServicesArgsForCall[i].arg2
}

func (fake *FakeRuntime) StartServicesReturns(result1 error) {
	fake.StartServicesStub = nil
	fake.startServicesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) DestroyServices(arg1 delmo.TestOutput, arg2 ...string) error {
	fake.destroyServicesMutex.Lock()
	fake.destroyServicesArgsForCall = append(fake.destroyServicesArgsForCall, struct {
		arg1 delmo.TestOutput
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("DestroyServices", []interface{}{arg1, arg2})
	fake.destroyServicesMutex.Unlock()
	if fake.DestroyServicesStub != nil {
		return fake.DestroyServicesStub(arg1, arg2...)
	} else {
		return fake.destroyServicesReturns.result1
	}
}

func (fake *FakeRuntime) DestroyServicesCallCount() int {
	fake.destroyServicesMutex.RLock()
	defer fake.destroyServicesMutex.RUnlock()
	return len(fake.destroyServicesArgsForCall)
}

func (fake *FakeRuntime) DestroyServicesArgsForCall(i int) (delmo.TestOutput, []string) {
	fake.destroyServicesMutex.RLock()
	defer fake.destroyServicesMutex.RUnlock()
	return fake.destroyServicesArgsForCall[i].arg1, fake.destroyServicesArgsForCall[i].arg2
}

func (fake *FakeRuntime) DestroyServicesReturns(result1 error) {
	fake.DestroyServicesStub = nil
	fake.destroyServicesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) SystemOutput() ([]byte, error) {
	fake.systemOutputMutex.Lock()
	fake.systemOutputArgsForCall = append(fake.systemOutputArgsForCall, struct{}{})
	fake.recordInvocation("SystemOutput", []interface{}{})
	fake.systemOutputMutex.Unlock()
	if fake.SystemOutputStub != nil {
		return fake.SystemOutputStub()
	} else {
		return fake.systemOutputReturns.result1, fake.systemOutputReturns.result2
	}
}

func (fake *FakeRuntime) SystemOutputCallCount() int {
	fake.systemOutputMutex.RLock()
	defer fake.systemOutputMutex.RUnlock()
	return len(fake.systemOutputArgsForCall)
}

func (fake *FakeRuntime) SystemOutputReturns(result1 []byte, result2 error) {
	fake.SystemOutputStub = nil
	fake.systemOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntime) ExecuteTask(arg1 string, arg2 delmo.TaskConfig, arg3 delmo.TaskEnvironment, arg4 delmo.TestOutput) error {
	fake.executeTaskMutex.Lock()
	fake.executeTaskArgsForCall = append(fake.executeTaskArgsForCall, struct {
		arg1 string
		arg2 delmo.TaskConfig
		arg3 delmo.TaskEnvironment
		arg4 delmo.TestOutput
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ExecuteTask", []interface{}{arg1, arg2, arg3, arg4})
	fake.executeTaskMutex.Unlock()
	if fake.ExecuteTaskStub != nil {
		return fake.ExecuteTaskStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.executeTaskReturns.result1
	}
}

func (fake *FakeRuntime) ExecuteTaskCallCount() int {
	fake.executeTaskMutex.RLock()
	defer fake.executeTaskMutex.RUnlock()
	return len(fake.executeTaskArgsForCall)
}

func (fake *FakeRuntime) ExecuteTaskArgsForCall(i int) (string, delmo.TaskConfig, delmo.TaskEnvironment, delmo.TestOutput) {
	fake.executeTaskMutex.RLock()
	defer fake.executeTaskMutex.RUnlock()
	return fake.executeTaskArgsForCall[i].arg1, fake.executeTaskArgsForCall[i].arg2, fake.executeTaskArgsForCall[i].arg3, fake.executeTaskArgsForCall[i].arg4
}

func (fake *FakeRuntime) ExecuteTaskReturns(result1 error) {
	fake.ExecuteTaskStub = nil
	fake.executeTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	} else {
		return fake.cleanupReturns.result1
	}
}

func (fake *FakeRuntime) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeRuntime) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startAllMutex.RLock()
	defer fake.startAllMutex.RUnlock()
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	fake.stopServicesMutex.RLock()
	defer fake.stopServicesMutex.RUnlock()
	fake.startServicesMutex.RLock()
	defer fake.startServicesMutex.RUnlock()
	fake.destroyServicesMutex.RLock()
	defer fake.destroyServicesMutex.RUnlock()
	fake.systemOutputMutex.RLock()
	defer fake.systemOutputMutex.RUnlock()
	fake.executeTaskMutex.RLock()
	defer fake.executeTaskMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRuntime) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ delmo.Runtime = new(FakeRuntime)
