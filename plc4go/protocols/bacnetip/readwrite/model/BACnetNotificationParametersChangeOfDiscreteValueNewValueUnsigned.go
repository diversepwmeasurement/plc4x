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


// BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
type BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
        UnsignedValue BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned)  GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned( unsignedValue BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned{
		UnsignedValue: unsignedValue,
    	_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned {
    if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParse(theBytes []byte, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
_unsignedValue, _unsignedValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned{
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
			TagNumber: tagNumber,
		},
		UnsignedValue: unsignedValue,
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned")
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

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) isBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



