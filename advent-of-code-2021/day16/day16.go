package day16

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToBin(hexString string) string {
	mapping := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	result := make([]string, len(hexString))

	for i, char := range hexString {
		bin, present := mapping[char]

		if !present {
			panic(fmt.Errorf("invalid input, cannot parse %c", char))
		}

		result[i] = bin
	}

	return strings.Join(result, "")
}

func parseBin(binStr string) int {
	val, err := strconv.ParseInt(binStr, 2, 64)

	if err != nil {
		panic(err)
	}

	return int(val)
}

type Packet interface {
	GetVersion() int
	GetTypeId() int
	GetValue() int
}

type BasePacket struct {
	version int
	typeId  int
}

func (p BasePacket) GetVersion() int {
	return p.version
}

func (p BasePacket) GetTypeId() int {
	return p.typeId
}

type ValuePacket struct {
	BasePacket
	value int
}

func (p ValuePacket) GetValue() int {
	return p.value
}

type PacketWithSubpackets interface {
	Packet
	GetSubpackets() []Packet
}

type BaseOperatorPacket struct {
	BasePacket
	subpackets []Packet
}

func (p BaseOperatorPacket) GetSubpackets() []Packet {
	return p.subpackets
}

type SumOperatorPacket struct {
	BaseOperatorPacket
}

func (p SumOperatorPacket) GetValue() int {
	sum := 0

	for _, subpacket := range p.GetSubpackets() {
		sum += subpacket.GetValue()
	}

	return sum
}

type ProductOperatorPacket struct {
	BaseOperatorPacket
}

func (p ProductOperatorPacket) GetValue() int {
	prod := 1

	for _, subpacket := range p.subpackets {
		prod *= subpacket.GetValue()
	}

	return prod
}

type MinimumOperatorPacket struct {
	BaseOperatorPacket
}

func (p MinimumOperatorPacket) GetValue() int {
	min := p.subpackets[0].GetValue()

	for _, subpacket := range p.subpackets[1:] {
		if val := subpacket.GetValue(); val < min {
			min = val
		}
	}

	return min
}

type MaximumOperatorPacket struct {
	BaseOperatorPacket
}

func (p MaximumOperatorPacket) GetValue() int {
	max := p.subpackets[0].GetValue()

	for _, subpacket := range p.subpackets[1:] {
		if val := subpacket.GetValue(); val > max {
			max = val
		}
	}

	return max
}

type GreaterThanOperatorPacket struct {
	BaseOperatorPacket
}

func (p GreaterThanOperatorPacket) GetValue() int {
	if p.subpackets[0].GetValue() > p.subpackets[1].GetValue() {
		return 1
	}

	return 0
}

type LessThanOperatorPacket struct {
	BaseOperatorPacket
}

func (p LessThanOperatorPacket) GetValue() int {
	if p.subpackets[0].GetValue() < p.subpackets[1].GetValue() {
		return 1
	}

	return 0
}

type EqualOperatorPacket struct {
	BaseOperatorPacket
}

func (p EqualOperatorPacket) GetValue() int {
	if p.subpackets[0].GetValue() == p.subpackets[1].GetValue() {
		return 1
	}

	return 0
}

func readBits(data []rune, index *int, length int) int {
	result := parseBin(string(data[*index : *index+length]))
	*index += length
	return result
}

func readValue(data []rune, index *int) int {
	value := make([]rune, 0)

	// read by batches of 5, stop when the first bit is 0
	for data[*index] == '1' {
		value = append(value, data[*index+1:*index+5]...)
		*index += 5
	}

	// add the last batch (starting with 0)
	value = append(value, data[*index+1:*index+5]...)
	*index += 5

	return parseBin(string(value))
}

func readPacketRec(data []rune, index *int) Packet {
	version := readBits(data, index, 3)
	typeId := readBits(data, index, 3)

	base := BasePacket{version, typeId}

	if typeId == 4 {
		value := readValue(data, index)
		return ValuePacket{base, value}
	}

	lengthTypeId := readBits(data, index, 1)
	subpackets := make([]Packet, 0)

	if lengthTypeId == 0 {
		length := readBits(data, index, 15)
		stopAt := *index + length

		for *index < stopAt {
			subpackets = append(subpackets, readPacketRec(data, index))
		}
	}

	if lengthTypeId == 1 {
		length := readBits(data, index, 11)

		for len(subpackets) < length {
			subpackets = append(subpackets, readPacketRec(data, index))
		}
	}

	operatorBase := BaseOperatorPacket{
		base,
		subpackets,
	}

	switch base.typeId {
	case 0:
		return SumOperatorPacket{operatorBase}
	case 1:
		return ProductOperatorPacket{operatorBase}
	case 2:
		return MinimumOperatorPacket{operatorBase}
	case 3:
		return MaximumOperatorPacket{operatorBase}
	case 5:
		return GreaterThanOperatorPacket{operatorBase}
	case 6:
		return LessThanOperatorPacket{operatorBase}
	case 7:
		return EqualOperatorPacket{operatorBase}
	default:
		panic(fmt.Errorf("invalid typeId: %d", typeId))
	}

}

func parseInput(input string) Packet {
	binRepr := []rune(hexToBin(input))
	index := 0

	return readPacketRec(binRepr, &index)
}

func sumVersions(packet Packet) int {
	switch packet := packet.(type) {
	case ValuePacket:
		return packet.version
	case PacketWithSubpackets:
		tot := packet.GetVersion()

		for _, subpacket := range packet.GetSubpackets() {
			tot += sumVersions(subpacket)
		}

		return tot
	default:
		panic(fmt.Errorf("invalid packet type: %s", packet))
	}
}

func Run(input string) {
	pack := parseInput(input)

	fmt.Println("Sum of version numbers:", sumVersions(pack))
	fmt.Println("Value of outermost packet:", pack.GetValue())
}
