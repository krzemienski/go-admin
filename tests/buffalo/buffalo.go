package buffalo

import (
	_ "github.com/chenhg5/go-admin/adapter/buffalo"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/gobuffalo/buffalo"
	"net/http"
	"os"
)

func NewBuffaloHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	bu.GET("/admin", func(ctx buffalo.Context) error {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
		return nil
	})

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
