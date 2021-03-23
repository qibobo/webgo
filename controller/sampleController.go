package controller

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"

	"github.com/qibobo/webgo/models"
)

type SampleControler struct{}

func (s *SampleControler) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/demo/{id:int}", "GetDemoById")
	a.Handle("POST", "/demo", "CreateDemo")
}

func (s *SampleControler) GetDemoById(id int) mvc.Response {
	return mvc.Response{

		Code:        iris.StatusOK,
		ContentType: context.ContentJSONHeaderValue,
		Content:     []byte(fmt.Sprintf(`{"id":%d,"name":"测试1"}`, id)),
	}
}
func (s *SampleControler) CreateDemo(d models.Demo) mvc.Response {
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
