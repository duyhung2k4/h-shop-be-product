package router

import (
	"app/config"
	"app/controller"
	"app/middlewares"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

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

	middlewares := middlewares.NewMiddlewares()
	productController := controller.NewProductController()
	typeInWarehouseController := controller.NewTypeInWarehouseController()
	warehouseController := controller.NewWarehouseController()

	app.Route("/product/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			res := map[string]interface{}{
				"mess": "done",
			}
			render.JSON(w, r, res)
		})
		r.Route("/public", func(public chi.Router) {
			public.Route("/product", func(product chi.Router) {
				product.Get("/", productController.GetProductById)
			})
		})
		r.Route("/protected", func(protected chi.Router) {
			protected.Use(jwtauth.Verifier(config.GetJWT()))
			protected.Use(jwtauth.Authenticator(config.GetJWT()))
			protected.Use(middlewares.ValidateExpAccessToken())

			protected.Route("/product", func(product chi.Router) {
				product.Get("/all", productController.GetProductByProfileId)
				product.Get("/detail", productController.GetProductById)
				product.Post("/", productController.CreateProduct)
				product.Put("/", productController.UpdateProduct)
				product.Delete("/", productController.DeleteProduct)

				product.Post("/heart", productController.Heart)
				product.Get("/is-heart", productController.IsHeart)
				product.Get("/is-heart-list", productController.GetHeart)

				product.Post("/cart", productController.Cart)
				product.Get("/is-cart", productController.IsCart)
				product.Get("/is-cart-list", productController.GetCart)
			})

			protected.Route("/warehouse", func(warehouse chi.Router) {
				warehouse.Get("/", warehouseController.GetWarehouseByProductId)
			})

			protected.Route("/type-in-warehouse", func(typeInWarehouse chi.Router) {
				typeInWarehouse.Get("/", typeInWarehouseController.GetTypeInWarehouseByProductId)
				typeInWarehouse.Post("/", typeInWarehouseController.InsertTypeInWarehouse)
				typeInWarehouse.Put("/", typeInWarehouseController.UpdateTypeInWarehouse)
				typeInWarehouse.Delete("/", typeInWarehouseController.DeleteTypeInWarehouse)
			})
		})
	})

	log.Println("Sevice h-shop-be-product starting success!")

	return app
}
