// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package CAT

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CATCOLLECTION struct {
	_tab flatbuffers.Table
}

func GetRootAsCATCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *CATCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CATCOLLECTION{}
	x.Init(buf, n+offset)
	return x
}

func FinishCATCOLLECTIONBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsCATCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *CATCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CATCOLLECTION{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedCATCOLLECTIONBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *CATCOLLECTION) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CATCOLLECTION) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CATCOLLECTION) Records(obj *CAT, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *CATCOLLECTION) RecordsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CATCOLLECTIONStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CATCOLLECTIONAddRecords(builder *flatbuffers.Builder, records flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(records), 0)
}
func CATCOLLECTIONStartRecordsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CATCOLLECTIONEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
