package handler

import (
	"book-rest-api-showcase/helper"
	"book-rest-api-showcase/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
//	@Summary		Create Book
//	@Description	Create Book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Book
//	@Failure		404	{object}	model.Book
//	@Failure		500	{object}	model.Book
//	@Router			/books [post]

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// GetAllBooks godoc
//	@Summary		Get list all of the books
//	@Description	Get list all of the books
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Success		201	"Created"
//	@Router			/books [get]

func (h HttpServer) GetAllBooks(c *gin.Context) {
	res, err := h.app.GetAllBooks()
	if err != nil {
		if err.Error() == helper.ErrorNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

// GetBookByID godoc
//	@Summary		Get detail book for fiven Id
//	@Description	Get Details of car corresponding to Input ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			ID	path		int	true	"ID of the book"
//	@Success		200	{object}	model.Book
//	@Router			/books/{Id} [get]

func (h HttpServer) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetBookByID(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrorNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

// UpdateBoopk godoc
//	@Summary		Update book identified by the given ID
//	@Description	Update book identified by ID
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID book to update"
//	@Success		200	{object}	model.Book
//	@Router			/books/{id} {put}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	in.ID = idInt
	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

// DeleteBook godoc
//	@Summary		Delete book identified with id
//	@Description	Delete the book with id
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"ID book to be deleted"
//	@Success		200	"Successfully deleted book"
//	@Routeer		/books/{id} [delete]

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		if err.Error() == helper.ErrorNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Book deleted successfully")
}
