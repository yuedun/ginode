package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	ojwt "github.com/golang-jwt/jwt/v4"
)

type loginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var ErrAccountLocked = errors.New("账号已锁定")
var ErrAccountDisabled = errors.New("账号已失效")

var identityKey = "username"

func Jwt() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "accord-admin-b",
		Key:         []byte(g_config.JWTSecret),
		Timeout:     time.Hour * 24 * 7, // 7天
		MaxRefresh:  time.Hour * 24 * 7,
		IdentityKey: identityKey,
		// 第一步：首次通过用户名密码登录认证
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginReq
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			account := loginVals.Username
			pwdB64 := loginVals.Password

			tenantService := NewUserService(PGDB, nil)
			userObj := AdminUser{
				UserName: account,
			}
			adminUser, err := tenantService.getAdminUserInfo(&userObj)
			if err != nil {
				log.Println("登录查询数据", err)
				return nil, jwt.ErrFailedAuthentication
			}
			password, err := AesECBPkcs7Decrypt(pwdB64)
			if err != nil {
				log.Println("登录查询数据", err)
				return nil, jwt.ErrFailedAuthentication
			}
			if adminUser.Password != GetMD5(password, g_config.Salt) {
				return nil, jwt.ErrFailedAuthentication
			}

			// 返回的数据用在上面定义的PayloadFunc函数中
			c.Set(identityKey, adminUser) // 用户信息保存到context
			return adminUser, nil
		},
		// 第二步：登录验证成功后存储用户信息，data从第一步获取到，转换成jwt.MapClaim数据
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(AdminUser); ok {
				return jwt.MapClaims{
					"uid":       v.ID,
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		//第三步：此函数的目的是从嵌入在 jwt 令牌中的声明中获取用户身份，并将此身份值传递给 Authorizator
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			token := jwt.GetToken(c)
			cUser := AdminUser{
				// ID:       uint(claims["tid"].(float64)),
				UserName:  claims[identityKey].(string),
				AuthToken: token,
			}
			cUser.ID = uint(claims["uid"].(float64))

			tenantService := NewUserService(PGDB, nil)
			dbUser, err := tenantService.getAdminUserInfo(&cUser)
			if err != nil {
				log.Printf("无法获取用户信息，err=%v\n", err)
				return nil
			}
			if cUser.AuthToken != dbUser.AuthToken {
				log.Println("token不相符, 可能重新登陆过")
				return nil
			}
			return &dbUser
		},

		//第四步：登录以后通过token来获取用户标识，检测是否通过认证
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return data != nil
		},

		// 遇到err时处理返回信息
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			if err, ok := e.(*ojwt.ValidationError); ok {
				switch err.Errors {
				case ojwt.ValidationErrorExpired:
					return "登录已过期, 请重新登录"
				default:
					return "登录验证失败"
				}
			}
			switch {
			case errors.Is(e, jwt.ErrMissingSecretKey):
				return "秘钥未配置"
			case errors.Is(e, jwt.ErrForbidden):
				return "您没有访问此资源的权限"
			case errors.Is(e, jwt.ErrMissingLoginValues):
				return "缺少用户名或密码"
			case errors.Is(e, jwt.ErrFailedAuthentication):
				return "用户名或密码不正确"
			case errors.Is(e, jwt.ErrFailedTokenCreation):
				return "无法创建JWT令牌"
			case errors.Is(e, jwt.ErrExpiredToken):
				return "登录已过期, 请重新登录"
			case errors.Is(e, jwt.ErrEmptyAuthHeader):
				return "缺少令牌头信息"
			case errors.Is(e, jwt.ErrMissingExpField):
				return "无效的令牌"
			case errors.Is(e, jwt.ErrWrongFormatOfExp):
				return "无效的令牌"
			case errors.Is(e, jwt.ErrInvalidAuthHeader):
				return "无效的令牌头"
			case errors.Is(e, jwt.ErrEmptyQueryToken):
				return "令牌为空"
			case errors.Is(e, jwt.ErrEmptyCookieToken):
				return "cookie令牌为空"
			case errors.Is(e, jwt.ErrEmptyParamToken):
				return "参数令牌为空"
			case errors.Is(e, jwt.ErrInvalidSigningAlgorithm):
				return "无效的令牌算法"
			case errors.Is(e, jwt.ErrNoPrivKeyFile):
				return "无效的私钥"
			case errors.Is(e, jwt.ErrNoPubKeyFile):
				return "无效的公钥"
			case errors.Is(e, jwt.ErrInvalidPrivKey):
				return "无效的私钥"
			case errors.Is(e, jwt.ErrInvalidPubKey):
				return "无效的公钥"
			case errors.Is(e, ErrAccountLocked):
				return "出于安全原因，您的账号已锁定，请十分钟后重试"
			case errors.Is(e, ErrAccountDisabled):
				return "账号已失效"
			default:
				return "鉴权失败，请重新登录!"
			}
		},
		// 获取不到token或解析token失败时如何返回信息
		Unauthorized: func(c *gin.Context, code int, message string) {
			// 特殊处理，当token被刷新时，提示令牌失效，status = 401
			if code == http.StatusForbidden {
				if user, ok := c.Get(identityKey); !ok || user == nil {
					code = http.StatusUnauthorized
					message = "登录失效, 请重新登录"
				}
			}

			log.Println("token验证失败：", message)
			c.JSON(code, gin.H{
				"code": code,
				"msg":  message,
			})
		},
		// 获取jwt token的方法，从header中获取，从query中获取，从cookie中获取
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, cookie: token",
		// TokenLookup: "cookie:jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		//LoginHandler,LogoutHandler等handler中间件会默认提供，但其返回的数据格式并不一定符合项目规范，也可以在此处自定义，像上面Unauthorized这样
		SendCookie:     true,
		CookieName:     "token", // default jwt
		CookieHTTPOnly: true,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			// 刷新token
			tuser, _ := c.Get(identityKey)
			uObj := tuser.(AdminUser)
			uObj.AuthToken = token
			userervice := NewUserService(PGDB, nil)
			userervice.adminUserLoginSucc(uObj, token)

			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "ok",
				"result": map[string]string{
					"token": token,
				},
				"data": map[string]string{
					"token": token,
				},
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "ok",
				"result":  nil,
			})
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
