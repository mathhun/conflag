package conflag

import "fmt"

type conf map[string]interface{}

func (c conf) toArgs(positions ...string) []string {

	nowConf := c
	for _, p := range positions {
		nextConf, ok := nowConf[p]
		if !ok {
			return []string{}
		}
		n, ok := nextConf.(map[string]interface{})
		if !ok {
			break
		}
		nowConf = n
	}

	var args []string
	for k, v := range nowConf {
		switch v.(type) {
		case map[string]interface{}: // nested configuration
			continue
		case bool:
			args = append(args, fmt.Sprintf("-%s=%t", k, v.(bool)))
		default:
			args = append(args, "-"+k, fmt.Sprintf("%v", v))
		}
	}

	return args
}
