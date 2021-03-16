// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/ent/namespace"
	"github.com/vorteil/direktiv/ent/workflow"
	"github.com/vorteil/direktiv/ent/workflowevents"
	"github.com/vorteil/direktiv/ent/workflowinstance"
)

// WorkflowCreate is the builder for creating a Workflow entity.
type WorkflowCreate struct {
	config
	mutation *WorkflowMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WorkflowCreate) SetName(s string) *WorkflowCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetCreated sets the "created" field.
func (wc *WorkflowCreate) SetCreated(t time.Time) *WorkflowCreate {
	wc.mutation.SetCreated(t)
	return wc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableCreated(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetCreated(*t)
	}
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkflowCreate) SetDescription(s string) *WorkflowCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableDescription(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetActive sets the "active" field.
func (wc *WorkflowCreate) SetActive(b bool) *WorkflowCreate {
	wc.mutation.SetActive(b)
	return wc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableActive(b *bool) *WorkflowCreate {
	if b != nil {
		wc.SetActive(*b)
	}
	return wc
}

// SetRevision sets the "revision" field.
func (wc *WorkflowCreate) SetRevision(i int) *WorkflowCreate {
	wc.mutation.SetRevision(i)
	return wc
}

// SetNillableRevision sets the "revision" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableRevision(i *int) *WorkflowCreate {
	if i != nil {
		wc.SetRevision(*i)
	}
	return wc
}

// SetWorkflow sets the "workflow" field.
func (wc *WorkflowCreate) SetWorkflow(b []byte) *WorkflowCreate {
	wc.mutation.SetWorkflow(b)
	return wc
}

// SetLogToEvents sets the "logToEvents" field.
func (wc *WorkflowCreate) SetLogToEvents(s string) *WorkflowCreate {
	wc.mutation.SetLogToEvents(s)
	return wc
}

// SetNillableLogToEvents sets the "logToEvents" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableLogToEvents(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetLogToEvents(*s)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkflowCreate) SetID(u uuid.UUID) *WorkflowCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (wc *WorkflowCreate) SetNamespaceID(id string) *WorkflowCreate {
	wc.mutation.SetNamespaceID(id)
	return wc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (wc *WorkflowCreate) SetNamespace(n *Namespace) *WorkflowCreate {
	return wc.SetNamespaceID(n.ID)
}

// AddInstanceIDs adds the "instances" edge to the WorkflowInstance entity by IDs.
func (wc *WorkflowCreate) AddInstanceIDs(ids ...int) *WorkflowCreate {
	wc.mutation.AddInstanceIDs(ids...)
	return wc
}

// AddInstances adds the "instances" edges to the WorkflowInstance entity.
func (wc *WorkflowCreate) AddInstances(w ...*WorkflowInstance) *WorkflowCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddInstanceIDs(ids...)
}

// AddWfeventIDs adds the "wfevents" edge to the WorkflowEvents entity by IDs.
func (wc *WorkflowCreate) AddWfeventIDs(ids ...int) *WorkflowCreate {
	wc.mutation.AddWfeventIDs(ids...)
	return wc
}

// AddWfevents adds the "wfevents" edges to the WorkflowEvents entity.
func (wc *WorkflowCreate) AddWfevents(w ...*WorkflowEvents) *WorkflowCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWfeventIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wc *WorkflowCreate) Mutation() *WorkflowMutation {
	return wc.mutation
}

// Save creates the Workflow in the database.
func (wc *WorkflowCreate) Save(ctx context.Context) (*Workflow, error) {
	var (
		err  error
		node *Workflow
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			node, err = wc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowCreate) SaveX(ctx context.Context) *Workflow {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wc *WorkflowCreate) defaults() {
	if _, ok := wc.mutation.Created(); !ok {
		v := workflow.DefaultCreated()
		wc.mutation.SetCreated(v)
	}
	if _, ok := wc.mutation.Description(); !ok {
		v := workflow.DefaultDescription
		wc.mutation.SetDescription(v)
	}
	if _, ok := wc.mutation.Active(); !ok {
		v := workflow.DefaultActive
		wc.mutation.SetActive(v)
	}
	if _, ok := wc.mutation.Revision(); !ok {
		v := workflow.DefaultRevision
		wc.mutation.SetRevision(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workflow.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := workflow.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New("ent: missing required field \"created\"")}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := workflow.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf("ent: validator failed for field \"description\": %w", err)}
		}
	}
	if _, ok := wc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New("ent: missing required field \"active\"")}
	}
	if _, ok := wc.mutation.Revision(); !ok {
		return &ValidationError{Name: "revision", err: errors.New("ent: missing required field \"revision\"")}
	}
	if _, ok := wc.mutation.Workflow(); !ok {
		return &ValidationError{Name: "workflow", err: errors.New("ent: missing required field \"workflow\"")}
	}
	if _, ok := wc.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New("ent: missing required edge \"namespace\"")}
	}
	return nil
}

func (wc *WorkflowCreate) sqlSave(ctx context.Context) (*Workflow, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (wc *WorkflowCreate) createSpec() (*Workflow, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflow{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workflow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workflow.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldName,
		})
		_node.Name = value
	}
	if value, ok := wc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workflow.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := wc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workflow.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := wc.mutation.Revision(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workflow.FieldRevision,
		})
		_node.Revision = value
	}
	if value, ok := wc.mutation.Workflow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: workflow.FieldWorkflow,
		})
		_node.Workflow = value
	}
	if value, ok := wc.mutation.LogToEvents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldLogToEvents,
		})
		_node.LogToEvents = value
	}
	if nodes := wc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_workflows = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WfeventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkflowCreateBulk is the builder for creating many Workflow entities in bulk.
type WorkflowCreateBulk struct {
	config
	builders []*WorkflowCreate
}

// Save creates the Workflow entities in the database.
func (wcb *WorkflowCreateBulk) Save(ctx context.Context) ([]*Workflow, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflow, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) SaveX(ctx context.Context) []*Workflow {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
