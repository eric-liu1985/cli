// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeV3PushVersionActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	PollStartStub        func(appGUID string, warningsChannel chan<- v3action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		appGUID         string
		warningsChannel chan<- v3action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	RestartApplicationStub        func(appGUID string) (v3action.Warnings, error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		appGUID string
	}
	restartApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3PushVersionActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3PushVersionActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3PushVersionActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PushVersionActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}{appName, spaceGUID, client})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{appName, spaceGUID, client})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(appName, spaceGUID, client)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result1, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result2, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result3, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result4
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v3action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].appName, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].spaceGUID, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].client
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3PushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v3action.LogMessage
			result2 <-chan error
			result3 v3action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3PushVersionActor) PollStart(appGUID string, warningsChannel chan<- v3action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		appGUID         string
		warningsChannel chan<- v3action.Warnings
	}{appGUID, warningsChannel})
	fake.recordInvocation("PollStart", []interface{}{appGUID, warningsChannel})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(appGUID, warningsChannel)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pollStartReturns.result1
}

func (fake *FakeV3PushVersionActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeV3PushVersionActor) PollStartArgsForCall(i int) (string, chan<- v3action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return fake.pollStartArgsForCall[i].appGUID, fake.pollStartArgsForCall[i].warningsChannel
}

func (fake *FakeV3PushVersionActor) PollStartReturns(result1 error) {
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushVersionActor) PollStartReturnsOnCall(i int, result1 error) {
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeV3PushVersionActor) RestartApplication(appGUID string) (v3action.Warnings, error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("RestartApplication", []interface{}{appGUID})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.restartApplicationReturns.result1, fake.restartApplicationReturns.result2
}

func (fake *FakeV3PushVersionActor) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakeV3PushVersionActor) RestartApplicationArgsForCall(i int) string {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return fake.restartApplicationArgsForCall[i].appGUID
}

func (fake *FakeV3PushVersionActor) RestartApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushVersionActor) RestartApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3PushVersionActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3PushVersionActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3PushVersionActor = new(FakeV3PushVersionActor)