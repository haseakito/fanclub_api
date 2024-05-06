// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hackgame-org/fanclub_api/api/ent/user"
	"github.com/hackgame-org/fanclub_api/api/ent/verificationtoken"
)

// VerificationTokenCreate is the builder for creating a VerificationToken entity.
type VerificationTokenCreate struct {
	config
	mutation *VerificationTokenMutation
	hooks    []Hook
}

// SetEmail sets the "email" field.
func (vtc *VerificationTokenCreate) SetEmail(s string) *VerificationTokenCreate {
	vtc.mutation.SetEmail(s)
	return vtc
}

// SetVerificationCode sets the "verification_code" field.
func (vtc *VerificationTokenCreate) SetVerificationCode(s string) *VerificationTokenCreate {
	vtc.mutation.SetVerificationCode(s)
	return vtc
}

// SetExpiresAt sets the "expires_at" field.
func (vtc *VerificationTokenCreate) SetExpiresAt(t time.Time) *VerificationTokenCreate {
	vtc.mutation.SetExpiresAt(t)
	return vtc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vtc *VerificationTokenCreate) SetUserID(id string) *VerificationTokenCreate {
	vtc.mutation.SetUserID(id)
	return vtc
}

// SetUser sets the "user" edge to the User entity.
func (vtc *VerificationTokenCreate) SetUser(u *User) *VerificationTokenCreate {
	return vtc.SetUserID(u.ID)
}

// Mutation returns the VerificationTokenMutation object of the builder.
func (vtc *VerificationTokenCreate) Mutation() *VerificationTokenMutation {
	return vtc.mutation
}

// Save creates the VerificationToken in the database.
func (vtc *VerificationTokenCreate) Save(ctx context.Context) (*VerificationToken, error) {
	return withHooks(ctx, vtc.sqlSave, vtc.mutation, vtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vtc *VerificationTokenCreate) SaveX(ctx context.Context) *VerificationToken {
	v, err := vtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtc *VerificationTokenCreate) Exec(ctx context.Context) error {
	_, err := vtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtc *VerificationTokenCreate) ExecX(ctx context.Context) {
	if err := vtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vtc *VerificationTokenCreate) check() error {
	if _, ok := vtc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "VerificationToken.email"`)}
	}
	if _, ok := vtc.mutation.VerificationCode(); !ok {
		return &ValidationError{Name: "verification_code", err: errors.New(`ent: missing required field "VerificationToken.verification_code"`)}
	}
	if v, ok := vtc.mutation.VerificationCode(); ok {
		if err := verificationtoken.VerificationCodeValidator(v); err != nil {
			return &ValidationError{Name: "verification_code", err: fmt.Errorf(`ent: validator failed for field "VerificationToken.verification_code": %w`, err)}
		}
	}
	if _, ok := vtc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "VerificationToken.expires_at"`)}
	}
	if _, ok := vtc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "VerificationToken.user"`)}
	}
	return nil
}

func (vtc *VerificationTokenCreate) sqlSave(ctx context.Context) (*VerificationToken, error) {
	if err := vtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vtc.mutation.id = &_node.ID
	vtc.mutation.done = true
	return _node, nil
}

func (vtc *VerificationTokenCreate) createSpec() (*VerificationToken, *sqlgraph.CreateSpec) {
	var (
		_node = &VerificationToken{config: vtc.config}
		_spec = sqlgraph.NewCreateSpec(verificationtoken.Table, sqlgraph.NewFieldSpec(verificationtoken.FieldID, field.TypeInt))
	)
	if value, ok := vtc.mutation.Email(); ok {
		_spec.SetField(verificationtoken.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := vtc.mutation.VerificationCode(); ok {
		_spec.SetField(verificationtoken.FieldVerificationCode, field.TypeString, value)
		_node.VerificationCode = value
	}
	if value, ok := vtc.mutation.ExpiresAt(); ok {
		_spec.SetField(verificationtoken.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if nodes := vtc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   verificationtoken.UserTable,
			Columns: []string{verificationtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_verification_token = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VerificationTokenCreateBulk is the builder for creating many VerificationToken entities in bulk.
type VerificationTokenCreateBulk struct {
	config
	err      error
	builders []*VerificationTokenCreate
}

// Save creates the VerificationToken entities in the database.
func (vtcb *VerificationTokenCreateBulk) Save(ctx context.Context) ([]*VerificationToken, error) {
	if vtcb.err != nil {
		return nil, vtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vtcb.builders))
	nodes := make([]*VerificationToken, len(vtcb.builders))
	mutators := make([]Mutator, len(vtcb.builders))
	for i := range vtcb.builders {
		func(i int, root context.Context) {
			builder := vtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VerificationTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vtcb *VerificationTokenCreateBulk) SaveX(ctx context.Context) []*VerificationToken {
	v, err := vtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vtcb *VerificationTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := vtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vtcb *VerificationTokenCreateBulk) ExecX(ctx context.Context) {
	if err := vtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
