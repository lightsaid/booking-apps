package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store 定义CRUD接口，结合 sqlc 生成的和自己实现的
type Store interface {
	Querier
}

// SQLStore SQLStore 结构体是对 sqlc 生成代码进行封装和组合，实现自定义操作DB句柄、执行事务
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore 创建一个SQLStore实例
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx 执行事务公共方法
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	t := time.Now()
	// 开启事务
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 使用事务的 tx 获取一个新的 *Queries, 因为 Queries 的实例 实现了Querier接口，
	// 接而可以使用Queries提供的基础CRUD操作
	qs := New(tx)

	err = fn(qs)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	fmt.Println("执行事务花费时间： ", time.Since(t))
	return tx.Commit()
}
