package nofox

import (
	"fmt"
	"io"
	"maps"
	"testing"
)

func TestNewLexer(t *testing.T) {
	justAnEmitter := make(chan Token)
	alphaMappings := map[rune]Token{
		'a': TokenMoveRight,
		'b': TokenMoveLeft,
		'c': TokenIncrement,
		'd': TokenDecrement,
		'e': TokenPrint,
		'f': TokenRead,
		'g': TokenLoopStart,
		'h': TokenLoopEnd,
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mappings map[rune]Token
		emitter  chan Token
		want     *Lexer
		wantErr  bool
	}{
		{
			name:    "just give emitter",
			emitter: justAnEmitter,
			want: &Lexer{
				Mappings: DefaultMapping,
				Emitter:  justAnEmitter,
			},
		},
		{
			name:    "don't even give emitter",
			wantErr: true,
		},
		{
			name:     "Give both things",
			mappings: alphaMappings,
			emitter:  justAnEmitter,
			want: &Lexer{
				Mappings: alphaMappings,
				Emitter:  justAnEmitter,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := NewLexer(tt.mappings, tt.emitter)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewLexer(%v, %v) failed: %v", tt.mappings, tt.emitter, gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("NewLexer(%v, %v) succeeded unexpectedly", tt.mappings, tt.emitter)
			}
			if got.Emitter != tt.want.Emitter || !maps.Equal(got.Mappings, tt.want.Mappings) {
				t.Errorf("NewLexer(%v, %v) = %v, want %v", tt.mappings, tt.emitter, got, tt.want)
			}
		})
	}
}

func Test_validateMappings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mappings map[rune]Token
		wantErr  bool
	}{
		{
			name:     "default mappings",
			mappings: DefaultMapping,
			wantErr:  false,
		},
		{
			name:    "nil mappings",
			wantErr: true,
		},
		{
			name: "alphabet mappings",
			mappings: map[rune]Token{
				'a': TokenMoveRight,
				'b': TokenMoveLeft,
				'c': TokenIncrement,
				'd': TokenDecrement,
				'e': TokenPrint,
				'f': TokenRead,
				'g': TokenLoopStart,
				'h': TokenLoopEnd,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := validateMappings(tt.mappings)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("validateMappings() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("validateMappings() succeeded unexpectedly")
			}
		})
	}
}

func TestLexer_Lex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		mappings map[rune]Token
		emitter  chan Token
		// Named input parameters for target function.
		r       io.Reader
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := NewLexer(tt.mappings, tt.emitter)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := l.Lex(tt.r)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Lex() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Lex() succeeded unexpectedly")
			}
		})
	}
}

func TestLexer_String(t *testing.T) {
	justAnEmitter := make(chan Token)
	tests := []struct {
		name     string // description of this test case
		mappings map[rune]Token
		emitter  chan Token
		want     string
	}{
		{
			name:    "default",
			emitter: justAnEmitter,
			want:    fmt.Sprintf("Lexer{Mappings: map[43:TokenIncrement 44:TokenRead 45:TokenDecrement 46:TokenPrint 60:TokenMoveLeft 62:TokenMoveRight 91:TokenLoopStart 93:TokenLoopEnd], Emitter: %v}", justAnEmitter),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := NewLexer(tt.mappings, tt.emitter)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := l.String()
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
