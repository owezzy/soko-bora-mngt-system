package amotel

import (
	"fmt"
	"github.com/owezzy/soko-bora-mngt-system/internal/ddd"
	"strconv"

	"go.opentelemetry.io/otel/propagation"
)

type MetadataCarrier ddd.Metadata

var _ propagation.TextMapCarrier = (*MetadataCarrier)(nil)

func (mc MetadataCarrier) Get(key string) string {
	switch v := ddd.Metadata(mc).Get(key).(type) {
	case nil:
		return ""
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (mc MetadataCarrier) Set(key, value string) {
	ddd.Metadata(mc).Set(key, value)
}

func (mc MetadataCarrier) Keys() []string {
	return ddd.Metadata(mc).Keys()
}
