package main

import (
	"bytes"
	"net/http"
	"net/url"
	"os"

	"text/template"

	"github.com/gin-gonic/gin"
)

var signupTemplate = template.Must(template.New("signup").Parse(`
	<!DOCTYPE html>
	<html>
		<head>
		<meta charset="UTF-8">
		<title>Homeautomation auth</title>
		</head>
		<body>
		Copy homeautomation client id found from your logs here<br />
		<form action="{{.redirect_uri}}" method="get">
			{{ range $key, $values := .params }}
				{{ range $value := $values }}
					<input type="hidden" name="{{$key}}" value="{{$value}}" />
				{{ end }}
			{{ end }}
			<input type="text" name="code" value="" />
			
			<input type="submit" value="OK!"/>
		</form>
		</body>
	</html>
	`))

func oauthAuthorize(c *gin.Context) {
	redirect := c.Query("redirect_uri")
	parsedRedirect, _ := url.Parse(redirect)

	query := parsedRedirect.Query()
	query.Add("state", c.Query("state"))

	b := &bytes.Buffer{}
	signupTemplate.Execute(b, map[string]interface{}{"redirect_uri": redirect, "params": query})

	c.Data(http.StatusOK, "text/html; charset=utf-8", b.Bytes())
}

func oauthAccessToken(c *gin.Context) {
	code := c.Query("code")

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": code,
		"token_type":   "bearer",
	})
}

func main() {
	r := gin.Default()
	oauth := r.Group("/oauth")
	{
		oauth.GET("/authorize", oauthAuthorize)
		oauth.POST("/access_token", oauthAccessToken)
	}
	r.Run(":" + os.Getenv("PORT"))
}
