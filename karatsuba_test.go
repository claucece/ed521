package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) TestKaratsuba(c *C) {
	x, _ := deserialize(serialized{
		0xf5, 0x81, 0x74, 0xd5, 0x7a, 0x33, 0x72,
		0x36, 0x3c, 0x0d, 0x9f, 0xcf, 0xaa, 0x3d,
		0xc1, 0x8b, 0x1e, 0xff, 0x7e, 0x89, 0xbf,
		0x76, 0x78, 0x63, 0x65, 0x80, 0xd1, 0x7d,
		0xd8, 0x4a, 0x87, 0x3b, 0x14, 0xb9, 0xc0,
		0xe1, 0x68, 0x0b, 0xbd, 0xc8, 0x76, 0x47,
		0xf3, 0xc3, 0x82, 0x90, 0x2d, 0x2f, 0x58,
		0xd2, 0x75, 0x4b, 0x39, 0xbc, 0xa8, 0x74,
	})

	y, _ := deserialize(serialized{
		0x74, 0xa8, 0xbc, 0x39, 0x4b, 0x75, 0xd2,
		0x58, 0x2f, 0x2d, 0x90, 0x82, 0xc3, 0xf3,
		0x47, 0x76, 0xc8, 0xbd, 0x0b, 0x68, 0xe1,
		0xc0, 0xb9, 0x14, 0x3b, 0x87, 0x4a, 0xd8,
		0x7d, 0xd1, 0x80, 0x65, 0x63, 0x78, 0x76,
		0xbf, 0x89, 0x7e, 0xff, 0x1e, 0x8b, 0xc1,
		0x3d, 0xaa, 0xcf, 0x9f, 0x0d, 0x3c, 0x36,
		0x72, 0x33, 0x7a, 0xd5, 0x74, 0x81, 0xf5,
	})

	z, _ := deserialize(serialized{
		0x11, 0x95, 0x9c, 0x2e, 0x91, 0x78, 0x6f,
		0xec, 0xff, 0x37, 0xe5, 0x8e, 0x2b, 0x50,
		0x9e, 0xf8, 0xfb, 0x41, 0x08, 0xc4, 0xa7,
		0x02, 0x1c, 0xbf, 0x5a, 0x9f, 0x18, 0xa7,
		0xec, 0x32, 0x65, 0x7e, 0xed, 0xdc, 0x81,
		0x81, 0x80, 0xa8, 0x4c, 0xdd, 0x95, 0x14,
		0xe6, 0x67, 0x26, 0xd3, 0xa1, 0x22, 0xdb,
		0xb3, 0x9f, 0x17, 0x7a, 0x85, 0x16, 0x6c,
	})

	result := new(bigNumber)
	karatsubaMul(result, x, y)
	c.Assert(result, DeepEquals, z)
}
