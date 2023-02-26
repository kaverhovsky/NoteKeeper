package httptool

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// ReturnOK returns wrapped payload
func ReturnOK(ctx *fasthttp.RequestCtx, status int, payload interface{}) {
	ctx.SetContentType(CONTENTTYPE_JSON)
	ctx.Response.SetStatusCode(status)
	encoder := getEncoderForCtx(ctx)
	if err := encoder.Encode(payload); err != nil {
		ReturnError(ctx, fasthttp.StatusInternalServerError, err)
	}
}

// ReturnError returns wrapped error
func ReturnError(ctx *fasthttp.RequestCtx, status int, err error) {
	ctx.SetContentType(CONTENTTYPE_JSON)
	ctx.Response.SetStatusCode(status)
	encoder := getEncoderForCtx(ctx)
	if err := encoder.Encode(err); err != nil {
		panic(err)
	}
}

func getEncoderForCtx(ctx *fasthttp.RequestCtx) *json.Encoder {
	encoder := json.NewEncoder(ctx)
	if ctx.QueryArgs().GetBool("pretty") {
		encoder.SetIndent("", "	")
	}
	return encoder
}

// func PutLoggerToRequestContext adds to logger Operator and TraceId values from Request
func PutLoggerToRequestContext(logger *zap.Logger, ctx *fasthttp.RequestCtx) *zap.Logger {
	operator := ctx.UserValue(HEADER_OPERATOR)
	if operator != nil {
		logger = logger.With(zap.String("operator", operator.(string)))
	}
	traceID := ctx.UserValue(HEADER_TRACEID)
	if traceID != nil {
		logger = logger.With(zap.String("traceId", traceID.(string)))
	}
	return logger
}
