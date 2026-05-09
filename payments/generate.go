package payments

//go:generate buf generate

//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c paymentsclient -m paymentsclient/models --with-flatten=remove-unused
