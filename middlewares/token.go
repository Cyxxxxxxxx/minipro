package middlewares

//
//import (
//	"BytesDanceProject/controller"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"time"
//)
//
//type UserClaims struct {
//	UserId string `json:"userId"`
//	//jwt-go提供的标准claim
//	jwt.StandardClaims
//}
//
//var (
//	//自定义的token密钥
//	secret = []byte("BackToSchool")
//	//路由下不校验token
//	noVerity = []interface{}{}
//	//token有效时间(纳秒)--7天
//	effectTime = 7 * 24 * time.Hour
//	//effectTime = 5 * time.Second
//)
//
//func GenerateToken(claims *UserClaims) (string, error) {
//	//设置token有效期
//	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
//	//生成token
//	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
//	//这里引入异常处
//	if err != nil {
//		return "", err
//	}
//	return sign, nil
//}
//
////验证token
//func JwtVerify(c *gin.Context) {
//	//过滤是否验证token
//	//点击登录的时候无需校验token，因为此时token不存在
//	//if c.Request.RequestURI == "/user/mine/login" || c.Request.RequestURI == "/ok" {
//	//	fmt.Println("跳过")
//	//	return
//	//}
//	//拿到token
//	token := c.GetHeader("token")
//
//	//如果token为空
//	if token == "" {
//		controller.ResponseError(c, controller.CodeErrGetToken)
//		c.Abort()
//	}
//	//校验token是否过期，未过期则插入上下文
//	userid, flag, err := Parsetoken(token)
//	if err != nil {
//		zap.L().Error("Err in parse token", zap.Error(err))
//		c.Abort()
//	}
//	if flag == false {
//		//zap.L().Info("Token is valid",zap.ByteStrings())
//		controller.ResponseSuccess(c, CodeTokenValid)
//	}
//
//	c.Set("userid", userid.UserId)
//
//	////校验token是否存在
//	//if token == "" {
//	//	panic("token not exist!")
//	//}
//	//验证token,并存储在请求中
//	//如何判断过期？
//	//c.Set("user", Parsetoken(token))
//}
//
////解析token
//func Parsetoken(tokenString string) (*UserClaims, bool, error) {
//	//解析token
//	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return secret, nil
//	})
//	if err != nil {
//		//fmt.Println("err in parse")
//		//panic(err)
//		return nil, true, err
//	}
//	claims, ok := token.Claims.(*UserClaims)
//	if !ok {
//		//fmt.Println("err in claim")
//		//panic("token is valid")
//		return claims, ok, nil
//	}
//	return claims, true, err
//}
//
////更新token
//func Refresh(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return secret, nil
//	})
//	if err != nil {
//		panic(err)
//	}
//	claims, ok := token.Claims.(*UserClaims)
//	if !ok {
//		panic("token is valid")
//	}
//	jwt.TimeFunc = time.Now
//	claims.StandardClaims.ExpiresAt = time.Now().Add(effectTime).Unix()
//	//返回生成一个token
//	return GenerateToken(claims)
//}
//
//func GetToken(c *gin.Context) {
//	user, _ := c.Get("user")
//	claims := user.(*UserClaims)
//	fmt.Println("userid:", claims.UserId)
//}
