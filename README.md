# Gin-简易用户管理系统

---

### 项目功能

按照NX的模板，补充了国庆二面的代码，实现了CRUD

1. 用户的注册登录  login.go/register.go
2. 在鉴权下修改密码  update.go
3. 注销用户 delete.go
4. 展现用户信息 read.go

### 项目结构

````cmd
│  main.go
│  text.log
├─cmd
│  └─server
│          server.go
│
├─configs
│      config.yaml
│      init.go
│      model.go
│      setting.go
│
├─internal
│  ├─global
│  │  ├─db
│  │  │      db.go
│  │  │
│  │  ├─errs
│  │  │      code.go
│  │  │      logic.go
│  │  │      response.go
│  │  │
│  │  ├─jwt
│  │  │      NewJWT.go
│  │  │      ParseToken.go
│  │  │
│  │  ├─log
│  │  │      logger.go
│  │  │
│  │  └─middleware
│  │          Auth.go
│  │          recovery.go
│  │
│  ├─model
│  │      user.go
│  │
│  └─module
│      │  module.go
│      │
│      └─user
│              delete.go
│              init.go
│              login.go
│              read.go
│              register.go
│              routers.go
│              update.go
│
└─tools  // 没有写任何内容
````

### 项目组成(下面代码大多数都是GPT加抄的)

* **cmd ** 服务

````go
// server.go
func Init() {
	configs.Init()
	db.Init()
	for _, m := range module.Modules {
		fmt.Println("Init Module: " + m.GetName())
		m.Init()
	}
}
````

实现了 configs,db,功能模块的  Init（现阶段这边是空的）

````go
func Run() {
	r := gin.New()
	r.Use(log.Init(), middleware.Recovery())

	for _, m := range module.Modules {
		fmt.Println("InitRouter: " + m.GetName())
		m.InitRouter(r.Group("/" + m.GetName()))
	}

	panic(r.Run(":9090"))
}
````

中间件使用自己抄的log和Recovery，然后启动路由

* **configs **  配置文件

`````go
// setting.go
// 配置一个基础
func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 配置基础信息
	vp.SetConfigName("config")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath(".") //如果在当前目录找不到的话，可以在根目录上找
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func SetUpSettings() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}
	//...可以有多个Section
	err1 := setting.ReadSection("Database", &DbSettings)
	if err1 != nil {
		return err1
	}

	err2 := setting.ReadSection("jwt", &JwtSettings)
	if err2 != nil {
		return err2
	}

	return nil
}

// 将yaml数据绑定到结构体上
func (s *Setting) ReadSection(name string, v interface{}) error {
	err := s.vp.UnmarshalKey(name, v)
	if err != nil {
		return err
	}

	return nil
}
//init.go
func Init() {
	err := SetUpSettings()
	if err != nil {
		log.SugarLogger.Error(err)
		return
	}
}
`````

通过viper库，将yaml的字段绑定到结构体中

````go
type DatabaseSettings struct {
	Root      string
	Password  string
	Host      string
	Port      int
	Dbname    string
	Charset   string
	ParseTime string
	Loc       string
}

type JWTSettings struct {
	Issuer    string
	Subject   string
	SecretKey string
}
````

这边就配置了一个数据库和jwt的信息

````yaml
#数据库信息
database:
  root: "root"
  password: "123456"
  host: "localhost"
  port: 3306
  dbname: "itcast"
  charset: "utf8mb4"
  parseTime: "True"
  loc: "Local"
#鉴权信息
jwt:
  Issuer: "Echin"
  Subject: "Tom"
  SecretKey: "ItIsSecret"
````

* **db **   数据库

````go
func Init() {
	DB = Connect()
}

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		configs.DbSettings.Root,
		configs.DbSettings.Password,
		configs.DbSettings.Host,
		configs.DbSettings.Port,
		configs.DbSettings.Dbname,
		configs.DbSettings.Charset,
		configs.DbSettings.ParseTime,
		configs.DbSettings.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.SugarLogger.Error(err)
		return nil
	}
	fmt.Println("连接数据库成功")
	return db
}
````

使用gorm最基础的CRUD，连接操作

* **errs**   错误处理 （大体是抄NX的）

