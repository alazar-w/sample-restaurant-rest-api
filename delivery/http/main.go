package main

import (
	"net/http"
	"github.com/sample-restaurant-rest-api/comment/repository"
	"github.com/sample-restaurant-rest-api/comment/service"
	"github.com/sample-restaurant-rest-api/delivery/http/handler"
	urepim "github.com/sample-restaurant-rest-api/user/repository"
	usrvim "github.com/sample-restaurant-rest-api/user/service"


	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:P@$$w0rdD2@localhost/restaurantdb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	roleRepo := urepim.NewRoleGormRepo(dbconn)
	roleSrv := usrvim.NewRoleService(roleRepo)
	adminRoleHandler := handler.NewAdminRoleHandler(roleSrv)

	commentRepo := repository.NewCommentGormRepo(dbconn)
	commentSrv := service.NewCommentService(commentRepo)
	adminCommentHandler := handler.NewAdminCommentHandler(commentSrv)

	router := httprouter.New()

	router.GET("/v1/admin/roles", adminRoleHandler.GetRoles)

	router.GET("/v1/admin/comments/:id", adminCommentHandler.GetSingleComment)
	router.GET("/v1/admin/comments", adminCommentHandler.GetComments)
	router.PUT("/v1/admin/comments/:id", adminCommentHandler.PutComment)
	router.POST("/v1/admin/comments", adminCommentHandler.PostComment)
	router.DELETE("/v1/admin/comments/:id", adminCommentHandler.DeleteComment)

	http.ListenAndServe(":8181", router)
}
