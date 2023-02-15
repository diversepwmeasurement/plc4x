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


// NLMRequestKeyUpdate is the corresponding interface of NLMRequestKeyUpdate
type NLMRequestKeyUpdate interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetSet1KeyRevision returns Set1KeyRevision (property field)
	GetSet1KeyRevision() byte
	// GetSet1ActivationTime returns Set1ActivationTime (property field)
	GetSet1ActivationTime() uint32
	// GetSet1ExpirationTime returns Set1ExpirationTime (property field)
	GetSet1ExpirationTime() uint32
	// GetSet2KeyRevision returns Set2KeyRevision (property field)
	GetSet2KeyRevision() byte
	// GetSet2ActivationTime returns Set2ActivationTime (property field)
	GetSet2ActivationTime() uint32
	// GetSet2ExpirationTime returns Set2ExpirationTime (property field)
	GetSet2ExpirationTime() uint32
	// GetDistributionKeyRevision returns DistributionKeyRevision (property field)
	GetDistributionKeyRevision() byte
}

// NLMRequestKeyUpdateExactly can be used when we want exactly this type and not a type which fulfills NLMRequestKeyUpdate.
// This is useful for switch cases.
type NLMRequestKeyUpdateExactly interface {
	NLMRequestKeyUpdate
	isNLMRequestKeyUpdate() bool
}

// _NLMRequestKeyUpdate is the data-structure of this message
type _NLMRequestKeyUpdate struct {
	*_NLM
        Set1KeyRevision byte
        Set1ActivationTime uint32
        Set1ExpirationTime uint32
        Set2KeyRevision byte
        Set2ActivationTime uint32
        Set2ExpirationTime uint32
        DistributionKeyRevision byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRequestKeyUpdate)  GetMessageType() uint8 {
return 0x0D}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRequestKeyUpdate) InitializeParent(parent NLM ) {}

func (m *_NLMRequestKeyUpdate)  GetParent() NLM {
	return m._NLM
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRequestKeyUpdate) GetSet1KeyRevision() byte {
	return m.Set1KeyRevision
}

func (m *_NLMRequestKeyUpdate) GetSet1ActivationTime() uint32 {
	return m.Set1ActivationTime
}

func (m *_NLMRequestKeyUpdate) GetSet1ExpirationTime() uint32 {
	return m.Set1ExpirationTime
}

func (m *_NLMRequestKeyUpdate) GetSet2KeyRevision() byte {
	return m.Set2KeyRevision
}

func (m *_NLMRequestKeyUpdate) GetSet2ActivationTime() uint32 {
	return m.Set2ActivationTime
}

func (m *_NLMRequestKeyUpdate) GetSet2ExpirationTime() uint32 {
	return m.Set2ExpirationTime
}

