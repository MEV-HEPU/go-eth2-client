// Code generated by fastssz. DO NOT EDIT.
// Hash: e2f2043cae205eca9d20c3884703ceb64f23fda19f13beff0138f7b9bffa2107
package phase0

import (
	ssz "github.com/MEV-HEPU/fastssz"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(2687369)

	// Field (0) 'GenesisTime'
	dst = ssz.MarshalUint64(dst, b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if size := len(b.GenesisValidatorsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconState.GenesisValidatorsRoot", size, 32)
		return
	}
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if dst, err = b.Fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	if size := len(b.BlockRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if size := len(b.BlockRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.BlockRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.BlockRoots[ii][:]...)
	}

	// Field (6) 'StateRoots'
	if size := len(b.StateRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if size := len(b.StateRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.StateRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.StateRoots[ii][:]...)
	}

	// Offset (7) 'HistoricalRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.HistoricalRoots) * 32

	// Field (8) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(ETH1Data)
	}
	if dst, err = b.ETH1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (9) 'ETH1DataVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ETH1DataVotes) * 72

	// Offset (10) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (11) 'Balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Balances) * 8

	// Field (12) 'RANDAOMixes'
	if size := len(b.RANDAOMixes); size != 65536 {
		err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
		return
	}
	for ii := 0; ii < 65536; ii++ {
		if size := len(b.RANDAOMixes[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.RANDAOMixes[ii]", size, 32)
			return
		}
		dst = append(dst, b.RANDAOMixes[ii][:]...)
	}

	// Field (13) 'Slashings'
	if size := len(b.Slashings); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Slashings[ii]))
	}

	// Offset (14) 'PreviousEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		offset += 4
		offset += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Offset (15) 'CurrentEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		offset += 4
		offset += b.CurrentEpochAttestations[ii].SizeSSZ()
	}

	// Field (16) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	dst = append(dst, b.JustificationBits...)

	// Field (17) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.PreviousJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (18) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.CurrentJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (19) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(Checkpoint)
	}
	if dst, err = b.FinalizedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'HistoricalRoots'
	if size := len(b.HistoricalRoots); size > 16777216 {
		err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
		return
	}
	for ii := 0; ii < len(b.HistoricalRoots); ii++ {
		if size := len(b.HistoricalRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.HistoricalRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.HistoricalRoots[ii][:]...)
	}

	// Field (9) 'ETH1DataVotes'
	if size := len(b.ETH1DataVotes); size > 1024 {
		err = ssz.ErrListTooBigFn("BeaconState.ETH1DataVotes", size, 1024)
		return
	}
	for ii := 0; ii < len(b.ETH1DataVotes); ii++ {
		if dst, err = b.ETH1DataVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (10) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (11) 'Balances'
	if size := len(b.Balances); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Balances); ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Balances[ii]))
	}

	// Field (14) 'PreviousEpochAttestations'
	if size := len(b.PreviousEpochAttestations); size > 4096 {
		err = ssz.ErrListTooBigFn("BeaconState.PreviousEpochAttestations", size, 4096)
		return
	}
	{
		offset = 4 * len(b.PreviousEpochAttestations)
		for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.PreviousEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		if dst, err = b.PreviousEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (15) 'CurrentEpochAttestations'
	if size := len(b.CurrentEpochAttestations); size > 4096 {
		err = ssz.ErrListTooBigFn("BeaconState.CurrentEpochAttestations", size, 4096)
		return
	}
	{
		offset = 4 * len(b.CurrentEpochAttestations)
		for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.CurrentEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		if dst, err = b.CurrentEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 2687369 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o9, o10, o11, o14, o15 uint64

	// Field (0) 'GenesisTime'
	b.GenesisTime = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[8:40])

	// Field (2) 'Slot'
	b.Slot = Slot(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(Fork)
	}
	if err = b.Fork.UnmarshalSSZ(buf[48:64]); err != nil {
		return err
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[64:176]); err != nil {
		return err
	}

	// Field (5) 'BlockRoots'
	b.BlockRoots = make([]Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.BlockRoots[ii][:], buf[176:262320][ii*32:(ii+1)*32])
	}

	// Field (6) 'StateRoots'
	b.StateRoots = make([]Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.StateRoots[ii][:], buf[262320:524464][ii*32:(ii+1)*32])
	}

	// Offset (7) 'HistoricalRoots'
	if o7 = ssz.ReadOffset(buf[524464:524468]); o7 > size {
		return ssz.ErrOffset
	}

	if o7 < 2687369 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (8) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(ETH1Data)
	}
	if err = b.ETH1Data.UnmarshalSSZ(buf[524468:524540]); err != nil {
		return err
	}

	// Offset (9) 'ETH1DataVotes'
	if o9 = ssz.ReadOffset(buf[524540:524544]); o9 > size || o7 > o9 {
		return ssz.ErrOffset
	}

	// Offset (10) 'Validators'
	if o10 = ssz.ReadOffset(buf[524544:524548]); o10 > size || o9 > o10 {
		return ssz.ErrOffset
	}

	// Offset (11) 'Balances'
	if o11 = ssz.ReadOffset(buf[524548:524552]); o11 > size || o10 > o11 {
		return ssz.ErrOffset
	}

	// Field (12) 'RANDAOMixes'
	b.RANDAOMixes = make([]Root, 65536)
	for ii := 0; ii < 65536; ii++ {
		copy(b.RANDAOMixes[ii][:], buf[524560:2621712][ii*32:(ii+1)*32])
	}

	// Field (13) 'Slashings'
	b.Slashings = make([]Gwei, 8192)
	for ii := 0; ii < 8192; ii++ {
		b.Slashings[ii] = Gwei(ssz.UnmarshallUint64(buf[2621712:2687248][ii*8 : (ii+1)*8]))
	}

	// Offset (14) 'PreviousEpochAttestations'
	if o14 = ssz.ReadOffset(buf[2687240:2687244]); o14 > size || o11 > o14 {
		return ssz.ErrOffset
	}

	// Offset (15) 'CurrentEpochAttestations'
	if o15 = ssz.ReadOffset(buf[2687244:2687248]); o15 > size || o14 > o15 {
		return ssz.ErrOffset
	}

	// Field (16) 'JustificationBits'
	if cap(b.JustificationBits) == 0 {
		b.JustificationBits = make([]byte, 0, len(buf[2687248:2687249]))
	}
	b.JustificationBits = append(b.JustificationBits, buf[2687248:2687249]...)

	// Field (17) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.UnmarshalSSZ(buf[2687249:2687289]); err != nil {
		return err
	}

	// Field (18) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.UnmarshalSSZ(buf[2687289:2687329]); err != nil {
		return err
	}

	// Field (19) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(Checkpoint)
	}
	if err = b.FinalizedCheckpoint.UnmarshalSSZ(buf[2687329:2687369]); err != nil {
		return err
	}

	// Field (7) 'HistoricalRoots'
	{
		buf = tail[o7:o9]
		num, err := ssz.DivideInt2(len(buf), 32, 16777216)
		if err != nil {
			return err
		}
		b.HistoricalRoots = make([]Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.HistoricalRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (9) 'ETH1DataVotes'
	{
		buf = tail[o9:o10]
		num, err := ssz.DivideInt2(len(buf), 72, 1024)
		if err != nil {
			return err
		}
		b.ETH1DataVotes = make([]*ETH1Data, num)
		for ii := 0; ii < num; ii++ {
			if b.ETH1DataVotes[ii] == nil {
				b.ETH1DataVotes[ii] = new(ETH1Data)
			}
			if err = b.ETH1DataVotes[ii].UnmarshalSSZ(buf[ii*72 : (ii+1)*72]); err != nil {
				return err
			}
		}
	}

	// Field (10) 'Validators'
	{
		buf = tail[o10:o11]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (11) 'Balances'
	{
		buf = tail[o11:o14]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Balances = make([]Gwei, num)
		for ii := 0; ii < num; ii++ {
			b.Balances[ii] = Gwei(ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8]))
		}
	}

	// Field (14) 'PreviousEpochAttestations'
	{
		buf = tail[o14:o15]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.PreviousEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.PreviousEpochAttestations[indx] == nil {
				b.PreviousEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.PreviousEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (15) 'CurrentEpochAttestations'
	{
		buf = tail[o15:]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.CurrentEpochAttestations = make([]*PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.CurrentEpochAttestations[indx] == nil {
				b.CurrentEpochAttestations[indx] = new(PendingAttestation)
			}
			if err = b.CurrentEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 2687369

	// Field (7) 'HistoricalRoots'
	size += len(b.HistoricalRoots) * 32

	// Field (9) 'ETH1DataVotes'
	size += len(b.ETH1DataVotes) * 72

	// Field (10) 'Validators'
	size += len(b.Validators) * 121

	// Field (11) 'Balances'
	size += len(b.Balances) * 8

	// Field (14) 'PreviousEpochAttestations'
	for ii := 0; ii < len(b.PreviousEpochAttestations); ii++ {
		size += 4
		size += b.PreviousEpochAttestations[ii].SizeSSZ()
	}

	// Field (15) 'CurrentEpochAttestations'
	for ii := 0; ii < len(b.CurrentEpochAttestations); ii++ {
		size += 4
		size += b.CurrentEpochAttestations[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the BeaconState object
func (b *BeaconState) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher
func (b *BeaconState) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisTime'
	hh.PutUint64(b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if size := len(b.GenesisValidatorsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconState.GenesisValidatorsRoot", size, 32)
		return
	}
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (2) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (3) 'Fork'
	if err = b.Fork.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	{
		if size := len(b.BlockRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (6) 'StateRoots'
	{
		if size := len(b.StateRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (7) 'HistoricalRoots'
	{
		if size := len(b.HistoricalRoots); size > 16777216 {
			err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.HistoricalRoots {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i[:])
		}
		numItems := uint64(len(b.HistoricalRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(16777216, numItems, 32))
	}

	// Field (8) 'ETH1Data'
	if err = b.ETH1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (9) 'ETH1DataVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.ETH1DataVotes))
		if num > 1024 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.ETH1DataVotes {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1024)
	}

	// Field (10) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (11) 'Balances'
	{
		if size := len(b.Balances); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Balances {
			hh.AppendUint64(uint64(i))
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Balances))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (12) 'RANDAOMixes'
	{
		if size := len(b.RANDAOMixes); size != 65536 {
			err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RANDAOMixes {
			if len(i) != 32 {
				err = ssz.ErrBytesLength
				return
			}
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (13) 'Slashings'
	{
		if size := len(b.Slashings); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Slashings {
			hh.AppendUint64(uint64(i))
		}
		hh.Merkleize(subIndx)
	}

	// Field (14) 'PreviousEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.PreviousEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.PreviousEpochAttestations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (15) 'CurrentEpochAttestations'
	{
		subIndx := hh.Index()
		num := uint64(len(b.CurrentEpochAttestations))
		if num > 4096 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.CurrentEpochAttestations {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 4096)
	}

	// Field (16) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	hh.PutBytes(b.JustificationBits)

	// Field (17) 'PreviousJustifiedCheckpoint'
	if err = b.PreviousJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (18) 'CurrentJustifiedCheckpoint'
	if err = b.CurrentJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (19) 'FinalizedCheckpoint'
	if err = b.FinalizedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconState object
func (b *BeaconState) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}