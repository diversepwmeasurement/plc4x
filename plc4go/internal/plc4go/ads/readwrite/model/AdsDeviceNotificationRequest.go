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
type AdsDeviceNotificationRequest struct {
	Length uint32
	Stamps uint32
	AdsStampHeaders []*AdsStampHeader
	Parent *AdsData
}

// The corresponding interface
type IAdsDeviceNotificationRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeviceNotificationRequest) CommandId() CommandId {
	return CommandId_ADS_DEVICE_NOTIFICATION
}

func (m *AdsDeviceNotificationRequest) Response() bool {
	return false
}


func (m *AdsDeviceNotificationRequest) InitializeParent(parent *AdsData) {
}

func NewAdsDeviceNotificationRequest(length uint32, stamps uint32, adsStampHeaders []*AdsStampHeader) *AdsData {
	child := &AdsDeviceNotificationRequest{
		Length: length,
		Stamps: stamps,
		AdsStampHeaders: adsStampHeaders,
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsDeviceNotificationRequest(structType interface{}) *AdsDeviceNotificationRequest {
	castFunc := func(typ interface{}) *AdsDeviceNotificationRequest {
		if casted, ok := typ.(AdsDeviceNotificationRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsDeviceNotificationRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsDeviceNotificationRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsDeviceNotificationRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsDeviceNotificationRequest) GetTypeName() string {
	return "AdsDeviceNotificationRequest"
}

func (m *AdsDeviceNotificationRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsDeviceNotificationRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (length)
	lengthInBits += 32;

	// Simple field (stamps)
	lengthInBits += 32;

	// Array field
	if len(m.AdsStampHeaders) > 0 {
		for i, element := range m.AdsStampHeaders {
			last := i == len(m.AdsStampHeaders) -1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}


func (m *AdsDeviceNotificationRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsDeviceNotificationRequestParse(readBuffer utils.ReadBuffer, ) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsDeviceNotificationRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (length)
length, _lengthErr := readBuffer.ReadUint32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (stamps)
stamps, _stampsErr := readBuffer.ReadUint32("stamps", 32)
	if _stampsErr != nil {
		return nil, errors.Wrap(_stampsErr, "Error parsing 'stamps' field")
	}

	// Array field (adsStampHeaders)
	if pullErr := readBuffer.PullContext("adsStampHeaders", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	adsStampHeaders := make([]*AdsStampHeader, stamps)
	for curItem := uint16(0); curItem < uint16(stamps); curItem++ {
		_item, _err := AdsStampHeaderParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'adsStampHeaders' field")
		}
		adsStampHeaders[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("adsStampHeaders", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsDeviceNotificationRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsDeviceNotificationRequest{
		Length: length,
		Stamps: stamps,
		AdsStampHeaders: adsStampHeaders,
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsDeviceNotificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeviceNotificationRequest"); pushErr != nil {
			return pushErr
		}

	// Simple Field (length)
	length := uint32(m.Length)
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (stamps)
	stamps := uint32(m.Stamps)
	_stampsErr := writeBuffer.WriteUint32("stamps", 32, (stamps))
	if _stampsErr != nil {
		return errors.Wrap(_stampsErr, "Error serializing 'stamps' field")
	}

	// Array Field (adsStampHeaders)
	if m.AdsStampHeaders != nil {
		if pushErr := writeBuffer.PushContext("adsStampHeaders", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.AdsStampHeaders {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'adsStampHeaders' field")
			}
		}
		if popErr := writeBuffer.PopContext("adsStampHeaders", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

		if popErr := writeBuffer.PopContext("AdsDeviceNotificationRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}



