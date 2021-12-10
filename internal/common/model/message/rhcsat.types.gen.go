// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package message

import "fmt"
import "reflect"
import "encoding/json"

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookSatRunResponseMessageYamlEventsElemType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PlaybookSatRunResponseMessageYamlEventsElemType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PlaybookSatRunResponseMessageYamlEventsElemType, v)
	}
	*j = PlaybookSatRunResponseMessageYamlEventsElemType(v)
	return nil
}

type PlaybookSatRunResponseMessageYaml struct {
	// Account corresponds to the JSON schema field "account".
	Account string `json:"account"`

	// B64Identity corresponds to the JSON schema field "b64_identity".
	B64Identity string `json:"b64_identity"`

	// Events corresponds to the JSON schema field "events".
	Events []PlaybookSatRunResponseMessageYamlEventsElem `json:"events"`

	// RequestId corresponds to the JSON schema field "request_id".
	RequestId string `json:"request_id"`

	// UploadTimestamp corresponds to the JSON schema field "upload_timestamp".
	UploadTimestamp string `json:"upload_timestamp"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookSatRunResponseMessageYamlEventsElemStatus) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PlaybookSatRunResponseMessageYamlEventsElemStatus {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PlaybookSatRunResponseMessageYamlEventsElemStatus, v)
	}
	*j = PlaybookSatRunResponseMessageYamlEventsElemStatus(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookSatRunResponseMessageYamlEventsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["correlation_id"]; !ok || v == nil {
		return fmt.Errorf("field correlation_id: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version: required")
	}
	type Plain PlaybookSatRunResponseMessageYamlEventsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PlaybookSatRunResponseMessageYamlEventsElem(plain)
	return nil
}

type PlaybookSatRunResponseMessageYamlEventsElem struct {
	// ConnectionCode corresponds to the JSON schema field "connection_code".
	ConnectionCode *int `json:"connection_code,omitempty"`

	// Console corresponds to the JSON schema field "console".
	Console *string `json:"console,omitempty"`

	// CorrelationId corresponds to the JSON schema field "correlation_id".
	CorrelationId string `json:"correlation_id"`

	// ExecutionCode corresponds to the JSON schema field "execution_code".
	ExecutionCode *int `json:"execution_code,omitempty"`

	// Host corresponds to the JSON schema field "host".
	Host *string `json:"host,omitempty"`

	// SatelliteConnectionCode corresponds to the JSON schema field
	// "satellite_connection_code".
	SatelliteConnectionCode *int `json:"satellite_connection_code,omitempty"`

	// SatelliteConnectionError corresponds to the JSON schema field
	// "satellite_connection_error".
	SatelliteConnectionError *string `json:"satellite_connection_error,omitempty"`

	// SatelliteInfrastructureCode corresponds to the JSON schema field
	// "satellite_infrastructure_code".
	SatelliteInfrastructureCode *int `json:"satellite_infrastructure_code,omitempty"`

	// SatelliteInfrastructureError corresponds to the JSON schema field
	// "satellite_infrastructure_error".
	SatelliteInfrastructureError *string `json:"satellite_infrastructure_error,omitempty"`

	// Sequence corresponds to the JSON schema field "sequence".
	Sequence *int `json:"sequence,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status *PlaybookSatRunResponseMessageYamlEventsElemStatus `json:"status,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type PlaybookSatRunResponseMessageYamlEventsElemType `json:"type"`

	// Version corresponds to the JSON schema field "version".
	Version int `json:"version"`
}

type PlaybookSatRunResponseMessageYamlEventsElemStatus string

const PlaybookSatRunResponseMessageYamlEventsElemStatusCanceled PlaybookSatRunResponseMessageYamlEventsElemStatus = "canceled"
const PlaybookSatRunResponseMessageYamlEventsElemStatusFailure PlaybookSatRunResponseMessageYamlEventsElemStatus = "failure"
const PlaybookSatRunResponseMessageYamlEventsElemStatusSuccess PlaybookSatRunResponseMessageYamlEventsElemStatus = "success"

type PlaybookSatRunResponseMessageYamlEventsElemType string

const PlaybookSatRunResponseMessageYamlEventsElemTypePlaybookRunCompleted PlaybookSatRunResponseMessageYamlEventsElemType = "playbook_run_completed"
const PlaybookSatRunResponseMessageYamlEventsElemTypePlaybookRunFinished PlaybookSatRunResponseMessageYamlEventsElemType = "playbook_run_finished"
const PlaybookSatRunResponseMessageYamlEventsElemTypePlaybookRunUpdate PlaybookSatRunResponseMessageYamlEventsElemType = "playbook_run_update"

var enumValues_PlaybookSatRunResponseMessageYamlEventsElemStatus = []interface{}{
	"success",
	"failure",
	"canceled",
}
var enumValues_PlaybookSatRunResponseMessageYamlEventsElemType = []interface{}{
	"playbook_run_update",
	"playbook_run_finished",
	"playbook_run_completed",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PlaybookSatRunResponseMessageYaml) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["account"]; !ok || v == nil {
		return fmt.Errorf("field account: required")
	}
	if v, ok := raw["b64_identity"]; !ok || v == nil {
		return fmt.Errorf("field b64_identity: required")
	}
	if v, ok := raw["events"]; !ok || v == nil {
		return fmt.Errorf("field events: required")
	}
	if v, ok := raw["request_id"]; !ok || v == nil {
		return fmt.Errorf("field request_id: required")
	}
	if v, ok := raw["upload_timestamp"]; !ok || v == nil {
		return fmt.Errorf("field upload_timestamp: required")
	}
	type Plain PlaybookSatRunResponseMessageYaml
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PlaybookSatRunResponseMessageYaml(plain)
	return nil
}
