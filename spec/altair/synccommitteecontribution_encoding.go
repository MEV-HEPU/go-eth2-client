// Code generated by fastssz. DO NOT EDIT.
// Hash: 29e1a28a877a613dfb7b84c72471b1e26d6f7f159bb6d1de6a7a3a2ef499aa7a
package altair

import (
	"github.com/MEV-HEPU/go-eth2-client/spec/phase0"
	ssz "github.com/MEV-HEPU/fastssz"
)

// MarshalSSZ ssz marshals the SyncCommitteeContribution object
func (s *SyncCommitteeContribution) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncCommitteeContribution object to a target array
func (s *SyncCommitteeContribution) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(s.Slot))

	// Field (1) 'BeaconBlockRoot'
	dst = append(dst, s.BeaconBlockRoot[:]...)

	// Field (2) 'SubcommitteeIndex'
	dst = ssz.MarshalUint64(dst, s.SubcommitteeIndex)

	// Field (3) 'AggregationBits'
	if size := len(s.AggregationBits); size != 16 {
		err = ssz.ErrBytesLengthFn("SyncCommitteeContribution.AggregationBits", size, 16)
		return
	}
	dst = append(dst, s.AggregationBits...)

	// Field (4) 'Signature'
	dst = append(dst, s.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncCommitteeContribution object
func (s *SyncCommitteeContribution) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 160 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	s.Slot = phase0.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'BeaconBlockRoot'
	copy(s.BeaconBlockRoot[:], buf[8:40])

	// Field (2) 'SubcommitteeIndex'
	s.SubcommitteeIndex = ssz.UnmarshallUint64(buf[40:48])

	// Field (3) 'AggregationBits'
	if cap(s.AggregationBits) == 0 {
		s.AggregationBits = make([]byte, 0, len(buf[48:64]))
	}
	s.AggregationBits = append(s.AggregationBits, buf[48:64]...)

	// Field (4) 'Signature'
	copy(s.Signature[:], buf[64:160])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncCommitteeContribution object
func (s *SyncCommitteeContribution) SizeSSZ() (size int) {
	size = 160
	return
}

// HashTreeRoot ssz hashes the SyncCommitteeContribution object
func (s *SyncCommitteeContribution) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncCommitteeContribution object with a hasher
func (s *SyncCommitteeContribution) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(s.Slot))

	// Field (1) 'BeaconBlockRoot'
	hh.PutBytes(s.BeaconBlockRoot[:])

	// Field (2) 'SubcommitteeIndex'
	hh.PutUint64(s.SubcommitteeIndex)

	// Field (3) 'AggregationBits'
	if size := len(s.AggregationBits); size != 16 {
		err = ssz.ErrBytesLengthFn("SyncCommitteeContribution.AggregationBits", size, 16)
		return
	}
	hh.PutBytes(s.AggregationBits)

	// Field (4) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SyncCommitteeContribution object
func (s *SyncCommitteeContribution) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}