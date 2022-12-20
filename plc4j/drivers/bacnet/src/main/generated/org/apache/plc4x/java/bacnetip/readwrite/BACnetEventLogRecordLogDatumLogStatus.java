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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetEventLogRecordLogDatumLogStatus extends BACnetEventLogRecordLogDatum
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetLogStatusTagged logStatus;

  // Arguments.
  protected final Short tagNumber;

  public BACnetEventLogRecordLogDatumLogStatus(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetLogStatusTagged logStatus,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.logStatus = logStatus;
    this.tagNumber = tagNumber;
  }

  public BACnetLogStatusTagged getLogStatus() {
    return logStatus;
  }

  @Override
  protected void serializeBACnetEventLogRecordLogDatumChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventLogRecordLogDatumLogStatus");

    // Simple Field (logStatus)
    writeSimpleField("logStatus", logStatus, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventLogRecordLogDatumLogStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventLogRecordLogDatumLogStatus _value = this;

    // Simple field (logStatus)
    lengthInBits += logStatus.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventLogRecordLogDatumLogStatusBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetEventLogRecordLogDatumLogStatus");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetLogStatusTagged logStatus =
        readSimpleField(
            "logStatus",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetLogStatusTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetEventLogRecordLogDatumLogStatus");
    // Create the instance
    return new BACnetEventLogRecordLogDatumLogStatusBuilder(logStatus, tagNumber);
  }

  public static class BACnetEventLogRecordLogDatumLogStatusBuilder
      implements BACnetEventLogRecordLogDatum.BACnetEventLogRecordLogDatumBuilder {
    private final BACnetLogStatusTagged logStatus;
    private final Short tagNumber;

    public BACnetEventLogRecordLogDatumLogStatusBuilder(
        BACnetLogStatusTagged logStatus, Short tagNumber) {

      this.logStatus = logStatus;
      this.tagNumber = tagNumber;
    }

    public BACnetEventLogRecordLogDatumLogStatus build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetEventLogRecordLogDatumLogStatus bACnetEventLogRecordLogDatumLogStatus =
          new BACnetEventLogRecordLogDatumLogStatus(
              openingTag, peekedTagHeader, closingTag, logStatus, tagNumber);
      return bACnetEventLogRecordLogDatumLogStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventLogRecordLogDatumLogStatus)) {
      return false;
    }
    BACnetEventLogRecordLogDatumLogStatus that = (BACnetEventLogRecordLogDatumLogStatus) o;
    return (getLogStatus() == that.getLogStatus()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLogStatus());
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