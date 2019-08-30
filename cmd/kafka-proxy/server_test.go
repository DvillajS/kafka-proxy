package server

import (
	"github.com/grepplabs/kafka-proxy/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupBootstrapServersMappingTest() {
	Server.ResetFlags()
	c = new(config.Config)
	initFlags()
	_ = os.Setenv("BOOTSTRAP_SERVER_MAPPING", "")
	_ = os.Setenv("EXTERNAL_SERVER_MAPPING", "")
	_ = os.Setenv("DIAL_ADDRESS_MAPPING", "")
}

func TestBootstrapServersMappingFromFlags(t *testing.T) {
	setupBootstrapServersMappingTest()

	args := []string{"cobra.test",
		"--bootstrap-server-mapping", "192.168.99.100:32401,0.0.0.0:32401",
		"--bootstrap-server-mapping", "192.168.99.100:32402,0.0.0.0:32402",
		"--bootstrap-server-mapping", "kafka-2.example.com:9092,0.0.0.0:32403,kafka-2.grepplabs.com:9092",
	}

	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Nil(err)
	a.Len(c.Proxy.BootstrapServers, 3)

	a.Equal(c.Proxy.BootstrapServers[0].BrokerAddress, "192.168.99.100:32401")
	a.Equal(c.Proxy.BootstrapServers[0].ListenerAddress, "0.0.0.0:32401")
	a.Equal(c.Proxy.BootstrapServers[0].AdvertisedAddress, "0.0.0.0:32401")

	a.Equal(c.Proxy.BootstrapServers[1].BrokerAddress, "192.168.99.100:32402")
	a.Equal(c.Proxy.BootstrapServers[1].ListenerAddress, "0.0.0.0:32402")
	a.Equal(c.Proxy.BootstrapServers[1].AdvertisedAddress, "0.0.0.0:32402")

	a.Equal(c.Proxy.BootstrapServers[2].BrokerAddress, "kafka-2.example.com:9092")
	a.Equal(c.Proxy.BootstrapServers[2].ListenerAddress, "0.0.0.0:32403")
	a.Equal(c.Proxy.BootstrapServers[2].AdvertisedAddress, "kafka-2.grepplabs.com:9092")

}

func TestDialMappingFromFlags(t *testing.T) {
	setupBootstrapServersMappingTest()

	args := []string{"cobra.test",
		"--bootstrap-server-mapping", "192.168.99.100:32401,0.0.0.0:32401",
		"--dial-address-mapping", "service-kafka-0.service-kafka-headless.service:9092,0.0.0.0:19092",
		"--dial-address-mapping", "192.168.99.100:32402,0.0.0.0:32402",
	}

	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Nil(err)
	a.Len(c.Proxy.DialAddressMappings, 2)

	a.Equal(c.Proxy.DialAddressMappings[0].SourceAddress, "service-kafka-0.service-kafka-headless.service:9092")
	a.Equal(c.Proxy.DialAddressMappings[0].DestinationAddress, "0.0.0.0:19092")

	a.Equal(c.Proxy.DialAddressMappings[1].SourceAddress, "192.168.99.100:32402")
	a.Equal(c.Proxy.DialAddressMappings[1].DestinationAddress, "0.0.0.0:32402")
}
func TestBootstrapServersMappingFromEnv(t *testing.T) {
	setupBootstrapServersMappingTest()

	_ = os.Setenv("BOOTSTRAP_SERVER_MAPPING", "192.168.99.100:32404,0.0.0.0:32404 kafka-5.example.com:9092,0.0.0.0:32405,kafka-5.grepplabs.com:9092")

	var args []string
	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Nil(err)
	a.Len(c.Proxy.BootstrapServers, 2)

	a.Equal(c.Proxy.BootstrapServers[0].BrokerAddress, "192.168.99.100:32404")
	a.Equal(c.Proxy.BootstrapServers[0].ListenerAddress, "0.0.0.0:32404")
	a.Equal(c.Proxy.BootstrapServers[0].AdvertisedAddress, "0.0.0.0:32404")

	a.Equal(c.Proxy.BootstrapServers[1].BrokerAddress, "kafka-5.example.com:9092")
	a.Equal(c.Proxy.BootstrapServers[1].ListenerAddress, "0.0.0.0:32405")
	a.Equal(c.Proxy.BootstrapServers[1].AdvertisedAddress, "kafka-5.grepplabs.com:9092")

}

func TestEmptyBootstrapServersMapping(t *testing.T) {
	setupBootstrapServersMappingTest()

	var args []string
	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Error(err, "list of bootstrap-server-mapping must not be empty")
}

func TestBootstrapServersMappingFromEnvWithWhiteSpaces(t *testing.T) {
	setupBootstrapServersMappingTest()

	_ = os.Setenv("BOOTSTRAP_SERVER_MAPPING", "   192.168.99.100:32404,0.0.0.0:32404   kafka-5.example.com:9092,0.0.0.0:32405,kafka-5.grepplabs.com:9092    ")

	var args []string
	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Nil(err)
	a.Len(c.Proxy.BootstrapServers, 2)

	a.Equal(c.Proxy.BootstrapServers[0].BrokerAddress, "192.168.99.100:32404")
	a.Equal(c.Proxy.BootstrapServers[0].ListenerAddress, "0.0.0.0:32404")
	a.Equal(c.Proxy.BootstrapServers[0].AdvertisedAddress, "0.0.0.0:32404")

	a.Equal(c.Proxy.BootstrapServers[1].BrokerAddress, "kafka-5.example.com:9092")
	a.Equal(c.Proxy.BootstrapServers[1].ListenerAddress, "0.0.0.0:32405")
	a.Equal(c.Proxy.BootstrapServers[1].AdvertisedAddress, "kafka-5.grepplabs.com:9092")

}

func TestExternalServersMappingFromEnv(t *testing.T) {
	setupBootstrapServersMappingTest()

	_ = os.Setenv("BOOTSTRAP_SERVER_MAPPING", "	192.168.99.100:32401,0.0.0.0:32401")
	_ = os.Setenv("EXTERNAL_SERVER_MAPPING", "	192.168.99.100:32404,0.0.0.0:32404	kafka-5.example.com:9092,0.0.0.0:32405,kafka-5.grepplabs.com:9092")

	var args []string
	_ = Server.ParseFlags(args)
	err := Server.PreRunE(nil, args)
	a := assert.New(t)
	a.Nil(err)
	a.Len(c.Proxy.BootstrapServers, 1)
	a.Len(c.Proxy.ExternalServers, 2)

	a.Equal(c.Proxy.BootstrapServers[0].BrokerAddress, "192.168.99.100:32401")
	a.Equal(c.Proxy.BootstrapServers[0].ListenerAddress, "0.0.0.0:32401")
	a.Equal(c.Proxy.BootstrapServers[0].AdvertisedAddress, "0.0.0.0:32401")

	a.Equal(c.Proxy.ExternalServers[0].BrokerAddress, "192.168.99.100:32404")
	a.Equal(c.Proxy.ExternalServers[0].ListenerAddress, "0.0.0.0:32404")
	a.Equal(c.Proxy.ExternalServers[0].AdvertisedAddress, "0.0.0.0:32404")

	a.Equal(c.Proxy.ExternalServers[1].BrokerAddress, "kafka-5.example.com:9092")
	a.Equal(c.Proxy.ExternalServers[1].ListenerAddress, "0.0.0.0:32405")
	a.Equal(c.Proxy.ExternalServers[1].AdvertisedAddress, "kafka-5.grepplabs.com:9092")

}
