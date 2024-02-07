package server

import "sv.notify/pb/protos/pb/notify"

type NotifyHandler struct {
	notify.UnimplementedNotifyServiceServer
}

func NewNotifyHandler() *NotifyHandler {
	return &NotifyHandler{}
}

var _ notify.NotifyServiceServer = NotifyHandler{}
