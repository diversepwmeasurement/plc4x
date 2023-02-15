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
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// TelephonyDataRecallLastNumber is the corresponding interface of TelephonyDataRecallLastNumber
type TelephonyDataRecallLastNumber interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetRecallLastNumberType returns RecallLastNumberType (property field)
	GetRecallLastNumberType() byte
	// GetNumber returns Number (property field)
	GetNumber() string
	// GetIsNumberOfLastOutgoingCall returns IsNumberOfLastOutgoingCall (virtual field)
	GetIsNumberOfLastOutgoingCall() bool
	// GetIsNumberOfLastIncomingCall returns IsNumberOfLastIncomingCall (virtual field)
	GetIsNumberOfLastIncomingCall() bool
}

// TelephonyDataRecallLastNumberExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRecallLastNumber.
// This is useful for switch cases.
type TelephonyDataRecallLastNumberExactly interface {
	TelephonyDataRecallLastNumber
	isTelephonyDataRecallLastNumber() bool
}

// _TelephonyDataRecallLastNumber is the data-structure of this message
type _TelephonyDataRecallLastNumber struct {
	*_TelephonyData
        RecallLastNumberType byte
        Number string
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRecallLastNumber) InitializeParent(parent TelephonyData , commandTypeContainer TelephonyCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataRecallLastNumber)  GetParent() TelephonyData {
	return m._TelephonyData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumber) GetRecallLastNumberType() byte {
	return m.RecallLastNumberType
}

func (m *_TelephonyDataRecallLastNumber) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TelephonyDataRecallLastNumber) GetIsNumberOfLastOutgoingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x01)))
}

func (m *_TelephonyDataRecallLastNumber) GetIsNumberOfLastIncomingCall() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetRecallLastNumberType()) == (0x02)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewTelephonyDataRecallLastNumber factory function for _TelephonyDataRecallLastNumber
func NewTelephonyDataRecallLastNumber( recallLastNumberType byte , number string , commandTypeContainer TelephonyCommandTypeContainer , argument byte ) *_TelephonyDataRecallLastNumber {
	_result := &_TelephonyDataRecallLastNumber{
		RecallLastNumberType: recallLastNumberType,
		Number: number,
    	_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRecallLastNumber(structType interface{}) TelephonyDataRecallLastNumber {
    if casted, ok := structType.(TelephonyDataRecallLastNumber); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRecallLastNumber); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRecallLastNumber) GetTypeName() string {
	return "TelephonyDataRecallLastNumber"
}

func (m *_TelephonyDataRecallLastNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (recallLastNumberType)
	lengthInBits += 8;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(2)))) * int32(int32(8)))

	return lengthInBits
}


func (m *_TelephonyDataRecallLastNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataRecallLastNumberParse(theBytes []byte, commandTypeContainer TelephonyCommandTypeContainer) (TelephonyDataRecallLastNumber, error) {
	return TelephonyDataRecallLastNumberParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func TelephonyDataRecallLastNumberParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer TelephonyCommandTypeContainer) (TelephonyDataRecallLastNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRecallLastNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRecallLastNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recallLastNumberType)
_recallLastNumberType, _recallLastNumberTypeErr := readBuffer.ReadByte("recallLastNumberType")
	if _recallLastNumberTypeErr != nil {
		return nil, errors.Wrap(_recallLastNumberTypeErr, "Error parsing 'recallLastNumberType' field of TelephonyDataRecallLastNumber")
	}
	recallLastNumberType := _recallLastNumberType

	// Virtual field
	_isNumberOfLastOutgoingCall := bool((recallLastNumberType) == (0x01))
	isNumberOfLastOutgoingCall := bool(_isNumberOfLastOutgoingCall)
	_ = isNumberOfLastOutgoingCall

	// Virtual field
	_isNumberOfLastIncomingCall := bool((recallLastNumberType) == (0x02))
	isNumberOfLastIncomingCall := bool(_isNumberOfLastIncomingCall)
	_ = isNumberOfLastIncomingCall

	// Simple Field (number)
_number, _numberErr := readBuffer.ReadString("number", uint32((((commandTypeContainer.NumBytes()) - ((2)))) * ((8))), "UTF-8")
	if _numberErr != nil {
		return nil, errors.Wrap(_numberErr, "Error parsing 'number' field of TelephonyDataRecallLastNumber")
	}
	number := _number

	if closeErr := readBuffer.CloseContext("TelephonyDataRecallLastNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRecallLastNumber")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataRecallLastNumber{
		_TelephonyData: &_TelephonyData{
		},
		RecallLastNumberType: recallLastNumberType,
		Number: number,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataRecallLastNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRecallLastNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRecallLastNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRecallLastNumber")
		}

	// Simple Field (recallLastNumberType)
	recallLastNumberType := byte(m.GetRecallLastNumberType())
	_recallLastNumberTypeErr := writeBuffer.WriteByte("recallLastNumberType", (recallLastNumberType))
	if _recallLastNumberTypeErr != nil {
		return errors.Wrap(_recallLastNumberTypeErr, "Error serializing 'recallLastNumberType' field")
	}
	// Virtual field
	if _isNumberOfLastOutgoingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastOutgoingCall", m.GetIsNumberOfLastOutgoingCall()); _isNumberOfLastOutgoingCallErr != nil {
		return errors.Wrap(_isNumberOfLastOutgoingCallErr, "Error serializing 'isNumberOfLastOutgoingCall' field")
	}
	// Virtual field
	if _isNumberOfLastIncomingCallErr := writeBuffer.WriteVirtual(ctx, "isNumberOfLastIncomingCall", m.GetIsNumberOfLastIncomingCall()); _isNumberOfLastIncomingCallErr != nil {
		return errors.Wrap(_isNumberOfLastIncomingCallErr, "Error serializing 'isNumberOfLastIncomingCall' field")
	}

	// Simple Field (number)
	number := string(m.GetNumber())
	_numberErr := writeBuffer.WriteString("number", uint32((((m.GetCommandTypeContainer().NumBytes()) - ((2)))) * ((8))), "UTF-8", (number))
	if _numberErr != nil {
		return errors.Wrap(_numberErr, "Error serializing 'number' field")
	}

		if popErr := writeBuffer.PopContext("TelephonyDataRecallLastNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRecallLastNumber")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_TelephonyDataRecallLastNumber) isTelephonyDataRecallLastNumber() bool {
	return true
}

func (m *_TelephonyDataRecallLastNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



