# File Editor Server

This is a server written in Go that is meant to serve as
the backend to a simple online file editor tool. It
provides a set of endpoints through which a client-side
Javascript based application might operate, and can also
serve that client side app as well.

## Dependencies

This depends on a small net/http add-on that I wrote
called httpapp, which enables the composing of HTTP
applications written using the middleware style
popularized by WSGI (python), Rack (ruby) and Plack
(perl). You can install it in your `$GOPATH` by
doing:

    go get github.com/stevan/httpapp


