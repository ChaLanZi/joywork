package auth

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func SetInternalHeaders(externalReq *http.Request, internalHeaders http.Header) {
	ProxyHeaders(externalReq.Header, internalHeaders)
	authorization := AuthorizationAnonymousWeb
	uuid, support, err := getSession(externalReq)
	if err == nil {
		if support {
			authorization = AuthorizationSupportUser
		} else {
			authorization = AuthorizationAuthenticatedUser
		}
		internalHeaders.Set(currentUserHeader, uuid)
	}
	internalHeaders.Set(AuthorizationHeader, authorization)
	return
}

func ProxyHeaders(from http.Header, to http.Header) {
	for k, v := range from {
		for _, x := range v {
			to.Add(k, x)
		}
	}
}

func GetCurrentUserUUIDFromMetadata(data metadata.MD) (uuid string, err error) {
	res, ok := data[currentUserMetadata]
	if !ok || len(res) == 0 {
		err = fmt.Errorf(" User not authenticated")
		return
	}
	uuid = res[0]
	return
}

func GetCurrentUserUUIDFromHeader(data http.Header) (uuid string, err error) {
	res, ok := data[currentUserHeader]
	if !ok || len(res) == 0 {
		err = fmt.Errorf(" User not anthenticated")
		return
	}
	uuid = res[0]
	return
}
