// Copyright (c) 2021 Amirhossein Movahedi (@qolzam)
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package router

import (
	"example.com/function/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	// authHMACMiddleware := func(hmacWithCookie bool) func(*fiber.Ctx) error {
	// 	var Next func(c *fiber.Ctx) bool
	// 	if hmacWithCookie {
	// 		Next = func(c *fiber.Ctx) bool {
	// 			if c.Get(types.HeaderHMACAuthenticate) != "" {
	// 				return false
	// 			}
	// 			return true
	// 		}
	// 	}
	// 	return authhmac.New(authhmac.Config{
	// 		Next:          Next,
	// 		PayloadSecret: *config.AppConfig.PayloadSecret,
	// 	})
	// }

	// authCookieMiddleware := func(hmacWithCookie bool) func(*fiber.Ctx) error {
	// 	var Next func(c *fiber.Ctx) bool
	// 	if hmacWithCookie {
	// 		Next = func(c *fiber.Ctx) bool {
	// 			if c.Get(types.HeaderHMACAuthenticate) != "" {
	// 				return true
	// 			}
	// 			return false
	// 		}
	// 	}
	// 	return authcookie.New(authcookie.Config{
	// 		Next:         Next,
	// 		JWTSecretKey: []byte(*config.AppConfig.PublicKey),
	// 	})
	// }

	// hmacCookieHandlers := []func(*fiber.Ctx) error{authHMACMiddleware(true), authCookieMiddleware(true)}

	// Routers
	app.Post("/create-vote", handlers.CreateVoteHandle)
	app.Put("/update-vote", handlers.UpdateVoteHandle)
	app.Delete("/id/:voteId", handlers.DeleteVoteHandle)
	app.Delete("/post/:postId", handlers.DeleteVoteByPostIdHandle)
	app.Get("/get-votes-by-post-id", handlers.GetVotesByPostIdHandle)
	app.Get("/:voteId", handlers.GetVoteHandle)
}
