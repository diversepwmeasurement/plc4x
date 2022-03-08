/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse struct {
	*S7PayloadUserDataItem
	Result     uint8
	Reserved01 uint8
	AlarmType  AlarmType
	Reserved02 uint8
	Reserved03 uint8
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse interface {
	IS7PayloadUserDataItem
	// GetResult returns Result (property field)
	GetResult() uint8
	// GetReserved01 returns Reserved01 (property field)
	GetReserved01() uint8
	// GetAlarmType returns AlarmType (property field)
	GetAlarmType() AlarmType
	// GetReserved02 returns Reserved02 (property field)
	GetReserved02() uint8
	// GetReserved03 returns Reserved03 (property field)
	GetReserved03() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetCpuSubfunction() uint8 {
	return 0x02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetDataLength() uint16 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetParent() *S7PayloadUserDataItem {
	return m.S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetResult() uint8 {
	return m.Result
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved01() uint8 {
	return m.Reserved01
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetAlarmType() AlarmType {
	return m.AlarmType
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved02() uint8 {
	return m.Reserved02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetReserved03() uint8 {
	return m.Reserved03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse factory function for S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
func NewS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(result uint8, reserved01 uint8, alarmType AlarmType, reserved02 uint8, reserved03 uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	_result := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		Result:                result,
		Reserved01:            reserved01,
		AlarmType:             alarmType,
		Reserved02:            reserved02,
		Reserved03:            reserved03,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(structType interface{}) *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse); ok {
		return casted
	}
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(casted.Child)
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return CastS7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 8

	// Simple field (reserved01)
	lengthInBits += 8

	// Simple field (alarmType)
	lengthInBits += 8

	// Simple field (reserved02)
	lengthInBits += 8

	// Simple field (reserved03)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (result)
	_result, _resultErr := readBuffer.ReadUint8("result", 8)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result

	// Simple Field (reserved01)
	_reserved01, _reserved01Err := readBuffer.ReadUint8("reserved01", 8)
	if _reserved01Err != nil {
		return nil, errors.Wrap(_reserved01Err, "Error parsing 'reserved01' field")
	}
	reserved01 := _reserved01

	// Simple Field (alarmType)
	if pullErr := readBuffer.PullContext("alarmType"); pullErr != nil {
		return nil, pullErr
	}
	_alarmType, _alarmTypeErr := AlarmTypeParse(readBuffer)
	if _alarmTypeErr != nil {
		return nil, errors.Wrap(_alarmTypeErr, "Error parsing 'alarmType' field")
	}
	alarmType := _alarmType
	if closeErr := readBuffer.CloseContext("alarmType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (reserved02)
	_reserved02, _reserved02Err := readBuffer.ReadUint8("reserved02", 8)
	if _reserved02Err != nil {
		return nil, errors.Wrap(_reserved02Err, "Error parsing 'reserved02' field")
	}
	reserved02 := _reserved02

	// Simple Field (reserved03)
	_reserved03, _reserved03Err := readBuffer.ReadUint8("reserved03", 8)
	if _reserved03Err != nil {
		return nil, errors.Wrap(_reserved03Err, "Error parsing 'reserved03' field")
	}
	reserved03 := _reserved03

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse{
		Result:                result,
		Reserved01:            reserved01,
		AlarmType:             alarmType,
		Reserved02:            reserved02,
		Reserved03:            reserved03,
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child, nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		result := uint8(m.Result)
		_resultErr := writeBuffer.WriteUint8("result", 8, (result))
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (reserved01)
		reserved01 := uint8(m.Reserved01)
		_reserved01Err := writeBuffer.WriteUint8("reserved01", 8, (reserved01))
		if _reserved01Err != nil {
			return errors.Wrap(_reserved01Err, "Error serializing 'reserved01' field")
		}

		// Simple Field (alarmType)
		if pushErr := writeBuffer.PushContext("alarmType"); pushErr != nil {
			return pushErr
		}
		_alarmTypeErr := m.AlarmType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmType"); popErr != nil {
			return popErr
		}
		if _alarmTypeErr != nil {
			return errors.Wrap(_alarmTypeErr, "Error serializing 'alarmType' field")
		}

		// Simple Field (reserved02)
		reserved02 := uint8(m.Reserved02)
		_reserved02Err := writeBuffer.WriteUint8("reserved02", 8, (reserved02))
		if _reserved02Err != nil {
			return errors.Wrap(_reserved02Err, "Error serializing 'reserved02' field")
		}

		// Simple Field (reserved03)
		reserved03 := uint8(m.Reserved03)
		_reserved03Err := writeBuffer.WriteUint8("reserved03", 8, (reserved03))
		if _reserved03Err != nil {
			return errors.Wrap(_reserved03Err, "Error serializing 'reserved03' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
