package sql

import (
	"context"
)

type txKey struct{}

func InjectTx(ctx context.Context, tx *Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) *Tx {
	if tx, ok := ctx.Value(txKey{}).(*Tx); ok {
		return tx
	}
	return nil
}
