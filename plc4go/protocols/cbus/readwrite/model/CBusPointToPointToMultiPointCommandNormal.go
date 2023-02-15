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


// CBusPointToPointToMultiPointCommandNormal is the corresponding interface of CBusPointToPointToMultiPointCommandNormal
type CBusPointToPointToMultiPointCommandNormal interface {
	utils.LengthAware
	utils.Serializable
	CBusPointToPointToMultiPointCommand
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
}

// CBusPointToPointToMultiPointCommandNormalExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointToMultiPointCommandNormal.
// This is useful for switch cases.
type CBusPointToPointToMultiPointCommandNormalExactly interface {
	CBusPointToPointToMultiPointCommandNormal
	isCBusPointToPointToMultiPointCommandNormal() bool
}

// _CBusPointToPointToMultiPointCommandNormal is the data-structure of this message
type _CBusPointToPointToMultiPointCommandNormal struct {
	*_CBusPointToPointToMultiPointCommand
        Application ApplicationIdContainer
        SalData SALData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToPointToMultiPointCommandNormal) InitializeParent(parent CBusPointToPointToMultiPointCommand , bridgeAddress BridgeAddress , networkRoute NetworkRoute , peekedApplication byte ) {	m.BridgeAddress = bridgeAddress
	m.NetworkRoute = networkRoute
	m.PeekedApplication = peekedApplication
}

func (m *_CBusPointToPointToMultiPointCommandNormal)  GetParent() CBusPointToPointToMultiPointCommand {
	return m._CBusPointToPointToMultiPointCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultiPointCommandNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CBusPointToPointToMultiPointCommandNormal) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCBusPointToPointToMultiPointCommandNormal factory function for _CBusPointToPointToMultiPointCommandNormal
func NewCBusPointToPointToMultiPointCommandNormal( application ApplicationIdContainer , salData SALData , bridgeAddress BridgeAddress , networkRoute NetworkRoute , peekedApplication byte , cBusOptions CBusOptions ) *_CBusPointToPointToMultiPointCommandNormal {
	_result := &_CBusPointToPointToMultiPointCommandNormal{
		Application: application,
		SalData: salData,
    	_CBusPointToPointToMultiPointCommand: NewCBusPointToPointToMultiPointCommand(bridgeAddress, networkRoute, peekedApplication, cBusOptions),
	}
	_result._CBusPointToPointToMultiPointCommand._CBusPointToPointToMultiPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultiPointCommandNormal(structType interface{}) CBusPointToPointToMultiPointCommandNormal {
    if casted, ok := structType.(CBusPointToPointToMultiPointCommandNormal); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultiPointCommandNormal); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultiPointCommandNormal) GetTypeName() string {
	return "CBusPointToPointToMultiPointCommandNormal"
}

func (m *_CBusPointToPointToMultiPointCommandNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (application)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_CBusPointToPointToMultiPointCommandNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToPointToMultiPointCommandNormalParse(theBytes []byte, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommandNormal, error) {
	return CBusPointToPointToMultiPointCommandNormalParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointToMultiPointCommandNormalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToPointToMultiPointCommandNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultiPointCommandNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultiPointCommandNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CBusPointToPointToMultiPointCommandNormal")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for salData")
	}
_salData, _salDataErr := SALDataParseWithBuffer(ctx, readBuffer , ApplicationId( application.ApplicationId() ) )
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field of CBusPointToPointToMultiPointCommandNormal")
	}
	salData := _salData.(SALData)
	if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for salData")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultiPointCommandNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultiPointCommandNormal")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToPointToMultiPointCommandNormal{
		_CBusPointToPointToMultiPointCommand: &_CBusPointToPointToMultiPointCommand{
			CBusOptions: cBusOptions,
		},
		Application: application,
		SalData: salData,
	}
	_child._CBusPointToPointToMultiPointCommand._CBusPointToPointToMultiPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToPointToMultiPointCommandNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToPointToMultiPointCommandNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToPointToMultiPointCommandNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultiPointCommandNormal")
		}

	// Simple Field (application)
	if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for application")
	}
	_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
	if popErr := writeBuffer.PopContext("application"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for application")
	}
	if _applicationErr != nil {
		return errors.Wrap(_applicationErr, "Error serializing 'application' field")
	}

	// Simple Field (salData)
	if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for salData")
	}
	_salDataErr := writeBuffer.WriteSerializable(ctx, m.GetSalData())
	if popErr := writeBuffer.PopContext("salData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for salData")
	}
	if _salDataErr != nil {
		return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
	}

		if popErr := writeBuffer.PopContext("CBusPointToPointToMultiPointCommandNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultiPointCommandNormal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_CBusPointToPointToMultiPointCommandNormal) isCBusPointToPointToMultiPointCommandNormal() bool {
	return true
}

func (m *_CBusPointToPointToMultiPointCommandNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



