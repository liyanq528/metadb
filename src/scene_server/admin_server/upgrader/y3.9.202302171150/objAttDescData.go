/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package y3_9_202302171150

import (
	"time"

	"configcenter/src/common"
	"configcenter/src/common/metadata"
	mCommon "configcenter/src/scene_server/admin_server/common"
)

// default group
var (
	groupBaseInfo = mCommon.BaseInfo
)

// Distribution init revision
var Distribution = "community" // could be community or enterprise

/*
	&Attribute{ObjectID: objID, PropertyID: "", PropertyName: "", IsRequired: , IsOnly: , PropertyGroup: , PropertyType: , Option: ""},
*/

// AppRow app structure
func AppRow() []*Attribute {
	objID := common.BKInnerObjIDApp

	groupAppRole := mCommon.AppRole

	lifeCycleOption := []metadata.EnumVal{{ID: "1", Name: "测试中", Type: "text"}, {ID: "2", Name: "已上线", Type: "text", IsDefault: true}, {ID: "3", Name: "停运", Type: "text"}}
	languageOption := []metadata.EnumVal{{ID: "1", Name: "中文", Type: "text", IsDefault: true}, {ID: "2", Name: "English", Type: "text"}}
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_name", PropertyName: "业务名", IsRequired: true, IsOnly: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "life_cycle", PropertyName: "生命周期", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: lifeCycleOption},

		//role
		&Attribute{ObjectID: objID, PropertyID: common.BKMaintainersField, PropertyName: "运维人员", IsRequired: true, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKProductPMField, PropertyName: "产品人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},

		&Attribute{ObjectID: objID, PropertyID: common.BKTesterField, PropertyName: "测试人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_developer", PropertyName: "开发人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKOperatorField, PropertyName: "操作人员", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupAppRole, PropertyType: common.FieldTypeUser, Option: ""},

		&Attribute{ObjectID: objID, PropertyID: "time_zone", PropertyName: "时区", IsRequired: true, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTimeZone, Option: "", IsReadOnly: true},
		&Attribute{ObjectID: objID, PropertyID: "language", PropertyName: "语言", IsRequired: true, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: languageOption, IsReadOnly: true},
	}

	return dataRows

}

// SetRow set structure
func SetRow() []*Attribute {
	//objID := common.BKInnerObjIDSet
	//serviceStatusOption := []metadata.EnumVal{{ID: "1", Name: "开放", Type: "text", IsDefault: true}, {ID: "2", Name: "关闭", Type: "text"}}

	dataRows := []*Attribute{
		&Attribute{ID: 152, ObjectID: "region", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 158, ObjectID: "region", PropertyID: "RegionID", PropertyIndex: 1, PropertyName: "地域ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 159, ObjectID: "region", PropertyID: "RegionName", PropertyIndex: 2, PropertyName: "地域名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 160, ObjectID: "region", PropertyID: "RegionPublicity", PropertyIndex: 3, PropertyName: "地域属性", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: true, IsOnly: false, IsEditable: true},
		&Attribute{ID: 161, ObjectID: "region", PropertyID: "StackEndpoint", PropertyIndex: 4, PropertyName: "stack端点", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 163, ObjectID: "region", PropertyID: "Description", PropertyIndex: 5, PropertyName: "描述", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 166, ObjectID: "region", PropertyID: "RegionState", PropertyIndex: 6, PropertyName: "状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"Ready", "Enabled", "Disabled", "Archived"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 167, ObjectID: "region", PropertyID: "Deleted", PropertyIndex: 7, PropertyName: "归档标记", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 168, ObjectID: "region", PropertyID: "CreateAt", PropertyIndex: 8, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 169, ObjectID: "region", PropertyID: "UpdateAt", PropertyIndex: 9, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 170, ObjectID: "region", PropertyID: "SyncAt", PropertyIndex: 10, PropertyName: "同步时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 171, ObjectID: "region", PropertyID: "ZoneNumber", PropertyIndex: 11, PropertyName: "可用区数量", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},

		&Attribute{ID: 190, ObjectID: "zone", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 191, ObjectID: "zone", PropertyID: "RegionID", PropertyIndex: 1, PropertyName: "地域ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 192, ObjectID: "zone", PropertyID: "RegionName", PropertyIndex: 2, PropertyName: "地域名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 193, ObjectID: "zone", PropertyID: "ZoneID", PropertyIndex: 3, PropertyName: "可用区ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 194, ObjectID: "zone", PropertyID: "ZoneName", PropertyIndex: 4, PropertyName: "可用区名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 195, ObjectID: "zone", PropertyID: "Description", PropertyIndex: 5, PropertyName: "描述", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 196, ObjectID: "zone", PropertyID: "ZoneState", PropertyIndex: 6, PropertyName: "状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"Ready", "Enabled", "Disabled", "Archived"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 197, ObjectID: "zone", PropertyID: "Deleted", PropertyIndex: 7, PropertyName: "归档标记", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 198, ObjectID: "zone", PropertyID: "CreateAt", PropertyIndex: 8, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 199, ObjectID: "zone", PropertyID: "UpdateAt", PropertyIndex: 9, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 200, ObjectID: "zone", PropertyID: "SyncAt", PropertyIndex: 10, PropertyName: "同步时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 201, ObjectID: "zone", PropertyID: "ReservedCpu", PropertyIndex: 11, PropertyName: "保留cpu资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 202, ObjectID: "zone", PropertyID: "ReservedMemory", PropertyIndex: 12, PropertyName: "保留内存资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 203, ObjectID: "zone", PropertyID: "ZonePublicity", PropertyIndex: 13, PropertyName: "可用区属性", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 204, ObjectID: "zone", PropertyID: "CpuRatio", PropertyIndex: 14, PropertyName: "cpu超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 205, ObjectID: "zone", PropertyID: "MemoryRatio", PropertyIndex: 15, PropertyName: "内存超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},

		&Attribute{ID: 210, ObjectID: "cluster", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 211, ObjectID: "cluster", PropertyID: "RegionID", PropertyIndex: 1, PropertyName: "地域ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 212, ObjectID: "cluster", PropertyID: "ZoneID", PropertyIndex: 2, PropertyName: "可用区ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 213, ObjectID: "cluster", PropertyID: "ClusterID", PropertyIndex: 3, PropertyName: "集群ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 214, ObjectID: "cluster", PropertyID: "ClusterName", PropertyIndex: 4, PropertyName: "集群名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 215, ObjectID: "cluster", PropertyID: "Description", PropertyIndex: 5, PropertyName: "集群描述", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 216, ObjectID: "cluster", PropertyID: "ReservedCpu", PropertyIndex: 6, PropertyName: "保留cpu资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 217, ObjectID: "cluster", PropertyID: "ReservedMemory", PropertyIndex: 7, PropertyName: "保留内存资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 218, ObjectID: "cluster", PropertyID: "CpuRatio", PropertyIndex: 8, PropertyName: "cpu超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 219, ObjectID: "cluster", PropertyID: "MemoryRatio", PropertyIndex: 9, PropertyName: "内存超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 220, ObjectID: "cluster", PropertyID: "State", PropertyIndex: 10, PropertyName: "集群状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"Ready", "Enabled", "Disabled", "Archived"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 221, ObjectID: "cluster", PropertyID: "Deleted", PropertyIndex: 11, PropertyName: "归档标记", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 222, ObjectID: "cluster", PropertyID: "PkID", PropertyIndex: 12, PropertyName: "主键", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 223, ObjectID: "cluster", PropertyID: "TagIDs", PropertyIndex: 13, PropertyName: "TagIDs", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 224, ObjectID: "cluster", PropertyID: "Tags", PropertyIndex: 14, PropertyName: "Tags", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 225, ObjectID: "cluster", PropertyID: "CreateAt", PropertyIndex: 15, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 226, ObjectID: "cluster", PropertyID: "UpdateAt", PropertyIndex: 16, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 227, ObjectID: "cluster", PropertyID: "RegionName", PropertyIndex: 17, PropertyName: "地域名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 228, ObjectID: "cluster", PropertyID: "ZoneName", PropertyIndex: 18, PropertyName: "可用区名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},

		&Attribute{ID: 250, ObjectID: "storageBackendPool", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 251, ObjectID: "storageBackendPool", PropertyID: "StorageBackendPoolID", PropertyIndex: 1, PropertyName: "存储后端池ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 252, ObjectID: "storageBackendPool", PropertyID: "StorageBackendID", PropertyIndex: 2, PropertyName: "存储后端ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 253, ObjectID: "storageBackendPool", PropertyID: "PoolName", PropertyIndex: 3, PropertyName: "池名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 254, ObjectID: "storageBackendPool", PropertyID: "TotalCapacity", PropertyIndex: 4, PropertyName: "池总存储大小", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 255, ObjectID: "storageBackendPool", PropertyID: "CreateAt", PropertyIndex: 5, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 256, ObjectID: "storageBackendPool", PropertyID: "UpdateAt", PropertyIndex: 6, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},

		&Attribute{ID: 270, ObjectID: "storageBackend", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 271, ObjectID: "storageBackend", PropertyID: "StorageBackendID", PropertyIndex: 1, PropertyName: "存储后端ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 272, ObjectID: "storageBackend", PropertyID: "StorageType", PropertyIndex: 2, PropertyName: "存储类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 273, ObjectID: "storageBackend", PropertyID: "ProtocolType", PropertyIndex: 3, PropertyName: "协议类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 274, ObjectID: "storageBackend", PropertyID: "IscsiType", PropertyIndex: 4, PropertyName: "iscsi类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 275, ObjectID: "storageBackend", PropertyID: "Url", PropertyIndex: 5, PropertyName: "存储后端url", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 276, ObjectID: "storageBackend", PropertyID: "Username", PropertyIndex: 6, PropertyName: "用户名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 277, ObjectID: "storageBackend", PropertyID: "Password", PropertyIndex: 7, PropertyName: "密码", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 278, ObjectID: "storageBackend", PropertyID: "CreateAt", PropertyIndex: 8, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 279, ObjectID: "storageBackend", PropertyID: "UpdateAt", PropertyIndex: 9, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},

		&Attribute{ID: 300, ObjectID: "cmdb_host", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		&Attribute{ID: 301, ObjectID: "cmdb_host", PropertyID: "HostID", PropertyIndex: 1, PropertyName: "主机ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 302, ObjectID: "cmdb_host", PropertyID: "HostName", PropertyIndex: 2, PropertyName: "主机名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 303, ObjectID: "cmdb_host", PropertyID: "RegionID", PropertyIndex: 3, PropertyName: "地域ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 304, ObjectID: "cmdb_host", PropertyID: "ZoneID", PropertyIndex: 4, PropertyName: "可用区ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 305, ObjectID: "cmdb_host", PropertyID: "ClusterID", PropertyIndex: 5, PropertyName: "集群ID", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 306, ObjectID: "cmdb_host", PropertyID: "Description", PropertyIndex: 6, PropertyName: "描述", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 307, ObjectID: "cmdb_host", PropertyID: "HypervisorType", PropertyIndex: 7, PropertyName: "虚拟化类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 308, ObjectID: "cmdb_host", PropertyID: "State", PropertyIndex: 8, PropertyName: "可用状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"Enabled", "Disabled", "PreMaintenance", "Maintenance"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 309, ObjectID: "cmdb_host", PropertyID: "Status", PropertyIndex: 9, PropertyName: "连接状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"Connecting", "Connected", "Disconnected"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 310, ObjectID: "cmdb_host", PropertyID: "ReservedCpu", PropertyIndex: 10, PropertyName: "保留cpu资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 311, ObjectID: "cmdb_host", PropertyID: "ReservedMemory", PropertyIndex: 11, PropertyName: "保留内存资源", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 312, ObjectID: "cmdb_host", PropertyID: "CpuRatio", PropertyIndex: 12, PropertyName: "cpu超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 313, ObjectID: "cmdb_host", PropertyID: "MemoryRatio", PropertyIndex: 13, PropertyName: "内存超分比", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 314, ObjectID: "cmdb_host", PropertyID: "ManagementIp", PropertyIndex: 14, PropertyName: "管理网卡ip", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 315, ObjectID: "cmdb_host", PropertyID: "ManagementMac", PropertyIndex: 15, PropertyName: "管理网卡mac地址", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 316, ObjectID: "cmdb_host", PropertyID: "BusinessIp", PropertyIndex: 16, PropertyName: "业务网卡ip", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 317, ObjectID: "cmdb_host", PropertyID: "BusinessMac", PropertyIndex: 17, PropertyName: "业务网卡mac地址", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 318, ObjectID: "cmdb_host", PropertyID: "StorageIp", PropertyIndex: 18, PropertyName: "存储网卡ip", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 319, ObjectID: "cmdb_host", PropertyID: "StorageMac", PropertyIndex: 19, PropertyName: "存储网卡mac地址", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 320, ObjectID: "cmdb_host", PropertyID: "Sn", PropertyIndex: 20, PropertyName: "序列号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 321, ObjectID: "cmdb_host", PropertyID: "AgentUrl", PropertyIndex: 21, PropertyName: "调用agent的url", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 322, ObjectID: "cmdb_host", PropertyID: "TotalCpu", PropertyIndex: 22, PropertyName: "TotalCpu", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 323, ObjectID: "cmdb_host", PropertyID: "TotalMemory", PropertyIndex: 23, PropertyName: "TotalMemory", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "", Max: ""}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 324, ObjectID: "cmdb_host", PropertyID: "Maintenance", PropertyIndex: 24, PropertyName: "Maintenance", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: false, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 325, ObjectID: "cmdb_host", PropertyID: "TotalExtendResource", PropertyIndex: 25, PropertyName: "TotalExtendResource", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 326, ObjectID: "cmdb_host", PropertyID: "CreateAt", PropertyIndex: 26, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 327, ObjectID: "cmdb_host", PropertyID: "UpdateAt", PropertyIndex: 27, PropertyName: "更新时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 328, ObjectID: "cmdb_host", PropertyID: "RegionName", PropertyIndex: 28, PropertyName: "地域名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 329, ObjectID: "cmdb_host", PropertyID: "ZoneName", PropertyIndex: 29, PropertyName: "可用区名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		&Attribute{ID: 330, ObjectID: "cmdb_host", PropertyID: "ClusterName", PropertyIndex: 30, PropertyName: "集群名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 257, ObjectID: "crm_dim_product", PropertyID: "name", PropertyIndex: 1, PropertyName: "产品分类名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 258, ObjectID: "crm_dim_product", PropertyID: "owner", PropertyIndex: 2, PropertyName: "负责人", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 259, ObjectID: "crm_dim_product", PropertyID: "field_bGEi1__c", PropertyIndex: 3, PropertyName: "产品分类编号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 260, ObjectID: "crm_dim_product", PropertyID: "field_gAUuC__c", PropertyIndex: 4, PropertyName: "五大产品线名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 261, ObjectID: "crm_dim_product", PropertyID: "field_04268__c", PropertyIndex: 5, PropertyName: "五大产品线", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 262, ObjectID: "crm_dim_product", PropertyID: "lock_status", PropertyIndex: 6, PropertyName: "锁定状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 263, ObjectID: "crm_dim_product", PropertyID: "owner_department", PropertyIndex: 7, PropertyName: "负责人主属部门", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 264, ObjectID: "crm_dim_product", PropertyID: "out_owner", PropertyIndex: 8, PropertyName: "外部负责人", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 267, ObjectID: "crm_dim_product", PropertyID: "last_modified_by", PropertyIndex: 11, PropertyName: "最后修改人", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 268, ObjectID: "crm_dim_product", PropertyID: "created_by", PropertyIndex: 12, PropertyName: "创建人", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 269, ObjectID: "crm_dim_product", PropertyID: "data_own_department", PropertyIndex: 13, PropertyName: "归属部门", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 270, ObjectID: "crm_dim_product", PropertyID: "record_type", PropertyIndex: 14, PropertyName: "业务类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 271, ObjectID: "crm_dim_product", PropertyID: "life_status", PropertyIndex: 15, PropertyName: "生命状态", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 265, ObjectID: "crm_dim_product", PropertyID: "last_modified_time", PropertyIndex: 9, PropertyName: "最后修改时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 272, ObjectID: "crm_dim_product", PropertyID: "createdate", PropertyIndex: 16, PropertyName: "创建日期", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeDate, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 266, ObjectID: "crm_dim_product", PropertyID: "create_time", PropertyIndex: 10, PropertyName: "创建时间", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 150, ObjectID: "water_cooled_chiller", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 170, ObjectID: "low_voltage_switchgear", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 172, ObjectID: "water_cooled_chiller", PropertyID: "SN", PropertyIndex: 3, PropertyName: "SN", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 173, ObjectID: "water_cooled_chiller", PropertyID: "manufacturer", PropertyIndex: 4, PropertyName: "厂商", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 174, ObjectID: "water_cooled_chiller", PropertyID: "model", PropertyIndex: 4, PropertyName: "型号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 175, ObjectID: "water_cooled_chiller", PropertyID: "rated_power", PropertyIndex: 0, PropertyName: "额定功率", Unit: "KW", PropertyGroup: "1b4717fe-7e71-4fd9-a725-70d454b80803", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 176, ObjectID: "water_cooled_chiller", PropertyID: "refrigerating_capacity", PropertyIndex: 1, PropertyName: "制冷量", Unit: "KW", PropertyGroup: "1b4717fe-7e71-4fd9-a725-70d454b80803", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 177, ObjectID: "water_cooled_chiller", PropertyID: "rated_temperature_of_chilled_water", PropertyIndex: 2, PropertyName: "冷冻水额定温度", Unit: "C", PropertyGroup: "1b4717fe-7e71-4fd9-a725-70d454b80803", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true},
		//&Attribute{ID: 179, ObjectID: "low_voltage_switchgear", PropertyID: "manufacturer", PropertyIndex: 1, PropertyName: "厂商", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 180, ObjectID: "low_voltage_switchgear", PropertyID: "model", PropertyIndex: 2, PropertyName: "型号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 181, ObjectID: "low_voltage_switchgear", PropertyID: "SN", PropertyIndex: 3, PropertyName: "SN", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 182, ObjectID: "low_voltage_switchgear", PropertyID: "nominal_voltage", PropertyIndex: 0, PropertyName: "标称电压", Unit: "V", PropertyGroup: "1b99731e-bbb3-4258-b7fc-2cc3be29185b", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 184, ObjectID: "low_voltage_switchgear", PropertyID: "rated_power", PropertyIndex: 2, PropertyName: "额定功率", Unit: "KW", PropertyGroup: "1b99731e-bbb3-4258-b7fc-2cc3be29185b", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 186, ObjectID: "water_cooled_chiller", PropertyID: "system_design_pressure", PropertyIndex: 3, PropertyName: "系统设计压力", Unit: "MPa", PropertyGroup: "1b4717fe-7e71-4fd9-a725-70d454b80803", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 187, ObjectID: "idc", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 188, ObjectID: "idc", PropertyID: "name", PropertyIndex: 1, PropertyName: "数据中心名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 189, ObjectID: "idc", PropertyID: "type", PropertyIndex: 2, PropertyName: "类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "自建", Type: "text", IsDefault: true}, {ID: "2", Name: "代维", Type: "text", IsDefault: false}}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 190, ObjectID: "idc", PropertyID: "build_time", PropertyIndex: 3, PropertyName: "建设日期", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeDate, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 191, ObjectID: "idc", PropertyID: "prod_time", PropertyIndex: 4, PropertyName: "投产日期", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeDate, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 192, ObjectID: "idc", PropertyID: "external_power_capacity", PropertyIndex: 5, PropertyName: "外电容量", Unit: "KVA", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 193, ObjectID: "idc", PropertyID: "build_area", PropertyIndex: 6, PropertyName: "建筑面积", Unit: "平方", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 194, ObjectID: "idc", PropertyID: "address", PropertyIndex: 7, PropertyName: "地址", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 195, ObjectID: "idc", PropertyID: "is_independent_park", PropertyIndex: 8, PropertyName: "是否是独立园区", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 196, ObjectID: "idc", PropertyID: "construction_property_rights", PropertyIndex: 9, PropertyName: "建筑产权", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "自有", Type: "text", IsDefault: true}, {ID: "2", Name: "租赁", Type: "text", IsDefault: false}}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 197, ObjectID: "idc", PropertyID: "around_traffic_routes", PropertyIndex: 10, PropertyName: "周边交通线路", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeList, Option: []string{"101路公交", "102路公交", "103路公交", "104路公交"}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 198, ObjectID: "low_voltage_switchgear", PropertyID: "output_voltage", PropertyIndex: 3, PropertyName: "输出电压", Unit: "v", PropertyGroup: "1b99731e-bbb3-4258-b7fc-2cc3be29185b", PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 199, ObjectID: "building", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 200, ObjectID: "building", PropertyID: "name", PropertyIndex: 1, PropertyName: "楼栋名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 201, ObjectID: "building", PropertyID: "area", PropertyIndex: 2, PropertyName: "建筑面积", Unit: "平方", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 202, ObjectID: "building", PropertyID: "floor_high", PropertyIndex: 3, PropertyName: "层高", Unit: "m", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 203, ObjectID: "refrigeration_pump", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 204, ObjectID: "refrigeration_pump", PropertyID: "manufacturer", PropertyIndex: 1, PropertyName: "厂商", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 205, ObjectID: "building", PropertyID: "building_structure_type", PropertyIndex: 4, PropertyName: "建筑结构类型", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{IsDefault: true, ID: "1", Type: "text", Name: "钢筋混凝土结构"}}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 206, ObjectID: "refrigeration_pump", PropertyID: "model", PropertyIndex: 2, PropertyName: "型号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 207, ObjectID: "refrigeration_pump", PropertyID: "SN", PropertyIndex: 3, PropertyName: "SN", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 208, ObjectID: "refrigeration_pump", PropertyID: "nominal_voltage", PropertyIndex: 2, PropertyName: "标称电压", PropertyGroup: "18746577-f5f5-4b8f-972f-82d24a3a5139", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 209, ObjectID: "building", PropertyID: "seismic_grade", PropertyIndex: 5, PropertyName: "抗震等级", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{IsDefault: false, ID: "1", Type: "text", Name: "1级"}, {IsDefault: false, ID: "2", Type: "text", Name: "2级"}, {IsDefault: false, ID: "3", Type: "text", Name: "3级"}, {IsDefault: false, ID: "4", Type: "text", Name: "4级"}, {IsDefault: false, ID: "5", Type: "text", Name: "5级"}, {IsDefault: false, ID: "6", Type: "text", Name: "6级"}, {IsDefault: false, ID: "7", Type: "text", Name: "7级"}, {IsDefault: true, ID: "8", Type: "text", Name: "8级"}, {IsDefault: false, ID: "9", Type: "text", Name: "9级"}}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 210, ObjectID: "refrigeration_pump", PropertyID: "rated_speed", PropertyIndex: 1, PropertyName: "额定转速", PropertyGroup: "18746577-f5f5-4b8f-972f-82d24a3a5139", PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 211, ObjectID: "floor", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: false, IsEditable: true, IsPre: true},
		//&Attribute{ID: 212, ObjectID: "floor", PropertyID: "name", PropertyIndex: 1, PropertyName: "楼层名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 213, ObjectID: "floor", PropertyID: "build_time", PropertyIndex: 2, PropertyName: "建设日期", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeDate, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 214, ObjectID: "floor", PropertyID: "prod_time", PropertyIndex: 3, PropertyName: "投产日期", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeDate, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 215, ObjectID: "modular_UPS", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: false, IsEditable: true, IsPre: true},
		//&Attribute{ID: 216, ObjectID: "floor", PropertyID: "handle_weight", PropertyIndex: 4, PropertyName: "楼层承重", Unit: "吨", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 217, ObjectID: "modular_UPS", PropertyID: "manufacturer", PropertyIndex: 1, PropertyName: "厂商", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 218, ObjectID: "floor", PropertyID: "function", PropertyIndex: 5, PropertyName: "功能", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 219, ObjectID: "modular_UPS", PropertyID: "model", PropertyIndex: 2, PropertyName: "型号", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 220, ObjectID: "room", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 221, ObjectID: "modular_UPS", PropertyID: "SN", PropertyIndex: 3, PropertyName: "SN", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 222, ObjectID: "room", PropertyID: "name", PropertyIndex: 1, PropertyName: "房间名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 224, ObjectID: "room", PropertyID: "is_cold_channel", PropertyIndex: 3, PropertyName: "是否冷通道", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeBool, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 225, ObjectID: "room", PropertyID: "customer_name", PropertyIndex: 4, PropertyName: "主要客户名称", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeOrganization, Option: "", IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 226, ObjectID: "line", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 227, ObjectID: "cabinet", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
		//&Attribute{ID: 230, ObjectID: "line", PropertyID: "manufacturer", PropertyIndex: 1, PropertyName: "厂商", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 231, ObjectID: "cabinet", PropertyID: "depth", PropertyIndex: 1, PropertyName: "机柜深度", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 232, ObjectID: "cabinet", PropertyID: "hight", PropertyIndex: 2, PropertyName: "机柜高度", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeFloat, Option: metadata.FloatOption{}, IsAPI: false, IsRequired: false, IsOnly: false, IsEditable: true, IsPre: false},
		//&Attribute{ID: 234, ObjectID: "ww", PropertyID: "bk_inst_name", PropertyIndex: 0, PropertyName: "实例名", PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: nil, IsAPI: false, IsRequired: true, IsOnly: true, IsEditable: true, IsPre: true},
	}
	return dataRows
}

// ModuleRow module structure
func ModuleRow() []*Attribute {
	objID := common.BKInnerObjIDModule
	moduleTypeOption := []metadata.EnumVal{{ID: "1", Name: "普通", Type: "text", IsDefault: true}, {ID: "2", Name: "数据库", Type: "text"}}

	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: common.BKAppIDField, PropertyName: "业务ID", IsAPI: true, IsRequired: false, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKSetIDField, PropertyName: "集群ID", IsAPI: true, IsRequired: false, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKModuleNameField, PropertyName: "模块名", IsRequired: true, IsOnly: true, IsEditable: true, PropertyType: common.FieldTypeSingleChar, Option: "", PropertyGroup: groupBaseInfo},
		//&Attribute{ObjectID: objID, PropertyID: "bk_childid", PropertyName: "", IsRequired: false, IsOnly: false, IsSystem: true, PropertyType: "", Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_module_type", PropertyName: "模块类型", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: moduleTypeOption},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "主要维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_bak_operator", PropertyName: "备份维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
	}
	return dataRows
}

// PlatRow plat structure
func PlatRow() []*Attribute {
	objID := common.BKInnerObjIDPlat
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: common.BKCloudNameField, PropertyName: "云区域", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKOwnerIDField, PropertyName: "供应商", IsRequired: true, IsOnly: true, IsPre: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
	}
	return dataRows
}

// HostRow host structure
func HostRow() []*Attribute {
	objID := common.BKInnerObjIDHost
	groupAgent := mCommon.HostAutoFields
	dataRows := []*Attribute{
		//基本信息分组
		&Attribute{ObjectID: objID, PropertyID: common.BKHostInnerIPField, PropertyName: "内网IP", IsRequired: true, IsOnly: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: common.BKHostOuterIPField, PropertyName: "外网IP", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "operator", PropertyName: "主要维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_bak_operator", PropertyName: "备份维护人", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeUser, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_asset_id", PropertyName: "固资编号", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_sn", PropertyName: "设备SN", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_comment", PropertyName: "备注", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_service_term", PropertyName: "质保年限", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "10"}},
		&Attribute{ObjectID: objID, PropertyID: "bk_sla", PropertyName: "SLA级别", IsRequired: false, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "L1", Type: "text"}, {ID: "2", Name: "L2", Type: "text"}, {ID: "3", Name: "L3", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: common.BKCloudIDField, PropertyName: "云区域", IsRequired: false, IsOnly: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: "singleasst", Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_state_name", PropertyName: "所在国家", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: stateEnum},
		&Attribute{ObjectID: objID, PropertyID: "bk_province_name", PropertyName: "所在省份", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: provincesEnum},
		&Attribute{ObjectID: objID, PropertyID: "bk_isp_name", PropertyName: "所属运营商", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: ispNameEnum},

		//自动发现分组
		&Attribute{ObjectID: objID, PropertyID: "bk_host_name", PropertyName: "主机名称", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: common.BKOSTypeField, PropertyName: "操作系统类型", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "Linux", Type: "text"}, {ID: "2", Name: "Windows", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: common.BKOSNameField, PropertyName: "操作系统名称", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_version", PropertyName: "操作系统版本", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_bit", PropertyName: "操作系统位数", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_cpu", PropertyName: "CPU逻辑核心数", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "1000000"}},
		&Attribute{ObjectID: objID, PropertyID: "bk_cpu_mhz", PropertyName: "CPU频率", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeInt, Unit: "Hz", Option: metadata.IntOption{Min: "1", Max: "100000000"}},
		&Attribute{ObjectID: objID, PropertyID: "bk_cpu_module", PropertyName: "CPU型号", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_mem", PropertyName: "内存容量", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeInt, Unit: "MB", Option: metadata.IntOption{Min: "1", Max: "100000000"}},
		&Attribute{ObjectID: objID, PropertyID: "bk_disk", PropertyName: "磁盘容量", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeInt, Unit: "GB", Option: metadata.IntOption{Min: "1", Max: "100000000"}},
		&Attribute{ObjectID: objID, PropertyID: "bk_mac", PropertyName: "内网MAC地址", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_outer_mac", PropertyName: "外网MAC", IsRequired: false, IsOnly: false, PropertyGroup: groupAgent, PropertyType: common.FieldTypeSingleChar, Option: ""},

		//agent default
		&Attribute{ObjectID: objID, PropertyID: common.CreateTimeField, PropertyName: "录入时间", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeTime, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "import_from", PropertyName: "录入方式", IsRequired: false, IsOnly: false, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "excel", Type: "text"}, {ID: "2", Name: "agent", Type: "text"}, {ID: "3", Name: "api", Type: "text"}}},
	}

	return dataRows
}

// ProcRow proc structure
func ProcRow() []*Attribute {
	objID := common.BKInnerObjIDProc
	groupPort := mCommon.ProcPort
	// groupGsekit := mCommon.Proc_gsekit_base_info
	// groupGsekitManage := mCommon.Proc_gsekit_manage_info
	dataRows := []*Attribute{
		//base info
		//&Attribute{ObjectID: objID, PropertyID: "bk_process_id", PropertyName: "进程ID", IsSystem: true, IsRequired: true, IsOnly: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: "{}"},
		&Attribute{ObjectID: objID, PropertyID: common.BKAppIDField, PropertyName: "业务ID", IsAPI: true, IsRequired: true, IsOnly: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{}},
		&Attribute{ObjectID: objID, PropertyID: common.BKProcessNameField, PropertyName: "进程名称", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "description", PropertyName: "进程描述", IsRequired: false, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},

		//监听端口分组
		&Attribute{ObjectID: objID, PropertyID: "bind_ip", PropertyName: "绑定IP", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "127.0.0.1", Type: "text"}, {ID: "2", Name: "0.0.0.0", Type: "text"}, {ID: "3", Name: "第一内网IP", Type: "text"}, {ID: "4", Name: "第一外网IP", Type: "text"}}},
		&Attribute{ObjectID: objID, PropertyID: "port", PropertyName: "端口", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultiplePortRange, Placeholder: `单个端口：8080 </br>多个连续端口：8080-8089 </br>多个不连续端口：8080-8089,8199`},
		&Attribute{ObjectID: objID, PropertyID: "protocol", PropertyName: "协议", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: groupPort, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "TCP", Type: "text"}, {ID: "2", Name: "UDP", Type: "text"}}},

		//gsekit 基础信息
		&Attribute{ObjectID: objID, PropertyID: "bk_func_id", PropertyName: "功能ID", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_func_name", PropertyName: "功能名称", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "work_path", PropertyName: "工作路径", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "user", PropertyName: "启动用户", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "proc_num", PropertyName: "启动数量", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "1000000"}},
		&Attribute{ObjectID: objID, PropertyID: "priority", PropertyName: "启动优先级", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "100"}},
		&Attribute{ObjectID: objID, PropertyID: "timeout", PropertyName: "操作超时时长", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "1000000"}},

		//gsekit 进程信息
		&Attribute{ObjectID: objID, PropertyID: "start_cmd", PropertyName: "启动命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "stop_cmd", PropertyName: "停止命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "restart_cmd", PropertyName: "重启命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "face_stop_cmd", PropertyName: "强制停止命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "reload_cmd", PropertyName: "进程重载命令", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "pid_file", PropertyName: "PID文件路径", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "auto_start", PropertyName: "是否自动拉起", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeBool, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "auto_time_gap", PropertyName: "拉起间隔", IsRequired: false, IsOnly: false, IsEditable: true, PropertyGroup: mCommon.GroupNone, PropertyType: common.FieldTypeInt, Option: metadata.IntOption{Min: "1", Max: "1000000"}},
	}
	return dataRows
}

