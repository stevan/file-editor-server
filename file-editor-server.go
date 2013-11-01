package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/stevan/httpapp/component"
	"github.com/stevan/httpapp/middleware"
	"github.com/stevan/httpapp/middleware/auth"
	"github.com/stevan/httpapp/middleware/sessions"

	"github.com/stevan/file-editor-server/controller"
)

const (
	defaultHost = "localhost"
	defaultPort = "8080"
)

func main() {

	var host = flag.String("host", defaultHost, "the hostname")
	var port = flag.String("port", defaultPort, "the port")
	flag.Parse()

	session_state := sessions.NewCookieState("session")
	session_state.Path = "/"

	session_store := sessions.NewMemoryStore()

	urlmap := component.URLMapper()

	urlmap.AddApplication("/", component.Redirect("/editor/"))
	urlmap.AddApplication("/editor/", component.ServeFiles(os.Getenv("DD_EDITOR_ROOT")))
	urlmap.AddApplication("/data/", controller.NewContentController(os.Getenv("DD_DATA_ROOT")))

	http.ListenAndServe(
		*host+":"+*port,
		middleware.HandleSimpleLogging(
			middleware.HandleErrors(
				middleware.HandleSessions(
					middleware.HandleGoogleOAuthAuthentication(
						urlmap,
						auth.CreateGoogleOAuthConfig(
							os.Getenv("DD_OAUTH_CLIENT_ID"),
							os.Getenv("DD_OAUTH_CLIENT_SECRET"),
							fmt.Sprintf(os.Getenv("DD_OAUTH_CALLBACK_FORMAT"), *host, *port),
							os.Getenv("DD_OAUTH_CACHE_FILE"),
						),
					),
					session_state,
					session_store,
				),
			),
		),
	)

}
