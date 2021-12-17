package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBin(t *testing.T) {
	assert.Equal(t, "01010111", hexToBin("57"))
	assert.Equal(t, "10001010", hexToBin("8A"))
}

func TestParseBin(t *testing.T) {
	assert.Equal(t, 4, parseBin("100"))
	assert.Panics(t, func() {
		parseBin("1002")
	})
}

func TestParseInput(t *testing.T) {
	assert.Equal(
		t,
		ValuePacket{BasePacket{6, 4}, 2021},
		parseInput("D2FE28"),
	)

	assert.Equal(
		t,
		LessThanOperatorPacket{
			BaseOperatorPacket{
				BasePacket{1, 6},
				[]Packet{
					ValuePacket{BasePacket{6, 4}, 10},
					ValuePacket{BasePacket{2, 4}, 20},
				},
			},
		},
		parseInput("38006F45291200"),
	)
}

func TestSumVersions(t *testing.T) {
	assert.Equal(t, 16, sumVersions(parseInput("8A004A801A8002F478")))
	assert.Equal(t, 12, sumVersions(parseInput("620080001611562C8802118E34")))
	assert.Equal(t, 23, sumVersions(parseInput("C0015000016115A2E0802F182340")))
	assert.Equal(t, 31, sumVersions(parseInput("A0016C880162017C3686B18A3D4780")))
}

func TestComputeValue(t *testing.T) {
	assert.Equal(t, 3, parseInput("C200B40A82").GetValue())
	assert.Equal(t, 54, parseInput("04005AC33890").GetValue())
	assert.Equal(t, 7, parseInput("880086C3E88112").GetValue())
	assert.Equal(t, 9, parseInput("CE00C43D881120").GetValue())
	assert.Equal(t, 1, parseInput("D8005AC2A8F0").GetValue())
	assert.Equal(t, 0, parseInput("F600BC2D8F").GetValue())
	assert.Equal(t, 0, parseInput("9C005AC2F8F0").GetValue())
	assert.Equal(t, 1, parseInput("9C0141080250320F1802104A08").GetValue())
}
