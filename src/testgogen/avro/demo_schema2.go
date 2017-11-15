/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/alanctgardner/gogen-avro
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 *
 * SOURCES:
 *     example.avsc
 *     example2.avsc
 *     activity.avsc
 */

package avro

import (
	"github.com/alanctgardner/gogen-avro/container"
	"io"
)

type DemoSchema2 struct {
	IntField    int32
	DoubleField float64
	StringField string
	BoolField   bool
	BytesField  []byte
	Newfield    string
}

func DeserializeDemoSchema2(r io.Reader) (*DemoSchema2, error) {
	return readDemoSchema2(r)
}

func NewDemoSchema2Writer(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &DemoSchema2{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func NewDemoSchema2() *DemoSchema2 {
	v := &DemoSchema2{}

	return v
}

func (r *DemoSchema2) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"},{\"name\":\"Newfield\",\"type\":\"string\"}],\"name\":\"DemoSchema2\",\"type\":\"record\"}"
}

func (r *DemoSchema2) Serialize(w io.Writer) error {
	return writeDemoSchema2(r, w)
}
