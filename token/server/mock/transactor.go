// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/protos/token"
	"github.com/hyperledger/fabric/token/server"
)

type Transactor struct {
	DoneStub        func()
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
	}
	ListTokensStub        func() (*token.UnspentTokens, error)
	listTokensMutex       sync.RWMutex
	listTokensArgsForCall []struct {
	}
	listTokensReturns struct {
		result1 *token.UnspentTokens
		result2 error
	}
	listTokensReturnsOnCall map[int]struct {
		result1 *token.UnspentTokens
		result2 error
	}
	RequestRedeemStub        func(*token.RedeemRequest) (*token.TokenTransaction, error)
	requestRedeemMutex       sync.RWMutex
	requestRedeemArgsForCall []struct {
		arg1 *token.RedeemRequest
	}
	requestRedeemReturns struct {
		result1 *token.TokenTransaction
		result2 error
	}
	requestRedeemReturnsOnCall map[int]struct {
		result1 *token.TokenTransaction
		result2 error
	}
	RequestTokenOperationStub        func([]*token.TokenId, *token.TokenOperation) (*token.TokenTransaction, int, error)
	requestTokenOperationMutex       sync.RWMutex
	requestTokenOperationArgsForCall []struct {
		arg1 []*token.TokenId
		arg2 *token.TokenOperation
	}
	requestTokenOperationReturns struct {
		result1 *token.TokenTransaction
		result2 int
		result3 error
	}
	requestTokenOperationReturnsOnCall map[int]struct {
		result1 *token.TokenTransaction
		result2 int
		result3 error
	}
	RequestTransferStub        func(*token.TransferRequest) (*token.TokenTransaction, error)
	requestTransferMutex       sync.RWMutex
	requestTransferArgsForCall []struct {
		arg1 *token.TransferRequest
	}
	requestTransferReturns struct {
		result1 *token.TokenTransaction
		result2 error
	}
	requestTransferReturnsOnCall map[int]struct {
		result1 *token.TokenTransaction
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Transactor) Done() {
	fake.doneMutex.Lock()
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
	}{})
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		fake.DoneStub()
	}
}

