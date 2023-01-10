// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package resource

import (
	"sync"
)

// Ensure, that ManagerMock does implement Manager.
// If this is not the case, regenerate this file with moq.
var _ Manager = &ManagerMock{}

// ManagerMock is a mock implementation of Manager.
//
//	func TestSomethingThatUsesManager(t *testing.T) {
//
//		// make and configure a mocked Manager
//		mockedManager := &ManagerMock{
//			GetCudaDriverVersionFunc: func() (*uint, *uint, error) {
//				panic("mock out the GetCudaDriverVersion method")
//			},
//			GetDevicesFunc: func() ([]Device, error) {
//				panic("mock out the GetDevices method")
//			},
//			GetDriverVersionFunc: func() (string, error) {
//				panic("mock out the GetDriverVersion method")
//			},
//			InitFunc: func() error {
//				panic("mock out the Init method")
//			},
//			ShutdownFunc: func() error {
//				panic("mock out the Shutdown method")
//			},
//		}
//
//		// use mockedManager in code that requires Manager
//		// and then make assertions.
//
//	}
type ManagerMock struct {
	// GetCudaDriverVersionFunc mocks the GetCudaDriverVersion method.
	GetCudaDriverVersionFunc func() (*uint, *uint, error)

	// GetDevicesFunc mocks the GetDevices method.
	GetDevicesFunc func() ([]Device, error)

	// GetDriverVersionFunc mocks the GetDriverVersion method.
	GetDriverVersionFunc func() (string, error)

	// InitFunc mocks the Init method.
	InitFunc func() error

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// GetCudaDriverVersion holds details about calls to the GetCudaDriverVersion method.
		GetCudaDriverVersion []struct {
		}
		// GetDevices holds details about calls to the GetDevices method.
		GetDevices []struct {
		}
		// GetDriverVersion holds details about calls to the GetDriverVersion method.
		GetDriverVersion []struct {
		}
		// Init holds details about calls to the Init method.
		Init []struct {
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
		}
	}
	lockGetCudaDriverVersion sync.RWMutex
	lockGetDevices           sync.RWMutex
	lockGetDriverVersion     sync.RWMutex
	lockInit                 sync.RWMutex
	lockShutdown             sync.RWMutex
}

// GetCudaDriverVersion calls GetCudaDriverVersionFunc.
func (mock *ManagerMock) GetCudaDriverVersion() (*uint, *uint, error) {
	if mock.GetCudaDriverVersionFunc == nil {
		panic("ManagerMock.GetCudaDriverVersionFunc: method is nil but Manager.GetCudaDriverVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCudaDriverVersion.Lock()
	mock.calls.GetCudaDriverVersion = append(mock.calls.GetCudaDriverVersion, callInfo)
	mock.lockGetCudaDriverVersion.Unlock()
	return mock.GetCudaDriverVersionFunc()
}

// GetCudaDriverVersionCalls gets all the calls that were made to GetCudaDriverVersion.
// Check the length with:
//
//	len(mockedManager.GetCudaDriverVersionCalls())
func (mock *ManagerMock) GetCudaDriverVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCudaDriverVersion.RLock()
	calls = mock.calls.GetCudaDriverVersion
	mock.lockGetCudaDriverVersion.RUnlock()
	return calls
}

// GetDevices calls GetDevicesFunc.
func (mock *ManagerMock) GetDevices() ([]Device, error) {
	if mock.GetDevicesFunc == nil {
		panic("ManagerMock.GetDevicesFunc: method is nil but Manager.GetDevices was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDevices.Lock()
	mock.calls.GetDevices = append(mock.calls.GetDevices, callInfo)
	mock.lockGetDevices.Unlock()
	return mock.GetDevicesFunc()
}

// GetDevicesCalls gets all the calls that were made to GetDevices.
// Check the length with:
//
//	len(mockedManager.GetDevicesCalls())
func (mock *ManagerMock) GetDevicesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDevices.RLock()
	calls = mock.calls.GetDevices
	mock.lockGetDevices.RUnlock()
	return calls
}

// GetDriverVersion calls GetDriverVersionFunc.
func (mock *ManagerMock) GetDriverVersion() (string, error) {
	if mock.GetDriverVersionFunc == nil {
		panic("ManagerMock.GetDriverVersionFunc: method is nil but Manager.GetDriverVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDriverVersion.Lock()
	mock.calls.GetDriverVersion = append(mock.calls.GetDriverVersion, callInfo)
	mock.lockGetDriverVersion.Unlock()
	return mock.GetDriverVersionFunc()
}

// GetDriverVersionCalls gets all the calls that were made to GetDriverVersion.
// Check the length with:
//
//	len(mockedManager.GetDriverVersionCalls())
func (mock *ManagerMock) GetDriverVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDriverVersion.RLock()
	calls = mock.calls.GetDriverVersion
	mock.lockGetDriverVersion.RUnlock()
	return calls
}

// Init calls InitFunc.
func (mock *ManagerMock) Init() error {
	if mock.InitFunc == nil {
		panic("ManagerMock.InitFunc: method is nil but Manager.Init was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInit.Lock()
	mock.calls.Init = append(mock.calls.Init, callInfo)
	mock.lockInit.Unlock()
	return mock.InitFunc()
}

// InitCalls gets all the calls that were made to Init.
// Check the length with:
//
//	len(mockedManager.InitCalls())
func (mock *ManagerMock) InitCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInit.RLock()
	calls = mock.calls.Init
	mock.lockInit.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *ManagerMock) Shutdown() error {
	if mock.ShutdownFunc == nil {
		panic("ManagerMock.ShutdownFunc: method is nil but Manager.Shutdown was just called")
	}
	callInfo := struct {
	}{}
	mock.lockShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	mock.lockShutdown.Unlock()
	return mock.ShutdownFunc()
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//
//	len(mockedManager.ShutdownCalls())
func (mock *ManagerMock) ShutdownCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockShutdown.RLock()
	calls = mock.calls.Shutdown
	mock.lockShutdown.RUnlock()
	return calls
}