// SwitchRow proc structure
func SwitchRow() []*Attribute {
	objID := common.BKInnerObjIDSwitch
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "bk_asset_id", PropertyName: "固资编号", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_sn", PropertyName: "SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_func", PropertyName: "用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "bk_operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_status", PropertyName: "运营状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "待运营", Type: "text"}, {ID: "2", Name: "运营中", Type: "text"}, {ID: "3", Name: "已下架", Type: "text"}}},
	}
	return dataRows
}

// RouterRow proc structure
func RouterRow() []*Attribute {
	objID := common.BKInnerObjIDRouter
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "bk_asset_id", PropertyName: "固资编号", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_sn", PropertyName: "SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_func", PropertyName: "用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "bk_operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_status", PropertyName: "运营状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "待运营", Type: "text"}, {ID: "2", Name: "运营中", Type: "text"}, {ID: "3", Name: "已下架", Type: "text"}}},
	}
	return dataRows
}

// LoadBalanceRow proc structure
func LoadBalanceRow() []*Attribute {
	objID := common.BKInnerObjIDBlance
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "bk_asset_id", PropertyName: "固资编号", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_sn", PropertyName: "SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_func", PropertyName: "用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "bk_operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_status", PropertyName: "运营状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "待运营", Type: "text"}, {ID: "2", Name: "运营中", Type: "text"}, {ID: "3", Name: "已下架", Type: "text"}}},
	}
	return dataRows
}

