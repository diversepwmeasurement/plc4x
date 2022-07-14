/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataWrite is the corresponding interface of CALDataWrite
type CALDataWrite interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() Parameter
	// GetData returns Data (property field)
	GetData() []byte
}

// CALDataWriteExactly can be used when we want exactly this type and not a type which fulfills CALDataWrite.
// This is useful for switch cases.
type CALDataWriteExactly interface {
	CALDataWrite
	isCALDataWrite() bool
}

// _CALDataWrite is the data-structure of this message
type _CALDataWrite struct {
	*_CALData
	ParamNo Parameter
	Data    []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataWrite) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataWrite) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataWrite) GetParamNo() Parameter {
	return m.ParamNo
}

func (m *_CALDataWrite) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataWrite factory function for _CALDataWrite
func NewCALDataWrite(paramNo Parameter, data []byte, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataWrite {
	_result := &_CALDataWrite{
		ParamNo:  paramNo,
		Data:     data,
		_CALData: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataWrite(structType interface{}) CALDataWrite {
	if casted, ok := structType.(CALDataWrite); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataWrite); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataWrite) GetTypeName() string {
	return "CALDataWrite"
}

func (m *_CALDataWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CALDataWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataWriteParse(readBuffer utils.ReadBuffer, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	if pullErr := readBuffer.PullContext("paramNo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for paramNo")
	}
	_paramNo, _paramNoErr := ParameterParse(readBuffer)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataWrite")
	}
	paramNo := _paramNo
	if closeErr := readBuffer.CloseContext("paramNo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for paramNo")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(commandTypeContainer.NumBytes()) - uint16(uint16(1)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CALDataWrite")
	}

	if closeErr := readBuffer.CloseContext("CALDataWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataWrite")
	}

	// Create a partially initialized instance
	_child := &_CALDataWrite{
		ParamNo: paramNo,
		Data:    data,
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataWrite")
		}

		// Simple Field (paramNo)
		if pushErr := writeBuffer.PushContext("paramNo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for paramNo")
		}
		_paramNoErr := writeBuffer.WriteSerializable(m.GetParamNo())
		if popErr := writeBuffer.PopContext("paramNo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for paramNo")
		}
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("CALDataWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataWrite) isCALDataWrite() bool {
	return true
}

func (m *_CALDataWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
