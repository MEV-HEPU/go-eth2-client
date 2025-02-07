// Code generated by fastssz. DO NOT EDIT.
// Hash: 83015bdd2d9bc91a6f5ec24ad922c254bc9bec576b11b7bbd734ae1ba22f4e72
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBlockContents object
func (s *SignedBlockContents) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBlockContents object to a target array
func (s *SignedBlockContents) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'SignedBlock'
	dst = ssz.WriteOffset(dst, offset)
	if s.SignedBlock == nil {
		s.SignedBlock = new(deneb.SignedBeaconBlock)
	}
	offset += s.SignedBlock.SizeSSZ()

	// Offset (1) 'SignedBlobSidecars'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.SignedBlobSidecars) * 131352

	// Field (0) 'SignedBlock'
	if dst, err = s.SignedBlock.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'SignedBlobSidecars'
	if size := len(s.SignedBlobSidecars); size > 6 {
		err = ssz.ErrListTooBigFn("SignedBlockContents.SignedBlobSidecars", size, 6)
		return
	}
	for ii := 0; ii < len(s.SignedBlobSidecars); ii++ {
		if dst, err = s.SignedBlobSidecars[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBlockContents object
func (s *SignedBlockContents) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'SignedBlock'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'SignedBlobSidecars'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'SignedBlock'
	{
		buf = tail[o0:o1]
		if s.SignedBlock == nil {
			s.SignedBlock = new(deneb.SignedBeaconBlock)
		}
		if err = s.SignedBlock.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'SignedBlobSidecars'
	{
		buf = tail[o1:]
		num, err := ssz.DivideInt2(len(buf), 131352, 6)
		if err != nil {
			return err
		}
		s.SignedBlobSidecars = make([]*deneb.SignedBlobSidecar, num)
		for ii := 0; ii < num; ii++ {
			if s.SignedBlobSidecars[ii] == nil {
				s.SignedBlobSidecars[ii] = new(deneb.SignedBlobSidecar)
			}
			if err = s.SignedBlobSidecars[ii].UnmarshalSSZ(buf[ii*131352 : (ii+1)*131352]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBlockContents object
func (s *SignedBlockContents) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'SignedBlock'
	if s.SignedBlock == nil {
		s.SignedBlock = new(deneb.SignedBeaconBlock)
	}
	size += s.SignedBlock.SizeSSZ()

	// Field (1) 'SignedBlobSidecars'
	size += len(s.SignedBlobSidecars) * 131352

	return
}

// HashTreeRoot ssz hashes the SignedBlockContents object
func (s *SignedBlockContents) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBlockContents object with a hasher
func (s *SignedBlockContents) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'SignedBlock'
	if err = s.SignedBlock.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'SignedBlobSidecars'
	{
		subIndx := hh.Index()
		num := uint64(len(s.SignedBlobSidecars))
		if num > 6 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.SignedBlobSidecars {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 6)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedBlockContents object
func (s *SignedBlockContents) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
