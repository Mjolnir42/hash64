/*-
 * Copyright (c) 2016, Jörg Pernfuß <code.jpe@gmail.com>
 * No rights reserved.
 *
 * This code is available licensed as CC0. Please see the included
 * LICENSE file for more information.
 */

package hash64

import "testing"

type testvector struct {
	value    []byte
	standard string
	padded   string
}

var vectors = []testvector{
	{
		[]byte(`OrpheanBeholderScryDoubt`),
		`Hr7kO4JVPY7ZO4xgN4JmIqBmSIFjRK7o`,
		`Hr7kO4JVPY7ZO4xgN4JmIqBmSIFjRK7o`,
	},
	{
		[]byte(`correct horse battery staple`),
		`MqxmQaJXR0/cPr7nNG/WMLFoNL7t65BoML/gNE`,
		`MqxmQaJXR0/cPr7nNG/WMLFoNL7t65BoML/gNE==`,
	},
	{
		[]byte(`Tr0ub4dor&3`),
		`J56kRK6oN4xm7XA`,
		`J56kRK6oN4xm7XA=`,
	},
	{
		[]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA},
		`.E61/.I4/kU70U`,
		`.E61/.I4/kU70U==`,
	},
	{
		[]byte{0x00, 0x17, 0x2A, 0xFF},
		`./Qezk`,
		`./Qezk==`,
	},
}

func TestStdEncoding(t *testing.T) {
	for _, vc := range vectors {
		v := StdEncoding.EncodeToString(vc.value)
		if v != vc.standard {
			t.Error(
				"For", vc.value,
				"expected", vc.standard,
				"got", v,
			)
		}
	}
}

func TestPadEncoding(t *testing.T) {
	for _, vc := range vectors {
		v := PadEncoding.EncodeToString(vc.value)
		if v != vc.padded {
			t.Error(
				"For", vc.value,
				"expected", vc.padded,
				"got", v,
			)
		}
	}
}
