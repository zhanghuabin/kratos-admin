package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/position"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"
)

type PositionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPositionRepo(data *Data, logger log.Logger) *PositionRepo {
	l := log.NewHelper(log.With(logger, "module", "position/repo/admin-service"))
	return &PositionRepo{
		data: data,
		log:  l,
	}
}

func (r *PositionRepo) convertEntToProto(in *ent.Position) *userV1.Position {
	if in == nil {
		return nil
	}
	return &userV1.Position{
		Id:         trans.Ptr(in.ID),
		Name:       &in.Name,
		Code:       &in.Code,
		Remark:     in.Remark,
		OrderNo:    &in.OrderNo,
		ParentId:   &in.ParentID,
		Status:     (*string)(in.Status),
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *PositionRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Position.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *PositionRepo) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListPositionResponse, error) {
	builder := r.data.db.Client().Position.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), position.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*userV1.Position, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &userV1.ListPositionResponse{
		Total: int32(count),
		Items: items,
	}, err
}

func (r *PositionRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	return r.data.db.Client().Position.Query().
		Where(position.IDEQ(id)).
		Exist(ctx)
}

func (r *PositionRepo) Get(ctx context.Context, req *userV1.GetPositionRequest) (*userV1.Position, error) {
	ret, err := r.data.db.Client().Position.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *PositionRepo) Create(ctx context.Context, req *userV1.CreatePositionRequest) error {
	if req.Position == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Position.Create().
		SetNillableName(req.Position.Name).
		SetNillableParentID(req.Position.ParentId).
		SetNillableOrderNo(req.Position.OrderNo).
		SetNillableCode(req.Position.Code).
		SetNillableStatus((*position.Status)(req.Position.Status)).
		SetNillableRemark(req.Position.Remark).
		SetCreateBy(req.GetOperatorId())

	if req.Position.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	} else {
		builder.SetCreateTime(*timeutil.TimestamppbToTime(req.Position.CreateTime))
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *PositionRepo) Update(ctx context.Context, req *userV1.UpdatePositionRequest) error {
	if req.Position == nil {
		return errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetPosition().GetId())
		if err != nil {
			return err
		}
		if !exist {
			return r.Create(ctx, &userV1.CreatePositionRequest{Position: req.Position, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Position) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetPosition(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().Position.UpdateOneID(req.Position.GetId()).
		SetNillableName(req.Position.Name).
		SetNillableParentID(req.Position.ParentId).
		SetNillableOrderNo(req.Position.OrderNo).
		SetNillableCode(req.Position.Code).
		SetNillableRemark(req.Position.Remark).
		SetNillableStatus((*position.Status)(req.Position.Status))

	if req.Position.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	} else {
		builder.SetUpdateTime(*timeutil.TimestamppbToTime(req.Position.UpdateTime))
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Position, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return err
}

func (r *PositionRepo) Delete(ctx context.Context, req *userV1.DeletePositionRequest) (bool, error) {
	err := r.data.db.Client().Position.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	return err != nil, err
}