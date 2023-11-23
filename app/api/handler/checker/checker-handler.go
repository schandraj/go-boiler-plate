package checker

import (
	"base-plate/domain/checker"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/v2"
)

type HandlerChecker struct {
	service checker.Service
}

func NewCheckerHandler(svc checker.Service) *HandlerChecker {
	return &HandlerChecker{
		service: svc,
	}
}

func (h *HandlerChecker) HealthCheck(c *fiber.Ctx) error {
	span, ctx := apm.StartSpan(c.Context(), "HealthCheck", "api")
	defer span.End()

	res, err := h.service.HealthCheck(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
