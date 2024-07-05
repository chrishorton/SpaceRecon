// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package OMM

import (
	flatbuffers "github.com/google/flatbuffers/go"

	MET "github.com/chrishorton/spacerecon/pkg/flatbuffers/MET"
	RFM "github.com/chrishorton/spacerecon/pkg/flatbuffers/RFM"
	TIM "github.com/chrishorton/spacerecon/pkg/flatbuffers/TIM"
)

/// Orbit Mean Elements Message
type OMM struct {
	_tab flatbuffers.Table
}

const OMMIdentifier = "$OMM"

func GetRootAsOMM(buf []byte, offset flatbuffers.UOffsetT) *OMM {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OMM{}
	x.Init(buf, n+offset)
	return x
}

func FinishOMMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(OMMIdentifier)
	builder.FinishWithFileIdentifier(offset, identifierBytes)
}

func OMMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.BufferHasIdentifier(buf, OMMIdentifier)
}

func GetSizePrefixedRootAsOMM(buf []byte, offset flatbuffers.UOffsetT) *OMM {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OMM{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedOMMBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	identifierBytes := []byte(OMMIdentifier)
	builder.FinishSizePrefixedWithFileIdentifier(offset, identifierBytes)
}

func SizePrefixedOMMBufferHasIdentifier(buf []byte) bool {
	return flatbuffers.SizePrefixedBufferHasIdentifier(buf, OMMIdentifier)
}

func (rcv *OMM) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OMM) Table() flatbuffers.Table {
	return rcv._tab
}

