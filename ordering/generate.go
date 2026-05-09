package ordering

//go:generate buf generate

//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c orderingclient -m orderingclient/models --with-flatten=remove-unused
