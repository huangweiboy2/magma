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
	"github.com/facebookincubator/symphony/graph/ent/technician"
)

// Technician is the model entity for the Technician schema.
type Technician struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Technician) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullTime{},
		&sql.NullTime{},
		&sql.NullString{},
		&sql.NullString{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Technician fields.
func (t *Technician) assignValues(values ...interface{}) error {
	if m, n := len(values), len(technician.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = strconv.FormatInt(value.Int64, 10)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		t.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		t.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[3])
	} else if value.Valid {
		t.Email = value.String
	}
	return nil
}

// QueryWorkOrders queries the work_orders edge of the Technician.
func (t *Technician) QueryWorkOrders() *WorkOrderQuery {
	return (&TechnicianClient{t.config}).QueryWorkOrders(t)
}

// Update returns a builder for updating this Technician.
// Note that, you need to call Technician.Unwrap() before calling this method, if this Technician
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Technician) Update() *TechnicianUpdateOne {
	return (&TechnicianClient{t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Technician) Unwrap() *Technician {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Technician is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Technician) String() string {
	var builder strings.Builder
	builder.WriteString("Technician(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", email=")
	builder.WriteString(t.Email)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (t *Technician) id() int {
	id, _ := strconv.Atoi(t.ID)
	return id
}

// Technicians is a parsable slice of Technician.
type Technicians []*Technician

func (t Technicians) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
