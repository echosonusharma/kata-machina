package colors

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#

var (
	escape_char                   = "\033"
	control_sequence_introducer   = "["
	parameter_separator           = ";"
	text_and_color_style_ansi_cmd = "m"
)

var (
	reset_code               uint8 = 0
	ansi_bold_code           uint8 = 1
	ansi_underline_code      uint8 = 4
	ansi_blink_code          uint8 = 5
	ansi_reverse_code        uint8 = 7
	ansi_strike_through_code uint8 = 9
)

var reset_cmd = fmt.Sprintf("%s%s%d%s", escape_char, control_sequence_introducer, reset_code, text_and_color_style_ansi_cmd)

type ansi_fg_color_code uint8
type ansi_bg_color_code uint8

type colorRegistry[T comparable] struct {
	Black         T
	Red           T
	Green         T
	Yellow        T
	Blue          T
	Magenta       T
	Cyan          T
	White         T
	BrightBlack   T
	BrightRed     T
	BrightGreen   T
	BrightYellow  T
	BrightBlue    T
	BrightMagenta T
	BrightCyan    T
	BrightWhite   T
}

var BgColorRegistry = &colorRegistry[ansi_bg_color_code]{
	Black:         40,
	Red:           41,
	Green:         42,
	Yellow:        43,
	Blue:          44,
	Magenta:       45,
	Cyan:          46,
	White:         47,
	BrightBlack:   100,
	BrightRed:     101,
	BrightGreen:   102,
	BrightYellow:  103,
	BrightBlue:    104,
	BrightMagenta: 105,
	BrightCyan:    106,
	BrightWhite:   107,
}

var FgColorRegistry = &colorRegistry[ansi_fg_color_code]{
	Black:         30,
	Red:           31,
	Green:         32,
	Yellow:        33,
	Blue:          34,
	Magenta:       35,
	Cyan:          36,
	White:         37,
	BrightBlack:   90,
	BrightRed:     91,
	BrightGreen:   92,
	BrightYellow:  93,
	BrightBlue:    94,
	BrightMagenta: 95,
	BrightCyan:    96,
	BrightWhite:   97,
}

type ColorProfileOpts struct {
	bg             ansi_bg_color_code
	fg             ansi_fg_color_code
	bold           bool
	blink          bool
	underline      bool
	reverse        bool
	reset          bool
	strike_through bool
}

type ColorProfile struct {
	ColorProfileOpts
}

type ColorProfileOptFn func(*ColorProfileOpts)

func defaultColorProfileOpts() ColorProfileOpts {
	return ColorProfileOpts{
		bold:           false,
		blink:          false,
		underline:      false,
		reverse:        false,
		reset:          true,
		strike_through: false,
	}
}

func WithBgColor(bg ansi_bg_color_code) ColorProfileOptFn {
	return func(cpo *ColorProfileOpts) {
		cpo.bg = bg
	}
}

func WithFgColor(fg ansi_fg_color_code) ColorProfileOptFn {
	return func(cpo *ColorProfileOpts) {
		cpo.fg = fg
	}
}

func WithBoldText(opt *ColorProfileOpts) {
	opt.bold = true
}

func WithBlinkingText(opt *ColorProfileOpts) {
	opt.blink = true
}

func WithUnderliningText(opt *ColorProfileOpts) {
	opt.underline = true
}

func WithReverse(opt *ColorProfileOpts) {
	opt.reverse = true
}

func WithoutReset(opt *ColorProfileOpts) {
	opt.reset = false
}

func WithStrikeThrough(opt *ColorProfileOpts) {
	opt.strike_through = true
}

func New(opts ...ColorProfileOptFn) *ColorProfile {
	o := defaultColorProfileOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &ColorProfile{
		ColorProfileOpts: o,
	}
}

func (cp *ColorProfile) Build(msg string) string {
	codes_to_apply := []uint8{}

	if cp.bg != 0 {
		codes_to_apply = append(codes_to_apply, uint8(cp.bg))
	}

	if cp.fg != 0 {
		codes_to_apply = append(codes_to_apply, uint8(cp.fg))
	}

	if cp.bold {
		codes_to_apply = append(codes_to_apply, ansi_bold_code)
	}

	if cp.underline {
		codes_to_apply = append(codes_to_apply, ansi_underline_code)
	}

	if cp.blink {
		codes_to_apply = append(codes_to_apply, ansi_blink_code)
	}

	if cp.reverse {
		codes_to_apply = append(codes_to_apply, ansi_reverse_code)
	}

	if cp.strike_through {
		codes_to_apply = append(codes_to_apply, ansi_strike_through_code)
	}

	var fmt_text string = fmt.Sprintf("%s%s", escape_char, control_sequence_introducer)

	for i, code := range codes_to_apply {
		fmt_text += fmt.Sprintf("%d", code)

		if len(codes_to_apply)-1 != i {
			fmt_text += parameter_separator
		}
	}

	fmt_text += text_and_color_style_ansi_cmd
	fmt_text += msg

	if cp.reset {
		fmt_text += reset_cmd
	}

	return fmt_text
}

func (cp *ColorProfile) Print(msg string) {
	str := cp.Build(msg)
	fmt.Print(str)
}
