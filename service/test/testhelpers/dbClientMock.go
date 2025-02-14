// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package testhelpers

import (
	"context"
	"github.com/cello-proj/cello/internal/types"
	"github.com/cello-proj/cello/service/internal/db"
	"sync"
)

// Ensure, that DBClientMock does implement db.Client.
// If this is not the case, regenerate this file with moq.
var _ db.Client = &DBClientMock{}

// DBClientMock is a mock implementation of db.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked db.Client
// 		mockedClient := &DBClientMock{
// 			CreateProjectEntryFunc: func(ctx context.Context, pe db.ProjectEntry) error {
// 				panic("mock out the CreateProjectEntry method")
// 			},
// 			CreateTokenEntryFunc: func(ctx context.Context, token types.Token) error {
// 				panic("mock out the CreateTokenEntry method")
// 			},
// 			DeleteProjectEntryFunc: func(ctx context.Context, project string) error {
// 				panic("mock out the DeleteProjectEntry method")
// 			},
// 			DeleteTokenEntryFunc: func(ctx context.Context, token string) error {
// 				panic("mock out the DeleteTokenEntry method")
// 			},
// 			ListTokenEntriesFunc: func(ctx context.Context, project string) ([]db.TokenEntry, error) {
// 				panic("mock out the ListTokenEntries method")
// 			},
// 			ReadProjectEntryFunc: func(ctx context.Context, project string) (db.ProjectEntry, error) {
// 				panic("mock out the ReadProjectEntry method")
// 			},
// 			ReadTokenEntryFunc: func(ctx context.Context, token string) (db.TokenEntry, error) {
// 				panic("mock out the ReadTokenEntry method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires db.Client
// 		// and then make assertions.
//
// 	}
type DBClientMock struct {
	// CreateProjectEntryFunc mocks the CreateProjectEntry method.
	CreateProjectEntryFunc func(ctx context.Context, pe db.ProjectEntry) error

	// CreateTokenEntryFunc mocks the CreateTokenEntry method.
	CreateTokenEntryFunc func(ctx context.Context, token types.Token) error

	// DeleteProjectEntryFunc mocks the DeleteProjectEntry method.
	DeleteProjectEntryFunc func(ctx context.Context, project string) error

	// DeleteTokenEntryFunc mocks the DeleteTokenEntry method.
	DeleteTokenEntryFunc func(ctx context.Context, token string) error

	// ListTokenEntriesFunc mocks the ListTokenEntries method.
	ListTokenEntriesFunc func(ctx context.Context, project string) ([]db.TokenEntry, error)

	// ReadProjectEntryFunc mocks the ReadProjectEntry method.
	ReadProjectEntryFunc func(ctx context.Context, project string) (db.ProjectEntry, error)

	// ReadTokenEntryFunc mocks the ReadTokenEntry method.
	ReadTokenEntryFunc func(ctx context.Context, token string) (db.TokenEntry, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateProjectEntry holds details about calls to the CreateProjectEntry method.
		CreateProjectEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Pe is the pe argument value.
			Pe db.ProjectEntry
		}
		// CreateTokenEntry holds details about calls to the CreateTokenEntry method.
		CreateTokenEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token types.Token
		}
		// DeleteProjectEntry holds details about calls to the DeleteProjectEntry method.
		DeleteProjectEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Project is the project argument value.
			Project string
		}
		// DeleteTokenEntry holds details about calls to the DeleteTokenEntry method.
		DeleteTokenEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token string
		}
		// ListTokenEntries holds details about calls to the ListTokenEntries method.
		ListTokenEntries []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Project is the project argument value.
			Project string
		}
		// ReadProjectEntry holds details about calls to the ReadProjectEntry method.
		ReadProjectEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Project is the project argument value.
			Project string
		}
		// ReadTokenEntry holds details about calls to the ReadTokenEntry method.
		ReadTokenEntry []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token string
		}
	}
	lockCreateProjectEntry sync.RWMutex
	lockCreateTokenEntry   sync.RWMutex
	lockDeleteProjectEntry sync.RWMutex
	lockDeleteTokenEntry   sync.RWMutex
	lockListTokenEntries   sync.RWMutex
	lockReadProjectEntry   sync.RWMutex
	lockReadTokenEntry     sync.RWMutex
}

