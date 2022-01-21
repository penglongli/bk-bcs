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

package resource

import (
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/common"
)

type ClusterConf struct {
	Rest      *rest.Config
	ClusterID string
}

// 生成测试用 ClusterConf 对象（默认是本地集群）
func newMockClusterConfig(clusterID string) *ClusterConf {
	kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	conf, _ := clientcmd.BuildConfigFromFlags("", kubeConfig)
	return &ClusterConf{conf, clusterID}
}

// NewClusterConfig 生成 ClusterConf 对象
func NewClusterConfig(clusterID string) *ClusterConf {
	if common.RunMode == common.Dev || common.RunMode == common.UnitTest {
		return newMockClusterConfig(clusterID)
	}
	return &ClusterConf{
		Rest: &rest.Config{
			Host:            common.BCSApiGWHost + "/clusters/" + clusterID,
			BearerToken:     common.BCSApiGWAuthToken,
			TLSClientConfig: rest.TLSClientConfig{Insecure: true},
		},
		ClusterID: clusterID,
	}
}
