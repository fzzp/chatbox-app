package v1

import (
	"chatbox-app/api"
	"chatbox-app/config"
	"chatbox-app/dto"
	"chatbox-app/lib/errs"
	"chatbox-app/lib/token"
	"chatbox-app/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

// SignUp godoc
// @Summary      用户注册
// @Description  使用邮箱密码注册
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /signup [post]
func (*AuthApi) SignUp(c *gin.Context) {
	var req dto.CreateUserInput
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	api.Fail(c, errs.ErrBadRequest.AsException(err))
	// 	return
	// }
	if ok := api.ShouldBind(c, &req); !ok {
		return
	}
	pswdHash, err := api.GenHashPassword(req.Password)
	if err != nil {
		api.Fail(c, errs.ErrBadRequest.AsException(err))
		return
	}
	user := models.User{
		Email:        req.Email,
		PasswordHash: pswdHash,
		Username:     strings.Split(req.Email, "@")[0],
	}
	userDao.CreateUser(&user)

	api.Succ(c, user)
}

func (*AuthApi) Login(c *gin.Context) {
	var input dto.UserLoginInput
	if ok := api.ShouldBind(c, &input); !ok {
		return
	}
	user := userDao.GetByEmail(input.Email)
	err := api.CheckPassword(input.Password, user.PasswordHash)
	if err != nil {
		log.Println(err)
		api.Fail(c, errs.ErrBadRequest.AsMessage("账号不存在或密码不匹配"))
		return
	}

	aPayload := token.NewPayload(int64(user.ID), config.App.Env.AccessTokenDur)
	rPayload := token.NewPayload(int64(user.ID), config.App.Env.RefreshTokenDur)

	aToken, err1 := config.App.JwtMaker.GenToken(aPayload)
	rToken, err2 := config.App.JwtMaker.GenToken(rPayload)
	if err1 != nil || err2 != nil {
		log.Println("err1: ", err1, "err2: ", err2)
		api.Fail(c, errs.ErrServerError)
		return
	}

	resp := dto.NewLoginUserOutput(user, aToken, rToken)

	api.Succ(c, resp)
}

func (*AuthApi) Logout(c *gin.Context) {
	api.Succ(c, c.Request.URL.Path)
}
