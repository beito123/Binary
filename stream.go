package binary

import "errors"

/*
 * Binary
 *
 * Copyright (c) 2018 beito
 *
 * This software is released under the MIT License.
 * http://opensource.org/licenses/mit-license.php
 */

var errNoEnough = errors.New("no enough buffer")

// NewStream returns new Stream
func NewStream() *Stream {
	return NewStreamBytes([]byte{})
}

// NewStreamBytes returns new Stream from bytes
func NewStreamBytes(b []byte) *Stream {
	return &Stream{
		buf:     b,
		correct: true,
	}
}

// Stream is basic binary stream.
type Stream struct {
	buf     []byte
	off     int
	correct bool
}

// Reset resets Buffer
func (bs *Stream) Reset() {
	bs.correct = true
	bs.off = 0
	bs.buf = []byte{}
}

// Off returns offset
func (bs *Stream) Off() int {
	return bs.off
}

// Get returns n bytes from Buffer with []byte
func (bs *Stream) Get(n int) []byte {
	off := bs.off
	if (n + off) >= bs.Len() {
		n = bs.Len()
	}

	bs.off += n

	return bs.buf[off : off+n]
}

// Put puts value to buffer
func (bs *Stream) Put(value []byte) error {
	_, err := bs.Write(value)

	return err
}

// Bytes returns the bytes left from Buffer.
func (bs *Stream) Bytes() []byte {
	return bs.buf[bs.off:]
}

// AllBytes return all bytes
func (bs *Stream) AllBytes() []byte {
	return bs.buf
}

// SetBytes sets bytes
func (bs *Stream) SetBytes(b []byte) {
	bs.Reset()

	bs.buf = b
}

// Len returns len the bytes left
func (bs *Stream) Len() int {
	return len(bs.buf[bs.off:])
}

// Skip skips n bytes on buffer
func (bs *Stream) Skip(n int) {
	if (n + bs.off) >= bs.Len() {
		n = bs.Len()
	}

	bs.off += n
}

//Pad puts empty bytes (0x00) of le (len).
func (bs *Stream) Pad(le int) error {
	return bs.Put(make([]byte, le))
}

// Read reads and sets p
func (bs *Stream) Read(p []byte) (n int, err error) {
	return copy(p, bs.Get(len(p))), nil
}

// Write writes p
func (bs *Stream) Write(p []byte) (n int, err error) {
	bs.buf = append(bs.buf, p...)

	return len(p), nil
}

/*
 * Data types
 * | name  | size | encode |                   range                   |
 *  Byte    1byte   Big                        0 - 255
 *  SByte   1byte   Big                     -128 - 127
 *  Short   2bytes  Big                        0 - 65535
 *  SShort  2bytes  Big                   -32768 - 32767
 *  LShort  2bytes  Little                     0 - 65535
 *  LSShort 2bytes  Little                -32768 - 32767
 *  Triad   3bytes  Little                     0 - 16777215
 *  Int     4bytes  Big              -2147483648 - 2147483647
 *  Long    8bytes  Big     -9223372036854775808 - 9223372036854775807
 *  String  ?bytes  Big                        ? - ?
 */

/*
 * Byte
 * SignedByte
 * Short
 * SignedShort
 * LShort
 * SignedLShort
 * Triad
 * LTriad
 * Int
 * Float
 * LFloat
 * Double
 * LDouble
 * Long
 * /////////////
 * Bool
 * String
 * HexString
 * Address
 * NBT
 * Item
 * UUID
 * Position
 * BlockPosition
 * EntityMetadata
 */

// Byte sets byte(unsign) got from buffer to value
func (bs *Stream) Byte() (byte, error) {
	return ReadEByte(bs.Get(ByteSize))
}

// SByte sets byte(sign) got from buffer to value
func (bs *Stream) SByte() (int8, error) {
	return ReadESByte(bs.Get(ByteSize))
}

// PutByte puts byte(unsign) from value to buffer
func (bs *Stream) PutByte(value byte) error {
	return Write(bs, BigEndian, value)
}

// PutSByte puts byte(sign) from value to buffer
func (bs *Stream) PutSByte(value int8) error {
	return Write(bs, BigEndian, value)
}

// Short sets short(unsign) got from buffer to value
func (bs *Stream) Short() (uint16, error) {
	return ReadEUShort(bs.Get(ShortSize))
}

// SShort sets short(sign) got from buffer to value
func (bs *Stream) SShort() (int16, error) {
	return ReadEShort(bs.Get(ShortSize))
}

// LShort sets short(unsign) got from buffer as LittleEndian to value
func (bs *Stream) LShort() (uint16, error) {
	return ReadELUShort(bs.Get(ShortSize))
}

// LSShort sets short(sign) got from buffer as LittleEndian to value
func (bs *Stream) LSShort() (int16, error) {
	return ReadELShort(bs.Get(ShortSize))
}

// PutShort puts short(unsign) from value to buffer
func (bs *Stream) PutShort(value uint16) error {
	return Write(bs, BigEndian, value)
}

// PutSShort puts short(sign) from value to buffer
func (bs *Stream) PutSShort(value int16) error {
	return Write(bs, BigEndian, value)
}

// PutLShort puts short(unsign) from value to buffer as LittleEndian
func (bs *Stream) PutLShort(value uint16) error {
	return Write(bs, LittleEndian, value)
}

// PutLSShort puts short(sign) from value to buffer as LittleEndian
func (bs *Stream) PutLSShort(value int16) error {
	return Write(bs, LittleEndian, value)
}

// Int sets int got from buffer to value
func (bs *Stream) Int() (int32, error) {
	return ReadEInt(bs.Get(IntSize))
}

// PutInt puts int from value to buffer
func (bs *Stream) PutInt(value int32) error {
	return Write(bs, BigEndian, value)
}

// Long sets long got from buffer to value
func (bs *Stream) Long() (int64, error) {
	return ReadELong(bs.Get(LongSize))
}

// PutLong puts long from value to buffer
func (bs *Stream) PutLong(value int64) error {
	return Write(bs, BigEndian, value)
}

// Bool sets byte got from buffer as bool to value
func (bs *Stream) Bool() (bool, error) {
	val, err := bs.Byte()
	if err != nil {
		return false, err
	}

	return val != 0, nil
}

//PutBool puts bool as byte from value to buffer
func (bs *Stream) PutBool(value bool) error {
	var val byte
	if value {
		val = 1 // true
	}

	return bs.PutByte(val)
}
