package clover

type IMiddleware func(HandlerFunc) HandlerFunc
