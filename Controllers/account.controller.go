package Controllers

import (
	"fmt"
	"gin-gonic-gom/Models"
	"gin-gonic-gom/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	AccountService Services.AccountService
}

func New(accountService Services.AccountService) AccountController {
	return AccountController{
		AccountService: accountService,
	}
}
func (account *AccountController) CreateAccount(ctx *gin.Context) {
	var user Models.AccountModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Print(user)
	err := account.AccountService.CreateAccount(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create Account Success!!!",
	})
}
func (account *AccountController) GetAccount(ctx *gin.Context) {
	var username string = ctx.Param("name")
	user, err := account.AccountService.GetAccount(&username)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	ctx.JSON(
		http.StatusOK,
		user,
	)
}

func (ac *AccountController) RegisterAccountRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/account")
	userroute.POST("/create", ac.CreateAccount)
	userroute.GET("/get/:name", ac.GetAccount)
	//userroute.GET("/getall", uc.GetAll)
	//userroute.PATCH("/update", uc.UpdateUser)
	//userroute.DELETE("/delete/:name", uc.DeleteUser)
}
