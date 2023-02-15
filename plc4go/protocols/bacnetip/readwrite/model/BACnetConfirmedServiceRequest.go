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


// BACnetConfirmedServiceRequest is the corresponding interface of BACnetConfirmedServiceRequest
type BACnetConfirmedServiceRequest interface {
	utils.LengthAware
	utils.Serializable
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetConfirmedServiceChoice
	// GetServiceRequestPayloadLength returns ServiceRequestPayloadLength (virtual field)
	GetServiceRequestPayloadLength() uint32
}

// BACnetConfirmedServiceRequestExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequest.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestExactly interface {
	BACnetConfirmedServiceRequest
	isBACnetConfirmedServiceRequest() bool
}

// _BACnetConfirmedServiceRequest is the data-structure of this message
type _BACnetConfirmedServiceRequest struct {
	_BACnetConfirmedServiceRequestChildRequirements

	// Arguments.
	ServiceRequestLength uint32
}

type _BACnetConfirmedServiceRequestChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetServiceChoice() BACnetConfirmedServiceChoice
}


type BACnetConfirmedServiceRequestParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetConfirmedServiceRequestChild interface {
	utils.Serializable
InitializeParent(parent BACnetConfirmedServiceRequest )
	GetParent() *BACnetConfirmedServiceRequest

	GetTypeName() string
	BACnetConfirmedServiceRequest
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequest) GetServiceRequestPayloadLength() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(utils.InlineIf((bool((m.ServiceRequestLength) > ((0)))), func() interface{} {return uint32((uint32(m.ServiceRequestLength) - uint32(uint32(1))))}, func() interface{} {return uint32(uint32(0))}).(uint32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConfirmedServiceRequest factory function for _BACnetConfirmedServiceRequest
func NewBACnetConfirmedServiceRequest( serviceRequestLength uint32 ) *_BACnetConfirmedServiceRequest {
return &_BACnetConfirmedServiceRequest{ ServiceRequestLength: serviceRequestLength }
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequest(structType interface{}) BACnetConfirmedServiceRequest {
    if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequest) GetTypeName() string {
	return "BACnetConfirmedServiceRequest"
}


func (m *_BACnetConfirmedServiceRequest) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8;

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestParse(theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequest, error) {
	return BACnetConfirmedServiceRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("serviceChoice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceChoice")
	}
	serviceChoice_temp, _serviceChoiceErr := BACnetConfirmedServiceChoiceParseWithBuffer(ctx, readBuffer)
	var serviceChoice BACnetConfirmedServiceChoice = serviceChoice_temp
	if closeErr := readBuffer.CloseContext("serviceChoice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceChoice")
	}
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field of BACnetConfirmedServiceRequest")
	}

	// Virtual field
	_serviceRequestPayloadLength := utils.InlineIf((bool((serviceRequestLength) > ((0)))), func() interface{} {return uint32((uint32(serviceRequestLength) - uint32(uint32(1))))}, func() interface{} {return uint32(uint32(0))}).(uint32)
	serviceRequestPayloadLength := uint32(_serviceRequestPayloadLength)
	_ = serviceRequestPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetConfirmedServiceRequestChildSerializeRequirement interface {
		BACnetConfirmedServiceRequest
		InitializeParent(BACnetConfirmedServiceRequest )
		GetParent() BACnetConfirmedServiceRequest
	}
	var _childTemp interface{}
	var _child BACnetConfirmedServiceRequestChildSerializeRequirement
	var typeSwitchError error
	switch {
case serviceChoice == BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM : // BACnetConfirmedServiceRequestAcknowledgeAlarm
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAcknowledgeAlarmParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION : // BACnetConfirmedServiceRequestConfirmedCOVNotification
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE : // BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_EVENT_NOTIFICATION : // BACnetConfirmedServiceRequestConfirmedEventNotification
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedEventNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY : // BACnetConfirmedServiceRequestGetEnrollmentSummary
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestGetEnrollmentSummaryParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION : // BACnetConfirmedServiceRequestGetEventInformation
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestGetEventInformationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION : // BACnetConfirmedServiceRequestLifeSafetyOperation
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestLifeSafetyOperationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV : // BACnetConfirmedServiceRequestSubscribeCOV
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY : // BACnetConfirmedServiceRequestSubscribeCOVProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE : // BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_READ_FILE : // BACnetConfirmedServiceRequestAtomicReadFile
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAtomicReadFileParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE : // BACnetConfirmedServiceRequestAtomicWriteFile
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAtomicWriteFileParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT : // BACnetConfirmedServiceRequestAddListElement
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAddListElementParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT : // BACnetConfirmedServiceRequestRemoveListElement
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestRemoveListElementParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT : // BACnetConfirmedServiceRequestCreateObject
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestCreateObjectParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_DELETE_OBJECT : // BACnetConfirmedServiceRequestDeleteObject
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestDeleteObjectParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY : // BACnetConfirmedServiceRequestReadProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE : // BACnetConfirmedServiceRequestReadPropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyMultipleParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_READ_RANGE : // BACnetConfirmedServiceRequestReadRange
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadRangeParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY : // BACnetConfirmedServiceRequestWriteProperty
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE : // BACnetConfirmedServiceRequestWritePropertyMultiple
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestWritePropertyMultipleParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_DEVICE_COMMUNICATION_CONTROL : // BACnetConfirmedServiceRequestDeviceCommunicationControl
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestDeviceCommunicationControlParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER : // BACnetConfirmedServiceRequestConfirmedPrivateTransfer
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedPrivateTransferParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_TEXT_MESSAGE : // BACnetConfirmedServiceRequestConfirmedTextMessage
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestConfirmedTextMessageParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_REINITIALIZE_DEVICE : // BACnetConfirmedServiceRequestReinitializeDevice
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReinitializeDeviceParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_VT_OPEN : // BACnetConfirmedServiceRequestVTOpen
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTOpenParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_VT_CLOSE : // BACnetConfirmedServiceRequestVTClose
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTCloseParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_VT_DATA : // BACnetConfirmedServiceRequestVTData
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestVTDataParseWithBuffer(ctx, readBuffer, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_AUTHENTICATE : // BACnetConfirmedServiceRequestAuthenticate
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestAuthenticateParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_REQUEST_KEY : // BACnetConfirmedServiceRequestRequestKey
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestRequestKeyParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL : // BACnetConfirmedServiceRequestReadPropertyConditional
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestReadPropertyConditionalParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
case 0==0 : // BACnetConfirmedServiceRequestUnknown
		_childTemp, typeSwitchError = BACnetConfirmedServiceRequestUnknownParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetConfirmedServiceRequest")
	}
	_child = _childTemp.(BACnetConfirmedServiceRequestChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequest")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_BACnetConfirmedServiceRequest) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetConfirmedServiceRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetConfirmedServiceRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequest")
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := BACnetConfirmedServiceChoice(child.GetServiceChoice())
	if pushErr := writeBuffer.PushContext("serviceChoice"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for serviceChoice")
	}
	_serviceChoiceErr := writeBuffer.WriteSerializable(ctx, serviceChoice)
	if popErr := writeBuffer.PopContext("serviceChoice"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for serviceChoice")
	}

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}
	// Virtual field
	if _serviceRequestPayloadLengthErr := writeBuffer.WriteVirtual(ctx, "serviceRequestPayloadLength", m.GetServiceRequestPayloadLength()); _serviceRequestPayloadLengthErr != nil {
		return errors.Wrap(_serviceRequestPayloadLengthErr, "Error serializing 'serviceRequestPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequest")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequest) GetServiceRequestLength() uint32 {
	return m.ServiceRequestLength
}
//
////

func (m *_BACnetConfirmedServiceRequest) isBACnetConfirmedServiceRequest() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



