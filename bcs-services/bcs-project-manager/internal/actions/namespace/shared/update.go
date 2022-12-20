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

package shared

import (
	"context"
	"math"

	"github.com/Tencent/bk-bcs/bcs-services/pkg/bcs-auth/middleware"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/component/itsm"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/logging"
	nsm "github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/store/namespace"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/proto/bcsproject"
)

// UpdateNamespace implement for UpdateNamespace interface
func (a *SharedNamespaceAction) UpdateNamespace(ctx context.Context,
	req *proto.UpdateNamespaceRequest, resp *proto.UpdateNamespaceResponse) error {
	var username string
	if authUser, err := middleware.GetUserFromContext(ctx); err == nil {
		username = authUser.GetUsername()
	}
	namespace := &nsm.Namespace{
		ProjectCode: req.GetProjectCode(),
		ClusterID:   req.GetClusterID(),
		Name:        req.GetNamespace(),
		Creator:     username,
		ResourceQuota: &nsm.Quota{
			CPURequests:    req.GetQuota().GetCpuRequests(),
			CPULimits:      req.GetQuota().GetCpuLimits(),
			MemoryRequests: req.GetQuota().GetMemoryRequests(),
			MemoryLimits:   req.GetQuota().GetMemoryLimits(),
		},
	}
	cpuLimits, err := resource.ParseQuantity(req.GetQuota().GetCpuLimits())
	if err != nil {
		logging.Error("parse quantity cpu limits failed, err: %s", err.Error())
		return err
	}
	memoryLimits, err := resource.ParseQuantity(req.GetQuota().GetMemoryLimits())
	if err != nil {
		logging.Error("parse quantity cpu limits failed, err: %s", err.Error())
		return err
	}
	// memoryLimits.Value() return unit is byte， needs to be converted to Gi (divide 2^30)
	itsmResp, err := itsm.SubmitUpdateNamespaceTicket(username,
		req.GetProjectCode(), req.GetClusterID(), req.GetNamespace(),
		int(cpuLimits.Value()), int(memoryLimits.Value()/int64(math.Pow(2, 30))))
	if err != nil {
		logging.Error("itsm create ticket failed, err: %s", err.Error())
		return err
	}
	namespace.ItsmTicketType = nsm.ItsmTicketTypeUpdate
	namespace.ItsmTicketURL = itsmResp.TicketURL
	namespace.ItsmTicketStatus = nsm.ItsmTicketStatusCreated
	namespace.ItsmTicketSN = itsmResp.SN
	if err := a.model.CreateNamespace(ctx, namespace); err != nil {
		logging.Error("create namespace %s/%s in db failed, err: %s", req.GetClusterID(), req.GetNamespace(), err.Error())
		return err
	}
	return nil
}