// FirewallRow proc structure
func FirewallRow() []*Attribute {
	objID := common.BKInnerObjIDFirewall
	dataRows := []*Attribute{
		&Attribute{ObjectID: objID, PropertyID: "bk_asset_id", PropertyName: "固资编号", IsRequired: true, IsOnly: true, IsPre: true, IsEditable: false, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_inst_name", PropertyName: "名称", IsRequired: true, IsOnly: false, IsPre: true, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_sn", PropertyName: "SN", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_func", PropertyName: "用途", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_vendor", PropertyName: "厂商", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_model", PropertyName: "设备型号", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_admin_ip", PropertyName: "管理IP", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: common.PatternMultipleIP},
		&Attribute{ObjectID: objID, PropertyID: "bk_operator", PropertyName: "维护人", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_os_detail", PropertyName: "操作系统详情", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeSingleChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_detail", PropertyName: "详细描述", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeLongChar, Option: ""},
		&Attribute{ObjectID: objID, PropertyID: "bk_biz_status", PropertyName: "运营状态", IsRequired: false, IsOnly: false, IsPre: false, IsEditable: true, PropertyGroup: groupBaseInfo, PropertyType: common.FieldTypeEnum, Option: []metadata.EnumVal{{ID: "1", Name: "待运营", Type: "text"}, {ID: "2", Name: "运营中", Type: "text"}, {ID: "3", Name: "已下架", Type: "text"}}},
	}
	return dataRows
}

