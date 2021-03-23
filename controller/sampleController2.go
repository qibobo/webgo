package controller

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"

	"github.com/qibobo/webgo/models"
)

type SampleControler2 struct{}

func (s *SampleControler2) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/{id:int}", "GetDemoById2")
	a.Handle("POST", "/create", "CreateDemo2")
}

func (s *SampleControler2) GetDemoById2(id int) mvc.Response {
	return mvc.Response{

		Code:        iris.StatusOK,
		ContentType: context.ContentJSONHeaderValue,
		Content:     []byte(fmt.Sprintf(`{"id":%d,"name":"测试2"}`, id)),
	}

}
func (s *SampleControler2) CreateDemo2(d models.Demo) mvc.Response {
	demoBytes, err := json.Marshal(d)
	if err != nil {
		return mvc.Response{

			Code:        iris.StatusInternalServerError,
			ContentType: context.ContentJSONHeaderValue,
			Content:     []byte(""),
		}
	}
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: "application/json",
		Content:     demoBytes,
	}
}
