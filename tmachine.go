package main

import "fmt"

const h = 3      // height of the instruction board
const w = 13     // width of the instruction board
const tlen = 32  // width of the tape

// tape and turing machine state
type tmachine struct {
	tape [tlen]byte
	pos int
	prog [w*h]byte // the 'instruction board'
	pc int
}

// 0xf0: instruction type
const (
	Move = (1+iota) << 4 // move the head
	Back                 // go left on the instruction board
	Climb                // go up/down on the instruction board
	Branch               // conditinoally go up/down on the instruction board
	Write                // write a blank, a zero, or a one
)

// 0x0f: value bits, with instruction-specific meaning
const (
	// for Move:
	Left = 0   
	Right = 1

	// For Branch and Write:
	Blank = 0
	Zero = 1
	One = 2

	// For Climb and Branch:
	Up = 0
	Down = 4

	// Noop
	None = 0
)

// sends rightmost bit to {-1, 1}
func sgn(n byte) int {
	return ((int(n) & 1) << 1) - 1
}

func (p *tmachine) Step() {
	t := p.tape[p.pos]
	s := p.prog[p.pc]
	i := s & 0xf0
	v := s & 0x0f

	switch i {
	case None:
		p.pc++
	case Move:
		p.pos += sgn(v)
		p.pc++
	case Back:
		p.pc  -= int(v) 
	case Climb:
		p.pc += sgn(v >> 2) * w
	case Branch:
		if t == v & 3 {
			p.pc += sgn(v >> 2) * w
		} else {
			p.pc++
		}
	case Write:
		p.tape[p.pos] = v & 3
		p.pc++
	default:
		panic("unknown instruction")
	}
}

var blank01 = []string{"□", "0", "1"}
var cross01 = []string{"×", "0", "1"}
var leftright = []string{"←", "→"}
var updown = []string{"↑", "↓"}

func (p *tmachine) String() string {

	// describe the tape
	var tapesyms [tlen]byte
	for i, t := range p.tape {
		tapesyms[i] = " 01"[t]
	}
	tape := string(tapesyms[:p.pos]) +
		string(0x035f) +
		string(tapesyms[p.pos:])

	// describe the instruction
	b := p.prog[p.pc]
	i := b & 0xf0
	v := b & 0x0f
	var sym string
	switch i {
	case None:   
	case Move:   sym = " " + leftright[v&1] 
	case Back:   sym = "⟲ " + string('0' + v)
	case Climb:  sym = " " + updown[(v>>2)&1]
	case Branch: sym = "?" + blank01[v&3] + updown[(v>>2)&1]
	case Write:  sym = " " +cross01[v&3]
	}
	
	return fmt.Sprintf(
		"%2d %2d [%s] %2d %s",
		p.pc, p.pos,
		tape, b, sym,
	)
}

func (p *tmachine) PrintRun(n int) {
	fmt.Println(p)
	for i := 0; i < n; i++ {
		p.Step()
		fmt.Println(p)
	}
}

func main() {
	prog := []byte{
		Write|One,  None,     Branch|One|Down, None,       None,              None,       None,       Write|Blank,     Move|Right, Branch|Blank|Down, Back|2,            None,      None,
		Move|Right, Climb|Up, Write|Blank,     Move|Right, Branch|Blank|Down, Back|2,     None,       Write|One,       Back|8,     Write|One,         Branch|Blank|Down, Move|Left, Back|2,
		Climb|Up,   None,     Back|2,          None,       Write|One,         Move|Right, Write|Zero, Branch|Blank|Up, Move|Left,  Back|2,            Write|Zero,        Back|9,    None,
	}

	var tm tmachine

	copy(tm.prog[:], prog)

	tm.pc = 0
	tm.pos = 2
	tm.tape[1] = 2

	tm.PrintRun(512)
}