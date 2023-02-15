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

// BACnetLifeSafetyMode is an enum
type BACnetLifeSafetyMode uint16

type IBACnetLifeSafetyMode interface {
	utils.Serializable
}

const(
	BACnetLifeSafetyMode_OFF BACnetLifeSafetyMode = 0
	BACnetLifeSafetyMode_ON BACnetLifeSafetyMode = 1
	BACnetLifeSafetyMode_TEST BACnetLifeSafetyMode = 2
	BACnetLifeSafetyMode_MANNED BACnetLifeSafetyMode = 3
	BACnetLifeSafetyMode_UNMANNED BACnetLifeSafetyMode = 4
	BACnetLifeSafetyMode_ARMED BACnetLifeSafetyMode = 5
	BACnetLifeSafetyMode_DISARMED BACnetLifeSafetyMode = 6
	BACnetLifeSafetyMode_PREARMED BACnetLifeSafetyMode = 7
	BACnetLifeSafetyMode_SLOW BACnetLifeSafetyMode = 8
	BACnetLifeSafetyMode_FAST BACnetLifeSafetyMode = 9
	BACnetLifeSafetyMode_DISCONNECTED BACnetLifeSafetyMode = 10
	BACnetLifeSafetyMode_ENABLED BACnetLifeSafetyMode = 11
	BACnetLifeSafetyMode_DISABLED BACnetLifeSafetyMode = 12
	BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED BACnetLifeSafetyMode = 13
	BACnetLifeSafetyMode_DEFAULT BACnetLifeSafetyMode = 14
	BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE BACnetLifeSafetyMode = 0XFFFF
)

var BACnetLifeSafetyModeValues []BACnetLifeSafetyMode

