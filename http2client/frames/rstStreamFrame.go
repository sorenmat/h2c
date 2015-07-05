package frames

import (
	"encoding/binary"
	"fmt"
)

type ErrorCode uint32

const (
	NO_ERROR            ErrorCode = 0x0
	PROTOCOL_ERROR      ErrorCode = 0x1
	INTERNAL_ERROR      ErrorCode = 0x2
	FLOW_CONTROL_ERROR  ErrorCode = 0x3
	SETTINGS_TIMEOUT    ErrorCode = 0x4
	STREAM_CLOSED       ErrorCode = 0x5
	FRAME_SIZE_ERROR    ErrorCode = 0x6
	REFUSED_STREAM      ErrorCode = 0x7
	CANCEL              ErrorCode = 0x8
	COMPRESSION_ERROR   ErrorCode = 0x9
	CONNECT_ERROR       ErrorCode = 0xa
	ENHANCE_YOUR_CALM   ErrorCode = 0xb
	INADEQUATE_SECURITY ErrorCode = 0xc
	HTTP_1_1_REQUIRED   ErrorCode = 0xd
)

func (e ErrorCode) String() string {
	switch e {
	case NO_ERROR:
		return "NO_ERROR"
	case PROTOCOL_ERROR:
		return "PROTOCOL_ERROR"
	case INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case FLOW_CONTROL_ERROR:
		return "FLOW_CONTROL_ERROR"
	case SETTINGS_TIMEOUT:
		return "SETTINGS_TIMEOUT"
	case STREAM_CLOSED:
		return "STREAM_CLOSED"
	case FRAME_SIZE_ERROR:
		return "FRAME_SIZE_ERROR"
	case REFUSED_STREAM:
		return "REFUSED_STREAM"
	case CANCEL:
		return "CANCEL"
	case COMPRESSION_ERROR:
		return "COMPRESSION_ERROR"
	case CONNECT_ERROR:
		return "CONNECT_ERROR"
	case ENHANCE_YOUR_CALM:
		return "ENHANCE_YOUR_CALM"
	case INADEQUATE_SECURITY:
		return "INADEQUATE_SECURITY"
	case HTTP_1_1_REQUIRED:
		return "HTTP_1_1_REQUIRED"
	default:
		return "UNKNOWN_ERROR"
	}
}

type RstStreamFrame struct {
	StreamId  uint32
	ErrorCode ErrorCode
}

func NewRstStreamFrame(streamId uint32, errorCode ErrorCode) *RstStreamFrame {
	return &RstStreamFrame{
		StreamId:  streamId,
		ErrorCode: errorCode,
	}
}

func DecodeRstStreamFrame(flags byte, streamId uint32, payload []byte, context *DecodingContext) (Frame, error) {
	if len(payload) != 4 {
		return nil, fmt.Errorf("FRAME_SIZE_ERROR: Received RST_STREAM frame of length %v", len(payload))
	}
	return NewRstStreamFrame(streamId, ErrorCode(binary.BigEndian.Uint32(payload))), nil
}

func (f *RstStreamFrame) Type() Type {
	return RST_STREAM_TYPE
}

func (f *RstStreamFrame) Encode(context *EncodingContext) ([]byte, error) {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(f.ErrorCode))
	return result, nil
}
