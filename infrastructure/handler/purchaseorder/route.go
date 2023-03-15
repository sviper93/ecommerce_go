package purchaseorder

import (
	"github.com/sviper93/ecommerce_go/domain/purchaseorder"
	// "github.com/alexyslozada/ecommerce/infrastructure/handler/middle"
	"github.com/jackc/pgx/v5/pgxpool"
	purchaseorderStorage "github.com/sviper93/ecommerce_go/infrastructure/postgres/purchaseorder"

	"github.com/labstack/echo/v4"
)

// NewRouter returns a router to handle model.PurchaseOrder requests
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	// authMiddleware := middle.New()
	// privateRoutes(e, h, authMiddleware.IsVali--d)

	privateRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbPool))
	return newHandler(useCase)
}

// privateRoutes handle the routes that requires a token
// func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
func privateRoutes(e *echo.Echo, h handler) {
	// route := e.Group("/api/v1/private/purchase-orders", middlewares...)

	route := e.Group("/api/v1/private/purchase-orders")
	route.POST("", h.Create)
}
