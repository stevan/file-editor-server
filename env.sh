
export DD_DATA_ROOT=$HOME/Desktop/dd_test/cms
export DD_EDITOR_ROOT=$HOME/Desktop/dd_test/static
export DD_OAUTH_CLIENT_ID=390328890773.apps.googleusercontent.com
export DD_OAUTH_CLIENT_SECRET=G27ypVUsllOiyqPmWd9ekZ-t
export DD_OAUTH_CACHE_FILE=./request.json

NAMESPACE="github.com/stevan/file-editor-server"
alias run-tests="go test $NAMESPACE/controller/contentserver \
                         $NAMESPACE/model/directory       \
                         $NAMESPACE/model/file"