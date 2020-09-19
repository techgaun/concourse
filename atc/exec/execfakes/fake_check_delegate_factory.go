// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/exec"
)

type FakeCheckDelegateFactory struct {
	CheckDelegateStub        func() exec.CheckDelegate
	checkDelegateMutex       sync.RWMutex
	checkDelegateArgsForCall []struct {
	}
	checkDelegateReturns struct {
		result1 exec.CheckDelegate
	}
	checkDelegateReturnsOnCall map[int]struct {
		result1 exec.CheckDelegate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckDelegateFactory) CheckDelegate() exec.CheckDelegate {
	fake.checkDelegateMutex.Lock()
	ret, specificReturn := fake.checkDelegateReturnsOnCall[len(fake.checkDelegateArgsForCall)]
	fake.checkDelegateArgsForCall = append(fake.checkDelegateArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckDelegate", []interface{}{})
	fake.checkDelegateMutex.Unlock()
	if fake.CheckDelegateStub != nil {
		return fake.CheckDelegateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkDelegateReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegateFactory) CheckDelegateCallCount() int {
	fake.checkDelegateMutex.RLock()
	defer fake.checkDelegateMutex.RUnlock()
	return len(fake.checkDelegateArgsForCall)
}

func (fake *FakeCheckDelegateFactory) CheckDelegateCalls(stub func() exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = stub
}

func (fake *FakeCheckDelegateFactory) CheckDelegateReturns(result1 exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = nil
	fake.checkDelegateReturns = struct {
		result1 exec.CheckDelegate
	}{result1}
}

func (fake *FakeCheckDelegateFactory) CheckDelegateReturnsOnCall(i int, result1 exec.CheckDelegate) {
	fake.checkDelegateMutex.Lock()
	defer fake.checkDelegateMutex.Unlock()
	fake.CheckDelegateStub = nil
	if fake.checkDelegateReturnsOnCall == nil {
		fake.checkDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.CheckDelegate
		})
	}
	fake.checkDelegateReturnsOnCall[i] = struct {
		result1 exec.CheckDelegate
	}{result1}
}

func (fake *FakeCheckDelegateFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkDelegateMutex.RLock()
	defer fake.checkDelegateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckDelegateFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.CheckDelegateFactory = new(FakeCheckDelegateFactory)