```go
var (
	INVALID_REQUEST = newError(40001, "无效的请求")
	NOTFOUND        = newError(40002, "目标不存在")
	HAS_EXIST       = newError(40003, "目标已存在")
	LOGIN_ERROR     = newError(40004, "d登陆失败")
	UNTHORIZATION   = newError(40005, "鉴权失败")
)

var (
	DB_LINK_ERROR = newError(50001, "连接数据库失败")
	DB_CRUD_ERROR = newError(50002, "数据库操作失败")
	DB_BASE_ERROR = newError(50003, "数据库内部错误")
)

var (
	SERVE_INTERNAL = newError(60001, "服务器内部故障")
)

```

基础的错误字段

````go
// logic.go
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Origin  string `json:"origin"` // Origin字段通常用于只是错误的来源或者上下文
}

// 绑定数字和文字之间的关联
func newError(code int64, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

// 表示出Error中的Message  但是整个项目好像没用到
func (e *Error) Error() string {
	return e.Message
}

// 比较是否是有相同的错误码
func (e *Error) Is(target error) bool {
	var t *Error
	// errors.As是把target转化成t的类型
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// 添加error类型
func (e *Error) WithOrigin(err error) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Origin:  fmt.Sprintf("%+v", err),
	}
}

// 添加字符串进入
func (e *Error) WithTips(details ...string) *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message + " " + fmt.Sprintf("%v", details),
	}
}
````

`````go
// response.go
// 错误返回
type responseBody struct {
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Origin string `json:"origin"`
	Data   any    `json:"data"`
}

// ...any可以当作切片来处理
func Success(c *gin.Context, data ...any) {
	response := responseBody{
		Code:   SUCCESS.Code,
		Msg:    SUCCESS.Message,
		Origin: SUCCESS.Origin,
		Data:   data,
	}
	//if len(data) > 0 {
	//	response.Data = data[0]
	//}
	c.JSON(http.StatusOK, response)
}

func Fail(c *gin.Context, err error) {
	var e *Error
	ok := errors.As(err, &e)
	if !ok {
		e = SERVE_INTERNAL.WithOrigin(err)
	}

	var resp responseBody
	resp.Code = e.Code
	resp.Msg = e.Message
	resp.Origin = e.Origin

	c.JSON(int(e.Code/100), resp)
	c.Abort()
}
`````

运行成功了就返回Success , 失败了就返回Fail

往后的返回中会将c.JSON抛弃，然后返回更加统一的Error类型的错误，

通过一开始预定义的错误字段来返回相应的基础错误，同时还可以通过Withtips/WtihWrap方法来增加错误信息

* **jwt**

生成+验证

`````go
// 把claims作为一个结构体
type Payload struct {
	Authorized bool   `json:"authorized"`
	User       string `json:"user"`
}

type MyCustomClaims struct {
	Payload
	jwt.RegisteredClaims
}

func NewToken(name string) (string, error) {

	//设置一些预定义  Payload
	claims := &MyCustomClaims{
		Payload: Payload{
			Authorized: true,
			User:       name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.JwtSettings.Issuer,
			Subject:   configs.JwtSettings.Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			// jwt.NewNumericDate 可以创建一个符合JWT标准的时间格式
		},
	}

	// 创建一个新的令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Header,token是一个对象

	//签名并获取完整的编码令牌作为字符串  Signature
	tokenString, err := token.SignedString([]byte(configs.JwtSettings.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
`````

payload定义成结构体，创建一个令牌

```go
// ParseToken 是 解析令牌
func ParseToken(bearerToken string) (*MyCustomClaims, error) {
	// 解析方式需要添加 Bearer token模式
	tokenParts := strings.Split(bearerToken, " ")                           //通过空格分隔出两个部分，并且存入数组之中
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" { // strings.ToLower会把字符串变成小写的形式
		return nil, errors.New("Invalid token format,you need add bearer")
	}
	tokenString := tokenParts[1]
	// 解析后续token
	claims := &MyCustomClaims{}
	// 是*token和string之间的转换
	// 这是一个回调函数具体结构就是 jwt.Parse(string,KeyFunc)
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法 HMAC-SHA56签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(configs.JwtSettings.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
```

* **log**