var stateEnum = []metadata.EnumVal{
	{ID: "AR", Name: "阿根廷", Type: "text"},
	{ID: "AD", Name: "安道尔", Type: "text"},
	{ID: "AE", Name: "阿联酋", Type: "text"},
	{ID: "AF", Name: "阿富汗", Type: "text"},
	{ID: "AG", Name: "安提瓜和巴布达", Type: "text"},
	{ID: "AI", Name: "安圭拉", Type: "text"},
	{ID: "AL", Name: "阿尔巴尼亚", Type: "text"},
	{ID: "AM", Name: "亚美尼亚", Type: "text"},
	{ID: "AO", Name: "安哥拉", Type: "text"},
	{ID: "AQ", Name: "南极洲", Type: "text"},
	{ID: "AS", Name: "美属萨摩亚", Type: "text"},
	{ID: "AT", Name: "奥地利", Type: "text"},
	{ID: "AU", Name: "澳大利亚", Type: "text"},
	{ID: "AW", Name: "阿鲁巴", Type: "text"},
	{ID: "AX", Name: "奥兰群岛", Type: "text"},
	{ID: "AZ", Name: "阿塞拜疆", Type: "text"},
	{ID: "BA", Name: "波黑", Type: "text"},
	{ID: "BB", Name: "巴巴多斯", Type: "text"},
	{ID: "BD", Name: "孟加拉", Type: "text"},
	{ID: "BE", Name: "比利时", Type: "text"},
	{ID: "BF", Name: "布基纳法索", Type: "text"},
	{ID: "BG", Name: "保加利亚", Type: "text"},
	{ID: "BH", Name: "巴林", Type: "text"},
	{ID: "BI", Name: "布隆迪", Type: "text"},
	{ID: "BJ", Name: "贝宁", Type: "text"},
	{ID: "BL", Name: "圣巴泰勒米岛", Type: "text"},
	{ID: "BM", Name: "百慕大", Type: "text"},
	{ID: "BN", Name: "文莱", Type: "text"},
	{ID: "BO", Name: "玻利维亚", Type: "text"},
	{ID: "BQ", Name: "荷兰加勒比区", Type: "text"},
	{ID: "BR", Name: "巴西", Type: "text"},
	{ID: "BS", Name: "巴哈马", Type: "text"},
	{ID: "BT", Name: "不丹", Type: "text"},
	{ID: "BV", Name: "布韦岛", Type: "text"},
	{ID: "BW", Name: "博茨瓦纳", Type: "text"},
	{ID: "BY", Name: "白俄罗斯", Type: "text"},
	{ID: "BZ", Name: "伯利兹", Type: "text"},
	{ID: "CA", Name: "加拿大", Type: "text"},
	{ID: "CC", Name: "科科斯群岛", Type: "text"},
	{ID: "CD", Name: "刚果（金）", Type: "text"},
	{ID: "CF", Name: "中非", Type: "text"},
	{ID: "CG", Name: "刚果（布）", Type: "text"},
	{ID: "CH", Name: "瑞士", Type: "text"},
	{ID: "CI", Name: "科特迪瓦", Type: "text"},
	{ID: "CK", Name: "库克群岛", Type: "text"},
	{ID: "CL", Name: "智利", Type: "text"},
	{ID: "CM", Name: "喀麦隆", Type: "text"},
	{ID: "CN", Name: "中国", Type: "text"},
	{ID: "CO", Name: "哥伦比亚", Type: "text"},
	{ID: "CR", Name: "哥斯达黎加", Type: "text"},
	{ID: "CU", Name: "古巴", Type: "text"},
	{ID: "CV", Name: "佛得角", Type: "text"},
	{ID: "CW", Name: "库拉索", Type: "text"},
	{ID: "CX", Name: "圣诞岛", Type: "text"},
	{ID: "CY", Name: "塞浦路斯", Type: "text"},
	{ID: "CZ", Name: "捷克", Type: "text"},
	{ID: "DE", Name: "德国", Type: "text"},
	{ID: "DJ", Name: "吉布提", Type: "text"},
	{ID: "DK", Name: "丹麦", Type: "text"},
	{ID: "DM", Name: "多米尼克", Type: "text"},
	{ID: "DO", Name: "多米尼加", Type: "text"},
	{ID: "DZ", Name: "阿尔及利亚", Type: "text"},
	{ID: "EC", Name: "厄瓜多尔", Type: "text"},
	{ID: "EE", Name: "爱沙尼亚", Type: "text"},
	{ID: "EG", Name: "埃及", Type: "text"},
	{ID: "EH", Name: "西撒哈拉", Type: "text"},
	{ID: "ER", Name: "厄立特里亚", Type: "text"},
	{ID: "ES", Name: "西班牙", Type: "text"},
	{ID: "ET", Name: "埃塞俄比亚", Type: "text"},
	{ID: "FI", Name: "芬兰", Type: "text"},
	{ID: "FJ", Name: "斐济群岛", Type: "text"},
	{ID: "FK", Name: "马尔维纳斯群岛（福克兰）", Type: "text"},
	{ID: "FM", Name: "密克罗尼西亚联邦", Type: "text"},
	{ID: "FO", Name: "法罗群岛", Type: "text"},
	{ID: "FR", Name: "法国", Type: "text"},
	{ID: "GA", Name: "加蓬", Type: "text"},
	{ID: "GB", Name: "英国", Type: "text"},
	{ID: "GD", Name: "格林纳达", Type: "text"},
	{ID: "GE", Name: "格鲁吉亚", Type: "text"},
	{ID: "GF", Name: "法属圭亚那", Type: "text"},
	{ID: "GG", Name: "根西岛", Type: "text"},
	{ID: "GH", Name: "加纳", Type: "text"},
	{ID: "GI", Name: "直布罗陀", Type: "text"},
	{ID: "GL", Name: "格陵兰", Type: "text"},
	{ID: "GM", Name: "冈比亚", Type: "text"},
	{ID: "GN", Name: "几内亚", Type: "text"},
	{ID: "GP", Name: "瓜德罗普", Type: "text"},
	{ID: "GQ", Name: "赤道几内亚", Type: "text"},
	{ID: "GR", Name: "希腊", Type: "text"},
	{ID: "GS", Name: "南乔治亚岛和南桑威奇群岛", Type: "text"},
	{ID: "GT", Name: "危地马拉", Type: "text"},
	{ID: "GU", Name: "关岛", Type: "text"},
	{ID: "GW", Name: "几内亚比绍", Type: "text"},
	{ID: "GY", Name: "圭亚那", Type: "text"},
	{ID: "HM", Name: "赫德岛和麦克唐纳群岛", Type: "text"},
	{ID: "HN", Name: "洪都拉斯", Type: "text"},
	{ID: "HR", Name: "克罗地亚", Type: "text"},
	{ID: "HT", Name: "海地", Type: "text"},
	{ID: "HU", Name: "匈牙利", Type: "text"},
	{ID: "ID", Name: "印尼", Type: "text"},
	{ID: "IE", Name: "爱尔兰", Type: "text"},
	{ID: "IL", Name: "以色列", Type: "text"},
	{ID: "IM", Name: "马恩岛", Type: "text"},
	{ID: "IN", Name: "印度", Type: "text"},
	{ID: "IO", Name: "英属印度洋领地", Type: "text"},
	{ID: "IQ", Name: "伊拉克", Type: "text"},
	{ID: "IR", Name: "伊朗", Type: "text"},
	{ID: "IS", Name: "冰岛", Type: "text"},
	{ID: "IT", Name: "意大利", Type: "text"},
	{ID: "JE", Name: "泽西岛", Type: "text"},
	{ID: "JM", Name: "牙买加", Type: "text"},
	{ID: "JO", Name: "约旦", Type: "text"},
	{ID: "JP", Name: "日本", Type: "text"},
	{ID: "KE", Name: "肯尼亚", Type: "text"},
	{ID: "KG", Name: "吉尔吉斯斯坦", Type: "text"},
	{ID: "KH", Name: "柬埔寨", Type: "text"},
	{ID: "KI", Name: "基里巴斯", Type: "text"},
	{ID: "KM", Name: "科摩罗", Type: "text"},
	{ID: "KN", Name: "圣基茨和尼维斯", Type: "text"},
	{ID: "KP", Name: "朝鲜", Type: "text"},
	{ID: "KR", Name: "韩国", Type: "text"},
	{ID: "KW", Name: "科威特", Type: "text"},
	{ID: "KY", Name: "开曼群岛", Type: "text"},
	{ID: "KZ", Name: "哈萨克斯坦", Type: "text"},
	{ID: "LA", Name: "老挝", Type: "text"},
	{ID: "LB", Name: "黎巴嫩", Type: "text"},
	{ID: "LC", Name: "圣卢西亚", Type: "text"},
	{ID: "LI", Name: "列支敦士登", Type: "text"},
	{ID: "LK", Name: "斯里兰卡", Type: "text"},
	{ID: "LR", Name: "利比里亚", Type: "text"},
	{ID: "LS", Name: "莱索托", Type: "text"},
	{ID: "LT", Name: "立陶宛", Type: "text"},
	{ID: "LU", Name: "卢森堡", Type: "text"},
	{ID: "LV", Name: "拉脱维亚", Type: "text"},
	{ID: "LY", Name: "利比亚", Type: "text"},
	{ID: "MA", Name: "摩洛哥", Type: "text"},
	{ID: "MC", Name: "摩纳哥", Type: "text"},
	{ID: "MD", Name: "摩尔多瓦", Type: "text"},
	{ID: "ME", Name: "黑山", Type: "text"},
	{ID: "MF", Name: "法属圣马丁", Type: "text"},
	{ID: "MG", Name: "马达加斯加", Type: "text"},
	{ID: "MH", Name: "马绍尔群岛", Type: "text"},
	{ID: "MK", Name: "马其顿", Type: "text"},
	{ID: "ML", Name: "马里", Type: "text"},
	{ID: "MM", Name: "缅甸", Type: "text"},
	{ID: "MN", Name: "蒙古国", Type: "text"},
	{ID: "MP", Name: "北马里亚纳群岛", Type: "text"},
	{ID: "MQ", Name: "马提尼克", Type: "text"},
	{ID: "MR", Name: "毛里塔尼亚", Type: "text"},
	{ID: "MS", Name: "蒙塞拉特岛", Type: "text"},
	{ID: "MT", Name: "马耳他", Type: "text"},
	{ID: "MU", Name: "毛里求斯", Type: "text"},
	{ID: "MV", Name: "马尔代夫", Type: "text"},
	{ID: "MW", Name: "马拉维", Type: "text"},
	{ID: "MX", Name: "墨西哥", Type: "text"},
	{ID: "MY", Name: "马来西亚", Type: "text"},
	{ID: "MZ", Name: "莫桑比克", Type: "text"},
	{ID: "NA", Name: "纳米比亚", Type: "text"},
	{ID: "NC", Name: "新喀里多尼亚", Type: "text"},
	{ID: "NE", Name: "尼日尔", Type: "text"},
	{ID: "NF", Name: "诺福克岛", Type: "text"},
	{ID: "NG", Name: "尼日利亚", Type: "text"},
	{ID: "NI", Name: "尼加拉瓜", Type: "text"},
	{ID: "NL", Name: "荷兰", Type: "text"},
	{ID: "NO", Name: "挪威", Type: "text"},
	{ID: "NP", Name: "尼泊尔", Type: "text"},
	{ID: "NR", Name: "瑙鲁", Type: "text"},
	{ID: "NU", Name: "纽埃", Type: "text"},
	{ID: "NZ", Name: "新西兰", Type: "text"},
	{ID: "OM", Name: "阿曼", Type: "text"},
	{ID: "PA", Name: "巴拿马", Type: "text"},
	{ID: "PE", Name: "秘鲁", Type: "text"},
	{ID: "PF", Name: "法属波利尼西亚", Type: "text"},
	{ID: "PG", Name: "巴布亚新几内亚", Type: "text"},
	{ID: "PH", Name: "菲律宾", Type: "text"},
	{ID: "PK", Name: "巴基斯坦", Type: "text"},
	{ID: "PL", Name: "波兰", Type: "text"},
	{ID: "PM", Name: "圣皮埃尔和密克隆", Type: "text"},
	{ID: "PN", Name: "皮特凯恩群岛", Type: "text"},
	{ID: "PR", Name: "波多黎各", Type: "text"},
	{ID: "PS", Name: "巴勒斯坦", Type: "text"},
	{ID: "PT", Name: "葡萄牙", Type: "text"},
	{ID: "PW", Name: "帕劳", Type: "text"},
	{ID: "PY", Name: "巴拉圭", Type: "text"},
	{ID: "QA", Name: "卡塔尔", Type: "text"},
	{ID: "RE", Name: "留尼汪", Type: "text"},
	{ID: "RO", Name: "罗马尼亚", Type: "text"},
	{ID: "RS", Name: "塞尔维亚", Type: "text"},
	{ID: "RU", Name: "俄罗斯", Type: "text"},
	{ID: "RW", Name: "卢旺达", Type: "text"},
	{ID: "SA", Name: "沙特阿拉伯", Type: "text"},
	{ID: "SB", Name: "所罗门群岛", Type: "text"},
	{ID: "SC", Name: "塞舌尔", Type: "text"},
	{ID: "SD", Name: "苏丹", Type: "text"},
	{ID: "SE", Name: "瑞典", Type: "text"},
	{ID: "SG", Name: "新加坡", Type: "text"},
	{ID: "SH", Name: "圣赫勒拿", Type: "text"},
	{ID: "SI", Name: "斯洛文尼亚", Type: "text"},
	{ID: "SJ", Name: "斯瓦尔巴群岛和扬马延岛", Type: "text"},
	{ID: "SK", Name: "斯洛伐克", Type: "text"},
	{ID: "SL", Name: "塞拉利昂", Type: "text"},
	{ID: "SM", Name: "圣马力诺", Type: "text"},
	{ID: "SN", Name: "塞内加尔", Type: "text"},
	{ID: "SO", Name: "索马里", Type: "text"},
	{ID: "SR", Name: "苏里南", Type: "text"},
	{ID: "SS", Name: "南苏丹", Type: "text"},
	{ID: "ST", Name: "圣多美和普林西比", Type: "text"},
	{ID: "SV", Name: "萨尔瓦多", Type: "text"},
	{ID: "SX", Name: "荷属圣马丁", Type: "text"},
	{ID: "SY", Name: "叙利亚", Type: "text"},
	{ID: "SZ", Name: "斯威士兰", Type: "text"},
	{ID: "TC", Name: "特克斯和凯科斯群岛", Type: "text"},
	{ID: "TD", Name: "乍得", Type: "text"},
	{ID: "TF", Name: "法属南部领地", Type: "text"},
	{ID: "TG", Name: "多哥", Type: "text"},
	{ID: "TH", Name: "泰国", Type: "text"},
	{ID: "TJ", Name: "塔吉克斯坦", Type: "text"},
	{ID: "TK", Name: "托克劳", Type: "text"},
	{ID: "TL", Name: "东帝汶", Type: "text"},
	{ID: "TM", Name: "土库曼斯坦", Type: "text"},
	{ID: "TN", Name: "突尼斯", Type: "text"},
	{ID: "TO", Name: "汤加", Type: "text"},
	{ID: "TR", Name: "土耳其", Type: "text"},
	{ID: "TT", Name: "特立尼达和多巴哥", Type: "text"},
	{ID: "TV", Name: "图瓦卢", Type: "text"},
	{ID: "TZ", Name: "坦桑尼亚", Type: "text"},
	{ID: "UA", Name: "乌克兰", Type: "text"},
	{ID: "UG", Name: "乌干达", Type: "text"},
	{ID: "UM", Name: "美国本土外小岛屿", Type: "text"},
	{ID: "UY", Name: "乌拉圭", Type: "text"},
	{ID: "UZ", Name: "乌兹别克斯坦", Type: "text"},
	{ID: "VA", Name: "梵蒂冈", Type: "text"},
	{ID: "VC", Name: "圣文森特和格林纳丁斯", Type: "text"},
	{ID: "VE", Name: "委内瑞拉", Type: "text"},
	{ID: "VG", Name: "英属维尔京群岛", Type: "text"},
	{ID: "VI", Name: "美属维尔京群岛", Type: "text"},
	{ID: "VN", Name: "越南", Type: "text"},
	{ID: "US", Name: "美国", Type: "text"},
	{ID: "VU", Name: "瓦努阿图", Type: "text"},
	{ID: "WF", Name: "瓦利斯和富图纳", Type: "text"},
	{ID: "WS", Name: "萨摩亚", Type: "text"},
	{ID: "YE", Name: "也门", Type: "text"},
	{ID: "YT", Name: "马约特", Type: "text"},
	{ID: "ZA", Name: "南非", Type: "text"},
	{ID: "ZM", Name: "赞比亚", Type: "text"},
	{ID: "ZW", Name: "津巴布韦", Type: "text"},
}

