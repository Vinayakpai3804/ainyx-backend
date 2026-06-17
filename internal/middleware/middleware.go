package middleware

import (
	"ainyx-backend/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()
	requestId := uuid.New().String()
	c.Set("X-Request-ID", requestId)
	err := c.Next()
	duration := time.Since(start)
	logger.Log.Info("Request",
		zap.String("requestId", requestId),
		zap.String("method", c.Method()),
		zap.String("path", c.Path()),
		zap.Int("status", c.Response().StatusCode()),
		zap.Duration("duration", duration),
	)
	return err
}