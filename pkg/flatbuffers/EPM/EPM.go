// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package EPM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Entity Profile Message
type EPM struct {
	_tab flatbuffers.Table
}

const EPMIdentifier = "$EPM"

func GetRootAsEPM(buf []byte, offset flatbuffers.UOffsetT) *EPM {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EPM{}
	x.Init(buf, n+offset)
	return x
}

func FinishEPMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(EPMIdentifier)
	builder.FinishWithFileIdentifier(offset, identifierBytes)
}

func EPMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.BufferHasIdentifier(buf, EPMIdentifier)
}

func GetSizePrefixedRootAsEPM(buf []byte, offset flatbuffers.UOffsetT) *EPM {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EPM{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedEPMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(EPMIdentifier)
	builder.FinishSizePrefixedWithFileIdentifier(offset, identifierBytes)
}

func SizePrefixedEPMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.SizePrefixedBufferHasIdentifier(buf, EPMIdentifier)
}

func (rcv *EPM) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EPM) Table() flatbuffers.Table {
	return rcv._tab
}

/// Distinguished Name of the entity
func (rcv *EPM) Dn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Distinguished Name of the entity
/// Common name of the entity (person or organization)
func (rcv *EPM) LegalName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Common name of the entity (person or organization)
/// Family name or surname of the person
func (rcv *EPM) FamilyName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Family name or surname of the person
/// Given name or first name of the person
func (rcv *EPM) GivenName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Given name or first name of the person
/// Additional name or middle name of the person
func (rcv *EPM) AdditionalName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Additional name or middle name of the person
/// Honorific prefix preceding the person's name (e.g., Mr., Dr.)
func (rcv *EPM) HonorificPrefix() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Honorific prefix preceding the person's name (e.g., Mr., Dr.)
/// Honorific suffix following the person's name (e.g., Jr., Sr.)
func (rcv *EPM) HonorificSuffix() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Honorific suffix following the person's name (e.g., Jr., Sr.)
/// Job title of the person
func (rcv *EPM) JobTitle() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Job title of the person
/// Occupation of the person
func (rcv *EPM) Occupation() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Occupation of the person
/// Physical Address
func (rcv *EPM) Address(obj *Address) *Address {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Address)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Physical Address
/// Alternate names for the entity
func (rcv *EPM) AlternateNames(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *EPM) AlternateNamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Alternate names for the entity
/// Email address of the entity
func (rcv *EPM) Email() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Email address of the entity
/// Telephone number of the entity
func (rcv *EPM) Telephone() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Telephone number of the entity
/// Cryptographic keys associated with the entity
func (rcv *EPM) Keys(obj *CryptoKey, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EPM) KeysLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Cryptographic keys associated with the entity
/// Multiformat addresses associated with the entity
func (rcv *EPM) MultiformatAddress(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *EPM) MultiformatAddressLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Multiformat addresses associated with the entity
func EPMStart(builder *flatbuffers.Builder) {
	builder.StartObject(15)
}
func EPMAddDn(builder *flatbuffers.Builder, dn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(dn), 0)
}
func EPMAddLegalName(builder *flatbuffers.Builder, legalName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(legalName), 0)
}
func EPMAddFamilyName(builder *flatbuffers.Builder, familyName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(familyName), 0)
}
func EPMAddGivenName(builder *flatbuffers.Builder, givenName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(givenName), 0)
}
func EPMAddAdditionalName(builder *flatbuffers.Builder, additionalName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(additionalName), 0)
}
func EPMAddHonorificPrefix(builder *flatbuffers.Builder, honorificPrefix flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(honorificPrefix), 0)
}
func EPMAddHonorificSuffix(builder *flatbuffers.Builder, honorificSuffix flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(honorificSuffix), 0)
}
func EPMAddJobTitle(builder *flatbuffers.Builder, jobTitle flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(jobTitle), 0)
}
func EPMAddOccupation(builder *flatbuffers.Builder, occupation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(occupation), 0)
}
func EPMAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(address), 0)
}
func EPMAddAlternateNames(builder *flatbuffers.Builder, alternateNames flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(alternateNames), 0)
}
func EPMStartAlternateNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EPMAddEmail(builder *flatbuffers.Builder, email flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(email), 0)
}
func EPMAddTelephone(builder *flatbuffers.Builder, telephone flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(telephone), 0)
}
func EPMAddKeys(builder *flatbuffers.Builder, keys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(keys), 0)
}
func EPMStartKeysVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EPMAddMultiformatAddress(builder *flatbuffers.Builder, multiformatAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(multiformatAddress), 0)
}
func EPMStartMultiformatAddressVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EPMEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
