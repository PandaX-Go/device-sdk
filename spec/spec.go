package spec

type Topic int

const (
	_rawTopic       = "v1/devices/me/raw"
	_telemetryTopic = "v1/devices/me/telemetry"
	_attrTopic      = "v1/devices/me/attributes"

	_gatewaAttrTopic       = "v1/gateway/attributes"
	_gatewayTelemetryTopic = "v1/gateway/telemetry"

	_rpcTopicReQ  = "v1/devices/me/rpc/request/+"
	_rpcTopicResp = "v1/devices/me/rpc/response/%s"
)

const (
	RawTopic Topic = 1 << iota
	TelemetryTopic
	AttributeTopic
	GatewayTelemetryTopic
	GatewayAttributeTopic
	RpcReqTopic
	RpcRespTopic
)

func (t Topic) String() string {
	switch t {
	case RawTopic:
		return _rawTopic
	case TelemetryTopic:
		return _telemetryTopic
	case AttributeTopic:
		return _attrTopic
	case GatewayTelemetryTopic:
		return _gatewayTelemetryTopic
	case GatewayAttributeTopic:
		return _gatewaAttrTopic
	case RpcReqTopic:
		return _rpcTopicReQ
	case RpcRespTopic:
		return _rpcTopicResp
	}
	return ""
}