var provincesEnum = []metadata.EnumVal{
	{ID: "110000", Name: "北京市", Type: "text"},
	{ID: "120000", Name: "天津市", Type: "text"},
	{ID: "130000", Name: "河北省", Type: "text"},
	{ID: "140000", Name: "山西省", Type: "text"},
	{ID: "150000", Name: "内蒙古自治区", Type: "text"},
	{ID: "210000", Name: "辽宁省", Type: "text"},
	{ID: "220000", Name: "吉林省", Type: "text"},
	{ID: "230000", Name: "黑龙江省", Type: "text"},
	{ID: "310000", Name: "上海市", Type: "text"},
	{ID: "320000", Name: "江苏省", Type: "text"},
	{ID: "330000", Name: "浙江省", Type: "text"},
	{ID: "340000", Name: "安徽省", Type: "text"},
	{ID: "350000", Name: "福建省", Type: "text"},
	{ID: "360000", Name: "江西省", Type: "text"},
	{ID: "370000", Name: "山东省", Type: "text"},
	{ID: "410000", Name: "河南省", Type: "text"},
	{ID: "420000", Name: "湖北省", Type: "text"},
	{ID: "430000", Name: "湖南省", Type: "text"},
	{ID: "440000", Name: "广东省", Type: "text"},
	{ID: "450000", Name: "广西壮族自治区", Type: "text"},
	{ID: "460000", Name: "海南省", Type: "text"},
	{ID: "500000", Name: "重庆市", Type: "text"},
	{ID: "510000", Name: "四川省", Type: "text"},
	{ID: "520000", Name: "贵州省", Type: "text"},
	{ID: "530000", Name: "云南省", Type: "text"},
	{ID: "540000", Name: "西藏自治区", Type: "text"},
	{ID: "610000", Name: "陕西省", Type: "text"},
	{ID: "620000", Name: "甘肃省", Type: "text"},
	{ID: "630000", Name: "青海省", Type: "text"},
	{ID: "640000", Name: "宁夏回族自治区", Type: "text"},
	{ID: "650000", Name: "新疆维吾尔自治区", Type: "text"},
	{ID: "710000", Name: "台湾省", Type: "text"},
	{ID: "810000", Name: "香港特别行政区", Type: "text"},
	{ID: "820000", Name: "澳门特别行政区", Type: "text"},
}

