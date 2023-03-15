package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sviper93/ecommerce_go/domain/user"
	storageUser "github.com/sviper93/ecommerce_go/infrastructure/postgres/user"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)
	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	userCase := user.New(storage)

	return newHandler(userCase)
}

func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")
	g.GET("", h.GetAll)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")
	g.GET("", h.Create)
}
