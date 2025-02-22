package bbTypes

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	dst = append(dst, s.Root[:]...)

	// Field (1) 'Domain'
	dst = append(dst, s.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	copy(s.Root[:], buf[0:32])

	// Field (1) 'Domain'
	copy(s.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	hh.PutBytes(s.Root[:])

	// Field (1) 'Domain'
	hh.PutBytes(s.Domain[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SigningData object
func (s *SigningData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ForkData object
func (f *ForkData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(f)
}