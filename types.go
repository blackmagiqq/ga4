package ga4

import (
	"net/http"
)

type GA4Client struct {
	apiSecret     string       // Measurement Protocol API secret value
	measurementID string       // MEASUREMENT ID, G-XXXXXXXXXX
	httpClient    *http.Client // http client session
	debug         bool         // send events for validation, used for debug
}

type ClientID string
type Event struct {
	// Required. The name for the event.
	Name string `json:"event_name"`
	// Optional. A unique identifier for a user
	// Optional. The parameters for the event.
	// engagement_time_msec/session_id
	Params map[string]interface{} `json:"params,omitempty"`
}

// payload docs reference:
// https://developers.google.com/analytics/devguides/collection/protocol/ga4/reference?client_type=gtag
type Payload struct {
	// Required. Uniquely identifies a user instance of a web client
	ClientID string `json:"client_id"`
	// Optional. A unique identifier for a user
	UserID string `json:"user_id,omitempty"`
	// Optional. A Unix timestamp (in microseconds) for the time to associate with the event.
	// This should only be set to record events that happened in the past.
	// This value can be overridden via user_property or event timestamps.
	// Events can be backdated up to 3 calendar days based on the property's timezone.
	TimestampMicros int64 `json:"timestamp_micros,omitempty"`
	// Optional. The user properties for the measurement.
	UserProperties map[string]string `json:"user_properties,omitempty"`
	// Optional. Set to true to indicate these events should not be used for personalized ads.
	NonPersonalizedAds bool `json:"non_personalized_ads,omitempty"`
	// Required. An array of event items. Up to 25 events can be sent per request.
	Events []Event `json:"events"`
}

// validation docs reference:
// https://developers.google.com/analytics/devguides/collection/protocol/ga4/validating-events?client_type=gtag
type ValidationResponse struct {
	ValidationMessages []ValidationMessage `json:"validationMessages"` // An array of validation messages.
}

type ValidationMessage struct {
	FieldPath      string         `json:"fieldPath"`      // The path to the field that was invalid.
	Description    string         `json:"description"`    // A description of the error.
	ValidationCode ValidationCode `json:"validationCode"` // A ValidationCode that corresponds to the error.
}

type ValidationCode string

const (
	VALUE_INVALID         ValidationCode = "VALUE_INVALID"         // The value provided for a fieldPath was invalid.
	VALUE_REQUIRED        ValidationCode = "VALUE_REQUIRED"        // A required value for a fieldPath was not provided.
	NAME_INVALID          ValidationCode = "NAME_INVALID"          // The name provided was invalid.
	NAME_RESERVED         ValidationCode = "NAME_RESERVED"         // The name provided was one of the reserved names.
	VALUE_OUT_OF_BOUNDS   ValidationCode = "VALUE_OUT_OF_BOUNDS"   // The value provided was too large.
	EXCEEDED_MAX_ENTITIES ValidationCode = "EXCEEDED_MAX_ENTITIES" // There were too many parameters in the request.
	NAME_DUPLICATED       ValidationCode = "NAME_DUPLICATED"       // The same name was provided more than once in the request.
)
