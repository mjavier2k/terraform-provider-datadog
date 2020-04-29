/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// WidgetVizType Whether to display the Alert Graph as a timeseries or a top list.
type WidgetVizType string

// List of WidgetVizType
const (
	WIDGETVIZTYPE_TIMESERIES WidgetVizType = "timeseries"
	WIDGETVIZTYPE_TOPLIST    WidgetVizType = "toplist"
)

// Ptr returns reference to WidgetVizType value
func (v WidgetVizType) Ptr() *WidgetVizType {
	return &v
}

type NullableWidgetVizType struct {
	value *WidgetVizType
	isSet bool
}

func (v NullableWidgetVizType) Get() *WidgetVizType {
	return v.value
}

func (v *NullableWidgetVizType) Set(val *WidgetVizType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetVizType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetVizType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetVizType(val *WidgetVizType) *NullableWidgetVizType {
	return &NullableWidgetVizType{value: val, isSet: true}
}

func (v NullableWidgetVizType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetVizType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
