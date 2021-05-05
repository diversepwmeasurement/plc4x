//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SzlDataTreeItem struct {
	ItemIndex    uint16
	Mlfb         []byte
	ModuleTypeId uint16
	Ausbg        uint16
	Ausbe        uint16
}

// The corresponding interface
type ISzlDataTreeItem interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewSzlDataTreeItem(itemIndex uint16, mlfb []byte, moduleTypeId uint16, ausbg uint16, ausbe uint16) *SzlDataTreeItem {
	return &SzlDataTreeItem{ItemIndex: itemIndex, Mlfb: mlfb, ModuleTypeId: moduleTypeId, Ausbg: ausbg, Ausbe: ausbe}
}

func CastSzlDataTreeItem(structType interface{}) *SzlDataTreeItem {
	castFunc := func(typ interface{}) *SzlDataTreeItem {
		if casted, ok := typ.(SzlDataTreeItem); ok {
			return &casted
		}
		if casted, ok := typ.(*SzlDataTreeItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SzlDataTreeItem) GetTypeName() string {
	return "SzlDataTreeItem"
}

func (m *SzlDataTreeItem) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SzlDataTreeItem) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (itemIndex)
	lengthInBits += 16

	// Array field
	if len(m.Mlfb) > 0 {
		lengthInBits += 8 * uint16(len(m.Mlfb))
	}

	// Simple field (moduleTypeId)
	lengthInBits += 16

	// Simple field (ausbg)
	lengthInBits += 16

	// Simple field (ausbe)
	lengthInBits += 16

	return lengthInBits
}

func (m *SzlDataTreeItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SzlDataTreeItemParse(io utils.ReadBuffer) (*SzlDataTreeItem, error) {
	if pullErr := io.PullContext("SzlDataTreeItem"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (itemIndex)
	itemIndex, _itemIndexErr := io.ReadUint16("itemIndex", 16)
	if _itemIndexErr != nil {
		return nil, errors.Wrap(_itemIndexErr, "Error parsing 'itemIndex' field")
	}
	// Byte Array field (mlfb)
	numberOfBytes := int(uint16(20))
	mlfb, _readArrayErr := io.ReadByteArray("mlfb", numberOfBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'mlfb' field")
	}

	// Simple Field (moduleTypeId)
	moduleTypeId, _moduleTypeIdErr := io.ReadUint16("moduleTypeId", 16)
	if _moduleTypeIdErr != nil {
		return nil, errors.Wrap(_moduleTypeIdErr, "Error parsing 'moduleTypeId' field")
	}

	// Simple Field (ausbg)
	ausbg, _ausbgErr := io.ReadUint16("ausbg", 16)
	if _ausbgErr != nil {
		return nil, errors.Wrap(_ausbgErr, "Error parsing 'ausbg' field")
	}

	// Simple Field (ausbe)
	ausbe, _ausbeErr := io.ReadUint16("ausbe", 16)
	if _ausbeErr != nil {
		return nil, errors.Wrap(_ausbeErr, "Error parsing 'ausbe' field")
	}

	if closeErr := io.CloseContext("SzlDataTreeItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewSzlDataTreeItem(itemIndex, mlfb, moduleTypeId, ausbg, ausbe), nil
}

func (m *SzlDataTreeItem) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("SzlDataTreeItem"); pushErr != nil {
		return pushErr
	}

	// Simple Field (itemIndex)
	itemIndex := uint16(m.ItemIndex)
	_itemIndexErr := io.WriteUint16("itemIndex", 16, (itemIndex))
	if _itemIndexErr != nil {
		return errors.Wrap(_itemIndexErr, "Error serializing 'itemIndex' field")
	}

	// Array Field (mlfb)
	if m.Mlfb != nil {
		// Byte Array field (mlfb)
		_writeArrayErr := io.WriteByteArray("mlfb", m.Mlfb)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'mlfb' field")
		}
	}

	// Simple Field (moduleTypeId)
	moduleTypeId := uint16(m.ModuleTypeId)
	_moduleTypeIdErr := io.WriteUint16("moduleTypeId", 16, (moduleTypeId))
	if _moduleTypeIdErr != nil {
		return errors.Wrap(_moduleTypeIdErr, "Error serializing 'moduleTypeId' field")
	}

	// Simple Field (ausbg)
	ausbg := uint16(m.Ausbg)
	_ausbgErr := io.WriteUint16("ausbg", 16, (ausbg))
	if _ausbgErr != nil {
		return errors.Wrap(_ausbgErr, "Error serializing 'ausbg' field")
	}

	// Simple Field (ausbe)
	ausbe := uint16(m.Ausbe)
	_ausbeErr := io.WriteUint16("ausbe", 16, (ausbe))
	if _ausbeErr != nil {
		return errors.Wrap(_ausbeErr, "Error serializing 'ausbe' field")
	}

	if popErr := io.PopContext("SzlDataTreeItem"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *SzlDataTreeItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "itemIndex":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ItemIndex = data
			case "mlfb":
				var data []byte
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Mlfb = data
			case "moduleTypeId":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ModuleTypeId = data
			case "ausbg":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Ausbg = data
			case "ausbe":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Ausbe = data
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *SzlDataTreeItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.s7.readwrite.SzlDataTreeItem"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ItemIndex, xml.StartElement{Name: xml.Name{Local: "itemIndex"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Mlfb, xml.StartElement{Name: xml.Name{Local: "mlfb"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ModuleTypeId, xml.StartElement{Name: xml.Name{Local: "moduleTypeId"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Ausbg, xml.StartElement{Name: xml.Name{Local: "ausbg"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Ausbe, xml.StartElement{Name: xml.Name{Local: "ausbe"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m SzlDataTreeItem) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m SzlDataTreeItem) Box(name string, width int) utils.AsciiBox {
	boxName := "SzlDataTreeItem"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ItemIndex", m.ItemIndex, -1))
	// Array Field (mlfb)
	if m.Mlfb != nil {
		// Simple array base type byte will be rendered one by one
		arrayBoxes := make([]utils.AsciiBox, 0)
		for _, _element := range m.Mlfb {
			arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
		}
		boxes = append(boxes, utils.BoxBox("Mlfb", utils.AlignBoxes(arrayBoxes, width-4), 0))
	}
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ModuleTypeId", m.ModuleTypeId, -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Ausbg", m.Ausbg, -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Ausbe", m.Ausbe, -1))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
