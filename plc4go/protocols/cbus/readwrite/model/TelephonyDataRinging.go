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


// TelephonyDataRinging is the corresponding interface of TelephonyDataRinging
type TelephonyDataRinging interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetNumber returns Number (property field)
	GetNumber() string
}

// TelephonyDataRingingExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataRinging.
// This is useful for switch cases.
type TelephonyDataRingingExactly interface {
	TelephonyDataRinging
	isTelephonyDataRinging() bool
}

// _TelephonyDataRinging is the data-structure of this message
type _TelephonyDataRinging struct {
	*_TelephonyData
        Number string
	// Reserved Fields
	reservedField0 *byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataRinging) InitializeParent(parent TelephonyData , commandTypeContainer TelephonyCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataRinging)  GetParent() TelephonyData {
	return m._TelephonyData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataRinging) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewTelephonyDataRinging factory function for _TelephonyDataRinging
func NewTelephonyDataRinging( number string , commandTypeContainer TelephonyCommandTypeContainer , argument byte ) *_TelephonyDataRinging {
	_result := &_TelephonyDataRinging{
		Number: number,
    	_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataRinging(structType interface{}) TelephonyDataRinging {
    if casted, ok := structType.(TelephonyDataRinging); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataRinging); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataRinging) GetTypeName() string {
	return "TelephonyDataRinging"
}

func (m *_TelephonyDataRinging) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(2)))) * int32(int32(8)))

	return lengthInBits
}


func (m *_TelephonyDataRinging) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataRingingParse(theBytes []byte, commandTypeContainer TelephonyCommandTypeContainer) (TelephonyDataRinging, error) {
	return TelephonyDataRingingParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), commandTypeContainer)
}

func TelephonyDataRingingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer TelephonyCommandTypeContainer) (TelephonyDataRinging, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataRinging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataRinging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TelephonyDataRinging")
		}
		if reserved != byte(0x01) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": byte(0x01),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (number)
_number, _numberErr := readBuffer.ReadString("number", uint32((((commandTypeContainer.NumBytes()) - ((2)))) * ((8))), "UTF-8")
	if _numberErr != nil {
		return nil, errors.Wrap(_numberErr, "Error parsing 'number' field of TelephonyDataRinging")
	}
	number := _number

	if closeErr := readBuffer.CloseContext("TelephonyDataRinging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataRinging")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataRinging{
		_TelephonyData: &_TelephonyData{
		},
		Number: number,
		reservedField0: reservedField0,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataRinging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataRinging) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataRinging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataRinging")
		}

	// Reserved Field (reserved)
	{
		var reserved byte = byte(0x01)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": byte(0x01),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteByte("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (number)
	number := string(m.GetNumber())
	_numberErr := writeBuffer.WriteString("number", uint32((((m.GetCommandTypeContainer().NumBytes()) - ((2)))) * ((8))), "UTF-8", (number))
	if _numberErr != nil {
		return errors.Wrap(_numberErr, "Error serializing 'number' field")
	}

		if popErr := writeBuffer.PopContext("TelephonyDataRinging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataRinging")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_TelephonyDataRinging) isTelephonyDataRinging() bool {
	return true
}

func (m *_TelephonyDataRinging) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