/// OMM Header
func (rcv *OMM) CcsdsOmmVers() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// OMM Header
func (rcv *OMM) MutateCcsdsOmmVers(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// Creation Date
func (rcv *OMM) CreationDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Creation Date
/// Originator
func (rcv *OMM) Originator() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Originator
/// OMM Metadata
/// Satellite Name(s)
func (rcv *OMM) ObjectName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// OMM Metadata
/// Satellite Name(s)
/// International Designator (YYYY-NNNAAA)
func (rcv *OMM) ObjectId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// International Designator (YYYY-NNNAAA)
/// Origin of reference frame (EARTH, MARS, MOON, etc.)
func (rcv *OMM) CenterName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Origin of reference frame (EARTH, MARS, MOON, etc.)
/// Name of the reference frame (TEME, EME2000, etc.)
func (rcv *OMM) ReferenceFrame() RFM.RefFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return RFM.RefFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

/// Name of the reference frame (TEME, EME2000, etc.)
func (rcv *OMM) MutateReferenceFrame(n RFM.RefFrame) bool {
	return rcv._tab.MutateInt8Slot(16, int8(n))
}

/// REFERENCE_FRAME_EPOCH
func (rcv *OMM) ReferenceFrameEpoch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// REFERENCE_FRAME_EPOCH
/// Time system used for the orbit state and covariance matrix. (UTC)
func (rcv *OMM) TimeSystem() TIM.TimeSystem {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return TIM.TimeSystem(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 11
}

/// Time system used for the orbit state and covariance matrix. (UTC)
func (rcv *OMM) MutateTimeSystem(n TIM.TimeSystem) bool {
	return rcv._tab.MutateInt8Slot(20, int8(n))
}

/// Description of the Mean Element Theory. (SGP4,DSST,USM)
func (rcv *OMM) MeanElementTheory() MET.MeanElementTheory {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return MET.MeanElementTheory(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Description of the Mean Element Theory. (SGP4,DSST,USM)
func (rcv *OMM) MutateMeanElementTheory(n MET.MeanElementTheory) bool {
	return rcv._tab.MutateInt8Slot(22, int8(n))
}

/// Mean Keplerian Elements in the Specified Reference Frame
/// Plain-Text Comment
func (rcv *OMM) Comment() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Mean Keplerian Elements in the Specified Reference Frame
/// Plain-Text Comment
/// Epoch time, in ISO 8601 UTC format
func (rcv *OMM) Epoch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Epoch time, in ISO 8601 UTC format
/// Semi-major axis in km or mean motion in rev/day
func (rcv *OMM) SemiMajorAxis() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Semi-major axis in km or mean motion in rev/day
func (rcv *OMM) MutateSemiMajorAxis(n float64) bool {
	return rcv._tab.MutateFloat64Slot(28, n)
}

/// Mean motion
func (rcv *OMM) MeanMotion() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Mean motion
func (rcv *OMM) MutateMeanMotion(n float64) bool {
	return rcv._tab.MutateFloat64Slot(30, n)
}

/// Eccentricity
func (rcv *OMM) Eccentricity() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Eccentricity
func (rcv *OMM) MutateEccentricity(n float64) bool {
	return rcv._tab.MutateFloat64Slot(32, n)
}

/// Inclination
func (rcv *OMM) Inclination() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Inclination
func (rcv *OMM) MutateInclination(n float64) bool {
	return rcv._tab.MutateFloat64Slot(34, n)
}

/// Right ascension of ascending node
func (rcv *OMM) RaOfAscNode() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Right ascension of ascending node
func (rcv *OMM) MutateRaOfAscNode(n float64) bool {
	return rcv._tab.MutateFloat64Slot(36, n)
}

/// Argument of pericenter
func (rcv *OMM) ArgOfPericenter() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Argument of pericenter
func (rcv *OMM) MutateArgOfPericenter(n float64) bool {
	return rcv._tab.MutateFloat64Slot(38, n)
}

/// Mean anomaly
func (rcv *OMM) MeanAnomaly() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Mean anomaly
func (rcv *OMM) MutateMeanAnomaly(n float64) bool {
	return rcv._tab.MutateFloat64Slot(40, n)
}

/// Gravitational Coefficient (Gravitational Constant x Central Mass)
func (rcv *OMM) Gm() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Gravitational Coefficient (Gravitational Constant x Central Mass)
func (rcv *OMM) MutateGm(n float64) bool {
	return rcv._tab.MutateFloat64Slot(42, n)
}

/// Spacecraft Parameters
/// S/C Mass
func (rcv *OMM) Mass() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Spacecraft Parameters
/// S/C Mass
func (rcv *OMM) MutateMass(n float64) bool {
	return rcv._tab.MutateFloat64Slot(44, n)
}

/// Solar Radiation Pressure Area (AR) m**2
func (rcv *OMM) SolarRadArea() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Solar Radiation Pressure Area (AR) m**2
func (rcv *OMM) MutateSolarRadArea(n float64) bool {
	return rcv._tab.MutateFloat64Slot(46, n)
}

/// Solar Radiation Pressure Coefficient (CR)
func (rcv *OMM) SolarRadCoeff() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Solar Radiation Pressure Coefficient (CR)
func (rcv *OMM) MutateSolarRadCoeff(n float64) bool {
	return rcv._tab.MutateFloat64Slot(48, n)
}

/// Drag Area (AD) m**2
func (rcv *OMM) DragArea() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Drag Area (AD) m**2
func (rcv *OMM) MutateDragArea(n float64) bool {
	return rcv._tab.MutateFloat64Slot(50, n)
}

/// Drag Coefficient (CD)
func (rcv *OMM) DragCoeff() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Drag Coefficient (CD)
func (rcv *OMM) MutateDragCoeff(n float64) bool {
	return rcv._tab.MutateFloat64Slot(52, n)
}

/// TLE Related Parameters (This section is only required if MEAN_ELEMENT_THEORY=SGP/SGP4)
/// Default value = 0
func (rcv *OMM) EphemerisType() ephemerisType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return ephemerisType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 1
}

/// TLE Related Parameters (This section is only required if MEAN_ELEMENT_THEORY=SGP/SGP4)
/// Default value = 0
func (rcv *OMM) MutateEphemerisType(n ephemerisType) bool {
	return rcv._tab.MutateInt8Slot(54, int8(n))
}

/// Default value = U
func (rcv *OMM) ClassificationType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Default value = U
/// NORAD Catalog Number (Satellite Number) an integer
func (rcv *OMM) NoradCatId() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// NORAD Catalog Number (Satellite Number) an integer
func (rcv *OMM) MutateNoradCatId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(58, n)
}

/// Element set number for this satellite
func (rcv *OMM) ElementSetNo() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Element set number for this satellite
func (rcv *OMM) MutateElementSetNo(n uint32) bool {
	return rcv._tab.MutateUint32Slot(60, n)
}

/// Revolution Number
func (rcv *OMM) RevAtEpoch() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Revolution Number
func (rcv *OMM) MutateRevAtEpoch(n float64) bool {
	return rcv._tab.MutateFloat64Slot(62, n)
}

/// SGP/SGP4 drag-like coefficient (in units 1/[Earth radii])
func (rcv *OMM) Bstar() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// SGP/SGP4 drag-like coefficient (in units 1/[Earth radii])
func (rcv *OMM) MutateBstar(n float64) bool {
	return rcv._tab.MutateFloat64Slot(64, n)
}

/// First Time Derivative of the Mean Motion
func (rcv *OMM) MeanMotionDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// First Time Derivative of the Mean Motion
func (rcv *OMM) MutateMeanMotionDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(66, n)
}

/// Second Time Derivative of Mean Motion
func (rcv *OMM) MeanMotionDdot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Second Time Derivative of Mean Motion
func (rcv *OMM) MutateMeanMotionDdot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(68, n)
}

