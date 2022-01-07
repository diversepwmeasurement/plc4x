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
type BACnetConstructedDataElement struct {
	*BACnetTag
	Identifier       BACnetPropertyIdentifier
	ProprietaryValue uint32
	Value            *BACnetTag
	IsProprietary    bool
}

// The corresponding interface
type IBACnetConstructedDataElement interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConstructedDataElement) TagClass() TagClass {
	return TagClass_CONTEXT_SPECIFIC_TAGS
}

func (m *BACnetConstructedDataElement) InitializeParent(parent *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isBoolean bool, isConstructed bool, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.TagNumber = tagNumber
	m.LengthValueType = lengthValueType
	m.ExtTagNumber = extTagNumber
	m.ExtLength = extLength
	m.ExtExtLength = extExtLength
	m.ExtExtExtLength = extExtExtLength
}

func NewBACnetConstructedDataElement(identifier BACnetPropertyIdentifier, proprietaryValue uint32, value *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTag {
	child := &BACnetConstructedDataElement{
		Identifier:       identifier,
		ProprietaryValue: proprietaryValue,
		Value:            value,
		BACnetTag:        NewBACnetTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Child = child
	return child.BACnetTag
}

func CastBACnetConstructedDataElement(structType interface{}) *BACnetConstructedDataElement {
	castFunc := func(typ interface{}) *BACnetConstructedDataElement {
		if casted, ok := typ.(BACnetConstructedDataElement); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConstructedDataElement); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetConstructedDataElement(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetConstructedDataElement(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConstructedDataElement) GetTypeName() string {
	return "BACnetConstructedDataElement"
}

func (m *BACnetConstructedDataElement) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConstructedDataElement) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Manual Field (identifier)
	lengthInBits += uint16(m.ActualLength * 8)

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(m.ActualLength * 8)

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataElement) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConstructedDataElementParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElement"); pullErr != nil {
		return nil, pullErr
	}

	// Manual Field (identifier)
	identifier, _identifierErr := ReadPropertyIdentifier(readBuffer, actualLength)
	if _identifierErr != nil {
		return nil, errors.Wrap(_identifierErr, "Error parsing 'identifier' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryPropertyIdentifier(readBuffer, identifier, actualLength)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((identifier) == (BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	_value, _valueErr := BACnetTagParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := CastBACnetTag(_value)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElement"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataElement{
		Identifier:       identifier,
		ProprietaryValue: proprietaryValue,
		Value:            CastBACnetTag(value),
		IsProprietary:    isProprietary,
		BACnetTag:        &BACnetTag{},
	}
	_child.BACnetTag.Child = _child
	return _child.BACnetTag, nil
}

func (m *BACnetConstructedDataElement) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataElement"); pushErr != nil {
			return pushErr
		}

		// Manual Field (identifier)
		_identifierErr := WritePropertyIdentifier(writeBuffer, m.Identifier)
		if _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		// Manual Field (proprietaryValue)
		_proprietaryValueErr := WriteProprietaryPropertyIdentifier(writeBuffer, m.Identifier, m.ProprietaryValue)
		if _proprietaryValueErr != nil {
			return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
		}
		// Virtual field
		if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.IsProprietary); _isProprietaryErr != nil {
			return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return pushErr
		}
		_valueErr := m.Value.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return popErr
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataElement"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataElement) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
