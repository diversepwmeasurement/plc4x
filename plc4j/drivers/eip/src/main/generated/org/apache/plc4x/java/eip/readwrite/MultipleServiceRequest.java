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

public class MultipleServiceRequest extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0x0A;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Boolean getConnected() {
    return false;
  }

  // Constant values.
  public static final Byte REQUESTPATHSIZE = 0x02;
  public static final Long REQUESTPATH = 0x01240220L;

  // Properties.
  protected final Services data;

  // Arguments.
  protected final Integer serviceLen;
  protected final IntegerEncoding order;

  public MultipleServiceRequest(Services data, Integer serviceLen, IntegerEncoding order) {
    super(serviceLen, order);
    this.data = data;
    this.serviceLen = serviceLen;
    this.order = order;
  }

  public Services getData() {
    return data;
  }

  public byte getRequestPathSize() {
    return REQUESTPATHSIZE;
  }

  public long getRequestPath() {
    return REQUESTPATH;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MultipleServiceRequest");

    // Const Field (requestPathSize)
    writeConstField(
        "requestPathSize",
        REQUESTPATHSIZE,
        writeSignedByte(writeBuffer, 8),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Const Field (requestPath)
    writeConstField(
        "requestPath",
        REQUESTPATH,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    // Simple Field (data)
    writeSimpleField(
        "data",
        data,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(
            (((order) == (IntegerEncoding.BIG_ENDIAN))
                ? ByteOrder.BIG_ENDIAN
                : ByteOrder.LITTLE_ENDIAN)));

    writeBuffer.popContext("MultipleServiceRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MultipleServiceRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (requestPathSize)
    lengthInBits += 8;

    // Const Field (requestPath)
    lengthInBits += 32;

    // Simple field (data)
    lengthInBits += data.getLengthInBits();

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen, IntegerEncoding order)
      throws ParseException {
    readBuffer.pullContext("MultipleServiceRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte requestPathSize =
        readConstField(
            "requestPathSize",
            readSignedByte(readBuffer, 8),
            MultipleServiceRequest.REQUESTPATHSIZE,
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    long requestPath =
        readConstField(
            "requestPath",
            readUnsignedLong(readBuffer, 32),
            MultipleServiceRequest.REQUESTPATH,
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    Services data =
        readSimpleField(
            "data",
            new DataReaderComplexDefault<>(
                () ->
                    Services.staticParse(
                        readBuffer, (int) ((serviceLen) - (6)), (IntegerEncoding) (order)),
                readBuffer),
            WithOption.WithByteOrder(
                (((order) == (IntegerEncoding.BIG_ENDIAN))
                    ? ByteOrder.BIG_ENDIAN
                    : ByteOrder.LITTLE_ENDIAN)));

    readBuffer.closeContext("MultipleServiceRequest");
    // Create the instance
    return new MultipleServiceRequestBuilderImpl(data, serviceLen, order);
  }

  public static class MultipleServiceRequestBuilderImpl implements CipService.CipServiceBuilder {
    private final Services data;
    private final Integer serviceLen;
    private final IntegerEncoding order;

    public MultipleServiceRequestBuilderImpl(
        Services data, Integer serviceLen, IntegerEncoding order) {
      this.data = data;
      this.serviceLen = serviceLen;
      this.order = order;
    }

    public MultipleServiceRequest build(Integer serviceLen, IntegerEncoding order) {

      MultipleServiceRequest multipleServiceRequest =
          new MultipleServiceRequest(data, serviceLen, order);
      return multipleServiceRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MultipleServiceRequest)) {
      return false;
    }
    MultipleServiceRequest that = (MultipleServiceRequest) o;
    return (getData() == that.getData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getData());
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