func (fake *Transactor) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *Transactor) DoneCalls(stub func()) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *Transactor) ListTokens() (*token.UnspentTokens, error) {
	fake.listTokensMutex.Lock()
	ret, specificReturn := fake.listTokensReturnsOnCall[len(fake.listTokensArgsForCall)]
	fake.listTokensArgsForCall = append(fake.listTokensArgsForCall, struct {
	}{})
	fake.recordInvocation("ListTokens", []interface{}{})
	fake.listTokensMutex.Unlock()
	if fake.ListTokensStub != nil {
		return fake.ListTokensStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listTokensReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Transactor) ListTokensCallCount() int {
	fake.listTokensMutex.RLock()
	defer fake.listTokensMutex.RUnlock()
	return len(fake.listTokensArgsForCall)
}

func (fake *Transactor) ListTokensCalls(stub func() (*token.UnspentTokens, error)) {
	fake.listTokensMutex.Lock()
	defer fake.listTokensMutex.Unlock()
	fake.ListTokensStub = stub
}

func (fake *Transactor) ListTokensReturns(result1 *token.UnspentTokens, result2 error) {
	fake.listTokensMutex.Lock()
	defer fake.listTokensMutex.Unlock()
	fake.ListTokensStub = nil
	fake.listTokensReturns = struct {
		result1 *token.UnspentTokens
		result2 error
	}{result1, result2}
}

func (fake *Transactor) ListTokensReturnsOnCall(i int, result1 *token.UnspentTokens, result2 error) {
	fake.listTokensMutex.Lock()
	defer fake.listTokensMutex.Unlock()
	fake.ListTokensStub = nil
	if fake.listTokensReturnsOnCall == nil {
		fake.listTokensReturnsOnCall = make(map[int]struct {
			result1 *token.UnspentTokens
			result2 error
		})
	}
	fake.listTokensReturnsOnCall[i] = struct {
		result1 *token.UnspentTokens
		result2 error
	}{result1, result2}
}

func (fake *Transactor) RequestRedeem(arg1 *token.RedeemRequest) (*token.TokenTransaction, error) {
	fake.requestRedeemMutex.Lock()
	ret, specificReturn := fake.requestRedeemReturnsOnCall[len(fake.requestRedeemArgsForCall)]
	fake.requestRedeemArgsForCall = append(fake.requestRedeemArgsForCall, struct {
		arg1 *token.RedeemRequest
	}{arg1})
	fake.recordInvocation("RequestRedeem", []interface{}{arg1})
	fake.requestRedeemMutex.Unlock()
	if fake.RequestRedeemStub != nil {
		return fake.RequestRedeemStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestRedeemReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Transactor) RequestRedeemCallCount() int {
	fake.requestRedeemMutex.RLock()
	defer fake.requestRedeemMutex.RUnlock()
	return len(fake.requestRedeemArgsForCall)
}

func (fake *Transactor) RequestRedeemCalls(stub func(*token.RedeemRequest) (*token.TokenTransaction, error)) {
	fake.requestRedeemMutex.Lock()
	defer fake.requestRedeemMutex.Unlock()
	fake.RequestRedeemStub = stub
}

func (fake *Transactor) RequestRedeemArgsForCall(i int) *token.RedeemRequest {
	fake.requestRedeemMutex.RLock()
	defer fake.requestRedeemMutex.RUnlock()
	argsForCall := fake.requestRedeemArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Transactor) RequestRedeemReturns(result1 *token.TokenTransaction, result2 error) {
	fake.requestRedeemMutex.Lock()
	defer fake.requestRedeemMutex.Unlock()
	fake.RequestRedeemStub = nil
	fake.requestRedeemReturns = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Transactor) RequestRedeemReturnsOnCall(i int, result1 *token.TokenTransaction, result2 error) {
	fake.requestRedeemMutex.Lock()
	defer fake.requestRedeemMutex.Unlock()
	fake.RequestRedeemStub = nil
	if fake.requestRedeemReturnsOnCall == nil {
		fake.requestRedeemReturnsOnCall = make(map[int]struct {
			result1 *token.TokenTransaction
			result2 error
		})
	}
	fake.requestRedeemReturnsOnCall[i] = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Transactor) RequestTokenOperation(arg1 []*token.TokenId, arg2 *token.TokenOperation) (*token.TokenTransaction, int, error) {
	var arg1Copy []*token.TokenId
	if arg1 != nil {
		arg1Copy = make([]*token.TokenId, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.requestTokenOperationMutex.Lock()
	ret, specificReturn := fake.requestTokenOperationReturnsOnCall[len(fake.requestTokenOperationArgsForCall)]
	fake.requestTokenOperationArgsForCall = append(fake.requestTokenOperationArgsForCall, struct {
		arg1 []*token.TokenId
		arg2 *token.TokenOperation
	}{arg1Copy, arg2})
	fake.recordInvocation("RequestTokenOperation", []interface{}{arg1Copy, arg2})
	fake.requestTokenOperationMutex.Unlock()
	if fake.RequestTokenOperationStub != nil {
		return fake.RequestTokenOperationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.requestTokenOperationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *Transactor) RequestTokenOperationCallCount() int {
	fake.requestTokenOperationMutex.RLock()
	defer fake.requestTokenOperationMutex.RUnlock()
	return len(fake.requestTokenOperationArgsForCall)
}

func (fake *Transactor) RequestTokenOperationCalls(stub func([]*token.TokenId, *token.TokenOperation) (*token.TokenTransaction, int, error)) {
	fake.requestTokenOperationMutex.Lock()
	defer fake.requestTokenOperationMutex.Unlock()
	fake.RequestTokenOperationStub = stub
}

func (fake *Transactor) RequestTokenOperationArgsForCall(i int) ([]*token.TokenId, *token.TokenOperation) {
	fake.requestTokenOperationMutex.RLock()
	defer fake.requestTokenOperationMutex.RUnlock()
	argsForCall := fake.requestTokenOperationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Transactor) RequestTokenOperationReturns(result1 *token.TokenTransaction, result2 int, result3 error) {
	fake.requestTokenOperationMutex.Lock()
	defer fake.requestTokenOperationMutex.Unlock()
	fake.RequestTokenOperationStub = nil
	fake.requestTokenOperationReturns = struct {
		result1 *token.TokenTransaction
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *Transactor) RequestTokenOperationReturnsOnCall(i int, result1 *token.TokenTransaction, result2 int, result3 error) {
	fake.requestTokenOperationMutex.Lock()
	defer fake.requestTokenOperationMutex.Unlock()
	fake.RequestTokenOperationStub = nil
	if fake.requestTokenOperationReturnsOnCall == nil {
		fake.requestTokenOperationReturnsOnCall = make(map[int]struct {
			result1 *token.TokenTransaction
			result2 int
			result3 error
		})
	}
	fake.requestTokenOperationReturnsOnCall[i] = struct {
		result1 *token.TokenTransaction
		result2 int
		result3 error
	}{result1, result2, result3}
}

func (fake *Transactor) RequestTransfer(arg1 *token.TransferRequest) (*token.TokenTransaction, error) {
	fake.requestTransferMutex.Lock()
	ret, specificReturn := fake.requestTransferReturnsOnCall[len(fake.requestTransferArgsForCall)]
	fake.requestTransferArgsForCall = append(fake.requestTransferArgsForCall, struct {
		arg1 *token.TransferRequest
	}{arg1})
	fake.recordInvocation("RequestTransfer", []interface{}{arg1})
	fake.requestTransferMutex.Unlock()
	if fake.RequestTransferStub != nil {
		return fake.RequestTransferStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requestTransferReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Transactor) RequestTransferCallCount() int {
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	return len(fake.requestTransferArgsForCall)
}

func (fake *Transactor) RequestTransferCalls(stub func(*token.TransferRequest) (*token.TokenTransaction, error)) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = stub
}

func (fake *Transactor) RequestTransferArgsForCall(i int) *token.TransferRequest {
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	argsForCall := fake.requestTransferArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Transactor) RequestTransferReturns(result1 *token.TokenTransaction, result2 error) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = nil
	fake.requestTransferReturns = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Transactor) RequestTransferReturnsOnCall(i int, result1 *token.TokenTransaction, result2 error) {
	fake.requestTransferMutex.Lock()
	defer fake.requestTransferMutex.Unlock()
	fake.RequestTransferStub = nil
	if fake.requestTransferReturnsOnCall == nil {
		fake.requestTransferReturnsOnCall = make(map[int]struct {
			result1 *token.TokenTransaction
			result2 error
		})
	}
	fake.requestTransferReturnsOnCall[i] = struct {
		result1 *token.TokenTransaction
		result2 error
	}{result1, result2}
}

func (fake *Transactor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	fake.listTokensMutex.RLock()
	defer fake.listTokensMutex.RUnlock()
	fake.requestRedeemMutex.RLock()
	defer fake.requestRedeemMutex.RUnlock()
	fake.requestTokenOperationMutex.RLock()
	defer fake.requestTokenOperationMutex.RUnlock()
	fake.requestTransferMutex.RLock()
	defer fake.requestTransferMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Transactor) recordInvocation(key string, args []interface{}) {
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

var _ server.Transactor = new(Transactor)
