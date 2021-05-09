package route

import (
	"github.com/gofiber/fiber/v2"
	admin_controller "github.com/takemo101/go-fiber/app/controller/admin"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRoute is struct
type AdminRoute struct {
	logger              pkg.Logger
	app                 pkg.Application
	path                pkg.Path
	csrf                middleware.Csrf
	auth                middleware.SessionAdminAuth
	render              *helper.ViewRender
	dashboardController admin_controller.DashboardController
	adminController     admin_controller.AdminController
	userController      admin_controller.UserController
	todoController      admin_controller.TodoController
	accountController   admin_controller.AccountController
	authController      admin_controller.SessionAuthController
}

// Setup is setup route
func (r AdminRoute) Setup() {
	r.logger.Info("setup admin-route")

	app := r.app.App

	// root redirect
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(r.path.URL("system"))
	})

	// admin http route
	http := app.Group(
		"/system",
		r.csrf.CreateHandler("form:csrf_token"),
		r.render.CreateHandler(r.ViewRenderCreateHandler),
	)
	{
		// auth login route
		auth := http.Group(
			"/auth",
			r.auth.CreateHandler(false, "system"),
		)
		{
			auth.Get("/login", r.authController.LoginForm)
			auth.Post("/login", r.authController.Login)
		}

		// after login route
		system := http.Group(
			"/",
			r.auth.CreateHandler(true, "system/auth/login"),
		)
		{
			// dashboard route
			system.Get("/", r.dashboardController.Dashboard)

			// admin route
			admin := system.Group("/admin")
			{
				admin.Get("/", r.adminController.Index)
				admin.Get("/create", r.adminController.Create)
				admin.Post("/store", r.adminController.Store)
				admin.Get("/:id/edit", r.adminController.Edit)
				admin.Post("/:id/update", r.adminController.Update)
				admin.Post("/:id/delete", r.adminController.Delete)
			}

			// user route
			user := system.Group("/user")
			{
				user.Get("/", r.userController.Index)
				user.Get("/create", r.userController.Create)
				user.Post("/store", r.userController.Store)
				user.Get("/:id/edit", r.userController.Edit)
				user.Post("/:id/update", r.userController.Update)
				user.Post("/:id/delete", r.userController.Delete)
			}

			// todo route
			todo := system.Group("/todo")
			{
				todo.Get("/", r.todoController.Index)
				todo.Post("/store", r.todoController.Store)
				todo.Post("/:id/update", r.todoController.Update)
				todo.Post("/:id/change", r.todoController.ChangeStatus) // ajax
				todo.Post("/:id/delete", r.todoController.Delete)
			}

			// account route
			account := system.Group("/account")
			{
				account.Get("/edit", r.accountController.Edit)
				account.Post("/update", r.accountController.Update)
			}

			// auth logout route
			system.Post("/logout", r.authController.Logout)
		}
	}
}

// NewAdminRoute create new admin route
func NewAdminRoute(
	logger pkg.Logger,
	app pkg.Application,
	path pkg.Path,
	csrf middleware.Csrf,
	auth middleware.SessionAdminAuth,
	render *helper.ViewRender,
	dashboardController admin_controller.DashboardController,
	adminController admin_controller.AdminController,
	userController admin_controller.UserController,
	todoController admin_controller.TodoController,
	accountController admin_controller.AccountController,
	authController admin_controller.SessionAuthController,
) AdminRoute {
	return AdminRoute{
		logger:              logger,
		app:                 app,
		path:                path,
		csrf:                csrf,
		auth:                auth,
		render:              render,
		dashboardController: dashboardController,
		adminController:     adminController,
		userController:      userController,
		todoController:      todoController,
		accountController:   accountController,
		authController:      authController,
	}
}

// ViewRenderCreateHandler middleware handler
func (r AdminRoute) ViewRenderCreateHandler(c *fiber.Ctx, vr *helper.ViewRender) {
	// load menu list
	adminlte, err := r.app.Config.Load("admin-lte")
	if err == nil {
		for k, v := range map[string]string{
			"adminlte_menus":   "menus",
			"adminlte_plugins": "plugins",
		} {
			if value, ok := adminlte[v]; ok {
				vr.SetData(helper.DataMap{
					k: value,
				})
			}
		}
	}

	// load meta
	meta, err := r.app.Config.Load("admin-meta")
	if err == nil {
		vr.SetData(helper.DataMap{
			"admin_meta": meta,
		})
	}

	// csrf-token
	csrfToken := middleware.GetCSRFToken(c)
	// session errors
	errors, _ := middleware.GetSessionErrors(c)
	// session inputs
	inputs, _ := middleware.GetSessionInputs(c)
	// session messages
	messages, _ := middleware.GetSessionMessages(c)

	// admin user
	auth := r.auth.Auth.Admin()

	vr.SetData(helper.DataMap{
		"csrf_token": csrfToken,
		"errors":     errors,
		"inputs":     inputs,
		"messages":   messages,
		"auth":       auth,
	})
}
