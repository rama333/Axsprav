package response

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type APIResponse struct {
	Error   int32  `json:"error"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func ResponseOkRequest(message string, c *gin.Context) {
	//resp := APIResponse{
	//	Code:    http.StatusOK,
	//	Error:   0,
	//	Message: message,
	//}
	//
	//
	//c.JSON(http.StatusOK, resp)

	data := message
	tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
	tmpl.Execute(c.Writer, data)

	//c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//	"title": "Users",
	//})

}

func ResponseBadRequest(message string, c *gin.Context) {

	//resp := APIResponse{
	//	Code:    http.StatusBadRequest,
	//	Error:   1,
	//	Message: message,
	//}
	//
	//reqBodyBytes := new(bytes.Buffer)
	//json.NewEncoder(reqBodyBytes).Encode(resp)
	//
	//c.Data(http.StatusOK, "text/html; charset=utf-8", reqBodyBytes.Bytes())

	data := message
	tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
	tmpl.Execute(c.Writer, data)

}

func ResponseInternalServerError(message string, c *gin.Context) {

	//resp := APIResponse{
	//	Code:    http.StatusInternalServerError,
	//	Error:   1,
	//	Message: message,
	//}
	//
	//c.JSON(http.StatusOK, resp)
	//reqBodyBytes := new(bytes.Buffer)
	//json.NewEncoder(reqBodyBytes).Encode(resp)
	//
	//c.Data(http.StatusOK, "text/html; charset=utf-8", reqBodyBytes.Bytes())

	data := message
	tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
	tmpl.Execute(c.Writer, data)

}

func ResponseStatusNotFound(message string, c *gin.Context) {

	resp := APIResponse{
		Code:    http.StatusNotFound,
		Error:   1,
		Message: message,
	}

	c.JSON(http.StatusOK, resp)
}
