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


// BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue
        RealValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMaxNormalValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal)  GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal( realValue BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal{
		RealValue: realValue,
    	_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal(structType interface{}) BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal {
    if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealParse(theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal, error) {
	return BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
_realValue, _realValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
	}
	realValue := _realValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal{
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &_BACnetFaultParameterFaultOutOfRangeMaxNormalValue{
			TagNumber: tagNumber,
		},
		RealValue: realValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
		}

	// Simple Field (realValue)
	if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for realValue")
	}
	_realValueErr := writeBuffer.WriteSerializable(ctx, m.GetRealValue())
	if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for realValue")
	}
	if _realValueErr != nil {
		return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) isBACnetFaultParameterFaultOutOfRangeMaxNormalValueReal() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



