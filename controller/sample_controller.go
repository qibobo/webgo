package controller

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"go.uber.org/zap"

	"github.com/qibobo/webgo/db"
	"github.com/qibobo/webgo/models"
)

type SampleControler struct {
	logger *zap.Logger
	demoDB db.DemoDB
}

func NewSampleController(logger *zap.Logger, demoDB db.DemoDB) *SampleControler {
	return &SampleControler{
		logger: logger,
		demoDB: demoDB,
	}
}
func (s *SampleControler) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/demo/{id:int}", "GetDemoById")
	a.Handle("POST", "/demo", "CreateDemo")
}

func (s *SampleControler) GetDemoById(id int) mvc.Response {
	s.logger.Info("get demo by id", zap.Int("id", id))
	demo, err := s.demoDB.GetDemo(id)
	if err != nil {
		return mvc.Response{

			Code:        iris.StatusInternalServerError,
			ContentType: context.ContentJSONHeaderValue,
			Content:     []byte(fmt.Sprintf(`{"error":%s`, err.Error())),
		}
	}
	// if len(demos) == 0 {
	// 	return mvc.Response{

	// 		Code:        iris.StatusOK,
	// 		ContentType: context.ContentJSONHeaderValue,
	// 		Content:     []byte(`{}`),
	// 	}
	// }
	respJson, err := json.Marshal(demo)
	if err != nil {
		return mvc.Response{

			Code:        iris.StatusInternalServerError,
			ContentType: context.ContentJSONHeaderValue,
			Content:     []byte(fmt.Sprintf(`{"error":%s`, err.Error())),
		}
	}
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: context.ContentJSONHeaderValue,
		Content:     respJson,
	}

}
func (s *SampleControler) CreateDemo(d *models.Demo) mvc.Response {
	// demoBytes, err := json.Marshal(d)
	// if err != nil {
	// 	return mvc.Response{

	// 		Code:        iris.StatusInternalServerError,
	// 		ContentType: context.ContentJSONHeaderValue,
	// 		Content:     []byte(""),
	// 	}
	// }
	s.logger.Info("create demo", zap.Any("demo", d))
	err := s.demoDB.CreateDemo(d)
	if err != nil {
		return mvc.Response{

			Code:        iris.StatusInternalServerError,
			ContentType: context.ContentJSONHeaderValue,
			Content:     []byte(fmt.Sprintf(`{"error":%s`, err.Error())),
		}
	}
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: context.ContentJSONHeaderValue,
		Content:     []byte(`{}`),
	}
}
