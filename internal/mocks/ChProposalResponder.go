// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "perun.network/go-perun/client"

	mock "github.com/stretchr/testify/mock"

	session "github.com/hyperledger-labs/perun-node/session"
)

// ChProposalResponder is an autogenerated mock type for the ChProposalResponder type
type ChProposalResponder struct {
	mock.Mock
}

// Accept provides a mock function with given fields: _a0, _a1
func (_m *ChProposalResponder) Accept(_a0 context.Context, _a1 *client.LedgerChannelProposalAccMsg) (session.PChannel, error) {
	ret := _m.Called(_a0, _a1)

	var r0 session.PChannel
	if rf, ok := ret.Get(0).(func(context.Context, *client.LedgerChannelProposalAccMsg) session.PChannel); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.PChannel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.LedgerChannelProposalAccMsg) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reject provides a mock function with given fields: ctx, reason
func (_m *ChProposalResponder) Reject(ctx context.Context, reason string) error {
	ret := _m.Called(ctx, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChProposalResponder interface {
	mock.TestingT
	Cleanup(func())
}

// NewChProposalResponder creates a new instance of ChProposalResponder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChProposalResponder(t mockConstructorTestingTNewChProposalResponder) *ChProposalResponder {
	mock := &ChProposalResponder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
