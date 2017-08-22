package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestEspaciosHuerfanos(t *testing.T) {
	query := "Id:1"
	data := strings.NewReader(query)
	r, _ := http.NewRequest("GET", "http://10.20.0.254/oikos_api/v1/espacio_fisico/EspacioFisicosHuerfanos/", data)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	fmt.Print(w)
	Convey("GET de espacios físicos que son huerfanos\n", t, func() {
		Convey("El código de estado debe ser 200", func() {
			So(w.Code, ShouldEqual, 404)
		})

		Convey("El resultado no deberia ser vacio", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestArbolDependencias(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://10.20.0.254/oikos_api/v1/dependencia_padre/ArbolDependencias/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	fmt.Print(w)
	Convey("Get árbol de dependencias\n", t, func() {
		Convey("El codigo de estado debe ser 200", func() {
			So(w.Code, ShouldEqual, 404)
		})

		Convey("El resultado no deberia ser vacio", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
