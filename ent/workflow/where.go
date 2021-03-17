// Code generated by entc, DO NOT EDIT.

package workflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// Workflow applies equality check predicate on the "workflow" field. It's identical to WorkflowEQ.
func Workflow(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflow), v))
	})
}

// LogToEvents applies equality check predicate on the "logToEvents" field. It's identical to LogToEventsEQ.
func LogToEvents(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogToEvents), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevision), v))
	})
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRevision), v...))
	})
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRevision), v...))
	})
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevision), v))
	})
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevision), v))
	})
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevision), v))
	})
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevision), v))
	})
}

// WorkflowEQ applies the EQ predicate on the "workflow" field.
func WorkflowEQ(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkflow), v))
	})
}

// WorkflowNEQ applies the NEQ predicate on the "workflow" field.
func WorkflowNEQ(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkflow), v))
	})
}

// WorkflowIn applies the In predicate on the "workflow" field.
func WorkflowIn(vs ...[]byte) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkflow), v...))
	})
}

// WorkflowNotIn applies the NotIn predicate on the "workflow" field.
func WorkflowNotIn(vs ...[]byte) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkflow), v...))
	})
}

// WorkflowGT applies the GT predicate on the "workflow" field.
func WorkflowGT(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkflow), v))
	})
}

// WorkflowGTE applies the GTE predicate on the "workflow" field.
func WorkflowGTE(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkflow), v))
	})
}

// WorkflowLT applies the LT predicate on the "workflow" field.
func WorkflowLT(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkflow), v))
	})
}

// WorkflowLTE applies the LTE predicate on the "workflow" field.
func WorkflowLTE(v []byte) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkflow), v))
	})
}

// LogToEventsEQ applies the EQ predicate on the "logToEvents" field.
func LogToEventsEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsNEQ applies the NEQ predicate on the "logToEvents" field.
func LogToEventsNEQ(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsIn applies the In predicate on the "logToEvents" field.
func LogToEventsIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogToEvents), v...))
	})
}

// LogToEventsNotIn applies the NotIn predicate on the "logToEvents" field.
func LogToEventsNotIn(vs ...string) predicate.Workflow {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Workflow(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogToEvents), v...))
	})
}

// LogToEventsGT applies the GT predicate on the "logToEvents" field.
func LogToEventsGT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsGTE applies the GTE predicate on the "logToEvents" field.
func LogToEventsGTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsLT applies the LT predicate on the "logToEvents" field.
func LogToEventsLT(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsLTE applies the LTE predicate on the "logToEvents" field.
func LogToEventsLTE(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsContains applies the Contains predicate on the "logToEvents" field.
func LogToEventsContains(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsHasPrefix applies the HasPrefix predicate on the "logToEvents" field.
func LogToEventsHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsHasSuffix applies the HasSuffix predicate on the "logToEvents" field.
func LogToEventsHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsIsNil applies the IsNil predicate on the "logToEvents" field.
func LogToEventsIsNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLogToEvents)))
	})
}

// LogToEventsNotNil applies the NotNil predicate on the "logToEvents" field.
func LogToEventsNotNil() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLogToEvents)))
	})
}

// LogToEventsEqualFold applies the EqualFold predicate on the "logToEvents" field.
func LogToEventsEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogToEvents), v))
	})
}

// LogToEventsContainsFold applies the ContainsFold predicate on the "logToEvents" field.
func LogToEventsContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogToEvents), v))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstances applies the HasEdge predicate on the "instances" edge.
func HasInstances() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstancesWith applies the HasEdge predicate on the "instances" edge with a given conditions (other predicates).
func HasInstancesWith(preds ...predicate.WorkflowInstance) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstancesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InstancesTable, InstancesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWfevents applies the HasEdge predicate on the "wfevents" edge.
func HasWfevents() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventsTable, WfeventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWfeventsWith applies the HasEdge predicate on the "wfevents" edge with a given conditions (other predicates).
func HasWfeventsWith(preds ...predicate.WorkflowEvents) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WfeventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WfeventsTable, WfeventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		p(s.Not())
	})
}
