package chshare

//ProtocolVersion of chisel. When backwards
//incompatible changes are made, this will
//be incremented to signify a protocol
//mismatch.
//
// ProtocolName is the protocol identifier prefix (defaults to "chisel").
// ProtocolVersion is derived as ProtocolName + "-v3". The "v3" suffix
// remains hardcoded so that incompatible changes can still be signalled by
// bumping the suffix in future versions.
var ProtocolName = "chisel"
var ProtocolVersion = ProtocolName + "-v3"

// SetProtocolName allows overriding the protocol name portion (e.g. to use
// "myapp" so the full protocol becomes "myapp-v3"). Passing an empty string
// is a no-op.
func SetProtocolName(name string) {
	if name == "" {
		return
	}
	ProtocolName = name
	ProtocolVersion = ProtocolName + "-v3"
}

var BuildVersion = "0.0.0-src"
