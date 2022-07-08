package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/bootstrap"
	"github.com/lezi-wiki/lezi-api/model"
	"github.com/lezi-wiki/lezi-api/pkg/response"
	"github.com/lezi-wiki/lezi-api/services"
	"math/rand"
)

func GlobalHandler(c *gin.Context) {
	var err error

	ns := c.Query("ns")
	speaker := c.Query("speaker")

	var arr = bootstrap.TextData

	if ns != "" {
		arr, err = services.GetTextByNamespace(ns)
		if err != nil {
			response.NotFoundError(c)
			return
		}
	}

	var newArr = arr

	if speaker != "" {
		newArr = []model.TextData{}

		for _, d := range arr {
			if d.Speaker == speaker {
				newArr = append(newArr, d)
			}
		}
	}

	text := newArr[rand.Intn(len(newArr))]

	response.JsonData(c, text)
}
