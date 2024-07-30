package mini

import "fmt"

func (c Commands) addPrefixToArgs() {
	for _, v := range c {
		v.ShortCmd = fmt.Sprintf("-%v", v.ShortCmd)
		v.LongCmd = fmt.Sprintf("--%v", v.LongCmd)
	}
}
