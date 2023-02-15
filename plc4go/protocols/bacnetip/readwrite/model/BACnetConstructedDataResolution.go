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


// BACnetConstructedDataResolution is the corresponding interface of BACnetConstructedDataResolution
type BACnetConstructedDataResolution interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataResolutionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataResolution.
// This is useful for switch cases.
type BACnetConstructedDataResolutionExactly interface {
	BACnetConstructedDataResolution
	isBACnetConstructedDataResolution() bool
}

// _BACnetConstructedDataResolution is the data-structure of this message
type _BACnetConstructedDataResolution struct {
	*_BACnetConstructedData
        Resolution BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataResolution)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataResolution)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_RESOLUTION}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataResolution) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataResolution)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataResolution) GetResolution() BACnetApplicationTagReal {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataResolution) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataResolution factory function for _BACnetConstructedDataResolution
func NewBACnetConstructedDataResolution( resolution BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataResolution {
	_result := &_BACnetConstructedDataResolution{
		Resolution: resolution,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataResolution(structType interface{}) BACnetConstructedDataResolution {
    if casted, ok := structType.(BACnetConstructedDataResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataResolution) GetTypeName() string {
	return "BACnetConstructedDataResolution"
}

func (m *_BACnetConstructedDataResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataResolutionParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataResolution, error) {
	return BACnetConstructedDataResolutionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataResolutionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataResolution, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (resolution)
	if pullErr := readBuffer.PullContext("resolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for resolution")
	}
_resolution, _resolutionErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _resolutionErr != nil {
		return nil, errors.Wrap(_resolutionErr, "Error parsing 'resolution' field of BACnetConstructedDataResolution")
	}
	resolution := _resolution.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("resolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for resolution")
	}

	// Virtual field
	_actualValue := resolution
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataResolution")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataResolution{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Resolution: resolution,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataResolution")
		}

	// Simple Field (resolution)
	if pushErr := writeBuffer.PushContext("resolution"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for resolution")
	}
	_resolutionErr := writeBuffer.WriteSerializable(ctx, m.GetResolution())
	if popErr := writeBuffer.PopContext("resolution"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for resolution")
	}
	if _resolutionErr != nil {
		return errors.Wrap(_resolutionErr, "Error serializing 'resolution' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataResolution")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataResolution) isBACnetConstructedDataResolution() bool {
	return true
}

func (m *_BACnetConstructedDataResolution) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