func (m *_NLMRequestKeyUpdate) GetDistributionKeyRevision() byte {
	return m.DistributionKeyRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNLMRequestKeyUpdate factory function for _NLMRequestKeyUpdate
func NewNLMRequestKeyUpdate( set1KeyRevision byte , set1ActivationTime uint32 , set1ExpirationTime uint32 , set2KeyRevision byte , set2ActivationTime uint32 , set2ExpirationTime uint32 , distributionKeyRevision byte , apduLength uint16 ) *_NLMRequestKeyUpdate {
	_result := &_NLMRequestKeyUpdate{
		Set1KeyRevision: set1KeyRevision,
		Set1ActivationTime: set1ActivationTime,
		Set1ExpirationTime: set1ExpirationTime,
		Set2KeyRevision: set2KeyRevision,
		Set2ActivationTime: set2ActivationTime,
		Set2ExpirationTime: set2ExpirationTime,
		DistributionKeyRevision: distributionKeyRevision,
    	_NLM: NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRequestKeyUpdate(structType interface{}) NLMRequestKeyUpdate {
    if casted, ok := structType.(NLMRequestKeyUpdate); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRequestKeyUpdate); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRequestKeyUpdate) GetTypeName() string {
	return "NLMRequestKeyUpdate"
}

func (m *_NLMRequestKeyUpdate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (set1KeyRevision)
	lengthInBits += 8;

	// Simple field (set1ActivationTime)
	lengthInBits += 32;

	// Simple field (set1ExpirationTime)
	lengthInBits += 32;

	// Simple field (set2KeyRevision)
	lengthInBits += 8;

	// Simple field (set2ActivationTime)
	lengthInBits += 32;

	// Simple field (set2ExpirationTime)
	lengthInBits += 32;

	// Simple field (distributionKeyRevision)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_NLMRequestKeyUpdate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMRequestKeyUpdateParse(theBytes []byte, apduLength uint16) (NLMRequestKeyUpdate, error) {
	return NLMRequestKeyUpdateParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMRequestKeyUpdateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMRequestKeyUpdate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMRequestKeyUpdate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRequestKeyUpdate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (set1KeyRevision)
_set1KeyRevision, _set1KeyRevisionErr := readBuffer.ReadByte("set1KeyRevision")
	if _set1KeyRevisionErr != nil {
		return nil, errors.Wrap(_set1KeyRevisionErr, "Error parsing 'set1KeyRevision' field of NLMRequestKeyUpdate")
	}
	set1KeyRevision := _set1KeyRevision

	// Simple Field (set1ActivationTime)
_set1ActivationTime, _set1ActivationTimeErr := readBuffer.ReadUint32("set1ActivationTime", 32)
	if _set1ActivationTimeErr != nil {
		return nil, errors.Wrap(_set1ActivationTimeErr, "Error parsing 'set1ActivationTime' field of NLMRequestKeyUpdate")
	}
	set1ActivationTime := _set1ActivationTime

	// Simple Field (set1ExpirationTime)
_set1ExpirationTime, _set1ExpirationTimeErr := readBuffer.ReadUint32("set1ExpirationTime", 32)
	if _set1ExpirationTimeErr != nil {
		return nil, errors.Wrap(_set1ExpirationTimeErr, "Error parsing 'set1ExpirationTime' field of NLMRequestKeyUpdate")
	}
	set1ExpirationTime := _set1ExpirationTime

	// Simple Field (set2KeyRevision)
_set2KeyRevision, _set2KeyRevisionErr := readBuffer.ReadByte("set2KeyRevision")
	if _set2KeyRevisionErr != nil {
		return nil, errors.Wrap(_set2KeyRevisionErr, "Error parsing 'set2KeyRevision' field of NLMRequestKeyUpdate")
	}
	set2KeyRevision := _set2KeyRevision

	// Simple Field (set2ActivationTime)
_set2ActivationTime, _set2ActivationTimeErr := readBuffer.ReadUint32("set2ActivationTime", 32)
	if _set2ActivationTimeErr != nil {
		return nil, errors.Wrap(_set2ActivationTimeErr, "Error parsing 'set2ActivationTime' field of NLMRequestKeyUpdate")
	}
	set2ActivationTime := _set2ActivationTime

	// Simple Field (set2ExpirationTime)
_set2ExpirationTime, _set2ExpirationTimeErr := readBuffer.ReadUint32("set2ExpirationTime", 32)
	if _set2ExpirationTimeErr != nil {
		return nil, errors.Wrap(_set2ExpirationTimeErr, "Error parsing 'set2ExpirationTime' field of NLMRequestKeyUpdate")
	}
	set2ExpirationTime := _set2ExpirationTime

	// Simple Field (distributionKeyRevision)
_distributionKeyRevision, _distributionKeyRevisionErr := readBuffer.ReadByte("distributionKeyRevision")
	if _distributionKeyRevisionErr != nil {
		return nil, errors.Wrap(_distributionKeyRevisionErr, "Error parsing 'distributionKeyRevision' field of NLMRequestKeyUpdate")
	}
	distributionKeyRevision := _distributionKeyRevision

	if closeErr := readBuffer.CloseContext("NLMRequestKeyUpdate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRequestKeyUpdate")
	}

	// Create a partially initialized instance
	_child := &_NLMRequestKeyUpdate{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		Set1KeyRevision: set1KeyRevision,
		Set1ActivationTime: set1ActivationTime,
		Set1ExpirationTime: set1ExpirationTime,
		Set2KeyRevision: set2KeyRevision,
		Set2ActivationTime: set2ActivationTime,
		Set2ExpirationTime: set2ExpirationTime,
		DistributionKeyRevision: distributionKeyRevision,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMRequestKeyUpdate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRequestKeyUpdate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRequestKeyUpdate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRequestKeyUpdate")
		}

	// Simple Field (set1KeyRevision)
	set1KeyRevision := byte(m.GetSet1KeyRevision())
	_set1KeyRevisionErr := writeBuffer.WriteByte("set1KeyRevision", (set1KeyRevision))
	if _set1KeyRevisionErr != nil {
		return errors.Wrap(_set1KeyRevisionErr, "Error serializing 'set1KeyRevision' field")
	}

	// Simple Field (set1ActivationTime)
	set1ActivationTime := uint32(m.GetSet1ActivationTime())
	_set1ActivationTimeErr := writeBuffer.WriteUint32("set1ActivationTime", 32, (set1ActivationTime))
	if _set1ActivationTimeErr != nil {
		return errors.Wrap(_set1ActivationTimeErr, "Error serializing 'set1ActivationTime' field")
	}

	// Simple Field (set1ExpirationTime)
	set1ExpirationTime := uint32(m.GetSet1ExpirationTime())
	_set1ExpirationTimeErr := writeBuffer.WriteUint32("set1ExpirationTime", 32, (set1ExpirationTime))
	if _set1ExpirationTimeErr != nil {
		return errors.Wrap(_set1ExpirationTimeErr, "Error serializing 'set1ExpirationTime' field")
	}

	// Simple Field (set2KeyRevision)
	set2KeyRevision := byte(m.GetSet2KeyRevision())
	_set2KeyRevisionErr := writeBuffer.WriteByte("set2KeyRevision", (set2KeyRevision))
	if _set2KeyRevisionErr != nil {
		return errors.Wrap(_set2KeyRevisionErr, "Error serializing 'set2KeyRevision' field")
	}

	// Simple Field (set2ActivationTime)
	set2ActivationTime := uint32(m.GetSet2ActivationTime())
	_set2ActivationTimeErr := writeBuffer.WriteUint32("set2ActivationTime", 32, (set2ActivationTime))
	if _set2ActivationTimeErr != nil {
		return errors.Wrap(_set2ActivationTimeErr, "Error serializing 'set2ActivationTime' field")
	}

	// Simple Field (set2ExpirationTime)
	set2ExpirationTime := uint32(m.GetSet2ExpirationTime())
	_set2ExpirationTimeErr := writeBuffer.WriteUint32("set2ExpirationTime", 32, (set2ExpirationTime))
	if _set2ExpirationTimeErr != nil {
		return errors.Wrap(_set2ExpirationTimeErr, "Error serializing 'set2ExpirationTime' field")
	}

	// Simple Field (distributionKeyRevision)
	distributionKeyRevision := byte(m.GetDistributionKeyRevision())
	_distributionKeyRevisionErr := writeBuffer.WriteByte("distributionKeyRevision", (distributionKeyRevision))
	if _distributionKeyRevisionErr != nil {
		return errors.Wrap(_distributionKeyRevisionErr, "Error serializing 'distributionKeyRevision' field")
	}

		if popErr := writeBuffer.PopContext("NLMRequestKeyUpdate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRequestKeyUpdate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_NLMRequestKeyUpdate) isNLMRequestKeyUpdate() bool {
	return true
}

func (m *_NLMRequestKeyUpdate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



