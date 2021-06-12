package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/takemo101/go-fiber/app/controller/admin"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRoute is struct
type AdminRoute struct {
	logger               pkg.Logger
	app                  pkg.Application
	path                 pkg.Path
	csrf                 middleware.Csrf
	value                support.RequestValue
	auth                 middleware.SessionAdminAuth
	render               middleware.ViewRender
	dashboardController  controller.DashboardController
	adminController      controller.AdminController
	userController       controller.UserController
	todoController       controller.TodoController
	tagController        controller.TagController
	categoryController   controller.CategoryController
	requestController    controller.RequestController
	suggestController    controller.SuggestController
	discussionController controller.DiscussionController
	accountController    controller.AccountController
	authController       controller.SessionAuthController
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
				admin.Put("/:id/update", r.adminController.Update)
				admin.Delete("/:id/delete", r.adminController.Delete)
			}

			// user route
			user := system.Group("/user")
			{
				user.Get("/", r.userController.Index)
				user.Get("/create", r.userController.Create)
				user.Post("/store", r.userController.Store)
				user.Get("/:id/edit", r.userController.Edit)
				user.Put("/:id/update", r.userController.Update)
				user.Delete("/:id/delete", r.userController.Delete)
			}

			// todo route
			todo := system.Group("/todo")
			{
				todo.Get("/", r.todoController.Index)
				todo.Get("/your", r.todoController.Your)
				todo.Post("/store", r.todoController.Store)
				todo.Patch("/:id/change", r.todoController.ChangeStatus) // ajax
				todo.Delete("/:id/delete", r.todoController.Delete)
			}

			// tag route
			tag := system.Group("/tag")
			{
				tag.Get("/", r.tagController.Index)
				tag.Post("/store", r.tagController.Store)
				tag.Patch("/sort", r.tagController.Sort)
				tag.Put("/:id/update", r.tagController.Update)
				tag.Delete("/:id/delete", r.tagController.Delete)
			}

			// category route
			category := system.Group("/category")
			{
				category.Get("/", r.categoryController.Index)
				category.Post("/store", r.categoryController.Store)
				category.Patch("/sort", r.categoryController.Sort)
				category.Put("/:id/update", r.categoryController.Update)
				category.Delete("/:id/delete", r.categoryController.Delete)
			}

			// request route
			request := system.Group("/request")
			{
				request.Get("/", r.requestController.Index)
				request.Get("/create/user/:id", r.requestController.Create)
				request.Post("/store/user/:id", r.requestController.Store)
				request.Get("/:id/detail", r.requestController.Detail)
				request.Get("/:id/edit", r.requestController.Edit)
				request.Put("/:id/update", r.requestController.Update)
				request.Delete("/:id/delete", r.requestController.Delete)
			}

			// suggest route
			suggest := system.Group("/suggest")
			{
				suggest.Get("/:id/detail", r.suggestController.Detail)
				suggest.Delete("/:id/delete", r.suggestController.Delete)
			}

			// discussion route
			discussion := system.Group("/discussion")
			{
				discussion.Get("/", r.discussionController.Index)
				discussion.Delete("/:id/delete", r.discussionController.Delete)
			}

			// account route
			account := system.Group("/account")
			{
				account.Get("/edit", r.accountController.Edit)
				account.Put("/update", r.accountController.Update)
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
	value support.RequestValue,
	render middleware.ViewRender,
	auth middleware.SessionAdminAuth,
	dashboardController controller.DashboardController,
	adminController controller.AdminController,
	userController controller.UserController,
	todoController controller.TodoController,
	tagController controller.TagController,
	categoryController controller.CategoryController,
	requestController controller.RequestController,
	suggestController controller.SuggestController,
	discussionController controller.DiscussionController,
	accountController controller.AccountController,
	authController controller.SessionAuthController,
) AdminRoute {
	return AdminRoute{
		logger:               logger,
		app:                  app,
		path:                 path,
		csrf:                 csrf,
		value:                value,
		auth:                 auth,
		render:               render,
		dashboardController:  dashboardController,
		adminController:      adminController,
		userController:       userController,
		todoController:       todoController,
		tagController:        tagController,
		categoryController:   categoryController,
		requestController:    requestController,
		suggestController:    suggestController,
		discussionController: discussionController,
		accountController:    accountController,
		authController:       authController,
	}
}

// ViewRenderCreateHandler middleware handler
func (r AdminRoute) ViewRenderCreateHandler(c *fiber.Ctx, vr *helper.ViewRender) {
	// load request list
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
	auth := r.value.GetSessionAdminAuth(c)
	admin := auth.Admin()

	vr.SetData(helper.DataMap{
		"csrf_token": csrfToken,
		"errors":     errors,
		"inputs":     inputs,
		"messages":   messages,
		"auth":       admin,
	})
	vr.SetJS(helper.DataMap{
		"csrfToken": csrfToken,
		"errors":    errors,
		"inputs":    inputs,
		"messages":  messages,
	})
}