func init() {
	_ = errors.New
	BACnetLifeSafetyModeValues = []BACnetLifeSafetyMode {
		BACnetLifeSafetyMode_OFF,
		BACnetLifeSafetyMode_ON,
		BACnetLifeSafetyMode_TEST,
		BACnetLifeSafetyMode_MANNED,
		BACnetLifeSafetyMode_UNMANNED,
		BACnetLifeSafetyMode_ARMED,
		BACnetLifeSafetyMode_DISARMED,
		BACnetLifeSafetyMode_PREARMED,
		BACnetLifeSafetyMode_SLOW,
		BACnetLifeSafetyMode_FAST,
		BACnetLifeSafetyMode_DISCONNECTED,
		BACnetLifeSafetyMode_ENABLED,
		BACnetLifeSafetyMode_DISABLED,
		BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED,
		BACnetLifeSafetyMode_DEFAULT,
		BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLifeSafetyModeByValue(value uint16) (enum BACnetLifeSafetyMode, ok bool) {
	switch value {
		case 0:
			return BACnetLifeSafetyMode_OFF, true
		case 0XFFFF:
			return BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE, true
		case 1:
			return BACnetLifeSafetyMode_ON, true
		case 10:
			return BACnetLifeSafetyMode_DISCONNECTED, true
		case 11:
			return BACnetLifeSafetyMode_ENABLED, true
		case 12:
			return BACnetLifeSafetyMode_DISABLED, true
		case 13:
			return BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED, true
		case 14:
			return BACnetLifeSafetyMode_DEFAULT, true
		case 2:
			return BACnetLifeSafetyMode_TEST, true
		case 3:
			return BACnetLifeSafetyMode_MANNED, true
		case 4:
			return BACnetLifeSafetyMode_UNMANNED, true
		case 5:
			return BACnetLifeSafetyMode_ARMED, true
		case 6:
			return BACnetLifeSafetyMode_DISARMED, true
		case 7:
			return BACnetLifeSafetyMode_PREARMED, true
		case 8:
			return BACnetLifeSafetyMode_SLOW, true
		case 9:
			return BACnetLifeSafetyMode_FAST, true
	}
	return 0, false
}

func BACnetLifeSafetyModeByName(value string) (enum BACnetLifeSafetyMode, ok bool) {
	switch value {
	case "OFF":
		return BACnetLifeSafetyMode_OFF, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE, true
	case "ON":
		return BACnetLifeSafetyMode_ON, true
	case "DISCONNECTED":
		return BACnetLifeSafetyMode_DISCONNECTED, true
	case "ENABLED":
		return BACnetLifeSafetyMode_ENABLED, true
	case "DISABLED":
		return BACnetLifeSafetyMode_DISABLED, true
	case "AUTOMATIC_RELEASE_DISABLED":
		return BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED, true
	case "DEFAULT":
		return BACnetLifeSafetyMode_DEFAULT, true
	case "TEST":
		return BACnetLifeSafetyMode_TEST, true
	case "MANNED":
		return BACnetLifeSafetyMode_MANNED, true
	case "UNMANNED":
		return BACnetLifeSafetyMode_UNMANNED, true
	case "ARMED":
		return BACnetLifeSafetyMode_ARMED, true
	case "DISARMED":
		return BACnetLifeSafetyMode_DISARMED, true
	case "PREARMED":
		return BACnetLifeSafetyMode_PREARMED, true
	case "SLOW":
		return BACnetLifeSafetyMode_SLOW, true
	case "FAST":
		return BACnetLifeSafetyMode_FAST, true
	}
	return 0, false
}

func BACnetLifeSafetyModeKnows(value uint16)  bool {
	for _, typeValue := range BACnetLifeSafetyModeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastBACnetLifeSafetyMode(structType interface{}) BACnetLifeSafetyMode {
	castFunc := func(typ interface{}) BACnetLifeSafetyMode {
		if sBACnetLifeSafetyMode, ok := typ.(BACnetLifeSafetyMode); ok {
			return sBACnetLifeSafetyMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLifeSafetyMode) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetLifeSafetyMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLifeSafetyModeParse(ctx context.Context, theBytes []byte) (BACnetLifeSafetyMode, error) {
	return BACnetLifeSafetyModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLifeSafetyModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLifeSafetyMode, error) {
	val, err := readBuffer.ReadUint16("BACnetLifeSafetyMode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLifeSafetyMode")
	}
	if enum, ok := BACnetLifeSafetyModeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLifeSafetyMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLifeSafetyMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLifeSafetyMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLifeSafetyMode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLifeSafetyMode) PLC4XEnumName() string {
	switch e {
	case BACnetLifeSafetyMode_OFF:
		return "OFF"
	case BACnetLifeSafetyMode_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLifeSafetyMode_ON:
		return "ON"
	case BACnetLifeSafetyMode_DISCONNECTED:
		return "DISCONNECTED"
	case BACnetLifeSafetyMode_ENABLED:
		return "ENABLED"
	case BACnetLifeSafetyMode_DISABLED:
		return "DISABLED"
	case BACnetLifeSafetyMode_AUTOMATIC_RELEASE_DISABLED:
		return "AUTOMATIC_RELEASE_DISABLED"
	case BACnetLifeSafetyMode_DEFAULT:
		return "DEFAULT"
	case BACnetLifeSafetyMode_TEST:
		return "TEST"
	case BACnetLifeSafetyMode_MANNED:
		return "MANNED"
	case BACnetLifeSafetyMode_UNMANNED:
		return "UNMANNED"
	case BACnetLifeSafetyMode_ARMED:
		return "ARMED"
	case BACnetLifeSafetyMode_DISARMED:
		return "DISARMED"
	case BACnetLifeSafetyMode_PREARMED:
		return "PREARMED"
	case BACnetLifeSafetyMode_SLOW:
		return "SLOW"
	case BACnetLifeSafetyMode_FAST:
		return "FAST"
	}
	return ""
}

func (e BACnetLifeSafetyMode) String() string {
	return e.PLC4XEnumName()
}

