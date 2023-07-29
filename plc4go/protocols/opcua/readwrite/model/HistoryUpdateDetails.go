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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// HistoryUpdateDetails is the corresponding interface of HistoryUpdateDetails
type HistoryUpdateDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
}

// HistoryUpdateDetailsExactly can be used when we want exactly this type and not a type which fulfills HistoryUpdateDetails.
// This is useful for switch cases.
type HistoryUpdateDetailsExactly interface {
	HistoryUpdateDetails
	isHistoryUpdateDetails() bool
}

// _HistoryUpdateDetails is the data-structure of this message
type _HistoryUpdateDetails struct {
	*_ExtensionObjectDefinition
	NodeId NodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateDetails) GetIdentifier() string {
	return "679"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateDetails) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_HistoryUpdateDetails) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateDetails) GetNodeId() NodeId {
	return m.NodeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryUpdateDetails factory function for _HistoryUpdateDetails
func NewHistoryUpdateDetails(nodeId NodeId) *_HistoryUpdateDetails {
	_result := &_HistoryUpdateDetails{
		NodeId:                     nodeId,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryUpdateDetails(structType any) HistoryUpdateDetails {
	if casted, ok := structType.(HistoryUpdateDetails); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateDetails); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateDetails) GetTypeName() string {
	return "HistoryUpdateDetails"
}

func (m *_HistoryUpdateDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_HistoryUpdateDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryUpdateDetailsParse(ctx context.Context, theBytes []byte, identifier string) (HistoryUpdateDetails, error) {
	return HistoryUpdateDetailsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryUpdateDetailsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryUpdateDetails, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryUpdateDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of HistoryUpdateDetails")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	if closeErr := readBuffer.CloseContext("HistoryUpdateDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateDetails")
	}

	// Create a partially initialized instance
	_child := &_HistoryUpdateDetails{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryUpdateDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateDetails")
		}

		// Simple Field (nodeId)
		if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeId")
		}
		_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
		if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeId")
		}
		if _nodeIdErr != nil {
			return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
		}

		if popErr := writeBuffer.PopContext("HistoryUpdateDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateDetails")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryUpdateDetails) isHistoryUpdateDetails() bool {
	return true
}

func (m *_HistoryUpdateDetails) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}