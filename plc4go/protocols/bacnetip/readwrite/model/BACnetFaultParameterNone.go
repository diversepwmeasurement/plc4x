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


// BACnetFaultParameterNone is the corresponding interface of BACnetFaultParameterNone
type BACnetFaultParameterNone interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
}

// BACnetFaultParameterNoneExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterNone.
// This is useful for switch cases.
type BACnetFaultParameterNoneExactly interface {
	BACnetFaultParameterNone
	isBACnetFaultParameterNone() bool
}

// _BACnetFaultParameterNone is the data-structure of this message
type _BACnetFaultParameterNone struct {
	*_BACnetFaultParameter
        None BACnetContextTagNull
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterNone) InitializeParent(parent BACnetFaultParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterNone)  GetParent() BACnetFaultParameter {
	return m._BACnetFaultParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterNone) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterNone factory function for _BACnetFaultParameterNone
func NewBACnetFaultParameterNone( none BACnetContextTagNull , peekedTagHeader BACnetTagHeader ) *_BACnetFaultParameterNone {
	_result := &_BACnetFaultParameterNone{
		None: none,
    	_BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterNone(structType interface{}) BACnetFaultParameterNone {
    if casted, ok := structType.(BACnetFaultParameterNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterNone) GetTypeName() string {
	return "BACnetFaultParameterNone"
}

func (m *_BACnetFaultParameterNone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetFaultParameterNone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterNoneParse(theBytes []byte) (BACnetFaultParameterNone, error) {
	return BACnetFaultParameterNoneParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterNoneParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterNone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (none)
	if pullErr := readBuffer.PullContext("none"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for none")
	}
_none, _noneErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_NULL ) )
	if _noneErr != nil {
		return nil, errors.Wrap(_noneErr, "Error parsing 'none' field of BACnetFaultParameterNone")
	}
	none := _none.(BACnetContextTagNull)
	if closeErr := readBuffer.CloseContext("none"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for none")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterNone")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterNone{
		_BACnetFaultParameter: &_BACnetFaultParameter{
		},
		None: none,
	}
	_child._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterNone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterNone")
		}

	// Simple Field (none)
	if pushErr := writeBuffer.PushContext("none"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for none")
	}
	_noneErr := writeBuffer.WriteSerializable(ctx, m.GetNone())
	if popErr := writeBuffer.PopContext("none"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for none")
	}
	if _noneErr != nil {
		return errors.Wrap(_noneErr, "Error serializing 'none' field")
	}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterNone")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterNone) isBACnetFaultParameterNone() bool {
	return true
}

func (m *_BACnetFaultParameterNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



