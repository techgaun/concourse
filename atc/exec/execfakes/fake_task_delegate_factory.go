// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/exec"
)

type FakeTaskDelegateFactory struct {
	TaskDelegateStub        func() exec.TaskDelegate
	taskDelegateMutex       sync.RWMutex
	taskDelegateArgsForCall []struct {
	}
	taskDelegateReturns struct {
		result1 exec.TaskDelegate
	}
	taskDelegateReturnsOnCall map[int]struct {
		result1 exec.TaskDelegate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskDelegateFactory) TaskDelegate() exec.TaskDelegate {
	fake.taskDelegateMutex.Lock()
	ret, specificReturn := fake.taskDelegateReturnsOnCall[len(fake.taskDelegateArgsForCall)]
	fake.taskDelegateArgsForCall = append(fake.taskDelegateArgsForCall, struct {
	}{})
	fake.recordInvocation("TaskDelegate", []interface{}{})
	fake.taskDelegateMutex.Unlock()
	if fake.TaskDelegateStub != nil {
		return fake.TaskDelegateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.taskDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeTaskDelegateFactory) TaskDelegateCallCount() int {
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	return len(fake.taskDelegateArgsForCall)
}

func (fake *FakeTaskDelegateFactory) TaskDelegateCalls(stub func() exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = stub
}

func (fake *FakeTaskDelegateFactory) TaskDelegateReturns(result1 exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = nil
	fake.taskDelegateReturns = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeTaskDelegateFactory) TaskDelegateReturnsOnCall(i int, result1 exec.TaskDelegate) {
	fake.taskDelegateMutex.Lock()
	defer fake.taskDelegateMutex.Unlock()
	fake.TaskDelegateStub = nil
	if fake.taskDelegateReturnsOnCall == nil {
		fake.taskDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.TaskDelegate
		})
	}
	fake.taskDelegateReturnsOnCall[i] = struct {
		result1 exec.TaskDelegate
	}{result1}
}

func (fake *FakeTaskDelegateFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.taskDelegateMutex.RLock()
	defer fake.taskDelegateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskDelegateFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.TaskDelegateFactory = new(FakeTaskDelegateFactory)