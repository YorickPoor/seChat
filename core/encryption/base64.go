// Package base64 provides Base64 Encoding
package encryption

const cb64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

type Base64 string

func _indexAt(d byte) byte {
	var j byte = 65
	for i := 0; i < len(cb64); i++ {
		if d == cb64[i] {
			j = byte(i)
		}
	}
	return j
}

func _encodeblock(in []byte, l int) string {
	var out []byte = make([]byte, 4)
	out[0] = cb64[in[0]>>2]
	out[1] = cb64[((in[0]&0x03)<<4)|((in[1]&0xf0)>>4)]
	if l > 1 {
		out[2] = cb64[(((in[1] & 0x0f) << 2) | ((in[2] & 0xc0) >> 6))]
	} else {
		out[2] = "="[0]
	}
	if l > 2 {
		out[3] = cb64[in[2]&0x3f]
	} else {
		out[3] = "="[0]
	}
	return string(out)
}

func (b *Base64) Base64Encode(data string) string {
	var (
		rs string
		in []byte = make([]byte, 3)
	)
	for i := 0; i < len(data); i += 3 {
		if i+2 < len(data) {
			in[0] = data[i]
			in[1] = data[i+1]
			in[2] = data[i+2]
			rs += _encodeblock(in, 3)
		} else {
			switch len(data) - 1 {
			case i:
				in[0] = data[i]
				in[1] = 0
				in[2] = 0
				rs += _encodeblock(in, 1)
			case i + 1:
				in[0] = data[i]
				in[1] = data[i+1]
				in[2] = 0
				rs += _encodeblock(in, 2)
			}
		}

	}
	return rs
}

func _decodeblock(in []byte) string {
	var (
		out    []byte = make([]byte, 3)
		outstr string
	)
	out[0] = (in[0] << 2) | (in[1] >> 4)
	outstr += string(out[0])
	if in[2] != 64 {
		out[1] = ((in[1] & 15) << 4) | (in[2] >> 2)
		outstr += string(out[1])
	}
	if in[3] != 64 {
		out[2] = ((in[2] & 3) << 6) | in[3]
		outstr += string(out[2])
	}
	return outstr

}

func (b *Base64) Base64Decode(data string) string {
	var (
		rs  string
		enc []byte = make([]byte, 4)
		i   int    = 0
	)

	for i < len(data) {
		for j := 0; j < 4; j++ {
			enc[j] = _indexAt(data[i])
			i++
		}
		rs += _decodeblock(enc)
	}
	return rs
}
