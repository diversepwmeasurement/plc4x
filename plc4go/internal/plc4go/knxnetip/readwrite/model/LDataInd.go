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
type LDataInd struct {
	*CEMI
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame

	// Arguments.
	Size uint16
}

// The corresponding interface
type ILDataInd interface {
	ICEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []*CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() *LDataFrame
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
func (m *LDataInd) GetMessageCode() uint8 {
	return 0x29
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *LDataInd) InitializeParent(parent *CEMI) {}

func (m *LDataInd) GetParent() *CEMI {
	return m.CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *LDataInd) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *LDataInd) GetAdditionalInformation() []*CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *LDataInd) GetDataFrame() *LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataInd factory function for LDataInd
func NewLDataInd(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame, size uint16) *LDataInd {
	_result := &LDataInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		CEMI:                        NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastLDataInd(structType interface{}) *LDataInd {
	if casted, ok := structType.(LDataInd); ok {
		return &casted
	}
	if casted, ok := structType.(*LDataInd); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastLDataInd(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastLDataInd(casted.Child)
	}
	return nil
}

func (m *LDataInd) GetTypeName() string {
	return "LDataInd"
}

func (m *LDataInd) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *LDataInd) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits()

	return lengthInBits
}

func (m *LDataInd) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LDataIndParse(readBuffer utils.ReadBuffer, size uint16) (*LDataInd, error) {
	if pullErr := readBuffer.PullContext("LDataInd"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (additionalInformationLength)
	_additionalInformationLength, _additionalInformationLengthErr := readBuffer.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}
	additionalInformationLength := _additionalInformationLength

	// Array field (additionalInformation)
	if pullErr := readBuffer.PullContext("additionalInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	{
		_additionalInformationLength := additionalInformationLength
		_additionalInformationEndPos := readBuffer.GetPos() + uint16(_additionalInformationLength)
		for readBuffer.GetPos() < _additionalInformationEndPos {
			_item, _err := CEMIAdditionalInformationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
			}
			additionalInformation = append(additionalInformation, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (dataFrame)
	if pullErr := readBuffer.PullContext("dataFrame"); pullErr != nil {
		return nil, pullErr
	}
	_dataFrame, _dataFrameErr := LDataFrameParse(readBuffer)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}
	dataFrame := CastLDataFrame(_dataFrame)
	if closeErr := readBuffer.CloseContext("dataFrame"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("LDataInd"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LDataInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   CastLDataFrame(dataFrame),
		CEMI:                        &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *LDataInd) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataInd"); pushErr != nil {
			return pushErr
		}

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := writeBuffer.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			if pushErr := writeBuffer.PushContext("additionalInformation", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
			if popErr := writeBuffer.PopContext("additionalInformation", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Simple Field (dataFrame)
		if pushErr := writeBuffer.PushContext("dataFrame"); pushErr != nil {
			return pushErr
		}
		_dataFrameErr := m.DataFrame.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dataFrame"); popErr != nil {
			return popErr
		}
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		if popErr := writeBuffer.PopContext("LDataInd"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *LDataInd) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
