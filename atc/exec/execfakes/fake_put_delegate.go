// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/vars"
)

type FakePutDelegate struct {
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FinishedStub        func(lager.Logger, exec.ExitStatus, runtime.VersionResult)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 runtime.VersionResult
	}
	ImageVersionDeterminedStub        func(db.UsedResourceCache) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 db.UsedResourceCache
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	imageVersionDeterminedReturnsOnCall map[int]struct {
		result1 error
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	RedactImageSourceStub        func(atc.Source) (atc.Source, error)
	redactImageSourceMutex       sync.RWMutex
	redactImageSourceArgsForCall []struct {
		arg1 atc.Source
	}
	redactImageSourceReturns struct {
		result1 atc.Source
		result2 error
	}
	redactImageSourceReturnsOnCall map[int]struct {
		result1 atc.Source
		result2 error
	}
	SaveOutputStub        func(lager.Logger, atc.PutPlan, atc.Source, atc.VersionedResourceTypes, runtime.VersionResult)
	saveOutputMutex       sync.RWMutex
	saveOutputArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 atc.Source
		arg4 atc.VersionedResourceTypes
		arg5 runtime.VersionResult
	}
	SelectedWorkerStub        func(lager.Logger, string)
	selectedWorkerMutex       sync.RWMutex
	selectedWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	VariablesStub        func() *vars.BuildVariables
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
	}
	variablesReturns struct {
		result1 *vars.BuildVariables
	}
	variablesReturnsOnCall map[int]struct {
		result1 *vars.BuildVariables
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePutDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if fake.ErroredStub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakePutDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakePutDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakePutDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePutDelegate) Finished(arg1 lager.Logger, arg2 exec.ExitStatus, arg3 runtime.VersionResult) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 runtime.VersionResult
	}{arg1, arg2, arg3})
	fake.recordInvocation("Finished", []interface{}{arg1, arg2, arg3})
	fake.finishedMutex.Unlock()
	if fake.FinishedStub != nil {
		fake.FinishedStub(arg1, arg2, arg3)
	}
}

func (fake *FakePutDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakePutDelegate) FinishedCalls(stub func(lager.Logger, exec.ExitStatus, runtime.VersionResult)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakePutDelegate) FinishedArgsForCall(i int) (lager.Logger, exec.ExitStatus, runtime.VersionResult) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePutDelegate) ImageVersionDetermined(arg1 db.UsedResourceCache) error {
	fake.imageVersionDeterminedMutex.Lock()
	ret, specificReturn := fake.imageVersionDeterminedReturnsOnCall[len(fake.imageVersionDeterminedArgsForCall)]
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 db.UsedResourceCache
	}{arg1})
	fake.recordInvocation("ImageVersionDetermined", []interface{}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.imageVersionDeterminedReturns
	return fakeReturns.result1
}

func (fake *FakePutDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakePutDelegate) ImageVersionDeterminedCalls(stub func(db.UsedResourceCache) error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = stub
}

