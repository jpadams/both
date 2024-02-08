package main

import (
	"context"
)

type Both struct {}

// Run with: dagger call test
func (m *Both) Test(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"sh", "-c", `(echo "This is stdout" ; echo "This is stderr" 1>&2)`}).
		Stdout(ctx)
		//Stderr(ctx)
}