使用了zap日志，直接看官方文档配置的，大体能用，好不好不知道

````go
var (
	SugarLogger *zap.SugaredLogger
)

// 做一个闭包
func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		initLogger()
		c.Next()
	}
}

func initLogger() {
	Encoder := getEncoder()
	WriterSyncer := getWriterSyncer()
	core := zapcore.NewCore(Encoder, WriterSyncer, zapcore.DebugLevel)
	//  zap.AddCaller()可以实现记录函数信息
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 写入位置
func getWriterSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./text.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

````

````go
// 文档中的之前部分错误呈现
{"level":"ERROR","ts":"2024-02-03T14:11:52.395+0800","caller":"db/db.go:29","msg":"dial tcp :0: connectex: The requested address is not valid in its context."}
{"level":"ERROR","ts":"2024-02-03T14:13:34.266+0800","caller":"config/config.go:12","msg":"open ./config.yaml: The system cannot find the file specified."}
````

* **middleware**

一个是鉴权，一个是异常捕获

````go
// Auth.go
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to fetch token",
			})
			return
		}
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			log.SugarLogger.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":   "Failed to Auth",
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("Payload", parseToken)
		c.Next()
	}
}
````

通过c.Set（）把Payload放入上下文中，然后后面修改密码可以通过c.Get()+类型断言获取相应值

````go
// recovery.go
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer errs.Recovery(c)
		c.Next()
	}
}
func Recovery(c *gin.Context) {
	if info := recover(); info != nil {
		err, ok := info.(error)
		if ok {
			Fail(c, SERVE_INTERNAL.WithOrigin(err))
		} else {
			Fail(c, errors.New(fmt.Sprintf("%+v", info)))
		}
		return
	}
}
````

这样写更加全面，说实话我感觉直接recover（）也行

* **module**

这块内容的布局方式是抄NX的，感觉非常神奇

主要通过定义一个结构实现三个方法来使得每个大功能能够独立（但是我这边只有一个用户登录注册的大功能，所以就只创建了一个user包）

然后通过遍历modules切片的方式来启动路由

model

````go
type User struct {
	gorm.Model
	Name      string    `gorm:"varchar(20);not null" json:"name"`     //姓名
	Gender    *string   `gorm:"char" json:"gender"`                   //性别
	Age       int       `json:"age"`                                  //年龄
	Birthday  time.Time `gorm:"not null" json:"birthday"`             //出生日期
	Telephone *string   `gorm:"varchar(20)" json:"telephone"`         //电话
	Password  string    `gorm:"varchar(20);not null" json:"password"` //密码
	Email     string    `gorm:"varchar(30);not null" json:"email"`    //邮箱
}
````



````go
// module.go
// 实现了两个方法
type Module interface{
    GetName()string // 对照InitGroup,传入一个大功能的路径
    Init() // 我感觉没啥用，NX说可以做一些提前配置
    InitRouter(r *gin.RouterGroup) 
}
var Modules []Module

// 添加项目功能模块
RegisterModule(m Module){
    Module = append(Module,m)
}

func init(){
    RegisterModule(&user.ModuleUser)
}

````

下面是具体功能的实现：

init

````go
// init.go
type ModuleUser struct{}

func (u *ModuleUser) GetName()string {
    return "user"
}

func (u *ModuleUser) Init(){}
````

路由：

```go
// routers.go
func (u *ModuleUser) InitRouter(r *gin.RouterGroup){
    r.POST("/register", Register)
	r.POST("/login", Login)
	r.PUT("/update", middleware.Auth(), Update)
	r.GET("/read", middleware.Auth(), Read)
	r.DELETE("/delete", middleware.Auth(), Delete)
}
```

注册：

```go
// 用户的登录
func Login(c *gin.Context) {
	// 用PostForm 传输数据
	name := c.PostForm("name")
	password := c.PostForm("password")

	var v2 model.User
	if tx := db.DB.Where(" name = ? ", name).First(&v2); tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(tx.Error))
		return
	}

	if password != v2.Password {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("密码错误"))
		return
	}

	token, err := jwt.NewToken(name)
	if err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.UNTHORIZATION.WithOrigin(err))
		return
	}

	errs2.Success(c, v2, map[string]string{"token": token})
}
```

