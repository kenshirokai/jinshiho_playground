package server

import (
	"fmt"
	"net/http"
	"os"

	"jinshiho/middleware"
	"jinshiho/server/generated"

	ginMiddleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

func New(port int) *http.Server {
	swagger, err := generated.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	r := gin.Default()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(ginMiddleware.OapiRequestValidator(swagger),
		middleware.Error()) // ここに独自のmiddelwareをセットしていく

	svr := NewAPI()
	generated.RegisterHandlers(r, svr)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("localhost:%d", port),
	}

	return s
}
