package v1

import (
	"NoteKeeper/internal/domain/dto"
	"NoteKeeper/internal/usecase"
	"NoteKeeper/pkg/httptool"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
)

type ApiV1 struct {
	logger  *zap.Logger
	usecase *usecase.NoteUsecase
}

func NewApiV1(logger *zap.Logger, usecase *usecase.NoteUsecase) *ApiV1 {
	return &ApiV1{
		logger:  logger,
		usecase: usecase,
	}
}

func (api *ApiV1) CreateHandler(ctx *fasthttp.RequestCtx) {
	httptool.PutLoggerToRequestContext(api.logger, ctx)

	body := ctx.PostBody()

	var noteDTO dto.NoteDTO
	if err := json.Unmarshal(body, &noteDTO); err != nil {
		api.logger.With(zap.NamedError("reason", err)).Error("failed to unmarshall request body")
		// TODO вернуть обернутую ошибку
		httptool.ReturnError(ctx, fasthttp.StatusBadRequest, err)
	}

	// TODO добавить валидацию

	// call usecase method
	note, err := api.usecase.Create(noteDTO)
	if err != nil {
		api.logger.With(zap.NamedError("reason", err)).Error("failed create note")
		// TODO вернуть обернутую ошибку (конвертированную из ошибки хранилища)
		httptool.ReturnError(ctx, fasthttp.StatusInternalServerError, err)
	}

	httptool.ReturnOK(ctx, http.StatusCreated, note)
}

func (api *ApiV1) AddRoutes(group *router.Group) {
	group.POST("/notes", api.CreateHandler)
}
