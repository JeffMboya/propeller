package middleware

import (
	"context"
	"log/slog"
	"time"

	"github.com/absmach/propeller/manager"
	"github.com/absmach/propeller/proplet"
	"github.com/absmach/propeller/task"
)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    manager.Service
}

func Logging(logger *slog.Logger, svc manager.Service) manager.Service {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}

func (lm *loggingMiddleware) CreateProplet(ctx context.Context, w proplet.Proplet) (resp proplet.Proplet, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("proplet",
				slog.String("name", w.Name),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Save proplet failed", args...)

			return
		}
		lm.logger.Info("Save proplet completed successfully", args...)
	}(time.Now())

	return lm.svc.CreateProplet(ctx, w)
}

func (lm *loggingMiddleware) GetProplet(ctx context.Context, id string) (resp proplet.Proplet, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("proplet",
				slog.String("id", id),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Get proplet failed", args...)

			return
		}
		lm.logger.Info("Get proplet completed successfully", args...)
	}(time.Now())

	return lm.svc.GetProplet(ctx, id)
}

func (lm *loggingMiddleware) ListProplets(ctx context.Context, offset, limit uint64) (resp proplet.PropletPage, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Uint64("offset", offset),
			slog.Uint64("limit", limit),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("List proplets failed", args...)

			return
		}
		lm.logger.Info("List proplets completed successfully", args...)
	}(time.Now())

	return lm.svc.ListProplets(ctx, offset, limit)
}

func (lm *loggingMiddleware) UpdateProplet(ctx context.Context, t proplet.Proplet) (resp proplet.Proplet, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("proplet",
				slog.String("name", resp.Name),
				slog.String("id", t.ID),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update proplet failed", args...)

			return
		}
		lm.logger.Info("Update proplet completed successfully", args...)
	}(time.Now())

	return lm.svc.UpdateProplet(ctx, t)
}

func (lm *loggingMiddleware) DeleteProplet(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("proplet",
				slog.String("id", id),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Delete proplet failed", args...)

			return
		}
		lm.logger.Info("Delete proplet completed successfully", args...)
	}(time.Now())

	return lm.svc.DeleteProplet(ctx, id)
}

func (lm *loggingMiddleware) SelectProplet(ctx context.Context, t task.Task) (w proplet.Proplet, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("task",
				slog.String("name", t.Name),
				slog.String("id", t.ID),
			),
			slog.Group("proplet",
				slog.String("name", w.Name),
				slog.String("id", w.ID),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Select proplet failed", args...)

			return
		}
		lm.logger.Info("Select proplet completed successfully", args...)
	}(time.Now())

	return lm.svc.SelectProplet(ctx, t)
}

func (lm *loggingMiddleware) CreateTask(ctx context.Context, t task.Task) (resp task.Task, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("task",
				slog.String("name", t.Name),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Save task failed", args...)

			return
		}
		lm.logger.Info("Save task completed successfully", args...)
	}(time.Now())

	return lm.svc.CreateTask(ctx, t)
}

func (lm *loggingMiddleware) GetTask(ctx context.Context, id string) (resp task.Task, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("task",
				slog.String("id", id),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Get task failed", args...)

			return
		}
		lm.logger.Info("Get task completed successfully", args...)
	}(time.Now())

	return lm.svc.GetTask(ctx, id)
}

func (lm *loggingMiddleware) ListTasks(ctx context.Context, offset, limit uint64) (resp task.TaskPage, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Uint64("offset", offset),
			slog.Uint64("limit", limit),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("List tasks failed", args...)

			return
		}
		lm.logger.Info("List tasks completed successfully", args...)
	}(time.Now())

	return lm.svc.ListTasks(ctx, offset, limit)
}

func (lm *loggingMiddleware) UpdateTask(ctx context.Context, t task.Task) (resp task.Task, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("task",
				slog.String("name", resp.Name),
				slog.String("id", t.ID),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Update task failed", args...)

			return
		}
		lm.logger.Info("Update task completed successfully", args...)
	}(time.Now())

	return lm.svc.UpdateTask(ctx, t)
}

func (lm *loggingMiddleware) DeleteTask(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.Group("task",
				slog.String("id", id),
			),
		}
		if err != nil {
			args = append(args, slog.Any("error", err))
			lm.logger.Warn("Delete task failed", args...)

			return
		}
		lm.logger.Info("Delete task completed successfully", args...)
	}(time.Now())

	return lm.svc.DeleteTask(ctx, id)
}
