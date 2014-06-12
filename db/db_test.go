package db

import (
	"net"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block", func() {
	It("Constructs from a record correctly", func() {
		record := []string{
			"::ffff:223.255.252.0",
			"119",
			"1814991",
			"1814991",
			"",
			"",
			"",
			"",
			"0",
			"0",
		}

		block := newBlockFromRecord(record)
		Expect(block.NetworkStartIp).To(Equal(net.ParseIP("223.255.252.0")))
		Expect(block.NetworkPrefixLength).To(Equal(119))
		Expect(block.GeonameId).To(Equal("1814991"))
		Expect(block.IsAnonymousProxy).To(Equal(false))
		Expect(block.IsSatelliteProvider).To(Equal(false))
	})

	It("Has the correct network end ip", func() {
		block := Block{
			NetworkStartIp:      net.ParseIP("::ffff:223.255.252.0"),
			NetworkPrefixLength: 119,
		}
		Expect(block.NetworkEndIp()).To(Equal(net.ParseIP("::ffff:223.255.253.255")))
	})
})
