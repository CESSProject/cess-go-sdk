package rule

// byte size
const (
	SIZE_1KiB  = 1024
	SIZE_1MiB  = 1024 * SIZE_1KiB
	SIZE_1GiB  = 1024 * SIZE_1MiB
	SIZE_SLICE = 512 * SIZE_1MiB
)

const SegmentSize = 16 * SIZE_1MiB
const FragmentSize = 8 * SIZE_1MiB

const DataShards = 2
const ParShards = 1
