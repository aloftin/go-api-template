package models

import (
	"database/sql"
	"encoding/json"
)

// PagingParameters stores parameters used to page query results
type PagingParameters struct {
	Limit   int
	Offset  int
	Sort    string
	SortDir string
}

/***** DB Types *****/

// NullString is an alias for sql.NullString data type
// swagger:strfmt string
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(data []byte) error {
	value := string(data)
	ns.Valid = value != "null"
	ns.String = value
	return json.Unmarshal(data, &ns.String)
}

// NullInt is an alias for sql.NullInt32 data type
type NullInt struct {
	sql.NullInt32
}

// MarshalJSON for NullInt
func (ns NullInt) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Int32)
}

// UnmarshalJSON for NullInt
func (ns *NullInt) UnmarshalJSON(data []byte) error {
	ns.Valid = string(data) != "null"
	return json.Unmarshal(data, &ns.Int32)
}

// NullTime is an alias for sql.NullTime data type
// swagger:strfmt date-time
type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullTime
func (ns NullTime) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Time)
}

//UnmarshalJSON for NullTime
func (ns *NullTime) UnmarshalJSON(data []byte) error {
	ns.Valid = string(data) != "null"
	return json.Unmarshal(data, &ns.Time)
}

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (ns NullBool) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Bool)
}

// UnmarshalJSON for NullBool
func (ns *NullBool) UnmarshalJSON(data []byte) error {
	ns.Valid = string(data) != "null"
	return json.Unmarshal(data, &ns.Bool)
}
