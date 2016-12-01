// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"github.com/stretchr/testify/mock"
)

// MockUserAdmApp is an autogenerated mock type for the UserAdmApp type
type mockUserAdmApp struct {
	mock.Mock

	sign SignFunc
}

// Login provides a mock function with given fields: email, pass
func (_m *mockUserAdmApp) Login(email string, pass string) (*Token, error) {
	ret := _m.Called(email, pass)
	tok, _ := ret.Get(0).(*Token)

	return tok, ret.Error(1)
}

func (_m mockUserAdmApp) SignToken() SignFunc {
	_m.Called()
	// this does not work with mock, because SignFunc is defined as:
	//   func(*Token) (string, error)
	// but Get() returns:
	//   func(*main.Token) (string, error)
	// sf, ok := ret.Get(0).(SignFunc)

	return _m.sign
}
