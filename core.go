package lime

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recover2 "github.com/gofiber/fiber/v3/middleware/recover"
)

type Fun = func()

type FiberFun = func(app *fiber.App)

type ILime interface {
	Init(envPath string)
	Run()
}

type Lime struct {
	InitEnv   func(envPath string)
	InitUtils Fun
	InitDB    Fun
	InitRepo  Fun

	app        *fiber.App
	Middleware FiberFun
	Handler    FiberFun
}

func (l *Lime) Init(envPath string) {
	log.Debug("InitEnv ---")
	l.InitEnv(envPath)
	log.Debug("InitUtils ---")
	l.InitUtils()
	log.Debug("InitDB ---")
	l.InitDB()
	log.Debug("InitRepo ---")
	l.InitRepo()

	//初始化fiberApp
	app := fiber.New()
	//添加一些默认的中间件
	//图标
	app.Use(favicon.New())
	//日志
	app.Use(logger.New())
	//跨域
	app.Use(cors.New())
	//重试
	app.Use(recover2.New())
	l.app = app
	l.Middleware(app)
	l.Handler(app)
}

func (l *Lime) Run() {
	//读取端口
	addr := GetEnvValue(EnvKeyAddr)
	if addr == "" {
		addr = ":8080"
	}
	err := l.app.Listen(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func NewApp2(utils, repo Fun, addMiddleware, addHandler FiberFun) *Lime {
	return NewApp1(InitEnvFile, utils, InitDB, repo, addMiddleware, addHandler)
}

func NewApp1(env func(envPath string), utils, db, repo Fun, addMiddleware, addHandler FiberFun) *Lime {
	if env == nil {
		env = func(envPath string) {}
	}
	if utils == nil {
		utils = func() {}
	}
	if db == nil {
		db = func() {}
	}
	if repo == nil {
		repo = func() {}
	}
	if addMiddleware == nil {
		addMiddleware = func(app *fiber.App) {}
	}
	if addHandler == nil {
		addHandler = func(app *fiber.App) {}
	}
	return &Lime{
		InitEnv:    env,
		InitUtils:  utils,
		InitDB:     db,
		InitRepo:   repo,
		Middleware: addMiddleware,
		Handler:    addHandler,
	}
}

func Run(envPath string, lime ILime) {
	lime.Init(envPath)
	lime.Run()
}
