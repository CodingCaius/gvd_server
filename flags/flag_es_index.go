package flags

import (
	"gvd_server/models"
	"gvd_server/service/es_service/indexs"
)

func ESIndex() {
	indexs.CreateIndex(models.FullTextModel{})
}
