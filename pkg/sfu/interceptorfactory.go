package sfu

import (
	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3"
)

type PeerInterceptorFactory interface {
	NewInterceptorFactory(sid, uid string) InterceptorFactory
}

type InterceptorFactory interface {
	NewInterceptorRegistry(*webrtc.MediaEngine, WebRTCTransportConfig) *interceptor.Registry
}
