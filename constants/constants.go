package constants

/*
| === Instruction classes === | https://www.kernel.org/doc/html/v5.17/bpf/instruction-set.html#instruction-classes
*/
const (
	BPF_LD    = iota // 0x00, non-standard load operations
	BPF_LDX          // 0x01, load into register operations
	BPF_ST           // 0x02, store from immediate operations
	BPF_STX          // 0x03, store from register operations
	BPF_ALU          // 0x04, 32-bit arithmetic operations
	BPF_JMP          // 0x05, 64-bit jump operations
	BPF_JMP32        // 0x06, 32-bit jump operations
	BPF_ALU64        // 0x07, 64-bit arithmetic operations
)

/*
| === Load and store instructions === | https://www.kernel.org/doc/html/v5.17/bpf/instruction-set.html#load-and-store-instructions
BPF_LD, BPF_LDX, BPF_ST and BPF_STX
Opcode is interpreted as follows:

(MSB)						(LSB)
7            4      2          0
+-----------+------+-----------+
| Mode      | Size | Class     |
+-----------+------+-----------+
*/
// Size modifier
const (
	BPF_W  = 0x00 // word (4 bytes)
	BPF_H  = 0x08 // half word (2 bytes)
	BPF_B  = 0x10 // byte
	BPF_DW = 0x18 // double word (8 bytes)
)

// Mode modifier
const (
	BPF_IMM    = 0x00 // used for 64-bit mov
	BPF_ABS    = 0x20 // legacy BPF packet access
	BPF_IND    = 0x40 // legacy BPF packet access
	BPF_MEM    = 0x60 // all normal load and store operations
	BPF_ATOMIC = 0xc0 // atomic operations
)

/*
| === Arithmetic and jump instructions === | https://www.kernel.org/doc/html/v5.17/bpf/instruction-set.html#arithmetic-and-jump-instructions

(MSB)						      (LSB)
8                 3        2          0
+----------------+--------+-----------+
| Operation code | Source | Class     |
+----------------+--------+-----------+
*/

// Source Operand
const (
	BPF_K     = 0x00 // use 32-bit immediate as source operand
	BPF_X     = 0x08 // use ‘src_reg’ register as source operand
	BPF_TO_LE = 0x00 // convert to little-endian
	BPF_TO_BE = 0x08 // convert to big-endian
)

// Operation code
const (
	BPF_ADD  = 0x00 // dst += src
	BPF_SUB  = 0x10 // dst -= src
	BPF_MUL  = 0x20 // dst *= src
	BPF_DIV  = 0x30 // dst /= src
	BPF_OR   = 0x40 // dst |= src
	BPF_AND  = 0x50 // dst &= src
	BPF_LSH  = 0x60 // dst <<= src
	BPF_RSH  = 0x70 // dst >>= src
	BPF_NEG  = 0x80 // dst = ~src
	BPF_MOD  = 0x90 // dst %= src
	BPF_XOR  = 0xa0 // dst ^= src
	BPF_MOV  = 0xb0 // dst = src
	BPF_ARSH = 0xc0 // sign extending shift right
	BPF_END  = 0xd0 // endianness conversion
)

// Jump instructions
const (
	BPF_JA   = 0x00 // PC += off (BPF_JMP only)
	BPF_JEQ  = 0x10 // PC += off if dst == src
	BPF_JGT  = 0x20 // PC += off if dst > src (unsigned)
	BPF_JGE  = 0x30 // PC += off if dst >= src (unsigned)
	BPF_JSET = 0x40 // PC += off if dst & src
	BPF_JNE  = 0x50 // PC += off if dst != src
	BPF_JSGT = 0x60 // PC += off if dst > src (signed)
	BPF_JSGE = 0x70 // PC += off if dst >= src (signed)
	BPF_CALL = 0x80 // function call
	BPF_EXIT = 0x90 // function / program return (BPF_JMP only)
	BPF_JLT  = 0xa0 // PC += off if dst < src (unsigned)
	BPF_JLE  = 0xb0 // PC += off if dst <= src (unsigned)
	BPF_JSLT = 0xc0 // PC += off if dst < src (signed)
	BPF_JSLE = 0xd0 // PC += off if dst <= src (signed)
)
