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

#include <string.h>
#include <plc4c/driver_s7_static.h>

#include "controller_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_s7_read_write_controller_type plc4c_s7_read_write_controller_type_null_const;

plc4c_s7_read_write_controller_type plc4c_s7_read_write_controller_type_null() {
  return plc4c_s7_read_write_controller_type_null_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_controller_type_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_controller_type* _message) {
    plc4c_return_code _res = OK;

    uint32_t value;
    _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
    *_message = plc4c_s7_read_write_controller_type_for_value(value);

    return _res;
}

plc4c_return_code plc4c_s7_read_write_controller_type_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_controller_type* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, *_message);

    return _res;
}

plc4c_s7_read_write_controller_type plc4c_s7_read_write_controller_type_for_value(uint32_t value) {
    for(int i = 0; i < plc4c_s7_read_write_controller_type_num_values(); i++) {
        if(plc4c_s7_read_write_controller_type_value_for_index(i) == value) {
            return plc4c_s7_read_write_controller_type_value_for_index(i);
        }
    }
    return -1;
}

plc4c_s7_read_write_controller_type plc4c_s7_read_write_controller_type_value_of(char* value_string) {
    if(strcmp(value_string, "ANY") == 0) {
        return plc4c_s7_read_write_controller_type_ANY;
    }
    if(strcmp(value_string, "S7_200") == 0) {
        return plc4c_s7_read_write_controller_type_S7_200;
    }
    if(strcmp(value_string, "S7_300") == 0) {
        return plc4c_s7_read_write_controller_type_S7_300;
    }
    if(strcmp(value_string, "S7_400") == 0) {
        return plc4c_s7_read_write_controller_type_S7_400;
    }
    if(strcmp(value_string, "S7_1200") == 0) {
        return plc4c_s7_read_write_controller_type_S7_1200;
    }
    if(strcmp(value_string, "S7_1500") == 0) {
        return plc4c_s7_read_write_controller_type_S7_1500;
    }
    if(strcmp(value_string, "LOGO") == 0) {
        return plc4c_s7_read_write_controller_type_LOGO;
    }
    return -1;
}

int plc4c_s7_read_write_controller_type_num_values() {
  return 7;
}

plc4c_s7_read_write_controller_type plc4c_s7_read_write_controller_type_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_controller_type_ANY;
      }
      case 1: {
        return plc4c_s7_read_write_controller_type_S7_200;
      }
      case 2: {
        return plc4c_s7_read_write_controller_type_S7_300;
      }
      case 3: {
        return plc4c_s7_read_write_controller_type_S7_400;
      }
      case 4: {
        return plc4c_s7_read_write_controller_type_S7_1200;
      }
      case 5: {
        return plc4c_s7_read_write_controller_type_S7_1500;
      }
      case 6: {
        return plc4c_s7_read_write_controller_type_LOGO;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_s7_read_write_controller_type_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_controller_type* _message) {
    return plc4c_s7_read_write_controller_type_length_in_bits(ctx, _message) / 8;
}

uint16_t plc4c_s7_read_write_controller_type_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_controller_type* _message) {
    return 32;
}