/// Position/Velocity Covariance Matrix
/// Reference frame for the covariance matrix
func (rcv *OMM) CovReferenceFrame() RFM.RefFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return RFM.RefFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 23
}

/// Position/Velocity Covariance Matrix
/// Reference frame for the covariance matrix
func (rcv *OMM) MutateCovReferenceFrame(n RFM.RefFrame) bool {
	return rcv._tab.MutateInt8Slot(70, int8(n))
}

/// Covariance matrix [1,1] km**2
func (rcv *OMM) CxX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [1,1] km**2
func (rcv *OMM) MutateCxX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(72, n)
}

/// Covariance matrix [2,1] km**2
func (rcv *OMM) CyX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [2,1] km**2
func (rcv *OMM) MutateCyX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(74, n)
}

/// Covariance matrix [2,2] km**2
func (rcv *OMM) CyY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [2,2] km**2
func (rcv *OMM) MutateCyY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(76, n)
}

/// Covariance matrix [3,1] km**2
func (rcv *OMM) CzX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [3,1] km**2
func (rcv *OMM) MutateCzX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(78, n)
}

/// Covariance matrix [3,2] km**2
func (rcv *OMM) CzY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [3,2] km**2
func (rcv *OMM) MutateCzY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(80, n)
}

/// Covariance matrix [3,3] km**2
func (rcv *OMM) CzZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [3,3] km**2
func (rcv *OMM) MutateCzZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(82, n)
}

/// Covariance matrix [4,1] km**2/s
func (rcv *OMM) CxDotX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [4,1] km**2/s
func (rcv *OMM) MutateCxDotX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(84, n)
}

/// Covariance matrix [4,2] km**2/s
func (rcv *OMM) CxDotY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [4,2] km**2/s
func (rcv *OMM) MutateCxDotY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(86, n)
}

/// Covariance matrix [4,3] km**2/s
func (rcv *OMM) CxDotZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(88))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [4,3] km**2/s
func (rcv *OMM) MutateCxDotZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(88, n)
}

/// Covariance matrix [4,4] km**2/s**2
func (rcv *OMM) CxDotXDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(90))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [4,4] km**2/s**2
func (rcv *OMM) MutateCxDotXDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(90, n)
}

/// Covariance matrix [5,1] km**2/s
func (rcv *OMM) CyDotX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(92))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [5,1] km**2/s
func (rcv *OMM) MutateCyDotX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(92, n)
}

/// Covariance matrix [5,2] km**2/s
func (rcv *OMM) CyDotY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(94))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [5,2] km**2/s
func (rcv *OMM) MutateCyDotY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(94, n)
}

/// Covariance matrix [5,3] km**2/s
func (rcv *OMM) CyDotZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(96))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [5,3] km**2/s
func (rcv *OMM) MutateCyDotZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(96, n)
}

/// Covariance matrix [5,4] km**2/s**2
func (rcv *OMM) CyDotXDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(98))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [5,4] km**2/s**2
func (rcv *OMM) MutateCyDotXDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(98, n)
}

