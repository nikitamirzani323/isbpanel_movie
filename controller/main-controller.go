package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_movie/config"
)

type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type responseinit struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

const PATH string = config.PATH_API

func Init(c *fiber.Ctx) error {
	hostname := c.Hostname()
	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinit{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
		}).
		Post(PATH + "api/init")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseinit)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"token":  result.Token,
		"time":   time.Since(render_page).String(),
	})
}
func Listmovie(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetAuthToken(token[1]).
		SetResult(response{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
		}).
		Post(PATH + "api/movie")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": result.Status,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
