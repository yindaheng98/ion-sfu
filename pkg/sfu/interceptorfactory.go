package sfu

import (
	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3"
)

type InterceptorRegistryFactoryBuilder func(sid, uid string) InterceptorRegistryFactory

type InterceptorRegistryFactory func(*webrtc.MediaEngine, WebRTCTransportConfig) *interceptor.Registry

func WithPubInterceptorRegistryFactoryBuilder(pub InterceptorRegistryFactoryBuilder) func(*PeerLocal) {
	return func(p *PeerLocal) {
		p.pubInterceptorRegistryFactoryBuilder = pub
	}
}

func WithSubInterceptorRegistryFactoryBuilder(sub InterceptorRegistryFactoryBuilder) func(*PeerLocal) {
	return func(p *PeerLocal) {
		p.subInterceptorRegistryFactoryBuilder = sub
	}
}
