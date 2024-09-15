package fileheader

import (
	"fmt"
)

// SetVersion sets the version using an unsigned 4-byte integer
func (fv *FileVersion) SetVersion(version uint32) {
	fv.mm = uint8((version & 0xff000000) >> 24)
	fv.nn = uint8((version & 0x00ff0000) >> 16)
	fv.pp = uint8((version & 0x0000ff00) >> 8)
	fv.rr = uint8(version & 0x000000ff)
}

// SetVersionWithDetails sets the version using individual components
func (fv *FileVersion) SetVersionWithDetails(mm, nn, pp, rr uint8) {
	fv.mm = mm
	fv.nn = nn
	fv.pp = pp
	fv.rr = rr
}

// GetVersion retrieves the version as an unsigned 4-byte integer
func (fv *FileVersion) GetVersion() uint32 {
	var version uint32
	version |= uint32(fv.mm) << 24
	version |= uint32(fv.nn) << 16
	version |= uint32(fv.pp) << 8
	version |= uint32(fv.rr)
	return version
}

// GetMM returns the MM component of the version
func (fv *FileVersion) GetMM() uint8 {
	return fv.mm
}

// GetNN returns the NN component of the version
func (fv *FileVersion) GetNN() uint8 {
	return fv.nn
}

// GetPP returns the PP component of the version
func (fv *FileVersion) GetPP() uint8 {
	return fv.pp
}

// GetRR returns the RR component of the version
func (fv *FileVersion) GetRR() uint8 {
	return fv.rr
}

// IsOver checks if the current version is greater than the given version
func (fv *FileVersion) IsOver(mm2, nn2, pp2, rr2 uint8) bool {
	if fv.mm > mm2 {
		return true
	}
	if fv.mm == mm2 {
		if fv.nn > nn2 {
			return true
		}
		if fv.nn == nn2 {
			if fv.pp > pp2 {
				return true
			}
			if fv.pp == pp2 {
				if fv.rr > rr2 {
					return true
				}
			}
		}
	}
	return false
}

// ToString returns the version as a string
func (fv *FileVersion) ToString() string {
	return fmt.Sprintf("%d.%d.%d.%d", fv.mm, fv.nn, fv.pp, fv.rr)
}

// Copy copies the version from another FileVersion instance
func (fv *FileVersion) Copy(from *FileVersion) {
	fv.mm = from.mm
	fv.nn = from.nn
	fv.pp = from.pp
	fv.rr = from.rr
}