var ispNameEnum = []metadata.EnumVal{
	{ID: "0", Name: "其他", Type: "text"},
	{ID: "1", Name: "电信", Type: "text"}, {ID: "2", Name: "联通", Type: "text"}, {ID: "3", Name: "移动", Type: "text"},
}

type Attribute struct {
	ID                int64       `field:"id" json:"id" bson:"id"`
	OwnerID           string      `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"`
	ObjectID          string      `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
	PropertyID        string      `field:"bk_property_id" json:"bk_property_id" bson:"bk_property_id"`
	PropertyName      string      `field:"bk_property_name" json:"bk_property_name" bson:"bk_property_name"`
	PropertyGroup     string      `field:"bk_property_group" json:"bk_property_group" bson:"bk_property_group"`
	PropertyGroupName string      `field:"bk_property_group_name,ignoretomap" json:"bk_property_group_name" bson:"-"`
	PropertyIndex     int64       `field:"bk_property_index" json:"bk_property_index" bson:"bk_property_index"`
	Unit              string      `field:"unit" json:"unit" bson:"unit"`
	Placeholder       string      `field:"placeholder" json:"placeholder" bson:"placeholder"`
	IsEditable        bool        `field:"editable" json:"editable" bson:"editable"`
	IsPre             bool        `field:"ispre" json:"ispre" bson:"ispre"`
	IsRequired        bool        `field:"isrequired" json:"isrequired" bson:"isrequired"`
	IsReadOnly        bool        `field:"isreadonly" json:"isreadonly" bson:"isreadonly"`
	IsOnly            bool        `field:"isonly" json:"isonly" bson:"isonly"`
	IsSystem          bool        `field:"bk_issystem" json:"bk_issystem" bson:"bk_issystem"`
	IsAPI             bool        `field:"bk_isapi" json:"bk_isapi" bson:"bk_isapi"`
	PropertyType      string      `field:"bk_property_type" json:"bk_property_type" bson:"bk_property_type"`
	Option            interface{} `field:"option" json:"option" bson:"option"`
	Description       string      `field:"description" json:"description" bson:"description"`
	Creator           string      `field:"creator" json:"creator" bson:"creator"`
	CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
	LastTime          *time.Time  `json:"last_time" bson:"last_time"`
	BizID             int64       `bson:"bk_biz_id" json:"bk_biz_id" mapstructure:"bk_biz_id"`
}