登录：

````go
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.INVALID_REQUEST.WithOrigin(err))
		return
	}

	if len(user.Name) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("Username, password or email is required..."))
		return
	}

	var v1 model.User
	result := db.DB.Where("name = ?", user.Name).First(&v1)
	if result.Error == nil {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("姓名重复"))
	}

	if err := db.DB.Create(&user).Error; err != nil {
		err1 := db.DB.AutoMigrate(&user)
		if err1 != nil {
			log.SugarLogger.Error(err1)
			return
		}
		log.SugarLogger.Error(err)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(err))
		return
	}

	errs2.Success(c, "注册成功")
}
````

修改密码

`````go
func Update(c *gin.Context) {
	NewPassword := c.PostForm("NewPassword")
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.LOGIN_ERROR.WithTips("无法get你的payload"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	tx := db.DB.Model(&model.User{}).Where("name = ? ", load.User).Update("password", NewPassword)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
	}
	errs2.Success(c, load.User, NewPassword, "请妥善保存你的密码")
}
// 通过Payload中的name字段的比较，来鉴权
`````

注销账户

````go
// 注销账户
func Delete(c *gin.Context) {
	password := c.PostForm("password")
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.UNTHORIZATION.WithTips("没有获取到payload"))
		return
	}

	load := payload.(*jwt.MyCustomClaims)

	var user model.User
	db.DB.Where("name = ?", load.User).First(&user)
	if user.Password != password {
		fmt.Println("user.Password", user.Password)
		fmt.Println(password)
		errs2.Fail(c, errs2.INVALID_REQUEST.WithTips(password))
		return
	}

	result := db.DB.Delete(&user)
	if result.Error != nil {
		log.SugarLogger.Error(result.Error)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(result.Error))
		return
	}

	errs2.Success(c, "注销成功")
}
// 与 Update同理
````

查看用户信息

````go
func Read(c *gin.Context) {
	payload, exists := c.Get("Payload")
	if !exists {
		errs2.Fail(c, errs2.UNTHORIZATION.WithTips("Failed to Auth"))
		return
	}
	load := payload.(*jwt.MyCustomClaims)
	var user model.User
	tx := db.DB.Where("name = ?", load.User).First(&user)
	if tx.Error != nil {
		log.SugarLogger.Error(tx.Error)
		errs2.Fail(c, errs2.DB_CRUD_ERROR.WithOrigin(tx.Error))
	}
	errs2.Success(c, user)
}
````

## 运行

1. 注册(略)

````JSON
{
    "name": "string",
    "gender": "string",
    "age": "string",
    "birthday": "string",
    "telephone": "string",
    "password": "string",
    "email": "string"
}
````



2. 登录

在x-www-form-urlencoded输入账号密码正确即可

````json
// 返回
{
    "code": 200,
    "msg": "Success",
    "origin": "",
    "data": [
        {
            "ID": 4,
            "CreatedAt": "2024-02-03T20:17:01.557+08:00",
            "UpdatedAt": "2024-02-05T16:03:17.1+08:00",
            "DeletedAt": null,
            "name": "John",
            "gender": "Male",
            "age": 30,
            "birthday": "1994-02-03T08:00:00+08:00",
            "telephone": "1234567890",
            "password": "wordwordword",
            "email": "johndoe@example.com"
        },
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjoiSm9obiIsImlzcyI6IkVjaGluIiwic3ViIjoiVG9tIiwiZXhwIjoxNzA3Mzc1NTMyfQ.LRFwitASeeHpNRUKpQrH_GytsVn53YIor59uKdRbtUs"
        }
    ]
}
````

修改密码

输入密码和令牌就行

````json
{
    "code": 200,
    "msg": "Success",
    "origin": "",
    "data": [
        "John",
        "",
        "请妥善保存你的密码"
    ]
}
````

注销/用户信息展示(略，跟修改密码差不多)

## 总结

功能的实现很简单，上面的很多（基本上全部）代码也都是抄别人/GPT的

让我对项目整体的结构有了个相对完整的了解，曾经我以为只需要CRUD就行

但是以我现在的认知不知道还能够补充什么。

报错的过程还是很痛苦的。。