/// Covariance matrix [5,5] km**2/s**2
func (rcv *OMM) CyDotYDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(100))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [5,5] km**2/s**2
func (rcv *OMM) MutateCyDotYDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(100, n)
}

/// Covariance matrix [6,1] km**2/s
func (rcv *OMM) CzDotX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(102))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,1] km**2/s
func (rcv *OMM) MutateCzDotX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(102, n)
}

/// Covariance matrix [6,2] km**2/s
func (rcv *OMM) CzDotY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(104))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,2] km**2/s
func (rcv *OMM) MutateCzDotY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(104, n)
}

/// Covariance matrix [6,3] km**2/s
func (rcv *OMM) CzDotZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,3] km**2/s
func (rcv *OMM) MutateCzDotZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(106, n)
}

/// Covariance matrix [6,4] km**2/s**2
func (rcv *OMM) CzDotXDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,4] km**2/s**2
func (rcv *OMM) MutateCzDotXDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(108, n)
}

/// Covariance matrix [6,5] km**2/s**2
func (rcv *OMM) CzDotYDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,5] km**2/s**2
func (rcv *OMM) MutateCzDotYDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(110, n)
}

/// Covariance matrix [6,6] km**2/s**2
func (rcv *OMM) CzDotZDot() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance matrix [6,6] km**2/s**2
func (rcv *OMM) MutateCzDotZDot(n float64) bool {
	return rcv._tab.MutateFloat64Slot(112, n)
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) UserDefinedBip0044Type() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) MutateUserDefinedBip0044Type(n uint32) bool {
	return rcv._tab.MutateUint32Slot(114, n)
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) UserDefinedObjectDesignator() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// User defined parameter, must be described in an ICD
/// User defined parameter, must be described in an ICD
func (rcv *OMM) UserDefinedEarthModel() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// User defined parameter, must be described in an ICD
/// User defined parameter, must be described in an ICD
func (rcv *OMM) UserDefinedEpochTimestamp() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) MutateUserDefinedEpochTimestamp(n float64) bool {
	return rcv._tab.MutateFloat64Slot(120, n)
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) UserDefinedMicroseconds() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(122))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// User defined parameter, must be described in an ICD
func (rcv *OMM) MutateUserDefinedMicroseconds(n float64) bool {
	return rcv._tab.MutateFloat64Slot(122, n)
}

