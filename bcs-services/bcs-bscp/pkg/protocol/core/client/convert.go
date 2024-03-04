/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package pbclient xxx
package pbclient

import (
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

// ClientSpec convert pb ClientSpec to table ClientSpec
func (c *ClientSpec) ClientSpec() *table.ClientSpec {
	if c == nil {
		return nil
	}
	resource := table.Resource{}
	if c.Resource != nil {
		resource = table.Resource{
			CpuMaxUsage:    c.Resource.CpuMaxUsage,
			MemoryMaxUsage: c.Resource.MemoryMaxUsage,
			CpuUsage:       c.Resource.CpuUsage,
			MemoryUsage:    c.Resource.MemoryUsage,
		}
	}
	return &table.ClientSpec{
		ClientVersion:             c.ClientVersion,
		ClientType:                table.ClientType(c.ClientType),
		Ip:                        c.Ip,
		Labels:                    c.Labels,
		Annotations:               c.Annotations,
		FirstConnectTime:          c.FirstConnectTime.AsTime(),
		LastHeartbeatTime:         c.LastHeartbeatTime.AsTime(),
		OnlineStatus:              table.OnlineStatus(c.OnlineStatus),
		Resource:                  resource,
		CurrentReleaseID:          c.CurrentReleaseId,
		TargetReleaseID:           c.TargetReleaseId,
		ReleaseChangeStatus:       table.Status(c.ReleaseChangeStatus),
		ReleaseChangeFailedReason: table.FailedReason(c.ReleaseChangeFailedReason),
		FailedDetailReason:        c.FailedDetailReason,
	}
}

// ClientAttachment convert pb ClientAttachment to table ClientAttachment
func (c *ClientAttachment) ClientAttachment() *table.ClientAttachment {
	return &table.ClientAttachment{
		UID:   c.Uid,
		BizID: c.BizId,
		AppID: c.AppId,
	}
}

// PbClientAttachment convert table ClientAttachment to pb ClientAttachment
func PbClientAttachment(attachment *table.ClientAttachment) *ClientAttachment { // nolint
	if attachment == nil {
		return nil
	}
	return &ClientAttachment{
		Uid:   attachment.UID,
		BizId: attachment.BizID,
		AppId: attachment.AppID,
	}
}
