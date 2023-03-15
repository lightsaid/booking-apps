package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// Store 定义CRUD接口，结合 sqlc 生成的和自己实现的
type Store interface {
	Querier

	// TEST:
	TestRoleTx(ctx context.Context) (int64, error)

	BatchInsertSeats(context.Context, []*CreateSeatParams) error
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

// TestRoleTx 测试自定义事物
func (store *SQLStore) TestRoleTx(ctx context.Context) (int64, error) {
	var id int64
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		role, err := q.CreateRole(ctx, CreateRoleParams{Name: "See", Code: "SEE"})
		if err != nil {
			return err
		}
		id = role.ID
		desc := "查看后台管理系统权限"
		_, err = q.UpdateRole(ctx, UpdateRoleParams{ID: role.ID, Name: role.Name, Description: &desc})
		return err
	})

	return id, err
}

func (store *SQLStore) BatchInsertSeats(ctx context.Context, rows []*CreateSeatParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		stmt, err := q.db.PrepareContext(ctx, pq.CopyIn("tb_seats", "hall_id", "row_number", "col_number", "status"))
		if err != nil {
			return err
		}

		for _, row := range rows {
			_, err := stmt.Exec(row.HallID, row.RowNumber, row.ColNumber, row.Status)
			if err != nil {
				return err
			}
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}

		err = stmt.Close()
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