func OMMStart(builder *flatbuffers.Builder) {
	builder.StartObject(60)
}
func OMMAddCcsdsOmmVers(builder *flatbuffers.Builder, ccsdsOmmVers float64) {
	builder.PrependFloat64Slot(0, ccsdsOmmVers, 0.0)
}
func OMMAddCreationDate(builder *flatbuffers.Builder, creationDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(creationDate), 0)
}
func OMMAddOriginator(builder *flatbuffers.Builder, originator flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(originator), 0)
}
func OMMAddObjectName(builder *flatbuffers.Builder, objectName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(objectName), 0)
}
func OMMAddObjectId(builder *flatbuffers.Builder, objectId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(objectId), 0)
}
func OMMAddCenterName(builder *flatbuffers.Builder, centerName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(centerName), 0)
}
func OMMAddReferenceFrame(builder *flatbuffers.Builder, referenceFrame RFM.RefFrame) {
	builder.PrependInt8Slot(6, int8(referenceFrame), 2)
}
func OMMAddReferenceFrameEpoch(builder *flatbuffers.Builder, referenceFrameEpoch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(referenceFrameEpoch), 0)
}
func OMMAddTimeSystem(builder *flatbuffers.Builder, timeSystem TIM.TimeSystem) {
	builder.PrependInt8Slot(8, int8(timeSystem), 11)
}
func OMMAddMeanElementTheory(builder *flatbuffers.Builder, meanElementTheory MET.MeanElementTheory) {
	builder.PrependInt8Slot(9, int8(meanElementTheory), 0)
}
func OMMAddComment(builder *flatbuffers.Builder, comment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(comment), 0)
}
func OMMAddEpoch(builder *flatbuffers.Builder, epoch flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(epoch), 0)
}
func OMMAddSemiMajorAxis(builder *flatbuffers.Builder, semiMajorAxis float64) {
	builder.PrependFloat64Slot(12, semiMajorAxis, 0.0)
}
func OMMAddMeanMotion(builder *flatbuffers.Builder, meanMotion float64) {
	builder.PrependFloat64Slot(13, meanMotion, 0.0)
}
func OMMAddEccentricity(builder *flatbuffers.Builder, eccentricity float64) {
	builder.PrependFloat64Slot(14, eccentricity, 0.0)
}
func OMMAddInclination(builder *flatbuffers.Builder, inclination float64) {
	builder.PrependFloat64Slot(15, inclination, 0.0)
}
func OMMAddRaOfAscNode(builder *flatbuffers.Builder, raOfAscNode float64) {
	builder.PrependFloat64Slot(16, raOfAscNode, 0.0)
}
func OMMAddArgOfPericenter(builder *flatbuffers.Builder, argOfPericenter float64) {
	builder.PrependFloat64Slot(17, argOfPericenter, 0.0)
}
func OMMAddMeanAnomaly(builder *flatbuffers.Builder, meanAnomaly float64) {
	builder.PrependFloat64Slot(18, meanAnomaly, 0.0)
}
func OMMAddGm(builder *flatbuffers.Builder, gm float64) {
	builder.PrependFloat64Slot(19, gm, 0.0)
}
func OMMAddMass(builder *flatbuffers.Builder, mass float64) {
	builder.PrependFloat64Slot(20, mass, 0.0)
}
func OMMAddSolarRadArea(builder *flatbuffers.Builder, solarRadArea float64) {
	builder.PrependFloat64Slot(21, solarRadArea, 0.0)
}
func OMMAddSolarRadCoeff(builder *flatbuffers.Builder, solarRadCoeff float64) {
	builder.PrependFloat64Slot(22, solarRadCoeff, 0.0)
}
func OMMAddDragArea(builder *flatbuffers.Builder, dragArea float64) {
	builder.PrependFloat64Slot(23, dragArea, 0.0)
}
func OMMAddDragCoeff(builder *flatbuffers.Builder, dragCoeff float64) {
	builder.PrependFloat64Slot(24, dragCoeff, 0.0)
}
func OMMAddEphemerisType(builder *flatbuffers.Builder, ephemerisType ephemerisType) {
	builder.PrependInt8Slot(25, int8(ephemerisType), 1)
}
func OMMAddClassificationType(builder *flatbuffers.Builder, classificationType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(26, flatbuffers.UOffsetT(classificationType), 0)
}
func OMMAddNoradCatId(builder *flatbuffers.Builder, noradCatId uint32) {
	builder.PrependUint32Slot(27, noradCatId, 0)
}
func OMMAddElementSetNo(builder *flatbuffers.Builder, elementSetNo uint32) {
	builder.PrependUint32Slot(28, elementSetNo, 0)
}
func OMMAddRevAtEpoch(builder *flatbuffers.Builder, revAtEpoch float64) {
	builder.PrependFloat64Slot(29, revAtEpoch, 0.0)
}
func OMMAddBstar(builder *flatbuffers.Builder, bstar float64) {
	builder.PrependFloat64Slot(30, bstar, 0.0)
}
func OMMAddMeanMotionDot(builder *flatbuffers.Builder, meanMotionDot float64) {
	builder.PrependFloat64Slot(31, meanMotionDot, 0.0)
}
func OMMAddMeanMotionDdot(builder *flatbuffers.Builder, meanMotionDdot float64) {
	builder.PrependFloat64Slot(32, meanMotionDdot, 0.0)
}
func OMMAddCovReferenceFrame(builder *flatbuffers.Builder, covReferenceFrame RFM.RefFrame) {
	builder.PrependInt8Slot(33, int8(covReferenceFrame), 23)
}
func OMMAddCxX(builder *flatbuffers.Builder, cxX float64) {
	builder.PrependFloat64Slot(34, cxX, 0.0)
}
func OMMAddCyX(builder *flatbuffers.Builder, cyX float64) {
	builder.PrependFloat64Slot(35, cyX, 0.0)
}
func OMMAddCyY(builder *flatbuffers.Builder, cyY float64) {
	builder.PrependFloat64Slot(36, cyY, 0.0)
}
func OMMAddCzX(builder *flatbuffers.Builder, czX float64) {
	builder.PrependFloat64Slot(37, czX, 0.0)
}
func OMMAddCzY(builder *flatbuffers.Builder, czY float64) {
	builder.PrependFloat64Slot(38, czY, 0.0)
}
func OMMAddCzZ(builder *flatbuffers.Builder, czZ float64) {
	builder.PrependFloat64Slot(39, czZ, 0.0)
}
func OMMAddCxDotX(builder *flatbuffers.Builder, cxDotX float64) {
	builder.PrependFloat64Slot(40, cxDotX, 0.0)
}
func OMMAddCxDotY(builder *flatbuffers.Builder, cxDotY float64) {
	builder.PrependFloat64Slot(41, cxDotY, 0.0)
}
func OMMAddCxDotZ(builder *flatbuffers.Builder, cxDotZ float64) {
	builder.PrependFloat64Slot(42, cxDotZ, 0.0)
}
func OMMAddCxDotXDot(builder *flatbuffers.Builder, cxDotXDot float64) {
	builder.PrependFloat64Slot(43, cxDotXDot, 0.0)
}
func OMMAddCyDotX(builder *flatbuffers.Builder, cyDotX float64) {
	builder.PrependFloat64Slot(44, cyDotX, 0.0)
}
func OMMAddCyDotY(builder *flatbuffers.Builder, cyDotY float64) {
	builder.PrependFloat64Slot(45, cyDotY, 0.0)
}
func OMMAddCyDotZ(builder *flatbuffers.Builder, cyDotZ float64) {
	builder.PrependFloat64Slot(46, cyDotZ, 0.0)
}
func OMMAddCyDotXDot(builder *flatbuffers.Builder, cyDotXDot float64) {
	builder.PrependFloat64Slot(47, cyDotXDot, 0.0)
}
func OMMAddCyDotYDot(builder *flatbuffers.Builder, cyDotYDot float64) {
	builder.PrependFloat64Slot(48, cyDotYDot, 0.0)
}
func OMMAddCzDotX(builder *flatbuffers.Builder, czDotX float64) {
	builder.PrependFloat64Slot(49, czDotX, 0.0)
}
func OMMAddCzDotY(builder *flatbuffers.Builder, czDotY float64) {
	builder.PrependFloat64Slot(50, czDotY, 0.0)
}
func OMMAddCzDotZ(builder *flatbuffers.Builder, czDotZ float64) {
	builder.PrependFloat64Slot(51, czDotZ, 0.0)
}
func OMMAddCzDotXDot(builder *flatbuffers.Builder, czDotXDot float64) {
	builder.PrependFloat64Slot(52, czDotXDot, 0.0)
}
func OMMAddCzDotYDot(builder *flatbuffers.Builder, czDotYDot float64) {
	builder.PrependFloat64Slot(53, czDotYDot, 0.0)
}
func OMMAddCzDotZDot(builder *flatbuffers.Builder, czDotZDot float64) {
	builder.PrependFloat64Slot(54, czDotZDot, 0.0)
}
func OMMAddUserDefinedBip0044Type(builder *flatbuffers.Builder, userDefinedBip0044Type uint32) {
	builder.PrependUint32Slot(55, userDefinedBip0044Type, 0)
}
func OMMAddUserDefinedObjectDesignator(builder *flatbuffers.Builder, userDefinedObjectDesignator flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(56, flatbuffers.UOffsetT(userDefinedObjectDesignator), 0)
}
func OMMAddUserDefinedEarthModel(builder *flatbuffers.Builder, userDefinedEarthModel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(57, flatbuffers.UOffsetT(userDefinedEarthModel), 0)
}
func OMMAddUserDefinedEpochTimestamp(builder *flatbuffers.Builder, userDefinedEpochTimestamp float64) {
	builder.PrependFloat64Slot(58, userDefinedEpochTimestamp, 0.0)
}
func OMMAddUserDefinedMicroseconds(builder *flatbuffers.Builder, userDefinedMicroseconds float64) {
	builder.PrependFloat64Slot(59, userDefinedMicroseconds, 0.0)
}
func OMMEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
