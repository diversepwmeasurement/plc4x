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
package org.apache.plc4x.java.eip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class TypeId implements Message {

  // Abstract accessors for discriminator values.
  public abstract Integer getId();

  // Arguments.
  protected final IntegerEncoding order;

  public TypeId(IntegerEncoding order) {
    super();
    this.order = order;
  }

  protected abstract void serializeTypeIdChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("TypeId");

    // Discriminator Field (id) (Used as input to a switch field)
    writeDiscriminatorField(
        "id",
        getId(),
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Switch field (Serialize the sub-type)
    serializeTypeIdChild(writeBuffer);

    writeBuffer.popContext("TypeId");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    TypeId _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (id)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static TypeId staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    IntegerEncoding order;
    if (args[0] instanceof IntegerEncoding) {
      order = (IntegerEncoding) args[0];
    } else if (args[0] instanceof String) {
      order = IntegerEncoding.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type IntegerEncoding or a string which is parseable but was"
              + " "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, order);
  }

  public static TypeId staticParse(ReadBuffer readBuffer, IntegerEncoding order)
      throws ParseException {
    readBuffer.pullContext("TypeId");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int id =
        readDiscriminatorField(
            "id",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    TypeIdBuilder builder = null;
    if (EvaluationHelper.equals(id, (int) 0x0000)) {
      builder = NullAddressItem.staticParseTypeIdBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(id, (int) 0x0100)) {
      builder = ServicesResponse.staticParseTypeIdBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(id, (int) 0x00A1)) {
      builder = ConnectedAddressItem.staticParseTypeIdBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(id, (int) 0x00B1)) {
      builder = ConnectedDataItem.staticParseTypeIdBuilder(readBuffer, order);
    } else if (EvaluationHelper.equals(id, (int) 0x00B2)) {
      builder = UnConnectedDataItem.staticParseTypeIdBuilder(readBuffer, order);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "id=" + id + "]");
    }

    readBuffer.closeContext("TypeId");
    // Create the instance
    TypeId _typeId = builder.build(order);

    return _typeId;
  }

  public interface TypeIdBuilder {
    TypeId build(IntegerEncoding order);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TypeId)) {
      return false;
    }
    TypeId that = (TypeId) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
