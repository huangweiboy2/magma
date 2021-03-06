// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/serviceendpoint"
)

// ServiceEndpoint is the model entity for the ServiceEndpoint schema.
type ServiceEndpoint struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServiceEndpoint) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullTime{},
		&sql.NullTime{},
		&sql.NullString{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServiceEndpoint fields.
func (se *ServiceEndpoint) assignValues(values ...interface{}) error {
	if m, n := len(values), len(serviceendpoint.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	se.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		se.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		se.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field role", values[2])
	} else if value.Valid {
		se.Role = value.String
	}
	return nil
}

// QueryPort queries the port edge of the ServiceEndpoint.
func (se *ServiceEndpoint) QueryPort() *EquipmentPortQuery {
	return (&ServiceEndpointClient{se.config}).QueryPort(se)
}

// QueryService queries the service edge of the ServiceEndpoint.
func (se *ServiceEndpoint) QueryService() *ServiceQuery {
	return (&ServiceEndpointClient{se.config}).QueryService(se)
}

// Update returns a builder for updating this ServiceEndpoint.
// Note that, you need to call ServiceEndpoint.Unwrap() before calling this method, if this ServiceEndpoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (se *ServiceEndpoint) Update() *ServiceEndpointUpdateOne {
	return (&ServiceEndpointClient{se.config}).UpdateOne(se)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (se *ServiceEndpoint) Unwrap() *ServiceEndpoint {
	tx, ok := se.config.driver.(*txDriver)
	if !ok {
		panic("ent: ServiceEndpoint is not a transactional entity")
	}
	se.config.driver = tx.drv
	return se
}

// String implements the fmt.Stringer.
func (se *ServiceEndpoint) String() string {
	var builder strings.Builder
	builder.WriteString("ServiceEndpoint(")
	builder.WriteString(fmt.Sprintf("id=%v", se.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(se.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(se.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", role=")
	builder.WriteString(se.Role)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (se *ServiceEndpoint) id() int {
	id, _ := strconv.Atoi(se.ID)
	return id
}

// ServiceEndpoints is a parsable slice of ServiceEndpoint.
type ServiceEndpoints []*ServiceEndpoint

func (se ServiceEndpoints) config(cfg config) {
	for _i := range se {
		se[_i].config = cfg
	}
}