// CreateProjectEntry calls CreateProjectEntryFunc.
func (mock *DBClientMock) CreateProjectEntry(ctx context.Context, pe db.ProjectEntry) error {
	if mock.CreateProjectEntryFunc == nil {
		panic("DBClientMock.CreateProjectEntryFunc: method is nil but Client.CreateProjectEntry was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Pe  db.ProjectEntry
	}{
		Ctx: ctx,
		Pe:  pe,
	}
	mock.lockCreateProjectEntry.Lock()
	mock.calls.CreateProjectEntry = append(mock.calls.CreateProjectEntry, callInfo)
	mock.lockCreateProjectEntry.Unlock()
	return mock.CreateProjectEntryFunc(ctx, pe)
}

// CreateProjectEntryCalls gets all the calls that were made to CreateProjectEntry.
// Check the length with:
//     len(mockedClient.CreateProjectEntryCalls())
func (mock *DBClientMock) CreateProjectEntryCalls() []struct {
	Ctx context.Context
	Pe  db.ProjectEntry
} {
	var calls []struct {
		Ctx context.Context
		Pe  db.ProjectEntry
	}
	mock.lockCreateProjectEntry.RLock()
	calls = mock.calls.CreateProjectEntry
	mock.lockCreateProjectEntry.RUnlock()
	return calls
}

// CreateTokenEntry calls CreateTokenEntryFunc.
func (mock *DBClientMock) CreateTokenEntry(ctx context.Context, token types.Token) error {
	if mock.CreateTokenEntryFunc == nil {
		panic("DBClientMock.CreateTokenEntryFunc: method is nil but Client.CreateTokenEntry was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token types.Token
	}{
		Ctx:   ctx,
		Token: token,
	}
	mock.lockCreateTokenEntry.Lock()
	mock.calls.CreateTokenEntry = append(mock.calls.CreateTokenEntry, callInfo)
	mock.lockCreateTokenEntry.Unlock()
	return mock.CreateTokenEntryFunc(ctx, token)
}

// CreateTokenEntryCalls gets all the calls that were made to CreateTokenEntry.
// Check the length with:
//     len(mockedClient.CreateTokenEntryCalls())
func (mock *DBClientMock) CreateTokenEntryCalls() []struct {
	Ctx   context.Context
	Token types.Token
} {
	var calls []struct {
		Ctx   context.Context
		Token types.Token
	}
	mock.lockCreateTokenEntry.RLock()
	calls = mock.calls.CreateTokenEntry
	mock.lockCreateTokenEntry.RUnlock()
	return calls
}

// DeleteProjectEntry calls DeleteProjectEntryFunc.
func (mock *DBClientMock) DeleteProjectEntry(ctx context.Context, project string) error {
	if mock.DeleteProjectEntryFunc == nil {
		panic("DBClientMock.DeleteProjectEntryFunc: method is nil but Client.DeleteProjectEntry was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Project string
	}{
		Ctx:     ctx,
		Project: project,
	}
	mock.lockDeleteProjectEntry.Lock()
	mock.calls.DeleteProjectEntry = append(mock.calls.DeleteProjectEntry, callInfo)
	mock.lockDeleteProjectEntry.Unlock()
	return mock.DeleteProjectEntryFunc(ctx, project)
}

// DeleteProjectEntryCalls gets all the calls that were made to DeleteProjectEntry.
// Check the length with:
//     len(mockedClient.DeleteProjectEntryCalls())
func (mock *DBClientMock) DeleteProjectEntryCalls() []struct {
	Ctx     context.Context
	Project string
} {
	var calls []struct {
		Ctx     context.Context
		Project string
	}
	mock.lockDeleteProjectEntry.RLock()
	calls = mock.calls.DeleteProjectEntry
	mock.lockDeleteProjectEntry.RUnlock()
	return calls
}

// DeleteTokenEntry calls DeleteTokenEntryFunc.
func (mock *DBClientMock) DeleteTokenEntry(ctx context.Context, token string) error {
	if mock.DeleteTokenEntryFunc == nil {
		panic("DBClientMock.DeleteTokenEntryFunc: method is nil but Client.DeleteTokenEntry was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token string
	}{
		Ctx:   ctx,
		Token: token,
	}
	mock.lockDeleteTokenEntry.Lock()
	mock.calls.DeleteTokenEntry = append(mock.calls.DeleteTokenEntry, callInfo)
	mock.lockDeleteTokenEntry.Unlock()
	return mock.DeleteTokenEntryFunc(ctx, token)
}

// DeleteTokenEntryCalls gets all the calls that were made to DeleteTokenEntry.
// Check the length with:
//     len(mockedClient.DeleteTokenEntryCalls())
func (mock *DBClientMock) DeleteTokenEntryCalls() []struct {
	Ctx   context.Context
	Token string
} {
	var calls []struct {
		Ctx   context.Context
		Token string
	}
	mock.lockDeleteTokenEntry.RLock()
	calls = mock.calls.DeleteTokenEntry
	mock.lockDeleteTokenEntry.RUnlock()
	return calls
}

// ListTokenEntries calls ListTokenEntriesFunc.
func (mock *DBClientMock) ListTokenEntries(ctx context.Context, project string) ([]db.TokenEntry, error) {
	if mock.ListTokenEntriesFunc == nil {
		panic("DBClientMock.ListTokenEntriesFunc: method is nil but Client.ListTokenEntries was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Project string
	}{
		Ctx:     ctx,
		Project: project,
	}
	mock.lockListTokenEntries.Lock()
	mock.calls.ListTokenEntries = append(mock.calls.ListTokenEntries, callInfo)
	mock.lockListTokenEntries.Unlock()
	return mock.ListTokenEntriesFunc(ctx, project)
}

// ListTokenEntriesCalls gets all the calls that were made to ListTokenEntries.
// Check the length with:
//     len(mockedClient.ListTokenEntriesCalls())
func (mock *DBClientMock) ListTokenEntriesCalls() []struct {
	Ctx     context.Context
	Project string
} {
	var calls []struct {
		Ctx     context.Context
		Project string
	}
	mock.lockListTokenEntries.RLock()
	calls = mock.calls.ListTokenEntries
	mock.lockListTokenEntries.RUnlock()
	return calls
}

// ReadProjectEntry calls ReadProjectEntryFunc.
func (mock *DBClientMock) ReadProjectEntry(ctx context.Context, project string) (db.ProjectEntry, error) {
	if mock.ReadProjectEntryFunc == nil {
		panic("DBClientMock.ReadProjectEntryFunc: method is nil but Client.ReadProjectEntry was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Project string
	}{
		Ctx:     ctx,
		Project: project,
	}
	mock.lockReadProjectEntry.Lock()
	mock.calls.ReadProjectEntry = append(mock.calls.ReadProjectEntry, callInfo)
	mock.lockReadProjectEntry.Unlock()
	return mock.ReadProjectEntryFunc(ctx, project)
}

// ReadProjectEntryCalls gets all the calls that were made to ReadProjectEntry.
// Check the length with:
//     len(mockedClient.ReadProjectEntryCalls())
func (mock *DBClientMock) ReadProjectEntryCalls() []struct {
	Ctx     context.Context
	Project string
} {
	var calls []struct {
		Ctx     context.Context
		Project string
	}
	mock.lockReadProjectEntry.RLock()
	calls = mock.calls.ReadProjectEntry
	mock.lockReadProjectEntry.RUnlock()
	return calls
}

// ReadTokenEntry calls ReadTokenEntryFunc.
func (mock *DBClientMock) ReadTokenEntry(ctx context.Context, token string) (db.TokenEntry, error) {
	if mock.ReadTokenEntryFunc == nil {
		panic("DBClientMock.ReadTokenEntryFunc: method is nil but Client.ReadTokenEntry was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token string
	}{
		Ctx:   ctx,
		Token: token,
	}
	mock.lockReadTokenEntry.Lock()
	mock.calls.ReadTokenEntry = append(mock.calls.ReadTokenEntry, callInfo)
	mock.lockReadTokenEntry.Unlock()
	return mock.ReadTokenEntryFunc(ctx, token)
}

// ReadTokenEntryCalls gets all the calls that were made to ReadTokenEntry.
// Check the length with:
//     len(mockedClient.ReadTokenEntryCalls())
func (mock *DBClientMock) ReadTokenEntryCalls() []struct {
	Ctx   context.Context
	Token string
} {
	var calls []struct {
		Ctx   context.Context
		Token string
	}
	mock.lockReadTokenEntry.RLock()
	calls = mock.calls.ReadTokenEntry
	mock.lockReadTokenEntry.RUnlock()
	return calls
}
