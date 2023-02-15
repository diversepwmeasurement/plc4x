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


// BACnetConstructedDataLastCredentialRemoved is the corresponding interface of BACnetConstructedDataLastCredentialRemoved
type BACnetConstructedDataLastCredentialRemoved interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastCredentialRemoved returns LastCredentialRemoved (property field)
	GetLastCredentialRemoved() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataLastCredentialRemovedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastCredentialRemoved.
// This is useful for switch cases.
type BACnetConstructedDataLastCredentialRemovedExactly interface {
	BACnetConstructedDataLastCredentialRemoved
	isBACnetConstructedDataLastCredentialRemoved() bool
}

// _BACnetConstructedDataLastCredentialRemoved is the data-structure of this message
type _BACnetConstructedDataLastCredentialRemoved struct {
	*_BACnetConstructedData
        LastCredentialRemoved BACnetDeviceObjectReference
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLastCredentialRemoved)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastCredentialRemoved)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetLastCredentialRemoved() BACnetDeviceObjectReference {
	return m.LastCredentialRemoved
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetLastCredentialRemoved())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLastCredentialRemoved factory function for _BACnetConstructedDataLastCredentialRemoved
func NewBACnetConstructedDataLastCredentialRemoved( lastCredentialRemoved BACnetDeviceObjectReference , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLastCredentialRemoved {
	_result := &_BACnetConstructedDataLastCredentialRemoved{
		LastCredentialRemoved: lastCredentialRemoved,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialRemoved(structType interface{}) BACnetConstructedDataLastCredentialRemoved {
    if casted, ok := structType.(BACnetConstructedDataLastCredentialRemoved); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemoved); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemoved"
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastCredentialRemoved)
	lengthInBits += m.LastCredentialRemoved.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLastCredentialRemoved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastCredentialRemovedParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialRemoved, error) {
	return BACnetConstructedDataLastCredentialRemovedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastCredentialRemovedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCredentialRemoved, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemoved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialRemoved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCredentialRemoved)
	if pullErr := readBuffer.PullContext("lastCredentialRemoved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastCredentialRemoved")
	}
_lastCredentialRemoved, _lastCredentialRemovedErr := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
	if _lastCredentialRemovedErr != nil {
		return nil, errors.Wrap(_lastCredentialRemovedErr, "Error parsing 'lastCredentialRemoved' field of BACnetConstructedDataLastCredentialRemoved")
	}
	lastCredentialRemoved := _lastCredentialRemoved.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("lastCredentialRemoved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastCredentialRemoved")
	}

	// Virtual field
	_actualValue := lastCredentialRemoved
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemoved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialRemoved")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastCredentialRemoved{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastCredentialRemoved: lastCredentialRemoved,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemoved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialRemoved")
		}

	// Simple Field (lastCredentialRemoved)
	if pushErr := writeBuffer.PushContext("lastCredentialRemoved"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lastCredentialRemoved")
	}
	_lastCredentialRemovedErr := writeBuffer.WriteSerializable(ctx, m.GetLastCredentialRemoved())
	if popErr := writeBuffer.PopContext("lastCredentialRemoved"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lastCredentialRemoved")
	}
	if _lastCredentialRemovedErr != nil {
		return errors.Wrap(_lastCredentialRemovedErr, "Error serializing 'lastCredentialRemoved' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemoved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialRemoved")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLastCredentialRemoved) isBACnetConstructedDataLastCredentialRemoved() bool {
	return true
}

func (m *_BACnetConstructedDataLastCredentialRemoved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



