package amqp

import "time"

const (
    reconnectDelay  = 5 * time.Second // 连接断开后多久重连
    resendDelay     = 5 * time.Second // 消息发送失败后，多久重发
    resendTime      = 3               // 消息重发次数
    ContentTypeJson = "application/json"
)
