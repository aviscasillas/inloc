package controllers

import (
	"github.com/revel/revel"
	"github.com/wuicorp/inloc/app/models"
)

type Flags struct {
	*revel.Controller
}

var flags = []models.Flag{
	models.Flag{1, 0.001, 0.002, 10},
	models.Flag{1, 0.001, 0.002, 10},
	models.Flag{1, 0.001, 0.002, 10},
}

func (c Flags) List() revel.Result {
	return c.RenderJson(flags)
}

func (c Flags) Show(id int) revel.Result {
	var res models.Flag

	for _, flag := range flags {
		if flag.Id == id {
			res = flag
		}
	}

	if res.Id == 0 {
		return c.NotFound("Could not find flag")
	}

	return c.RenderJson(res)
}
