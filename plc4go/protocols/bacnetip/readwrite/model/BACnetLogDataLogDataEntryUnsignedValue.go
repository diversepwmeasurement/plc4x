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


// BACnetLogDataLogDataEntryUnsignedValue is the corresponding interface of BACnetLogDataLogDataEntryUnsignedValue
type BACnetLogDataLogDataEntryUnsignedValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetContextTagUnsignedInteger
}

// BACnetLogDataLogDataEntryUnsignedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryUnsignedValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryUnsignedValueExactly interface {
	BACnetLogDataLogDataEntryUnsignedValue
	isBACnetLogDataLogDataEntryUnsignedValue() bool
}

// _BACnetLogDataLogDataEntryUnsignedValue is the data-structure of this message
type _BACnetLogDataLogDataEntryUnsignedValue struct {
	*_BACnetLogDataLogDataEntry
        UnsignedValue BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryUnsignedValue) InitializeParent(parent BACnetLogDataLogDataEntry , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue)  GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetUnsignedValue() BACnetContextTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLogDataLogDataEntryUnsignedValue factory function for _BACnetLogDataLogDataEntryUnsignedValue
func NewBACnetLogDataLogDataEntryUnsignedValue( unsignedValue BACnetContextTagUnsignedInteger , peekedTagHeader BACnetTagHeader ) *_BACnetLogDataLogDataEntryUnsignedValue {
	_result := &_BACnetLogDataLogDataEntryUnsignedValue{
		UnsignedValue: unsignedValue,
    	_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryUnsignedValue(structType interface{}) BACnetLogDataLogDataEntryUnsignedValue {
    if casted, ok := structType.(BACnetLogDataLogDataEntryUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryUnsignedValue"
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogDataEntryUnsignedValueParse(theBytes []byte) (BACnetLogDataLogDataEntryUnsignedValue, error) {
	return BACnetLogDataLogDataEntryUnsignedValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogDataLogDataEntryUnsignedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryUnsignedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
_unsignedValue, _unsignedValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(3) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetLogDataLogDataEntryUnsignedValue")
	}
	unsignedValue := _unsignedValue.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryUnsignedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryUnsignedValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{
		},
		UnsignedValue: unsignedValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryUnsignedValue")
		}

	// Simple Field (unsignedValue)
	if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for unsignedValue")
	}
	_unsignedValueErr := writeBuffer.WriteSerializable(ctx, m.GetUnsignedValue())
	if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for unsignedValue")
	}
	if _unsignedValueErr != nil {
		return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryUnsignedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetLogDataLogDataEntryUnsignedValue) isBACnetLogDataLogDataEntryUnsignedValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



