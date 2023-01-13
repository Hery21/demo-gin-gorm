package testutils

//import (
//	"encoding/json"
//	" hery-ciaputra/demo-gin/mocks"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
//func ServeReq(opts *server.RouterOpts, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
//	router := server.NewRouter(opts)
//	rec := httptest.NewRecorder()
//	router.ServeHTTP(rec, req)
//	return router, rec
//}
//
//func MakeRequestBody(dto interface{}) *strings.Reader {
//	payload, _ := json.Marshal(dto)
//	return strings.NewReader(string(payload))
//}
//
//// sample usage in test
//t.Run("should return code 200 with list of books when no error", func(t *testing.T) {
//	mockBService := new(mocks.BookService)
//	mockBService.On("GetBooks").Return(books, nil)
//	opts := &server.RouterOpts{
//		BookService: mockBService,
//	}
//	expectedRes, _ := json.Marshal(books)
//
//	req, _ := http.NewRequest("GET", "/books", nil)
//	_, rec := testutils.ServeReq(opts, req)
//
//	assert.Equal(t, 200, rec.Code)
//	assert.Equal(t, string(expectedRes), rec.Body.String())
//})
