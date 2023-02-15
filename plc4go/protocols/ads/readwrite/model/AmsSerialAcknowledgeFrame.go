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


// AmsSerialAcknowledgeFrame is the corresponding interface of AmsSerialAcknowledgeFrame
type AmsSerialAcknowledgeFrame interface {
	utils.LengthAware
	utils.Serializable
	// GetMagicCookie returns MagicCookie (property field)
	GetMagicCookie() uint16
	// GetTransmitterAddress returns TransmitterAddress (property field)
	GetTransmitterAddress() int8
	// GetReceiverAddress returns ReceiverAddress (property field)
	GetReceiverAddress() int8
	// GetFragmentNumber returns FragmentNumber (property field)
	GetFragmentNumber() int8
	// GetLength returns Length (property field)
	GetLength() int8
	// GetCrc returns Crc (property field)
	GetCrc() uint16
}

// AmsSerialAcknowledgeFrameExactly can be used when we want exactly this type and not a type which fulfills AmsSerialAcknowledgeFrame.
// This is useful for switch cases.
type AmsSerialAcknowledgeFrameExactly interface {
	AmsSerialAcknowledgeFrame
	isAmsSerialAcknowledgeFrame() bool
}

// _AmsSerialAcknowledgeFrame is the data-structure of this message
type _AmsSerialAcknowledgeFrame struct {
        MagicCookie uint16
        TransmitterAddress int8
        ReceiverAddress int8
        FragmentNumber int8
        Length int8
        Crc uint16
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsSerialAcknowledgeFrame) GetMagicCookie() uint16 {
	return m.MagicCookie
}

func (m *_AmsSerialAcknowledgeFrame) GetTransmitterAddress() int8 {
	return m.TransmitterAddress
}

func (m *_AmsSerialAcknowledgeFrame) GetReceiverAddress() int8 {
	return m.ReceiverAddress
}

func (m *_AmsSerialAcknowledgeFrame) GetFragmentNumber() int8 {
	return m.FragmentNumber
}

func (m *_AmsSerialAcknowledgeFrame) GetLength() int8 {
	return m.Length
}

func (m *_AmsSerialAcknowledgeFrame) GetCrc() uint16 {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAmsSerialAcknowledgeFrame factory function for _AmsSerialAcknowledgeFrame
func NewAmsSerialAcknowledgeFrame( magicCookie uint16 , transmitterAddress int8 , receiverAddress int8 , fragmentNumber int8 , length int8 , crc uint16 ) *_AmsSerialAcknowledgeFrame {
return &_AmsSerialAcknowledgeFrame{ MagicCookie: magicCookie , TransmitterAddress: transmitterAddress , ReceiverAddress: receiverAddress , FragmentNumber: fragmentNumber , Length: length , Crc: crc }
}

// Deprecated: use the interface for direct cast
func CastAmsSerialAcknowledgeFrame(structType interface{}) AmsSerialAcknowledgeFrame {
    if casted, ok := structType.(AmsSerialAcknowledgeFrame); ok {
		return casted
	}
	if casted, ok := structType.(*AmsSerialAcknowledgeFrame); ok {
		return *casted
	}
	return nil
}

func (m *_AmsSerialAcknowledgeFrame) GetTypeName() string {
	return "AmsSerialAcknowledgeFrame"
}

func (m *_AmsSerialAcknowledgeFrame) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16;

	// Simple field (transmitterAddress)
	lengthInBits += 8;

	// Simple field (receiverAddress)
	lengthInBits += 8;

	// Simple field (fragmentNumber)
	lengthInBits += 8;

	// Simple field (length)
	lengthInBits += 8;

	// Simple field (crc)
	lengthInBits += 16;

	return lengthInBits
}


func (m *_AmsSerialAcknowledgeFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsSerialAcknowledgeFrameParse(theBytes []byte) (AmsSerialAcknowledgeFrame, error) {
	return AmsSerialAcknowledgeFrameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AmsSerialAcknowledgeFrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsSerialAcknowledgeFrame, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsSerialAcknowledgeFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsSerialAcknowledgeFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (magicCookie)
_magicCookie, _magicCookieErr := readBuffer.ReadUint16("magicCookie", 16)
	if _magicCookieErr != nil {
		return nil, errors.Wrap(_magicCookieErr, "Error parsing 'magicCookie' field of AmsSerialAcknowledgeFrame")
	}
	magicCookie := _magicCookie

	// Simple Field (transmitterAddress)
_transmitterAddress, _transmitterAddressErr := readBuffer.ReadInt8("transmitterAddress", 8)
	if _transmitterAddressErr != nil {
		return nil, errors.Wrap(_transmitterAddressErr, "Error parsing 'transmitterAddress' field of AmsSerialAcknowledgeFrame")
	}
	transmitterAddress := _transmitterAddress

	// Simple Field (receiverAddress)
_receiverAddress, _receiverAddressErr := readBuffer.ReadInt8("receiverAddress", 8)
	if _receiverAddressErr != nil {
		return nil, errors.Wrap(_receiverAddressErr, "Error parsing 'receiverAddress' field of AmsSerialAcknowledgeFrame")
	}
	receiverAddress := _receiverAddress

	// Simple Field (fragmentNumber)
_fragmentNumber, _fragmentNumberErr := readBuffer.ReadInt8("fragmentNumber", 8)
	if _fragmentNumberErr != nil {
		return nil, errors.Wrap(_fragmentNumberErr, "Error parsing 'fragmentNumber' field of AmsSerialAcknowledgeFrame")
	}
	fragmentNumber := _fragmentNumber

	// Simple Field (length)
_length, _lengthErr := readBuffer.ReadInt8("length", 8)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AmsSerialAcknowledgeFrame")
	}
	length := _length

	// Simple Field (crc)
_crc, _crcErr := readBuffer.ReadUint16("crc", 16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field of AmsSerialAcknowledgeFrame")
	}
	crc := _crc

	if closeErr := readBuffer.CloseContext("AmsSerialAcknowledgeFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsSerialAcknowledgeFrame")
	}

	// Create the instance
	return &_AmsSerialAcknowledgeFrame{
			MagicCookie: magicCookie,
			TransmitterAddress: transmitterAddress,
			ReceiverAddress: receiverAddress,
			FragmentNumber: fragmentNumber,
			Length: length,
			Crc: crc,
		}, nil
}

func (m *_AmsSerialAcknowledgeFrame) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsSerialAcknowledgeFrame) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("AmsSerialAcknowledgeFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsSerialAcknowledgeFrame")
	}

	// Simple Field (magicCookie)
	magicCookie := uint16(m.GetMagicCookie())
	_magicCookieErr := writeBuffer.WriteUint16("magicCookie", 16, (magicCookie))
	if _magicCookieErr != nil {
		return errors.Wrap(_magicCookieErr, "Error serializing 'magicCookie' field")
	}

	// Simple Field (transmitterAddress)
	transmitterAddress := int8(m.GetTransmitterAddress())
	_transmitterAddressErr := writeBuffer.WriteInt8("transmitterAddress", 8, (transmitterAddress))
	if _transmitterAddressErr != nil {
		return errors.Wrap(_transmitterAddressErr, "Error serializing 'transmitterAddress' field")
	}

	// Simple Field (receiverAddress)
	receiverAddress := int8(m.GetReceiverAddress())
	_receiverAddressErr := writeBuffer.WriteInt8("receiverAddress", 8, (receiverAddress))
	if _receiverAddressErr != nil {
		return errors.Wrap(_receiverAddressErr, "Error serializing 'receiverAddress' field")
	}

	// Simple Field (fragmentNumber)
	fragmentNumber := int8(m.GetFragmentNumber())
	_fragmentNumberErr := writeBuffer.WriteInt8("fragmentNumber", 8, (fragmentNumber))
	if _fragmentNumberErr != nil {
		return errors.Wrap(_fragmentNumberErr, "Error serializing 'fragmentNumber' field")
	}

	// Simple Field (length)
	length := int8(m.GetLength())
	_lengthErr := writeBuffer.WriteInt8("length", 8, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (crc)
	crc := uint16(m.GetCrc())
	_crcErr := writeBuffer.WriteUint16("crc", 16, (crc))
	if _crcErr != nil {
		return errors.Wrap(_crcErr, "Error serializing 'crc' field")
	}

	if popErr := writeBuffer.PopContext("AmsSerialAcknowledgeFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsSerialAcknowledgeFrame")
	}
	return nil
}


func (m *_AmsSerialAcknowledgeFrame) isAmsSerialAcknowledgeFrame() bool {
	return true
}

func (m *_AmsSerialAcknowledgeFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



