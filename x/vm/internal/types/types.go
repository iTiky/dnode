// Basic constants and function to work with types.
package types

import (
	"encoding/hex"
	"fmt"
	"unicode/utf8"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/crypto/sha3"

	"github.com/dfinance/dvm-proto/go/vm_grpc"

	"github.com/dfinance/dnode/x/common_vm"
)

const (
	ModuleName = "vm"

	StoreKey  = ModuleName
	RouterKey = ModuleName

	Codespace         types.CodespaceType = ModuleName
	DefaultParamspace                     = ModuleName

	VmAddressLength = 32
	VmGasPrice      = 1
	VmUnknowTagType = -1
	zeroBytes       = 12
)

// VM related variables.
var (
	KeyGenesis = []byte("gen") // used to save genesis
)

// Type of Move contract (bytes).
type Contract []byte

// Convert bech32 to libra hex.
func Bech32ToLibra(acc types.AccAddress) string {
	prefix := types.GetConfig().GetBech32AccountAddrPrefix()
	zeros := make([]byte, zeroBytes-len(prefix))

	bytes := make([]byte, 0)
	bytes = append(bytes, []byte(prefix)...)
	bytes = append(bytes, zeros...)
	bytes = append(bytes, acc...)

	return hex.EncodeToString(bytes)
}

// Convert VMAccessPath to hex string
func PathToHex(path vm_grpc.VMAccessPath) string {
	return fmt.Sprintf("Access path: \n"+
		"\tAddress: %s\n"+
		"\tPath:    %s\n"+
		"\tKey:     %s\n", hex.EncodeToString(path.Address), hex.EncodeToString(path.Path), hex.EncodeToString(common_vm.MakePathKey(path)))
}

// Get TypeTag by string TypeTag representation.
func GetVMTypeByString(typeTag string) (vm_grpc.VMTypeTag, error) {
	if val, ok := vm_grpc.VMTypeTag_value[typeTag]; !ok {
		return VmUnknowTagType, fmt.Errorf("can't find tag type %s, check correctness of type value", typeTag)
	} else {
		return vm_grpc.VMTypeTag(val), nil
	}
}

// Convert TypeTag to string representation.
func VMTypeToString(tag vm_grpc.VMTypeTag) (string, error) {
	if val, ok := vm_grpc.VMTypeTag_name[int32(tag)]; !ok {
		return "", fmt.Errorf("can't find string representation of type %d, check correctness of type value", tag)
	} else {
		return val, nil
	}
}

// Convert TypeTag to string representation with panic.
func VMTypeToStringPanic(tag vm_grpc.VMTypeTag) string {
	if val, ok := vm_grpc.VMTypeTag_name[int32(tag)]; !ok {
		panic(fmt.Errorf("can't find string representation of type %d, check correctness of type value", tag))
	} else {
		return val
	}
}

// Convert asset code to libra path.
func AssetCodeToPath(assetCode string) []byte {
	asciiBytes := StringToAsciiBytes(assetCode)
	for i, b := range asciiBytes {
		if b >= 'A' && b <= 'Z' {
			asciiBytes[i] = b | ((IsAsciiUpperCase(b)) << 5)
		}
	}

	hasher := sha3.New256()
	hasher.Write(asciiBytes)
	hash := hasher.Sum(nil)

	res := make([]byte, len(hash)+1)
	res[0] = 255
	for i, b := range hash {
		res[i+1] = b
	}

	return res
}

// If ascii is upper case to u8.
func IsAsciiUpperCase(b byte) uint8 {
	if b >= 'A' && b <= 'Z' {
		return 1
	} else {
		return 0
	}
}

// Convert string to ascii bytes
func StringToAsciiBytes(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		i++
	}
	return t
}
