package day13

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

func parseInput(input string) ([]Packet, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	packets := []Packet{}
	for {
		packet, err := parsePacket(s)
		if err != nil {
			if err == ErrEOF {
				break
			}
			return nil, errors.Wrap(err, "parsePacket")
		}
		packets = append(packets, packet)
	}
	return packets, nil
}

var ErrEOF = errors.New("end of file")

func parsePacket(s *bufio.Scanner) (Packet, error) {
	ok := s.Scan()
	if !ok {
		return Packet{}, ErrEOF
	}
	l1, err := parseLine(s.Text())
	if err != nil {
		return Packet{}, errors.Wrap(err, "parseLine")
	}
	s.Scan()
	l2, err := parseLine(s.Text())
	if err != nil {
		return Packet{}, errors.Wrap(err, "parseLine")
	}
	s.Scan()
	return Packet{l1, l2}, nil
}

func parseLine(line string) ([]interface{}, error) {
	out := []interface{}{}
	return out, json.Unmarshal([]byte(line), &out)
}
