/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"errors"
	"strconv"
)

// ArgocdServerError describe the errors and its messages for bcs argocd server
type ArgocdServerError uint32

const (
	ErrArgocdServerSuccess ArgocdServerError = iota
	ErrArgocdServerReqOrRespEmpty
	ErrActionFailed
	ErrProjectNotExist
	ErrInstanceNotFound
	ErrPluginNotFound
)

// Int32 return ArgocdServerError's value
func (ase ArgocdServerError) Int32() uint32 {
	code := uint32(ase)
	return code
}

// Error return ArgocdServerError's builtin error message
func (ase ArgocdServerError) Error() string {
	s, ok := errorCodeMapping[ase]
	if ok {
		return s
	}

	return "unknown error code " + strconv.Itoa(int(ase))
}

// ErrorMessage return ArgocdServerError's builtin error message and add the custom message
func (ase ArgocdServerError) ErrorMessage(message string) string {
	if message == "" {
		return ase.Error()
	}
	msg := "[" + ase.Error() + "]" + message
	return msg
}

// GenError return an error generated by ArgocdServerError
func (ase ArgocdServerError) GenError() error {
	return errors.New(ase.Error())
}

// OK check if ArgocdServerError is success
func (ase ArgocdServerError) OK() *bool {
	ok := ase == ErrArgocdServerSuccess
	return &ok
}

var errorCodeMapping = map[ArgocdServerError]string{
	ErrArgocdServerSuccess:        "success",
	ErrArgocdServerReqOrRespEmpty: "grpc req or resp is empty",
	ErrActionFailed:               "action failed",
	ErrProjectNotExist:            "project not exist",
	ErrInstanceNotFound:           "instance not exist",
	ErrPluginNotFound:             "plugin not exist",
}