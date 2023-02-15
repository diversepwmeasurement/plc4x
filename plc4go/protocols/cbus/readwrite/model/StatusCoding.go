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

// StatusCoding is an enum
type StatusCoding byte

type IStatusCoding interface {
	utils.Serializable
}

const(
	StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE StatusCoding = 0x00
	StatusCoding_BINARY_BY_ELSEWHERE StatusCoding = 0x40
	StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE StatusCoding = 0x07
	StatusCoding_LEVEL_BY_ELSEWHERE StatusCoding = 0x47
)

var StatusCodingValues []StatusCoding

func init() {
	_ = errors.New
	StatusCodingValues = []StatusCoding {
		StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE,
		StatusCoding_BINARY_BY_ELSEWHERE,
		StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE,
		StatusCoding_LEVEL_BY_ELSEWHERE,
	}
}

func StatusCodingByValue(value byte) (enum StatusCoding, ok bool) {
	switch value {
		case 0x00:
			return StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE, true
		case 0x07:
			return StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE, true
		case 0x40:
			return StatusCoding_BINARY_BY_ELSEWHERE, true
		case 0x47:
			return StatusCoding_LEVEL_BY_ELSEWHERE, true
	}
	return 0, false
}

func StatusCodingByName(value string) (enum StatusCoding, ok bool) {
	switch value {
	case "BINARY_BY_THIS_SERIAL_INTERFACE":
		return StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE, true
	case "LEVEL_BY_THIS_SERIAL_INTERFACE":
		return StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE, true
	case "BINARY_BY_ELSEWHERE":
		return StatusCoding_BINARY_BY_ELSEWHERE, true
	case "LEVEL_BY_ELSEWHERE":
		return StatusCoding_LEVEL_BY_ELSEWHERE, true
	}
	return 0, false
}

func StatusCodingKnows(value byte)  bool {
	for _, typeValue := range StatusCodingValues {
		if byte(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastStatusCoding(structType interface{}) StatusCoding {
	castFunc := func(typ interface{}) StatusCoding {
		if sStatusCoding, ok := typ.(StatusCoding); ok {
			return sStatusCoding
		}
		return 0
	}
	return castFunc(structType)
}

func (m StatusCoding) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m StatusCoding) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusCodingParse(ctx context.Context, theBytes []byte) (StatusCoding, error) {
	return StatusCodingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StatusCodingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StatusCoding, error) {
	val, err := readBuffer.ReadByte("StatusCoding")
	if err != nil {
		return 0, errors.Wrap(err, "error reading StatusCoding")
	}
	if enum, ok := StatusCodingByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return StatusCoding(val), nil
	} else {
		return enum, nil
	}
}

func (e StatusCoding) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e StatusCoding) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteByte("StatusCoding", byte(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e StatusCoding) PLC4XEnumName() string {
	switch e {
	case StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE:
		return "BINARY_BY_THIS_SERIAL_INTERFACE"
	case StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE:
		return "LEVEL_BY_THIS_SERIAL_INTERFACE"
	case StatusCoding_BINARY_BY_ELSEWHERE:
		return "BINARY_BY_ELSEWHERE"
	case StatusCoding_LEVEL_BY_ELSEWHERE:
		return "LEVEL_BY_ELSEWHERE"
	}
	return ""
}

func (e StatusCoding) String() string {
	return e.PLC4XEnumName()
}