func (fake *FakePutDelegate) ImageVersionDeterminedArgsForCall(i int) db.UsedResourceCache {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	argsForCall := fake.imageVersionDeterminedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePutDelegate) ImageVersionDeterminedReturnsOnCall(i int, result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	if fake.imageVersionDeterminedReturnsOnCall == nil {
		fake.imageVersionDeterminedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.imageVersionDeterminedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePutDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if fake.InitializingStub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakePutDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakePutDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakePutDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) RedactImageSource(arg1 atc.Source) (atc.Source, error) {
	fake.redactImageSourceMutex.Lock()
	ret, specificReturn := fake.redactImageSourceReturnsOnCall[len(fake.redactImageSourceArgsForCall)]
	fake.redactImageSourceArgsForCall = append(fake.redactImageSourceArgsForCall, struct {
		arg1 atc.Source
	}{arg1})
	fake.recordInvocation("RedactImageSource", []interface{}{arg1})
	fake.redactImageSourceMutex.Unlock()
	if fake.RedactImageSourceStub != nil {
		return fake.RedactImageSourceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.redactImageSourceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePutDelegate) RedactImageSourceCallCount() int {
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	return len(fake.redactImageSourceArgsForCall)
}

func (fake *FakePutDelegate) RedactImageSourceCalls(stub func(atc.Source) (atc.Source, error)) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = stub
}

func (fake *FakePutDelegate) RedactImageSourceArgsForCall(i int) atc.Source {
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	argsForCall := fake.redactImageSourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) RedactImageSourceReturns(result1 atc.Source, result2 error) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = nil
	fake.redactImageSourceReturns = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakePutDelegate) RedactImageSourceReturnsOnCall(i int, result1 atc.Source, result2 error) {
	fake.redactImageSourceMutex.Lock()
	defer fake.redactImageSourceMutex.Unlock()
	fake.RedactImageSourceStub = nil
	if fake.redactImageSourceReturnsOnCall == nil {
		fake.redactImageSourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
			result2 error
		})
	}
	fake.redactImageSourceReturnsOnCall[i] = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakePutDelegate) SaveOutput(arg1 lager.Logger, arg2 atc.PutPlan, arg3 atc.Source, arg4 atc.VersionedResourceTypes, arg5 runtime.VersionResult) {
	fake.saveOutputMutex.Lock()
	fake.saveOutputArgsForCall = append(fake.saveOutputArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 atc.Source
		arg4 atc.VersionedResourceTypes
		arg5 runtime.VersionResult
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("SaveOutput", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.saveOutputMutex.Unlock()
	if fake.SaveOutputStub != nil {
		fake.SaveOutputStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakePutDelegate) SaveOutputCallCount() int {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	return len(fake.saveOutputArgsForCall)
}

func (fake *FakePutDelegate) SaveOutputCalls(stub func(lager.Logger, atc.PutPlan, atc.Source, atc.VersionedResourceTypes, runtime.VersionResult)) {
	fake.saveOutputMutex.Lock()
	defer fake.saveOutputMutex.Unlock()
	fake.SaveOutputStub = stub
}

func (fake *FakePutDelegate) SaveOutputArgsForCall(i int) (lager.Logger, atc.PutPlan, atc.Source, atc.VersionedResourceTypes, runtime.VersionResult) {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	argsForCall := fake.saveOutputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakePutDelegate) SelectedWorker(arg1 lager.Logger, arg2 string) {
	fake.selectedWorkerMutex.Lock()
	fake.selectedWorkerArgsForCall = append(fake.selectedWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SelectedWorker", []interface{}{arg1, arg2})
	fake.selectedWorkerMutex.Unlock()
	if fake.SelectedWorkerStub != nil {
		fake.SelectedWorkerStub(arg1, arg2)
	}
}

func (fake *FakePutDelegate) SelectedWorkerCallCount() int {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	return len(fake.selectedWorkerArgsForCall)
}

func (fake *FakePutDelegate) SelectedWorkerCalls(stub func(lager.Logger, string)) {
	fake.selectedWorkerMutex.Lock()
	defer fake.selectedWorkerMutex.Unlock()
	fake.SelectedWorkerStub = stub
}

func (fake *FakePutDelegate) SelectedWorkerArgsForCall(i int) (lager.Logger, string) {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	argsForCall := fake.selectedWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePutDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if fake.StartingStub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakePutDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakePutDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakePutDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stderrReturns
	return fakeReturns.result1
}

func (fake *FakePutDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakePutDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakePutDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stdoutReturns
	return fakeReturns.result1
}

func (fake *FakePutDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakePutDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakePutDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) Variables() *vars.BuildVariables {
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
	}{})
	fake.recordInvocation("Variables", []interface{}{})
	fake.variablesMutex.Unlock()
	if fake.VariablesStub != nil {
		return fake.VariablesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.variablesReturns
	return fakeReturns.result1
}

func (fake *FakePutDelegate) VariablesCallCount() int {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return len(fake.variablesArgsForCall)
}

func (fake *FakePutDelegate) VariablesCalls(stub func() *vars.BuildVariables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = stub
}

func (fake *FakePutDelegate) VariablesReturns(result1 *vars.BuildVariables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	fake.variablesReturns = struct {
		result1 *vars.BuildVariables
	}{result1}
}

func (fake *FakePutDelegate) VariablesReturnsOnCall(i int, result1 *vars.BuildVariables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	if fake.variablesReturnsOnCall == nil {
		fake.variablesReturnsOnCall = make(map[int]struct {
			result1 *vars.BuildVariables
		})
	}
	fake.variablesReturnsOnCall[i] = struct {
		result1 *vars.BuildVariables
	}{result1}
}

func (fake *FakePutDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.redactImageSourceMutex.RLock()
	defer fake.redactImageSourceMutex.RUnlock()
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePutDelegate) recordInvocation(key string, args []interface{}) {
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

var _ exec.PutDelegate = new(FakePutDelegate)
