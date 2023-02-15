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


// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
        TimeValue BACnetApplicationTagTime
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime)  GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime( timeValue BACnetApplicationTagTime , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime{
		TimeValue: timeValue,
    	_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime {
    if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeParse(theBytes []byte, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeValue)
	if pullErr := readBuffer.PullContext("timeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeValue")
	}
_timeValue, _timeValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _timeValueErr != nil {
		return nil, errors.Wrap(_timeValueErr, "Error parsing 'timeValue' field of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
	}
	timeValue := _timeValue.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("timeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime{
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
			TagNumber: tagNumber,
		},
		TimeValue: timeValue,
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
		}

	// Simple Field (timeValue)
	if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeValue")
	}
	_timeValueErr := writeBuffer.WriteSerializable(ctx, m.GetTimeValue())
	if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeValue")
	}
	if _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) isBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



