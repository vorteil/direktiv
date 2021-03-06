// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/vorteil/direktiv/ent/namespace"
	"github.com/vorteil/direktiv/ent/schema"
	"github.com/vorteil/direktiv/ent/workflow"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	namespaceFields := schema.Namespace{}.Fields()
	_ = namespaceFields
	// namespaceDescCreated is the schema descriptor for created field.
	namespaceDescCreated := namespaceFields[1].Descriptor()
	// namespace.DefaultCreated holds the default value on creation for the created field.
	namespace.DefaultCreated = namespaceDescCreated.Default.(func() time.Time)
	// namespaceDescID is the schema descriptor for id field.
	namespaceDescID := namespaceFields[0].Descriptor()
	// namespace.IDValidator is a validator for the "id" field. It is called by the builders before save.
	namespace.IDValidator = func() func(string) error {
		validators := namespaceDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	workflowFields := schema.Workflow{}.Fields()
	_ = workflowFields
	// workflowDescName is the schema descriptor for name field.
	workflowDescName := workflowFields[1].Descriptor()
	// workflow.NameValidator is a validator for the "name" field. It is called by the builders before save.
	workflow.NameValidator = workflowDescName.Validators[0].(func(string) error)
	// workflowDescCreated is the schema descriptor for created field.
	workflowDescCreated := workflowFields[2].Descriptor()
	// workflow.DefaultCreated holds the default value on creation for the created field.
	workflow.DefaultCreated = workflowDescCreated.Default.(func() time.Time)
	// workflowDescDescription is the schema descriptor for description field.
	workflowDescDescription := workflowFields[3].Descriptor()
	// workflow.DefaultDescription holds the default value on creation for the description field.
	workflow.DefaultDescription = workflowDescDescription.Default.(string)
	// workflow.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	workflow.DescriptionValidator = workflowDescDescription.Validators[0].(func(string) error)
	// workflowDescActive is the schema descriptor for active field.
	workflowDescActive := workflowFields[4].Descriptor()
	// workflow.DefaultActive holds the default value on creation for the active field.
	workflow.DefaultActive = workflowDescActive.Default.(bool)
	// workflowDescRevision is the schema descriptor for revision field.
	workflowDescRevision := workflowFields[5].Descriptor()
	// workflow.DefaultRevision holds the default value on creation for the revision field.
	workflow.DefaultRevision = workflowDescRevision.Default.(int)
	// workflowDescID is the schema descriptor for id field.
	workflowDescID := workflowFields[0].Descriptor()
	// workflow.DefaultID holds the default value on creation for the id field.
	workflow.DefaultID = workflowDescID.Default.(func() uuid.UUID)
}
