package router

import (
	"app/controller"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type TypeReturn struct {
	Key   string
	Value interface{}
}

func Router() http.Handler {
	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Timeout(60 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	app.Use(cors.Handler)

	// middlewares := middlewares.NewMiddlewares()
	productController := controller.NewProductController()

	app.Route("/product/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			res := map[string]interface{}{
				"mess": "done",
			}
			render.JSON(w, r, res)
		})
		r.Route("/protected", func(protected chi.Router) {
			// protected.Use(jwtauth.Verifier(config.GetJWT()))
			// protected.Use(jwtauth.Authenticator(config.GetJWT()))
			// protected.Use(middlewares.ValidateExpAccessToken())

			protected.Route("/product", func(product chi.Router) {
				product.Post("/", productController.CreateProduct)
				product.Put("/", productController.UpdateProduct)
				product.Delete("/", productController.DeleteProduct)
			})
		})
	})

	log.Println("Sevice h-shop-be-product starting success!")

	return app
}
