// Generated by github.com/davyxu/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: core.proto

package gamedef

import (
	"github.com/davyxu/cellnet"
	"reflect"
)

func init() {

	// core.proto
	cellnet.RegisterMessageMeta("gamedef.SessionAccepted", reflect.TypeOf((*SessionAccepted)(nil)).Elem(), 4151444293)
	cellnet.RegisterMessageMeta("gamedef.SessionConnected", reflect.TypeOf((*SessionConnected)(nil)).Elem(), 459632229)
	cellnet.RegisterMessageMeta("gamedef.SessionAcceptFailed", reflect.TypeOf((*SessionAcceptFailed)(nil)).Elem(), 399042283)
	cellnet.RegisterMessageMeta("gamedef.SessionConnectFailed", reflect.TypeOf((*SessionConnectFailed)(nil)).Elem(), 1644962508)
	cellnet.RegisterMessageMeta("gamedef.SessionClosed", reflect.TypeOf((*SessionClosed)(nil)).Elem(), 1412646790)
	cellnet.RegisterMessageMeta("gamedef.RemoteCallACK", reflect.TypeOf((*RemoteCallACK)(nil)).Elem(), 1020080612)
	cellnet.RegisterMessageMeta("gamedef.TestEchoACK", reflect.TypeOf((*TestEchoACK)(nil)).Elem(), 1899977859)

}