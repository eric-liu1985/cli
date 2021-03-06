// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeDeleteOrphanedRoutesActor struct {
	DeleteRouteStub        func(string) (v2action.Warnings, error)
	deleteRouteMutex       sync.RWMutex
	deleteRouteArgsForCall []struct {
		arg1 string
	}
	deleteRouteReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	deleteRouteReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	GetOrphanedRoutesBySpaceStub        func(string) ([]v2action.Route, v2action.Warnings, error)
	getOrphanedRoutesBySpaceMutex       sync.RWMutex
	getOrphanedRoutesBySpaceArgsForCall []struct {
		arg1 string
	}
	getOrphanedRoutesBySpaceReturns struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	getOrphanedRoutesBySpaceReturnsOnCall map[int]struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRoute(arg1 string) (v2action.Warnings, error) {
	fake.deleteRouteMutex.Lock()
	ret, specificReturn := fake.deleteRouteReturnsOnCall[len(fake.deleteRouteArgsForCall)]
	fake.deleteRouteArgsForCall = append(fake.deleteRouteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteRoute", []interface{}{arg1})
	fake.deleteRouteMutex.Unlock()
	if fake.DeleteRouteStub != nil {
		return fake.DeleteRouteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteRouteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteCallCount() int {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return len(fake.deleteRouteArgsForCall)
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteCalls(stub func(string) (v2action.Warnings, error)) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = stub
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteArgsForCall(i int) string {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	argsForCall := fake.deleteRouteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteReturns(result1 v2action.Warnings, result2 error) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = nil
	fake.deleteRouteReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) DeleteRouteReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = nil
	if fake.deleteRouteReturnsOnCall == nil {
		fake.deleteRouteReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.deleteRouteReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpace(arg1 string) ([]v2action.Route, v2action.Warnings, error) {
	fake.getOrphanedRoutesBySpaceMutex.Lock()
	ret, specificReturn := fake.getOrphanedRoutesBySpaceReturnsOnCall[len(fake.getOrphanedRoutesBySpaceArgsForCall)]
	fake.getOrphanedRoutesBySpaceArgsForCall = append(fake.getOrphanedRoutesBySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrphanedRoutesBySpace", []interface{}{arg1})
	fake.getOrphanedRoutesBySpaceMutex.Unlock()
	if fake.GetOrphanedRoutesBySpaceStub != nil {
		return fake.GetOrphanedRoutesBySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrphanedRoutesBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceCallCount() int {
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	return len(fake.getOrphanedRoutesBySpaceArgsForCall)
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceCalls(stub func(string) ([]v2action.Route, v2action.Warnings, error)) {
	fake.getOrphanedRoutesBySpaceMutex.Lock()
	defer fake.getOrphanedRoutesBySpaceMutex.Unlock()
	fake.GetOrphanedRoutesBySpaceStub = stub
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceArgsForCall(i int) string {
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	argsForCall := fake.getOrphanedRoutesBySpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceReturns(result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.getOrphanedRoutesBySpaceMutex.Lock()
	defer fake.getOrphanedRoutesBySpaceMutex.Unlock()
	fake.GetOrphanedRoutesBySpaceStub = nil
	fake.getOrphanedRoutesBySpaceReturns = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDeleteOrphanedRoutesActor) GetOrphanedRoutesBySpaceReturnsOnCall(i int, result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.getOrphanedRoutesBySpaceMutex.Lock()
	defer fake.getOrphanedRoutesBySpaceMutex.Unlock()
	fake.GetOrphanedRoutesBySpaceStub = nil
	if fake.getOrphanedRoutesBySpaceReturnsOnCall == nil {
		fake.getOrphanedRoutesBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrphanedRoutesBySpaceReturnsOnCall[i] = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDeleteOrphanedRoutesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	fake.getOrphanedRoutesBySpaceMutex.RLock()
	defer fake.getOrphanedRoutesBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrphanedRoutesActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.DeleteOrphanedRoutesActor = new(FakeDeleteOrphanedRoutesActor)
