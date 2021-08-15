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

#include "memory_area.h"
#include <string.h>

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_null_const;

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_null() {
  return plc4c_s7_read_write_memory_area_null_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_memory_area_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_memory_area** _message) {
    plc4c_return_code _res = OK;
    // TODO: Implement
    return _res;
}

plc4c_return_code plc4c_s7_read_write_memory_area_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_memory_area* _message) {
    plc4c_return_code _res = OK;
    // TODO: Implement
    return _res;
}

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_value_of(char* value_string) {
    if(strcmp(value_string, "COUNTERS") == 0) {
        return plc4c_s7_read_write_memory_area_COUNTERS;
    }
    if(strcmp(value_string, "TIMERS") == 0) {
        return plc4c_s7_read_write_memory_area_TIMERS;
    }
    if(strcmp(value_string, "DIRECT_PERIPHERAL_ACCESS") == 0) {
        return plc4c_s7_read_write_memory_area_DIRECT_PERIPHERAL_ACCESS;
    }
    if(strcmp(value_string, "INPUTS") == 0) {
        return plc4c_s7_read_write_memory_area_INPUTS;
    }
    if(strcmp(value_string, "OUTPUTS") == 0) {
        return plc4c_s7_read_write_memory_area_OUTPUTS;
    }
    if(strcmp(value_string, "FLAGS_MARKERS") == 0) {
        return plc4c_s7_read_write_memory_area_FLAGS_MARKERS;
    }
    if(strcmp(value_string, "DATA_BLOCKS") == 0) {
        return plc4c_s7_read_write_memory_area_DATA_BLOCKS;
    }
    if(strcmp(value_string, "INSTANCE_DATA_BLOCKS") == 0) {
        return plc4c_s7_read_write_memory_area_INSTANCE_DATA_BLOCKS;
    }
    if(strcmp(value_string, "LOCAL_DATA") == 0) {
        return plc4c_s7_read_write_memory_area_LOCAL_DATA;
    }
    return -1;
}

int plc4c_s7_read_write_memory_area_num_values() {
  return 9;
}

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_memory_area_COUNTERS;
      }
      case 1: {
        return plc4c_s7_read_write_memory_area_TIMERS;
      }
      case 2: {
        return plc4c_s7_read_write_memory_area_DIRECT_PERIPHERAL_ACCESS;
      }
      case 3: {
        return plc4c_s7_read_write_memory_area_INPUTS;
      }
      case 4: {
        return plc4c_s7_read_write_memory_area_OUTPUTS;
      }
      case 5: {
        return plc4c_s7_read_write_memory_area_FLAGS_MARKERS;
      }
      case 6: {
        return plc4c_s7_read_write_memory_area_DATA_BLOCKS;
      }
      case 7: {
        return plc4c_s7_read_write_memory_area_INSTANCE_DATA_BLOCKS;
      }
      case 8: {
        return plc4c_s7_read_write_memory_area_LOCAL_DATA;
      }
      default: {
        return -1;
      }
    }
}

char* plc4c_s7_read_write_memory_area_get_short_name(plc4c_s7_read_write_memory_area value) {
  switch(value) {
    case plc4c_s7_read_write_memory_area_COUNTERS: { /* '0x1C' */
      return "C";
    }
    case plc4c_s7_read_write_memory_area_TIMERS: { /* '0x1D' */
      return "T";
    }
    case plc4c_s7_read_write_memory_area_DIRECT_PERIPHERAL_ACCESS: { /* '0x80' */
      return "D";
    }
    case plc4c_s7_read_write_memory_area_INPUTS: { /* '0x81' */
      return "I";
    }
    case plc4c_s7_read_write_memory_area_OUTPUTS: { /* '0x82' */
      return "Q";
    }
    case plc4c_s7_read_write_memory_area_FLAGS_MARKERS: { /* '0x83' */
      return "M";
    }
    case plc4c_s7_read_write_memory_area_DATA_BLOCKS: { /* '0x84' */
      return "DB";
    }
    case plc4c_s7_read_write_memory_area_INSTANCE_DATA_BLOCKS: { /* '0x85' */
      return "DBI";
    }
    case plc4c_s7_read_write_memory_area_LOCAL_DATA: { /* '0x86' */
      return "LD";
    }
    default: {
      return 0;
    }
  }
}

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_get_first_enum_for_field_short_name(char* value) {
    if (strcmp(value, "C") == 0) {
        return plc4c_s7_read_write_memory_area_COUNTERS;
    }
    if (strcmp(value, "D") == 0) {
        return plc4c_s7_read_write_memory_area_DIRECT_PERIPHERAL_ACCESS;
    }
    if (strcmp(value, "DB") == 0) {
        return plc4c_s7_read_write_memory_area_DATA_BLOCKS;
    }
    if (strcmp(value, "DBI") == 0) {
        return plc4c_s7_read_write_memory_area_INSTANCE_DATA_BLOCKS;
    }
    if (strcmp(value, "I") == 0) {
        return plc4c_s7_read_write_memory_area_INPUTS;
    }
    if (strcmp(value, "LD") == 0) {
        return plc4c_s7_read_write_memory_area_LOCAL_DATA;
    }
    if (strcmp(value, "M") == 0) {
        return plc4c_s7_read_write_memory_area_FLAGS_MARKERS;
    }
    if (strcmp(value, "Q") == 0) {
        return plc4c_s7_read_write_memory_area_OUTPUTS;
    }
    if (strcmp(value, "T") == 0) {
        return plc4c_s7_read_write_memory_area_TIMERS;
    }
}

uint16_t plc4c_s7_read_write_memory_area_length_in_bytes(plc4c_s7_read_write_memory_area* _message) {
    return plc4c_s7_read_write_memory_area_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_memory_area_length_in_bits(plc4c_s7_read_write_memory_area* _message) {
    // TODO: Implement
    return 0;
}
