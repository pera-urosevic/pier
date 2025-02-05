package database

import (
	"pier/api/database/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller.GetDatabaseIndex(r)
	controller.GetDocs(r)
	controller.GetFacets(r)
	controller.UpdateFacets(r)
	controller.GetDocsCount(r)
	controller.GetAutocompletes(r)
	controller.GetDoc(r)
	controller.GetDuplicate(r)
	controller.CreateDoc(r)
	controller.UpdateDoc(r)
	controller.RemoveDoc(r)
	controller.SetCache(r)
	controller.GetCache(r)
	controller.RemoveCache(r)
}
