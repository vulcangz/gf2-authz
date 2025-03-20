package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/user/v1"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	User = cUser{}

	// ErrCannotDeleteOwnAccount is returned when a user tries to delete their own account.
	ErrCannotDeleteOwnAccount = errors.New("a user cannot delete their own account")
)

type cUser struct{}

// Creates a new user
//
//	@security	Authentication
//	@Summary	Creates a new user
//	@Tags		User
//	@Produce	json
//	@Param		default	body		UserCreateRequest	true	"User creation request"
//	@Success	200		{object}	model.User
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/users [Post]
func (c *cUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	user, err := service.UserManager().Create(ctx, req.Username, "")
	if err != nil {
		return
	}

	res = new(v1.CreateRes)
	res.User = user

	return
}

// Lists users.
//
//	@security	Authentication
//	@Summary	Lists users
//	@Tags		User
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(name:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(name:desc)
//	@Success	200		{object}	[]model.User
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/users [Get]
func (c *cUser) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		err = gerror.NewCode(gcode.CodeInternalError)
		return
	}

	// List actions
	users, total, err := service.UserManager().GetRepository().Find(
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
	)
	if err != nil {
		err = gerror.Wrap(err, gcode.CodeInternalError.Message())
		return
	}

	res = new(v1.GetListRes)
	res.Paginated = orm.NewPaginated(users, total, page, size)

	return
}

// Retrieve a user.
//
//	@security	Authentication
//	@Summary	Retrieve a user
//	@Tags		User
//	@Produce	json
//	@Success	200	{object}	model.User
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/users/{identifier} [Get]
func (c *cUser) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve user
	user, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: identifier},
	})
	if err != nil {
		statusCode := gcode.CodeInternalError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = gcode.CodeNotFound
		}

		err = gerror.WrapCodef(statusCode, err, "cannot retrieve user: %v", err)
		return
	}

	res = new(v1.GetOneRes)
	res.User = user

	return
}

// Deletes a user.
//
//	@security	Authentication
//	@Summary	Deletes a user
//	@Tags		User
//	@Produce	json
//	@Success	200	{object}	model.User
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/users/{identifier} [Delete]
func (c *cUser) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	userID := ctx.Value(consts.UserIdentifierKey).(string)
	if userID == identifier {
		response.ReturnError(r, http.StatusBadRequest, ErrCannotDeleteOwnAccount, "")
		return
	}

	if err = service.UserManager().Delete(ctx, identifier); err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
