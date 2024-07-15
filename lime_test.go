package lime

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"testing"
)

func addHandler(app *fiber.App) {
	app.Get("/ni", func(ctx fiber.Ctx) error {
		return ctx.SendString("你好")
	})
}

func addMiddleware(app *fiber.App) {

}

func testUtils() {
	log.Debug("===Utils")
}

func testRepo() {
	log.Debug("===Repo")
}

func TestWebApp(t *testing.T) {
	app := NewApp2(testUtils, testRepo, addHandler, nil)
	Run("test.env", app)
}
