// Code generated by goctl. DO NOT EDIT.

package umsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	umsMemberLevelFieldNames          = builder.RawFieldNames(&UmsMemberLevel{})
	umsMemberLevelRows                = strings.Join(umsMemberLevelFieldNames, ",")
	umsMemberLevelRowsExpectAutoSet   = strings.Join(stringx.Remove(umsMemberLevelFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	umsMemberLevelRowsWithPlaceHolder = strings.Join(stringx.Remove(umsMemberLevelFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	umsMemberLevelModel interface {
		Insert(ctx context.Context, data *UmsMemberLevel) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsMemberLevel, error)
		Update(ctx context.Context, data *UmsMemberLevel) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsMemberLevelModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsMemberLevel struct {
		Id                    int64          `db:"id"`
		Name                  string         `db:"name"`
		GrowthPoint           int64          `db:"growth_point"`
		DefaultStatus         int64          `db:"default_status"`          // 是否为默认等级：0->不是；1->是
		FreeFreightPoint      float64        `db:"free_freight_point"`      // 免运费标准
		CommentGrowthPoint    int64          `db:"comment_growth_point"`    // 每次评价获取的成长值
		PriviledgeFreeFreight int64          `db:"priviledge_free_freight"` // 是否有免邮特权
		PriviledgeSignIn      int64          `db:"priviledge_sign_in"`      // 是否有签到特权
		PriviledgeComment     int64          `db:"priviledge_comment"`      // 是否有评论获奖励特权
		PriviledgePromotion   int64          `db:"priviledge_promotion"`    // 是否有专享活动特权
		PriviledgeMemberPrice int64          `db:"priviledge_member_price"` // 是否有会员价格特权
		PriviledgeBirthday    int64          `db:"priviledge_birthday"`     // 是否有生日特权
		Note                  sql.NullString `db:"note"`
	}
)

func newUmsMemberLevelModel(conn sqlx.SqlConn) *defaultUmsMemberLevelModel {
	return &defaultUmsMemberLevelModel{
		conn:  conn,
		table: "`ums_member_level`",
	}
}

func (m *defaultUmsMemberLevelModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUmsMemberLevelModel) FindOne(ctx context.Context, id int64) (*UmsMemberLevel, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsMemberLevelRows, m.table)
	var resp UmsMemberLevel
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUmsMemberLevelModel) Insert(ctx context.Context, data *UmsMemberLevel) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, umsMemberLevelRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.GrowthPoint, data.DefaultStatus, data.FreeFreightPoint, data.CommentGrowthPoint, data.PriviledgeFreeFreight, data.PriviledgeSignIn, data.PriviledgeComment, data.PriviledgePromotion, data.PriviledgeMemberPrice, data.PriviledgeBirthday, data.Note)
	return ret, err
}

func (m *defaultUmsMemberLevelModel) Update(ctx context.Context, data *UmsMemberLevel) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsMemberLevelRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.GrowthPoint, data.DefaultStatus, data.FreeFreightPoint, data.CommentGrowthPoint, data.PriviledgeFreeFreight, data.PriviledgeSignIn, data.PriviledgeComment, data.PriviledgePromotion, data.PriviledgeMemberPrice, data.PriviledgeBirthday, data.Note, data.Id)
	return err
}

func (m *defaultUmsMemberLevelModel) tableName() string {
	return m.table
}
