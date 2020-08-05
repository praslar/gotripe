package handler

import (
	"net/http"

	"github.com/gotripe/utils"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/price"
)

type (
	// Config hold MongoDB configuration information
	Config struct {
		PubKey string `envconfig:"STRIPE_PUBLISHABLE_KEY" json:"publishableKey" default:"pk_test_51H7kvALDFMQnQqVAwSbb6Q7PZOjuYEPHH4L1b3uoSLCBX6yEBwggv9BmGegOxRL7IWWBRSHketVDc8cT2GWkaTS800sRGP9dFJ"`
		SecKey string `envconfig:"STRIPE_SECRET_KEY" json:"secret_key" default:"sk_test_51H7kvALDFMQnQqVAAfjHjmW0RqEaKIptRJnwAJanMAuHLjP1zjBAjxsQWCym0NiRaiVRcuKhJSGOEh2WhZNEAqDF004RF22ANb"`
	}
)

func (h *Handler) GetPubKey(c echo.Context) error {
	var conf Config
	utils.Load(&conf)
	return c.JSON(http.StatusCreated, &conf.PubKey)
}

func (h *Handler) GetPlan(c echo.Context) error {
	var conf Config
	utils.Load(&conf)

	stripe.Key = conf.SecKey
	params := &stripe.PriceListParams{
		LookupKeys: []*string{
			stripe.String("Premium"),
			stripe.String("Standard"),
			stripe.String("Low"),
		},
	}
	i := price.List(params)
	res := make([]*stripe.Price, 0)
	for i.Next() {
		res = append(res, i.Price())
	}
	return c.JSON(http.StatusCreated, &res)
}
