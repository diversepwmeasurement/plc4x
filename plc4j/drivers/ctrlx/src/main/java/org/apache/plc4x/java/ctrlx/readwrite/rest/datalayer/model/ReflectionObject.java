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

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.StringJoiner;

/**
 * ReflectionObject
 */
@JsonPropertyOrder({
  ReflectionObject.JSON_PROPERTY_NAME,
  ReflectionObject.JSON_PROPERTY_FIELDS,
  ReflectionObject.JSON_PROPERTY_IS_STRUCT,
  ReflectionObject.JSON_PROPERTY_MINALIGN,
  ReflectionObject.JSON_PROPERTY_BYTESIZE,
  ReflectionObject.JSON_PROPERTY_ATTRIBUTES,
  ReflectionObject.JSON_PROPERTY_DOCUMENTATION
})
@JsonTypeName("reflection_Object")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class ReflectionObject {
  public static final String JSON_PROPERTY_NAME = "name";
  private String name;

  public static final String JSON_PROPERTY_FIELDS = "fields";
  private List<ReflectionField> fields = new ArrayList<>();

  public static final String JSON_PROPERTY_IS_STRUCT = "is_struct";
  private Boolean isStruct;

  public static final String JSON_PROPERTY_MINALIGN = "minalign";
  private Integer minalign;

  public static final String JSON_PROPERTY_BYTESIZE = "bytesize";
  private Integer bytesize;

  public static final String JSON_PROPERTY_ATTRIBUTES = "attributes";
  private List<ReflectionKeyValue> attributes = new ArrayList<>();

  public static final String JSON_PROPERTY_DOCUMENTATION = "documentation";
  private List<String> documentation = new ArrayList<>();

  public ReflectionObject() {
  }

  public ReflectionObject name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getName() {
    return name;
  }


  @JsonProperty(JSON_PROPERTY_NAME)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setName(String name) {
    this.name = name;
  }


  public ReflectionObject fields(List<ReflectionField> fields) {
    
    this.fields = fields;
    return this;
  }

  public ReflectionObject addFieldsItem(ReflectionField fieldsItem) {
    this.fields.add(fieldsItem);
    return this;
  }

   /**
   * Get fields
   * @return fields
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_FIELDS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public List<ReflectionField> getFields() {
    return fields;
  }


  @JsonProperty(JSON_PROPERTY_FIELDS)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setFields(List<ReflectionField> fields) {
    this.fields = fields;
  }


  public ReflectionObject isStruct(Boolean isStruct) {
    
    this.isStruct = isStruct;
    return this;
  }

   /**
   * Get isStruct
   * @return isStruct
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_IS_STRUCT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Boolean getIsStruct() {
    return isStruct;
  }


  @JsonProperty(JSON_PROPERTY_IS_STRUCT)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setIsStruct(Boolean isStruct) {
    this.isStruct = isStruct;
  }


  public ReflectionObject minalign(Integer minalign) {
    
    this.minalign = minalign;
    return this;
  }

   /**
   * Get minalign
   * minimum: -2147483648
   * maximum: 2147483647
   * @return minalign
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_MINALIGN)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getMinalign() {
    return minalign;
  }


  @JsonProperty(JSON_PROPERTY_MINALIGN)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMinalign(Integer minalign) {
    this.minalign = minalign;
  }


  public ReflectionObject bytesize(Integer bytesize) {
    
    this.bytesize = bytesize;
    return this;
  }

   /**
   * Get bytesize
   * minimum: -2147483648
   * maximum: 2147483647
   * @return bytesize
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_BYTESIZE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getBytesize() {
    return bytesize;
  }


  @JsonProperty(JSON_PROPERTY_BYTESIZE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setBytesize(Integer bytesize) {
    this.bytesize = bytesize;
  }


  public ReflectionObject attributes(List<ReflectionKeyValue> attributes) {
    
    this.attributes = attributes;
    return this;
  }

  public ReflectionObject addAttributesItem(ReflectionKeyValue attributesItem) {
    if (this.attributes == null) {
      this.attributes = new ArrayList<>();
    }
    this.attributes.add(attributesItem);
    return this;
  }

   /**
   * Get attributes
   * @return attributes
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_ATTRIBUTES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<ReflectionKeyValue> getAttributes() {
    return attributes;
  }


  @JsonProperty(JSON_PROPERTY_ATTRIBUTES)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setAttributes(List<ReflectionKeyValue> attributes) {
    this.attributes = attributes;
  }


  public ReflectionObject documentation(List<String> documentation) {
    
    this.documentation = documentation;
    return this;
  }

  public ReflectionObject addDocumentationItem(String documentationItem) {
    if (this.documentation == null) {
      this.documentation = new ArrayList<>();
    }
    this.documentation.add(documentationItem);
    return this;
  }

   /**
   * Get documentation
   * @return documentation
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DOCUMENTATION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<String> getDocumentation() {
    return documentation;
  }


  @JsonProperty(JSON_PROPERTY_DOCUMENTATION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDocumentation(List<String> documentation) {
    this.documentation = documentation;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ReflectionObject reflectionObject = (ReflectionObject) o;
    return Objects.equals(this.name, reflectionObject.name) &&
        Objects.equals(this.fields, reflectionObject.fields) &&
        Objects.equals(this.isStruct, reflectionObject.isStruct) &&
        Objects.equals(this.minalign, reflectionObject.minalign) &&
        Objects.equals(this.bytesize, reflectionObject.bytesize) &&
        Objects.equals(this.attributes, reflectionObject.attributes) &&
        Objects.equals(this.documentation, reflectionObject.documentation);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, fields, isStruct, minalign, bytesize, attributes, documentation);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ReflectionObject {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    fields: ").append(toIndentedString(fields)).append("\n");
    sb.append("    isStruct: ").append(toIndentedString(isStruct)).append("\n");
    sb.append("    minalign: ").append(toIndentedString(minalign)).append("\n");
    sb.append("    bytesize: ").append(toIndentedString(bytesize)).append("\n");
    sb.append("    attributes: ").append(toIndentedString(attributes)).append("\n");
    sb.append("    documentation: ").append(toIndentedString(documentation)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

  /**
   * Convert the instance into URL query string.
   *
   * @return URL query string
   */
  public String toUrlQueryString() {
    return toUrlQueryString(null);
  }

  /**
   * Convert the instance into URL query string.
   *
   * @param prefix prefix of the query string
   * @return URL query string
   */
  public String toUrlQueryString(String prefix) {
    String suffix = "";
    String containerSuffix = "";
    String containerPrefix = "";
    if (prefix == null) {
      // style=form, explode=true, e.g. /pet?name=cat&type=manx
      prefix = "";
    } else {
      // deepObject style e.g. /pet?id[name]=cat&id[type]=manx
      prefix = prefix + "[";
      suffix = "]";
      containerSuffix = "]";
      containerPrefix = "[";
    }

    StringJoiner joiner = new StringJoiner("&");

    // add `name` to the URL query string
    if (getName() != null) {
      try {
        joiner.add(String.format("%sname%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getName()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `fields` to the URL query string
    if (getFields() != null) {
      for (int i = 0; i < getFields().size(); i++) {
        if (getFields().get(i) != null) {
          joiner.add(getFields().get(i).toUrlQueryString(String.format("%sfields%s%s", prefix, suffix,
              "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, i, containerSuffix))));
        }
      }
    }

    // add `is_struct` to the URL query string
    if (getIsStruct() != null) {
      try {
        joiner.add(String.format("%sis_struct%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getIsStruct()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `minalign` to the URL query string
    if (getMinalign() != null) {
      try {
        joiner.add(String.format("%sminalign%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getMinalign()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `bytesize` to the URL query string
    if (getBytesize() != null) {
      try {
        joiner.add(String.format("%sbytesize%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getBytesize()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `attributes` to the URL query string
    if (getAttributes() != null) {
      for (int i = 0; i < getAttributes().size(); i++) {
        if (getAttributes().get(i) != null) {
          joiner.add(getAttributes().get(i).toUrlQueryString(String.format("%sattributes%s%s", prefix, suffix,
              "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, i, containerSuffix))));
        }
      }
    }

    // add `documentation` to the URL query string
    if (getDocumentation() != null) {
      for (int i = 0; i < getDocumentation().size(); i++) {
        try {
          joiner.add(String.format("%sdocumentation%s%s=%s", prefix, suffix,
              "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, i, containerSuffix),
              URLEncoder.encode(String.valueOf(getDocumentation().get(i)), "UTF-8").replaceAll("\\+", "%20")));
        } catch (UnsupportedEncodingException e) {
          // Should never happen, UTF-8 is always supported
          throw new RuntimeException(e);
        }
      }
    }

    return joiner.toString();
  }

}

