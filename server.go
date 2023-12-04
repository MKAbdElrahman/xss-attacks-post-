package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type ReviewData struct {
	Username string
	Rating   int
	Comment  string
}

func main() {
	e := echo.New()

	e.GET("/review", GetReviewForm)
	e.POST("/review", SubmitReview)

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":1323"))
}

func GetReviewForm(c echo.Context) error {
	return c.Render(http.StatusOK, "review_form.html", nil)
}

func SubmitReview(c echo.Context) error {
	// Retrieve form values
	username := c.FormValue("username")
	ratingStr := c.FormValue("rating")
	comment := c.FormValue("comment")

	// Convert rating to int
	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid rating")
	}

	// Create a struct to hold the review data
	reviewData := ReviewData{
		Username: username,
		Rating:   rating,
		Comment:  comment,
	}

	// Render a confirmation page with the submitted review data
	return c.Render(http.StatusOK, "review_confirmation.html", reviewData)
}
