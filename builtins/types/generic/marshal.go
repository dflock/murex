package generic

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"strings"
)

func marshal(_ *lang.Process, iface interface{}) (b []byte, err error) {
	switch v := iface.(type) {
	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case [][]string:
		for i := range v {
			b = append(b, []byte(strings.Join(v[i], "\t")+utils.NewLineString)...)
		}
		return

	case []interface{}:
		for i := range v {
			b = append(b, iface2str(&v[i])...)
		}
		return

	case map[string]string:
		for s := range v {
			b = append(b, []byte(s+": "+v[s]+utils.NewLineString)...)
		}
		return

	case map[string]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", s, fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return

	case interface{}:
		return []byte(fmt.Sprintln(iface)), nil

	default:
		err = errors.New("I don't know how to marshal that data into a `*`. Data possibly too complex?")
		return
	}
}

func iface2str(v *interface{}) (b []byte) {
	switch t := (*v).(type) {
	case string:
		return []byte((*v).(string) + utils.NewLineString)
	case int, uint, float64:
		s := fmt.Sprintln(*v)
		return []byte(s)
	default:
		s := fmt.Sprintf("%s: %s%s", t, *v, utils.NewLineString)
		return []byte(s)
	}
}

func unmarshal(p *lang.Process) (interface{}, error) {
	table := make([][]string, 1)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		table = append(table, rxWhitespace.Split(scanner.Text(), -1))
	}

	if len(table) > 1 {
		table = table[1:]
	}

	err := scanner.Err()
	return table, err
}
