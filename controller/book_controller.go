package controller

import (
	"Manu/data/request"
	"Manu/data/response"
	"Manu/helper"
	"Manu/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct{
	BookService service.BookService
}

func NewBookController(bookService service.BookService) * BookController{
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create( writer http.ResponseWriter, requests *http.Request, params httprouter.Params){
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestedBody(requests, &bookCreateRequest)

	controller.BookService.Create(requests.Context(), bookCreateRequest)
	WebResponse := response.WebResponse{
		Code: 200, 
		Status: "Ok",
		Data: nil,
	}
	helper.WriteResponseBody(writer, WebResponse)
}



func (controller *BookController) Update( writer http.ResponseWriter, requests *http.Request, params httprouter.Params){
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestedBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	controller.BookService.Update(requests.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}


func (controller *BookController)  Delete( writer http.ResponseWriter, requests *http.Request, params httprouter.Params){
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(requests.Context(), id)
	webResposne := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}

	helper.WriteResponseBody(writer, webResposne)
}

func (controller *BookController)  FindAll( writer http.ResponseWriter, requests *http.Request, params httprouter.Params){
	result := controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController)  FindById( writer http.ResponseWriter, requests *http.Request, params httprouter.Params){
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	result := controller.BookService.FindById(requests.Context(),id)
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
