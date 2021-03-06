// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	sync "sync"
	time "time"

	lager "code.cloudfoundry.org/lager"
	atc "github.com/concourse/concourse/atc"
	db "github.com/concourse/concourse/atc/db"
	algorithm "github.com/concourse/concourse/atc/db/algorithm"
	scheduler "github.com/concourse/concourse/atc/scheduler"
)

type FakeBuildScheduler struct {
	ScheduleStub        func(lager.Logger, *algorithm.VersionsDB, []db.Job, db.Resources, atc.VersionedResourceTypes) (map[string]time.Duration, error)
	scheduleMutex       sync.RWMutex
	scheduleArgsForCall []struct {
		arg1 lager.Logger
		arg2 *algorithm.VersionsDB
		arg3 []db.Job
		arg4 db.Resources
		arg5 atc.VersionedResourceTypes
	}
	scheduleReturns struct {
		result1 map[string]time.Duration
		result2 error
	}
	scheduleReturnsOnCall map[int]struct {
		result1 map[string]time.Duration
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildScheduler) Schedule(arg1 lager.Logger, arg2 *algorithm.VersionsDB, arg3 []db.Job, arg4 db.Resources, arg5 atc.VersionedResourceTypes) (map[string]time.Duration, error) {
	var arg3Copy []db.Job
	if arg3 != nil {
		arg3Copy = make([]db.Job, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.scheduleMutex.Lock()
	ret, specificReturn := fake.scheduleReturnsOnCall[len(fake.scheduleArgsForCall)]
	fake.scheduleArgsForCall = append(fake.scheduleArgsForCall, struct {
		arg1 lager.Logger
		arg2 *algorithm.VersionsDB
		arg3 []db.Job
		arg4 db.Resources
		arg5 atc.VersionedResourceTypes
	}{arg1, arg2, arg3Copy, arg4, arg5})
	fake.recordInvocation("Schedule", []interface{}{arg1, arg2, arg3Copy, arg4, arg5})
	fake.scheduleMutex.Unlock()
	if fake.ScheduleStub != nil {
		return fake.ScheduleStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scheduleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildScheduler) ScheduleCallCount() int {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	return len(fake.scheduleArgsForCall)
}

func (fake *FakeBuildScheduler) ScheduleCalls(stub func(lager.Logger, *algorithm.VersionsDB, []db.Job, db.Resources, atc.VersionedResourceTypes) (map[string]time.Duration, error)) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = stub
}

func (fake *FakeBuildScheduler) ScheduleArgsForCall(i int) (lager.Logger, *algorithm.VersionsDB, []db.Job, db.Resources, atc.VersionedResourceTypes) {
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	argsForCall := fake.scheduleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeBuildScheduler) ScheduleReturns(result1 map[string]time.Duration, result2 error) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = nil
	fake.scheduleReturns = struct {
		result1 map[string]time.Duration
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildScheduler) ScheduleReturnsOnCall(i int, result1 map[string]time.Duration, result2 error) {
	fake.scheduleMutex.Lock()
	defer fake.scheduleMutex.Unlock()
	fake.ScheduleStub = nil
	if fake.scheduleReturnsOnCall == nil {
		fake.scheduleReturnsOnCall = make(map[int]struct {
			result1 map[string]time.Duration
			result2 error
		})
	}
	fake.scheduleReturnsOnCall[i] = struct {
		result1 map[string]time.Duration
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildScheduler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scheduleMutex.RLock()
	defer fake.scheduleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildScheduler) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.BuildScheduler = new(FakeBuildScheduler)
