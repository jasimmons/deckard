package deckard

import (
	"encoding/json"
	"fmt"

	cron "github.com/robfig/cron/v3"
)

type Identifier struct {
	// ID is the Identifier's unique ID.
	ID int `json:"id,omitempty"`

	// Name is the human-readable name of the Identifier. Uniqueness is
	// not a requirement.
	Name string `json:"name,omitempty"`

	// Description is a long-form description of the Identifier.
	Description string `json:"description,omitempty"`

	// Tags is a list of keys that are used for targetting a Identifier for
	// execution. When a request is made to execute a Identifier or set of
	// Identifiers, the Tags field is used as a filter for determining if a
	// Identifier should be executed.
	Tags map[string]struct{} `json:"tags"`

	// Endpoint is the full URL of the remote service to call this
	// Identifier.
	Endpoint string `json:"endpoint,omitempty"`

	// Schedule is a cron representation of when this Identifier should
	// be called. For additional information on format, see:
	// https://en.wikipedia.org/wiki/Cron#Overview
	Schedule string `json:"schedule,omitempty"`
}

func (i *Identifier) hasTag(tagKey string) bool {
	_, ok := i.Tags[tagKey]
	return ok
}

func (i *Identifier) HasTags(tagKeys ...string) bool {
	for _, tagKey := range tagKeys {
		if !i.hasTag(tagKey) {
			return false
		}
	}
	return true
}

// MarshalJSON implements the json.Marshaler interface so Identifier
// values can be converted into strings.
// Specifically, we want to:
//  - Convert Endpoint values from Endpoint structs into strings.
//  - Convert Tags values from string-slices to maps.
func (i *Identifier) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = i.ID
	m["name"] = i.Name
	m["endpoint"] = i.Endpoint
	m["schedule"] = i.Schedule
	m["tags"] = make([]string, 0, len(i.Tags))
	for tag := range i.Tags {
		m["tags"] = append(m["tags"].([]string), tag)
	}
	return json.Marshal(m)
}

// UnmarshalJSON implements the json.Unmarshaler interface so Identifier
// values can be unmarshaled as strings.
// Specifically, we want to:
//  - Convert Endpoint values from strings into Endpoint structs.
//  - Convert Tags values from maps to string-slices.
func (i *Identifier) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	id := &Identifier{
		Tags: make(map[string]struct{}),
	}
	if _, ok := m["name"].(string); !ok {
		return fmt.Errorf("bad name value %v", m["name"])
	}
	id.Name = m["name"].(string)

	if _, ok := m["endpoint"].(string); !ok {
		if err != nil {
			return fmt.Errorf("bad endpoint value %v", m["endpoint"])
		}
	}
	id.Endpoint = m["endpoint"].(string)

	if _, ok := m["schedule"].(string); !ok {
		if err != nil {
			return fmt.Errorf("bad schedule value %v", m["schedule"])
		}
	}
	id.Schedule = m["schedule"].(string)

	if _, ok := m["tags"].([]interface{}); !ok {
		return fmt.Errorf("bad Tags value %v", m["tags"])
	}
	for _, tag := range m["tags"].([]interface{}) {
		if _, ok := tag.(string); !ok {
			return fmt.Errorf("bad tag value %v", tag)
		}
		id.Tags[tag.(string)] = struct{}{}
	}
	*i = *id

	return nil
}

func parseSchedule(cronStr string) error {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	_, err := p.Parse(cronStr)
	return err
}
