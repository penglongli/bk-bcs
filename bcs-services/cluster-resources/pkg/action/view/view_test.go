/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package view

import (
	"testing"
)

func TestGetNewFileName(t *testing.T) {
	testCases := []struct {
		name         string
		filenames    []string
		expectedName string
	}{
		{"aaa copy", []string{}, "aaa copy"},
		{"aaa", []string{"aaa", "aaa copy", "aaa copy 2", "bbb", "ccc copy 3"}, "aaa copy 3"},
		{"bbb", []string{"aaa", "aaa copy", "aaa copy 2", "bbb", "ccc copy 3"}, "bbb copy"},
		{"ccc", []string{"aaa", "aaa copy", "aaa copy 2", "bbb", "ccc copy 3"}, "ccc copy 4"},
		{"ddd", []string{"aaa", "aaa copy", "aaa copy 2", "bbb", "ccc copy 3"}, "ddd copy"},
		{"aaa", []string{"aaa", "aaa copy", "aaa copy 2", "aaa copy 3"}, "aaa copy 4"},
		{"aaa copy", []string{"aaa", "aaa copy", "aaa copy 2", "aaa copy 3"}, "aaa copy 4"},
		{"aaa copy 2", []string{"aaa", "aaa copy", "aaa copy 2", "aaa copy 3"}, "aaa copy 4"},
		{"aaa copy 3", []string{"aaa", "aaa copy", "aaa copy 2", "aaa copy 3"}, "aaa copy 4"},
	}

	for _, tc := range testCases {
		newFileName := getNewFileName(tc.name, tc.filenames)
		if newFileName != tc.expectedName {
			t.Errorf("For name: %s, expected: %s, but got: %s", tc.name, tc.expectedName, newFileName)
		}
	}
}