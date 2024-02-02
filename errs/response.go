package errs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseBody struct {
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Origin string `json:"origin"`
	Data   any    `json:"data"`
}

// ...any可以当作切片来处理
func Success(c *gin.Context, data ...any) {
	response := responseBody{
		Code:   SUCCESS.Code,
		Msg:    SUCCESS.Message,
		Origin: SUCCESS.Origin,
		Data:   data,
	}
	//if len(data) > 0 {
	//	response.Data = data[0]
	//}
	c.JSON(http.StatusOK, response)
}

func Fail(c *gin.Context, err error) {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		e = SERVE_INTERNAL.WithOrigin(err)
	}

	var resp responseBody
	resp.Code = e.Code
	resp.Msg = e.Message
	resp.Origin = e.Origin

	c.JSON(int(e.Code/100), resp)
	c.Abort()
}

/*
func Recovery(c *gin.Context) {
	info := recover()
	if info != nil {
		err, ok := info.(error)
		if ok {
			Fail(c, errors.WithStack(err))
		} else {
			Fail(c, errors.New)
		}
	}
	return
}*/
