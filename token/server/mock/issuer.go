// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/protos/token"
	"github.com/hyperledger/fabric/token/server"
)

type Issuer struct {
	RequestIssueStub        func(tokensToIssue []*token.Token) (*token.TokenTransaction, error)
	requestIssueMutex       sync.RWMutex
	requestIssueArgsForCall []struct {
		tokensToIssue []*token.Token
	}
	requestIssueReturns struct {
		result1 *token.TokenTransaction
		result2 error
	}
	requestIssueReturnsOnCall map[int]struct {
		result1 *token.TokenTransaction
		result2 error
	}
	RequestExpectationStub        func(request *token.ExpectationRequest) (*token.TokenTransaction, error)
	requestExpectationMutex       sync.RWMutex
	requestExpectationArgsForCall []struct {
		request *token.ExpectationRequest
	}
	requestExpectationReturns struct {
		result1 *token.TokenTransaction
		result2 error
	}
	requestExpectationReturnsOnCall map[int]struct {
		result1 *token.TokenTransaction
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Issuer) RequestIssue(tokensToIssue []*token.Token) (*token.TokenTransaction, error) {
	var tokensToIssueCopy []*token.Token
	if tokensToIssue != nil {
		tokensToIssueCopy = make([]*token.Token, len(tokensToIssue))
		copy(tokensToIssueCopy, tokensToIssue)
	}
	fake.requestIssueMutex.Lock()
	ret, specificReturn := fake.requestIssueReturnsOnCall[len(fake.requestIssueArgsForCall)]
	fake.requestIssueArgsForCall = append(fake.requestIssueArgsForCall, struct {
		tokensToIssue []*token.Token
	}{tokensToIssueCopy})
	fake.recordInvocation("RequestIssue", []interface{}{tokensToIssueCopy})
	fake.requestIssueMutex.Unlock()
	if fake.RequestIssueStub != nil {
		return fake.RequestIssueStub(tokensToIssue)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.requestIssueReturns.result1, fake.requestIssueReturns.result2
}

func (fake *Issuer) RequestIssueCallCount() int {
	fake.requestIssueMutex.RLock()
	defer fake.requestIssueMutex.RUnlock()
	return len(fake.requestIssueArgsForCall)
}

func (fake *Issuer) RequestIssueArgsForCall(i int) []*token.Token {
	fake.requestIssueMutex.RLock()
	defer fake.requestIssueMutex.RUnlock()
	return fake.requestIssueArgsForCall[i].tokensToIssue
}

func (fake *Issuer) RequestIssueReturns(result1 *token.TokenTransaction, result2 error) {
	fake.RequestIssueStub = nil
	fake.requestIssueReturns = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Issuer) RequestIssueReturnsOnCall(i int, result1 *token.TokenTransaction, result2 error) {
	fake.RequestIssueStub = nil
	if fake.requestIssueReturnsOnCall == nil {
		fake.requestIssueReturnsOnCall = make(map[int]struct {
			result1 *token.TokenTransaction
			result2 error
		})
	}
	fake.requestIssueReturnsOnCall[i] = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Issuer) RequestExpectation(request *token.ExpectationRequest) (*token.TokenTransaction, error) {
	fake.requestExpectationMutex.Lock()
	ret, specificReturn := fake.requestExpectationReturnsOnCall[len(fake.requestExpectationArgsForCall)]
	fake.requestExpectationArgsForCall = append(fake.requestExpectationArgsForCall, struct {
		request *token.ExpectationRequest
	}{request})
	fake.recordInvocation("RequestExpectation", []interface{}{request})
	fake.requestExpectationMutex.Unlock()
	if fake.RequestExpectationStub != nil {
		return fake.RequestExpectationStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.requestExpectationReturns.result1, fake.requestExpectationReturns.result2
}

func (fake *Issuer) RequestExpectationCallCount() int {
	fake.requestExpectationMutex.RLock()
	defer fake.requestExpectationMutex.RUnlock()
	return len(fake.requestExpectationArgsForCall)
}

func (fake *Issuer) RequestExpectationArgsForCall(i int) *token.ExpectationRequest {
	fake.requestExpectationMutex.RLock()
	defer fake.requestExpectationMutex.RUnlock()
	return fake.requestExpectationArgsForCall[i].request
}

func (fake *Issuer) RequestExpectationReturns(result1 *token.TokenTransaction, result2 error) {
	fake.RequestExpectationStub = nil
	fake.requestExpectationReturns = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Issuer) RequestExpectationReturnsOnCall(i int, result1 *token.TokenTransaction, result2 error) {
	fake.RequestExpectationStub = nil
	if fake.requestExpectationReturnsOnCall == nil {
		fake.requestExpectationReturnsOnCall = make(map[int]struct {
			result1 *token.TokenTransaction
			result2 error
		})
	}
	fake.requestExpectationReturnsOnCall[i] = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Issuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.requestIssueMutex.RLock()
	defer fake.requestIssueMutex.RUnlock()
	fake.requestExpectationMutex.RLock()
	defer fake.requestExpectationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Issuer) recordInvocation(key string, args []interface{}) {
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

var _ server.Issuer = new(Issuer)
