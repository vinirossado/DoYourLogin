package middlewares

import (
	"doYourLogin/source/configuration"
	"doYourLogin/source/domain/enumerations"
	"doYourLogin/source/domain/requests"
	"doYourLogin/source/repositories"
	"doYourLogin/source/utils"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginOK struct {
	Code   int    `json:"code" example:"200"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	Expire string `json:"expire" example:"2006-01-02T15:04:05Z07:00"`
}

type LoginError struct {
	StatusCode int    `json:"status_code" example:"401"`
	Message    string `json:"message" example:"Invalid username or password"`
}

type Claims struct {
	ID        uint
	Name      string
	Role      enumerations.Roles
	CompanyID uint
}

var identityKey = configuration.JWT_IDENTITY_KEY.ValueAsString()

func JwtMiddleware() *jwt.GinJWTMiddleware {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           configuration.JWT_REALM.ValueAsString(),
		Key:             configuration.JWT_SECRET_KEY.ValueAsByte(),
		Timeout:         configuration.JWT_TIMEOUT.ValueAsDuration(),
		MaxRefresh:      configuration.JWT_MAX_REFRESH.ValueAsDuration(),
		IdentityKey:     identityKey,
		Authenticator:   LoginHandler,
		PayloadFunc:     PayloadHandler,
		IdentityHandler: IdentityHandler,
		Authorizator:    AutorizatorHandler,
		Unauthorized:    UnauthorizedHandler,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the requests.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Initialization Error: " + err.Error())
	}

	return authMiddleware
}

// LoginHandler Login godoc
// @Summary      Login
// @Description  Login and generate jwt auth
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        auth  body      requests.Auth  true  "Auth Info"
// @Success      200   {object}  middlewares.LoginOK
// @Failure      400   {object}  exceptions.HttpException
// @Failure      401   {object}  middlewares.LoginError
// @Router       /login [post]
func LoginHandler(c *gin.Context) (interface{}, error) {
	auth := requests.Auth{}
	utils.ReadRequestBody(c, &auth)

	user, err := repositories.FindUserByEmail(auth.Email)

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	if string(hashedPassword) == user.Password {
		return nil, jwt.ErrFailedAuthentication
	}

	return &Claims{
		ID:        user.ID,
		Name:      user.Name,
		Role:      user.Role,
		CompanyID: user.CompanyID,
	}, nil

}

func PayloadHandler(data interface{}) jwt.MapClaims {
	user := data.(*Claims)

	return jwt.MapClaims{
		identityKey:  user.Name,
		"id":         float64(user.ID),
		"role":       int(user.Role),
		"company_id": float64(user.CompanyID),
	}
}

func IdentityHandler(c *gin.Context) interface{} {
	return ExtractClaims(c)
}

func AutorizatorHandler(data interface{}, c *gin.Context) bool {
	return true
}

func UnauthorizedHandler(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, LoginError{
		StatusCode: statusCode,
		Message:    message,
	})
}
