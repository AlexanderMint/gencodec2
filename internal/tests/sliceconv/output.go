// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package sliceconv

import (
	"encoding/json"
)

func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
		Func        []replacedInt
	}
	var enc X
	if x.Slice != nil {
		enc.Slice = make([]replacedInt, len(x.Slice))
		for k, v := range x.Slice {
			enc.Slice[k] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedSlice2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[k] = replacedInt(v)
		}
	}
	enc.ByteString = []byte(x.ByteString)
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	if x.Func() != nil {
		enc.Func = make([]replacedInt, len(x.Func()))
		for k, v := range x.Func() {
			enc.Func[k] = replacedInt(v)
		}
	}
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Slice != nil {
		x.Slice = make([]int, len(dec.Slice))
		for k, v := range dec.Slice {
			x.Slice[k] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedSlice, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[k] = int(v)
		}
	}
	if dec.ByteString != nil {
		x.ByteString = string(dec.ByteString)
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}

func (x X) MarshalYAML() (interface{}, error) {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
		Func        []replacedInt
	}
	var enc X
	if x.Slice != nil {
		enc.Slice = make([]replacedInt, len(x.Slice))
		for k, v := range x.Slice {
			enc.Slice[k] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedSlice2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[k] = replacedInt(v)
		}
	}
	enc.ByteString = []byte(x.ByteString)
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	if x.Func() != nil {
		enc.Func = make([]replacedInt, len(x.Func()))
		for k, v := range x.Func() {
			enc.Func[k] = replacedInt(v)
		}
	}
	return &enc, nil
}

func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Slice != nil {
		x.Slice = make([]int, len(dec.Slice))
		for k, v := range dec.Slice {
			x.Slice[k] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedSlice, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[k] = int(v)
		}
	}
	if dec.ByteString != nil {
		x.ByteString = string(dec.ByteString)
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}

func (x X) MarshalTOML() (interface{}, error) {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
		Func        []replacedInt
	}
	var enc X
	if x.Slice != nil {
		enc.Slice = make([]replacedInt, len(x.Slice))
		for k, v := range x.Slice {
			enc.Slice[k] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedSlice2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[k] = replacedInt(v)
		}
	}
	enc.ByteString = []byte(x.ByteString)
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	if x.Func() != nil {
		enc.Func = make([]replacedInt, len(x.Func()))
		for k, v := range x.Func() {
			enc.Func[k] = replacedInt(v)
		}
	}
	return &enc, nil
}

func (x *X) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type X struct {
		Slice       []replacedInt
		Named       namedSlice2
		ByteString  []byte
		NoConv      []int
		NoConvNamed namedSlice
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Slice != nil {
		x.Slice = make([]int, len(dec.Slice))
		for k, v := range dec.Slice {
			x.Slice[k] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedSlice, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[k] = int(v)
		}
	}
	if dec.ByteString != nil {
		x.ByteString = string(dec.ByteString)
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}
