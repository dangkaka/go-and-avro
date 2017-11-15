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

type Activity struct {
	Id   string
	Type string
	Data string
}

func DeserializeActivity(r io.Reader) (*Activity, error) {
	return readActivity(r)
}

func NewActivityWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &Activity{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func NewActivity() *Activity {
	v := &Activity{}

	return v
}

func (r *Activity) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Type\",\"type\":\"string\"},{\"name\":\"Data\",\"type\":\"string\"}],\"name\":\"Activity\",\"type\":\"record\"}"
}

func (r *Activity) Serialize(w io.Writer) error {
	return writeActivity(r, w)
}
