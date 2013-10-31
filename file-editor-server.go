package main

import (
    "os"
    "flag"
    "net/http"

    "github.com/stevan/httpapp/component"
    "github.com/stevan/httpapp/middleware"
    "github.com/stevan/httpapp/middleware/auth"
    "github.com/stevan/httpapp/middleware/sessions"

    "github.com/stevan/file-editor-server/controller"
)

const(
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

    urlmap.AddApplication("/editor/", component.ServeFiles(os.Getenv("DD_EDITOR_ROOT")))
    urlmap.AddApplication("/data/", controller.NewContentController(os.Getenv("DD_DATA_ROOT")))

    http.ListenAndServe(
        *host + ":" + *port,
        middleware.HandleSimpleLogging(
            middleware.HandleErrors(
                middleware.HandleSessions(
                    middleware.HandleGoogleOAuthAuthentication(
                        urlmap,
                        auth.CreateGoogleOAuthConfig(
                            os.Getenv("DD_OAUTH_CLIENT_ID"),
                            os.Getenv("DD_OAUTH_CLIENT_SECRET"),
                            "http://" + *host + ":" + *port + "/oauth2callback",
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

