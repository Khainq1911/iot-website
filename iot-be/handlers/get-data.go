package handlers

import (
	"fmt"
	"net/http"
	"web-ivsr-be/model"
	"web-ivsr-be/repository"

	"github.com/labstack/echo/v4"
)

type SiteHandler struct {
	Repo repository.Repo
}

func (u *SiteHandler) GetData(ctx echo.Context) error {
	result, err := u.Repo.GetDataRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error",
			"error":   err,
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "succesful",
		"objects": result,
	})
}

func (u *SiteHandler) PutData(ctx echo.Context) error {
	data := model.UpdateStatus{}
	node_id := ctx.Param("node_id")

	if err := ctx.Bind(&data); err != nil {
		fmt.Println("error o bind")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err := u.Repo.PutDataRepo(ctx.Request().Context(), node_id, data.Status)
	if err != nil {
		fmt.Println("error o put request")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "succesful",
	})
}

func (u *SiteHandler) GetDataId(ctx echo.Context) error {
	node_id := ctx.Param("node_id")

	result, err := u.Repo.GetDataIdRepo(ctx.Request().Context(), node_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error",
			"object":  err,
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "succesful",
		"object":  result,
	})
}
