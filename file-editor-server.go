package main

import (
    "os"
    "flag"
    "net/http"

    "code.google.com/p/goauth2/oauth"

    "github.com/stevan/httpapp/component"
    "github.com/stevan/httpapp/middleware"
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

    var oauth_config = &oauth.Config{
        ClientId:     os.Getenv("DD_OAUTH_CLIENT_ID"),
        ClientSecret: os.Getenv("DD_OAUTH_CLIENT_SECRET"),
        AuthURL:      "https://accounts.google.com/o/oauth2/auth",
        TokenURL:     "https://accounts.google.com/o/oauth2/token",
        RedirectURL:  "http://" + *host + ":" + *port + "/oauth2callback",
        Scope:        "https://www.googleapis.com/auth/userinfo.email",
        TokenCache:   oauth.CacheFile(os.Getenv("DD_OAUTH_CACHE_FILE")),
    }

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
                        oauth_config,
                    ),
                    session_state,
                    session_store,
                ),
            ),
        ),
    )

}

