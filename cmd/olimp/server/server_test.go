package server

/*import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type tFakeEngine struct{}

func (fe tFakeEngine) QueryInstitutions(instType, regCode string, valFields, reqFields map[string]string) ([]map[string]string, error) {
	return make([]map[string]string, 0), nil
}

func createFakeEngine() tFakeEngine {
	return struct{}{}
}

func TestGetInstList(t *testing.T) {
	cases := []struct {
		have struct{ it, rc, params string }
		want struct{ httpStatus int }
	}{
		{
			struct{ it, rc, params string }{"3", "80", ""},
			struct{ httpStatus int }{http.StatusOK},
		},
		{
			struct{ it, rc, params string }{"", "", ""},
			struct{ httpStatus int }{http.StatusBadRequest},
		},
	}
	for _, c := range cases {
		url := "http://example.com/api/v1/register/" + c.have.it + "/" + c.have.rc + c.have.params
		ctx := context.WithValue(context.Background(), "inst", c.have.it)
		ctx = context.WithValue(ctx, "reg", c.have.rc)
		req := httptest.NewRequest("GET", url, nil) //
		req = req.WithContext(ctx)
		w := httptest.NewRecorder()

		serv := TServer{
			Addr:   "http://example.com",
			Engine: createFakeEngine(),
		}
		fmt.Println(ctx)
		fmt.Println(req)
		serv.getInstList(w, req)
		//fmt.Println(w)

		fmt.Println(w.Code)
		fmt.Println(c.want.httpStatus)
		if w.Code != c.want.httpStatus {
			t.Fail()
		}
	}
}
*/
