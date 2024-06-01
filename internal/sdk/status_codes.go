package sdk

import "net/http"

var statusCodes = map[int]string{
	// 1xx
	http.StatusContinue:           "continue",
	http.StatusSwitchingProtocols: "switching protocols",
	http.StatusProcessing:         "processing",

	// 2xx
	http.StatusOK:                   "ok",
	http.StatusCreated:              "created",
	http.StatusAccepted:             "accepted",
	http.StatusNonAuthoritativeInfo: "non-authoritative information",
	http.StatusNoContent:            "no content",
	http.StatusResetContent:         "reset content",
	http.StatusPartialContent:       "partial content",
	http.StatusMultiStatus:          "multi-status",
	http.StatusAlreadyReported:      "already reported",
	http.StatusIMUsed:               "im used",

	// 3xx
	http.StatusMultipleChoices:   "multiple choices",
	http.StatusMovedPermanently:  "moved permanently",
	http.StatusFound:             "found",
	http.StatusSeeOther:          "see other",
	http.StatusNotModified:       "not modified",
	http.StatusUseProxy:          "use proxy",
	http.StatusTemporaryRedirect: "temporary redirect",
	http.StatusPermanentRedirect: "permanent redirect",

	// 4xx
	http.StatusBadRequest:                   "bad request",
	http.StatusUnauthorized:                 "unauthorized",
	http.StatusPaymentRequired:              "payment required",
	http.StatusForbidden:                    "forbidden",
	http.StatusNotFound:                     "not found",
	http.StatusMethodNotAllowed:             "method not allowed",
	http.StatusNotAcceptable:                "not acceptable",
	http.StatusProxyAuthRequired:            "proxy authentication required",
	http.StatusRequestTimeout:               "request timeout",
	http.StatusConflict:                     "conflict",
	http.StatusGone:                         "gone",
	http.StatusLengthRequired:               "length required",
	http.StatusPreconditionFailed:           "precondition failed",
	http.StatusRequestEntityTooLarge:        "request entity too large",
	http.StatusRequestURITooLong:            "request-uri too long",
	http.StatusUnsupportedMediaType:         "unsupported media type",
	http.StatusRequestedRangeNotSatisfiable: "requested range not satisfiable",
	http.StatusExpectationFailed:            "expectation failed",
	http.StatusTeapot:                       "i'm a teapot",
	http.StatusMisdirectedRequest:           "misdirected request",
	http.StatusUnprocessableEntity:          "unprocessable entity",
	http.StatusLocked:                       "locked",
	http.StatusFailedDependency:             "failed dependency",
	http.StatusTooEarly:                     "too early",
	http.StatusUpgradeRequired:              "upgrade required",
	http.StatusPreconditionRequired:         "precondition required",
	http.StatusTooManyRequests:              "too many requests",
	http.StatusRequestHeaderFieldsTooLarge:  "request header fields too large",
	http.StatusUnavailableForLegalReasons:   "unavailable for legal reasons",

	// 5xx
	http.StatusInternalServerError:           "internal server error",
	http.StatusNotImplemented:                "not implemented",
	http.StatusBadGateway:                    "bad gateway",
	http.StatusServiceUnavailable:            "service unavailable",
	http.StatusGatewayTimeout:                "gateway timeout",
	http.StatusHTTPVersionNotSupported:       "http version not supported",
	http.StatusVariantAlsoNegotiates:         "variant also negotiates",
	http.StatusInsufficientStorage:           "insufficient storage",
	http.StatusLoopDetected:                  "loop detected",
	http.StatusNotExtended:                   "not extended",
	http.StatusNetworkAuthenticationRequired: "network authentication required",
}
