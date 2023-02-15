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


// MeteringDataMeasureOtherWater is the corresponding interface of MeteringDataMeasureOtherWater
type MeteringDataMeasureOtherWater interface {
	utils.LengthAware
	utils.Serializable
	MeteringData
}

// MeteringDataMeasureOtherWaterExactly can be used when we want exactly this type and not a type which fulfills MeteringDataMeasureOtherWater.
// This is useful for switch cases.
type MeteringDataMeasureOtherWaterExactly interface {
	MeteringDataMeasureOtherWater
	isMeteringDataMeasureOtherWater() bool
}

// _MeteringDataMeasureOtherWater is the data-structure of this message
type _MeteringDataMeasureOtherWater struct {
	*_MeteringData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataMeasureOtherWater) InitializeParent(parent MeteringData , commandTypeContainer MeteringCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataMeasureOtherWater)  GetParent() MeteringData {
	return m._MeteringData
}


// NewMeteringDataMeasureOtherWater factory function for _MeteringDataMeasureOtherWater
func NewMeteringDataMeasureOtherWater( commandTypeContainer MeteringCommandTypeContainer , argument byte ) *_MeteringDataMeasureOtherWater {
	_result := &_MeteringDataMeasureOtherWater{
    	_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataMeasureOtherWater(structType interface{}) MeteringDataMeasureOtherWater {
    if casted, ok := structType.(MeteringDataMeasureOtherWater); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataMeasureOtherWater); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataMeasureOtherWater) GetTypeName() string {
	return "MeteringDataMeasureOtherWater"
}

func (m *_MeteringDataMeasureOtherWater) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}


func (m *_MeteringDataMeasureOtherWater) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MeteringDataMeasureOtherWaterParse(theBytes []byte) (MeteringDataMeasureOtherWater, error) {
	return MeteringDataMeasureOtherWaterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataMeasureOtherWaterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MeteringDataMeasureOtherWater, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataMeasureOtherWater"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataMeasureOtherWater")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MeteringDataMeasureOtherWater"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataMeasureOtherWater")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataMeasureOtherWater{
		_MeteringData: &_MeteringData{
		},
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataMeasureOtherWater) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataMeasureOtherWater) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataMeasureOtherWater"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataMeasureOtherWater")
		}

		if popErr := writeBuffer.PopContext("MeteringDataMeasureOtherWater"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataMeasureOtherWater")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_MeteringDataMeasureOtherWater) isMeteringDataMeasureOtherWater() bool {
	return true
}

func (m *_MeteringDataMeasureOtherWater) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



