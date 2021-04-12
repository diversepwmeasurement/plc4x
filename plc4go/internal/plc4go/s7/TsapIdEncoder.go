//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package s7

import "github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"

func encodeS7TsapId(deviceGroup model.DeviceGroup, rack int32, slot int32) int32 {
	firstByte := int32(deviceGroup) << 8
	secondByte := (rack << 4) | (slot & 0x0F)
	return firstByte | secondByte
}

func decodeDeviceGroup(tsapId int32) model.DeviceGroup {
	deviceGroupCode := int8((tsapId >> 8) & (0xFF))
	return model.DeviceGroupByValue(deviceGroupCode)
}

func decodeRack(tsapId int32) int64 {
	return int64((tsapId >> 4) & 0xF)
}

func decodeSlot(tsapId int32) int64 {
	return int64(tsapId & 0xF)
}
