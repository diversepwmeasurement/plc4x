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


// BACnetConstructedDataCarDriveStatus is the corresponding interface of BACnetConstructedDataCarDriveStatus
type BACnetConstructedDataCarDriveStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarDriveStatus returns CarDriveStatus (property field)
	GetCarDriveStatus() BACnetLiftCarDriveStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarDriveStatusTagged
}

// BACnetConstructedDataCarDriveStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarDriveStatus.
// This is useful for switch cases.
type BACnetConstructedDataCarDriveStatusExactly interface {
	BACnetConstructedDataCarDriveStatus
	isBACnetConstructedDataCarDriveStatus() bool
}

// _BACnetConstructedDataCarDriveStatus is the data-structure of this message
type _BACnetConstructedDataCarDriveStatus struct {
	*_BACnetConstructedData
        CarDriveStatus BACnetLiftCarDriveStatusTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarDriveStatus)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataCarDriveStatus)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_CAR_DRIVE_STATUS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarDriveStatus) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarDriveStatus)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarDriveStatus) GetCarDriveStatus() BACnetLiftCarDriveStatusTagged {
	return m.CarDriveStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarDriveStatus) GetActualValue() BACnetLiftCarDriveStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftCarDriveStatusTagged(m.GetCarDriveStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataCarDriveStatus factory function for _BACnetConstructedDataCarDriveStatus
func NewBACnetConstructedDataCarDriveStatus( carDriveStatus BACnetLiftCarDriveStatusTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataCarDriveStatus {
	_result := &_BACnetConstructedDataCarDriveStatus{
		CarDriveStatus: carDriveStatus,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarDriveStatus(structType interface{}) BACnetConstructedDataCarDriveStatus {
    if casted, ok := structType.(BACnetConstructedDataCarDriveStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarDriveStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarDriveStatus) GetTypeName() string {
	return "BACnetConstructedDataCarDriveStatus"
}

func (m *_BACnetConstructedDataCarDriveStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (carDriveStatus)
	lengthInBits += m.CarDriveStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataCarDriveStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCarDriveStatusParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarDriveStatus, error) {
	return BACnetConstructedDataCarDriveStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCarDriveStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarDriveStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarDriveStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarDriveStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carDriveStatus)
	if pullErr := readBuffer.PullContext("carDriveStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carDriveStatus")
	}
_carDriveStatus, _carDriveStatusErr := BACnetLiftCarDriveStatusTaggedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _carDriveStatusErr != nil {
		return nil, errors.Wrap(_carDriveStatusErr, "Error parsing 'carDriveStatus' field of BACnetConstructedDataCarDriveStatus")
	}
	carDriveStatus := _carDriveStatus.(BACnetLiftCarDriveStatusTagged)
	if closeErr := readBuffer.CloseContext("carDriveStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carDriveStatus")
	}

	// Virtual field
	_actualValue := carDriveStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarDriveStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarDriveStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarDriveStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CarDriveStatus: carDriveStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarDriveStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarDriveStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarDriveStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarDriveStatus")
		}

	// Simple Field (carDriveStatus)
	if pushErr := writeBuffer.PushContext("carDriveStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for carDriveStatus")
	}
	_carDriveStatusErr := writeBuffer.WriteSerializable(ctx, m.GetCarDriveStatus())
	if popErr := writeBuffer.PopContext("carDriveStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for carDriveStatus")
	}
	if _carDriveStatusErr != nil {
		return errors.Wrap(_carDriveStatusErr, "Error serializing 'carDriveStatus' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarDriveStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarDriveStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataCarDriveStatus) isBACnetConstructedDataCarDriveStatus() bool {
	return true
}

func (m *_BACnetConstructedDataCarDriveStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



