// UTILIDADES PARA CONVERSIONES

package controllers

import (
	"fmt"
)

func FilterKeys(in map[string]interface{}, neededKeys []string) (out map[string]interface{}, err error) {
	out = make(map[string]interface{})
	for _, key := range neededKeys {
		if v, ok := in[key]; ok {
			out[key] = v
		} else {
			err = fmt.Errorf("key '%s' not found in supplied object", key)
			return
		}
	}
	return
}
