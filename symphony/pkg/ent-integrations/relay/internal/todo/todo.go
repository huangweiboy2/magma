// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package todo

import (
	"context"

	"github.com/facebookincubator/symphony/pkg/ent-integrations/relay/internal/todo/ent"
)

func New(client *ent.Client) Config {
	return Config{
		Resolvers: &resolvers{
			client: client.Todo,
		},
	}
}

type resolvers struct{ client *ent.TodoClient }

func (r *resolvers) Todo(ctx context.Context, id int) (*ent.Todo, error) {
	return r.client.Get(ctx, id)
}

func (r *resolvers) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.TodoConnection, error) {
	return r.client.Query().Paginate(ctx, after, first, before, last)
}

func (r *resolvers) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
	return r.client.Create().
		SetText(todo.Text).
		Save(ctx)
}

func (r *resolvers) ClearTodos(ctx context.Context) (int, error) {
	return r.client.Delete().
		Exec(ctx)
}

func (r *resolvers) Query() QueryResolver       { return r }
func (r *resolvers) Mutation() MutationResolver { return r }
