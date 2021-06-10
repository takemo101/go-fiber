package service

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewAdminService),
	fx.Provide(NewTodoService),
	fx.Provide(NewTagService),
	fx.Provide(NewCategoryService),
	fx.Provide(NewRequestService),
	fx.Provide(NewSuggestService),
	fx.Provide(NewDiscussionService),
)
