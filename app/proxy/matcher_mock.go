// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package proxy

import (
	"sync"

	"github.com/umputun/reproxy/app/discovery"
)

// Ensure, that MatcherMock does implement Matcher.
// If this is not the case, regenerate this file with moq.
var _ Matcher = &MatcherMock{}

// MatcherMock is a mock implementation of Matcher.
//
// 	func TestSomethingThatUsesMatcher(t *testing.T) {
//
// 		// make and configure a mocked Matcher
// 		mockedMatcher := &MatcherMock{
// 			CheckHealthFunc: func() map[string]error {
// 				panic("mock out the CheckHealth method")
// 			},
// 			MappersFunc: func() []discovery.URLMapper {
// 				panic("mock out the Mappers method")
// 			},
// 			MatchFunc: func(srv string, src string) discovery.Matches {
// 				panic("mock out the Match method")
// 			},
// 			ServersFunc: func() []string {
// 				panic("mock out the Servers method")
// 			},
// 		}
//
// 		// use mockedMatcher in code that requires Matcher
// 		// and then make assertions.
//
// 	}
type MatcherMock struct {
	// CheckHealthFunc mocks the CheckHealth method.
	CheckHealthFunc func() map[string]error

	// MappersFunc mocks the Mappers method.
	MappersFunc func() []discovery.URLMapper

	// MatchFunc mocks the Match method.
	MatchFunc func(srv string, src string) discovery.Matches

	// ServersFunc mocks the Servers method.
	ServersFunc func() []string

	// calls tracks calls to the methods.
	calls struct {
		// CheckHealth holds details about calls to the CheckHealth method.
		CheckHealth []struct {
		}
		// Mappers holds details about calls to the Mappers method.
		Mappers []struct {
		}
		// Match holds details about calls to the Match method.
		Match []struct {
			// Srv is the srv argument value.
			Srv string
			// Src is the src argument value.
			Src string
		}
		// Servers holds details about calls to the Servers method.
		Servers []struct {
		}
	}
	lockCheckHealth sync.RWMutex
	lockMappers     sync.RWMutex
	lockMatch       sync.RWMutex
	lockServers     sync.RWMutex
}

// CheckHealth calls CheckHealthFunc.
func (mock *MatcherMock) CheckHealth() map[string]error {
	if mock.CheckHealthFunc == nil {
		panic("MatcherMock.CheckHealthFunc: method is nil but Matcher.CheckHealth was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCheckHealth.Lock()
	mock.calls.CheckHealth = append(mock.calls.CheckHealth, callInfo)
	mock.lockCheckHealth.Unlock()
	return mock.CheckHealthFunc()
}

// CheckHealthCalls gets all the calls that were made to CheckHealth.
// Check the length with:
//     len(mockedMatcher.CheckHealthCalls())
func (mock *MatcherMock) CheckHealthCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCheckHealth.RLock()
	calls = mock.calls.CheckHealth
	mock.lockCheckHealth.RUnlock()
	return calls
}

// Mappers calls MappersFunc.
func (mock *MatcherMock) Mappers() []discovery.URLMapper {
	if mock.MappersFunc == nil {
		panic("MatcherMock.MappersFunc: method is nil but Matcher.Mappers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockMappers.Lock()
	mock.calls.Mappers = append(mock.calls.Mappers, callInfo)
	mock.lockMappers.Unlock()
	return mock.MappersFunc()
}

// MappersCalls gets all the calls that were made to Mappers.
// Check the length with:
//     len(mockedMatcher.MappersCalls())
func (mock *MatcherMock) MappersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockMappers.RLock()
	calls = mock.calls.Mappers
	mock.lockMappers.RUnlock()
	return calls
}

// Match calls MatchFunc.
func (mock *MatcherMock) Match(srv string, src string) discovery.Matches {
	if mock.MatchFunc == nil {
		panic("MatcherMock.MatchFunc: method is nil but Matcher.Match was just called")
	}
	callInfo := struct {
		Srv string
		Src string
	}{
		Srv: srv,
		Src: src,
	}
	mock.lockMatch.Lock()
	mock.calls.Match = append(mock.calls.Match, callInfo)
	mock.lockMatch.Unlock()
	return mock.MatchFunc(srv, src)
}

// MatchCalls gets all the calls that were made to Match.
// Check the length with:
//     len(mockedMatcher.MatchCalls())
func (mock *MatcherMock) MatchCalls() []struct {
	Srv string
	Src string
} {
	var calls []struct {
		Srv string
		Src string
	}
	mock.lockMatch.RLock()
	calls = mock.calls.Match
	mock.lockMatch.RUnlock()
	return calls
}

// Servers calls ServersFunc.
func (mock *MatcherMock) Servers() []string {
	if mock.ServersFunc == nil {
		panic("MatcherMock.ServersFunc: method is nil but Matcher.Servers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockServers.Lock()
	mock.calls.Servers = append(mock.calls.Servers, callInfo)
	mock.lockServers.Unlock()
	return mock.ServersFunc()
}

// ServersCalls gets all the calls that were made to Servers.
// Check the length with:
//     len(mockedMatcher.ServersCalls())
func (mock *MatcherMock) ServersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockServers.RLock()
	calls = mock.calls.Servers
	mock.lockServers.RUnlock()
	return calls
}