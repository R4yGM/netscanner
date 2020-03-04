package netserv

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Protoent struct {
	Name    string
	Aliases []string
	Number  int
}

type Servent struct {
	Name     string
	Aliases  []string
	Port     int
	Protocol *Protoent
}

var (
	Protocols []*Protoent
	Services  []*Servent
)

func init() {
	protoMap := make(map[string]*Protoent)

	data, err := ioutil.ReadFile("/etc/protocols")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		split := strings.SplitN(line, "#", 2)
		fields := strings.Fields(split[0])
		if len(fields) < 2 {
			continue
		}

		num, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			panic(err)
		}

		protoent := &Protoent{
			Name:    fields[0],
			Aliases: fields[2:],
			Number:  int(num),
		}
		Protocols = append(Protocols, protoent)

		protoMap[fields[0]] = protoent
	}

	data, err = ioutil.ReadFile("/etc/services")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		split := strings.SplitN(line, "#", 2)
		fields := strings.Fields(split[0])
		if len(fields) < 2 {
			continue
		}

		name := fields[0]
		portproto := strings.SplitN(fields[1], "/", 2)
		port, err := strconv.ParseInt(portproto[0], 10, 32)
		if err != nil {
			panic(err)
		}

		proto := portproto[1]
		aliases := fields[2:]

		Services = append(Services, &Servent{
			Name:     name,
			Aliases:  aliases,
			Port:     int(port),
			Protocol: protoMap[proto],
		})
	}
}
func (this *Protoent) Equal(other *Protoent) bool {
	if this == nil && other == nil {
		return true
	}

	if this == nil || other == nil {
		return false
	}

	return this.Number == other.Number
}

func GetProtoByNumber(num int) (protoent *Protoent) {
	for _, protoent := range Protocols {
		if protoent.Number == num {
			return protoent
		}
	}
	return nil
}

func GetProtoByName(name string) (protoent *Protoent) {
	for _, protoent := range Protocols {
		if protoent.Name == name {
			return protoent
		}

		for _, alias := range protoent.Aliases {
			if alias == name {
				return protoent
			}
		}
	}

	return nil
}

func GetServByName(name string, protocol *Protoent) (servent *Servent) {
	for _, servent := range Services {
		if !servent.Protocol.Equal(protocol) {
			continue
		}

		if servent.Name == name {
			return servent
		}

		for _, alias := range servent.Aliases {
			if alias == name {
				return servent
			}
		}
	}

	return nil
}

func GetServByPort(port int, protocol *Protoent) *Servent {
	for _, servent := range Services {
		if servent.Port == port && servent.Protocol.Equal(protocol) {
			return servent
		}
	}

	return nil
}
