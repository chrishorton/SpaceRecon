// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package OEM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Orbit Ephemeris Message
type OEM struct {
	_tab flatbuffers.Table
}

const OEMIdentifier = "$OEM"

func GetRootAsOEM(buf []byte, offset flatbuffers.UOffsetT) *OEM {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OEM{}
	x.Init(buf, n+offset)
	return x
}

func FinishOEMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(OEMIdentifier)
	builder.FinishWithFileIdentifier(offset, identifierBytes)
}

func OEMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.BufferHasIdentifier(buf, OEMIdentifier)
}

func GetSizePrefixedRootAsOEM(buf []byte, offset flatbuffers.UOffsetT) *OEM {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OEM{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedOEMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(OEMIdentifier)
	builder.FinishSizePrefixedWithFileIdentifier(offset, identifierBytes)
}

func SizePrefixedOEMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.SizePrefixedBufferHasIdentifier(buf, OEMIdentifier)
}

func (rcv *OEM) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OEM) Table() flatbuffers.Table {
	return rcv._tab
}

/// OEM Header
/// OEM Version
func (rcv *OEM) CcsdsOemVers() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// OEM Header
/// OEM Version
func (rcv *OEM) MutateCcsdsOemVers(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// Creation Date
func (rcv *OEM) CreationDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Creation Date
/// Originator
func (rcv *OEM) Originator() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Originator
/// Array of ephemeris data blocks
func (rcv *OEM) EphemerisDataBlock(obj *ephemerisDataBlock, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *OEM) EphemerisDataBlockLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Array of ephemeris data blocks
func OEMStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func OEMAddCcsdsOemVers(builder *flatbuffers.Builder, ccsdsOemVers float64) {
	builder.PrependFloat64Slot(0, ccsdsOemVers, 0.0)
}
func OEMAddCreationDate(builder *flatbuffers.Builder, creationDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(creationDate), 0)
}
func OEMAddOriginator(builder *flatbuffers.Builder, originator flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(originator), 0)
}
func OEMAddEphemerisDataBlock(builder *flatbuffers.Builder, ephemerisDataBlock flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(ephemerisDataBlock), 0)
}
func OEMStartEphemerisDataBlockVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OEMEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}