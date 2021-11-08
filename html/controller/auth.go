package controller

import (
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/errors"
	"github.com/haleyrc/api/html/controller/mw"
	"github.com/haleyrc/api/html/template"
)

func NewAuthController(
	cookies CookieJar,
	errors ErrorFunc,
	renderer Renderer,
	users UserRepository,
) *AuthController {
	c := &AuthController{
		Cookies:     cookies,
		LoginPage:   template.New("fullscreen", "auth/login"),
		Renderer:    renderer,
		ReportError: errors,
		Users:       users,
	}
	return c
}

type AuthController struct {
	Renderer

	Cookies     CookieJar
	ReportError ErrorFunc
	Users       UserRepository

	LoginPage Template
}

type LoginPageData struct {
	Error string
}

func (c *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.LoginPage, LoginPageData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	user, err := c.Users.GetUser(ctx, api.UserQuery{
		Name: r.FormValue("name"),
	})
	if err != nil {
		switch errors.Kind(err) {
		case errors.KindResourceNotFound:
			c.Render(ctx, w, 401, c.LoginPage, LoginPageData{
				Error: "Invalid username or password.",
			})
		default:
			c.ReportError(ctx, err)
			c.Render(ctx, w, 500, c.LoginPage, LoginPageData{
				Error: "Oops, something went wrong!",
			})
		}
		return
	}

	c.Cookies.Set(w, "user", string(user.ID))
	c.Redirect(ctx, w, r, 301, "/library/books")
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := mw.UserFromContext(ctx)
	if user != nil {
		c.Redirect(ctx, w, r, 301, "/library/books")
		return
	}

	c.Render(ctx, w, 200, c.LoginPage, LoginPageData{})
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	c.Cookies.Clear(w, "user")
	c.Renderer.Redirect(r.Context(), w, r, 301, "/")
}
