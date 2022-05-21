package sfu

import (
	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3"
)

type InterceptorRegistryFactoryFactory interface {
	NewInterceptorRegistryFactory(sid, uid string) InterceptorRegistryFactory
}

type InterceptorRegistryFactory interface {
	NewInterceptorRegistry(*webrtc.MediaEngine, WebRTCTransportConfig) *interceptor.Registry
}
