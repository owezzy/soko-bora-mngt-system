package stores

//go:generate buf generate

//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c storesclient -m storesclient/models --with-flatten=remove-unused
