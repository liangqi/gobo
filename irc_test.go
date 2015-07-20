/*
 * Copyright (C) 2015 Robin Burchell <robin+git@viroteck.net>
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  - Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  - Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 * THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import "testing"
import "reflect"

func TestToString(t *testing.T) {
    tests := []string{
        ":w00t TEST", // with prefix, with no parameter
        ":w00t TEST hello", // with prefix, with short parameter
        ":w00t TEST hello world", // with prefix, with short parameters
        ":w00t TEST :foo bar", // with prefix, with long parameter
        ":w00t TEST hello world :foo bar", // with prefix, with multiple and long parameter

        "TEST", // without prefix, with no parameter
        "TEST hello", // without prefix, with short parameter
        "TEST hello world", // without prefix, with short parameters
        "TEST :foo bar", // without prefix, with long parameter
        "TEST hello world :foo bar", // without prefix, with multiple and long parameter
    }

    for _, test := range tests {
        // without prefix, with multiple and long parameter
        c := ParseLine(test)
        if (c.String() != test) {
            t.Error("Expected " + test + " got ", c.String())
        }
    }
}

type ParserTest struct {
    Input string
    Prefix string
    Command string
    Parameters []string
}

func TestParse(t *testing.T) {
    tests := []ParserTest{
        // single long
        {
            ":w00t TEST :hello world",
            "w00t",
            "TEST",
            []string{ "hello world", },
        },

        // multiple short
        {
            ":w00t TEST hello world",
            "w00t",
            "TEST",
            []string{ "hello", "world", },
        },

        // multiple short and long
        {
            ":w00t TEST hello world :how are you today",
            "w00t",
            "TEST",
            []string{ "hello", "world", "how are you today", },
        },
    }

    for _, test := range tests {
        c := ParseLine(test.Input)
        if (c.Prefix != test.Prefix) {
            t.Errorf("Expected: %#v, got %#v", test.Prefix, c.Prefix)
        }

        if (c.Command != test.Command) {
            t.Errorf("Expected: %#v, got %#v", test.Command, c.Command)
        }

        if (!reflect.DeepEqual(c.Parameters, test.Parameters)) {
            t.Errorf("Expected: %#v, got %#v", test.Parameters, c.Parameters)
        }
    }
}

func BenchmarkParseSingleLong(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ParseLine(":w00t TEST :hello world")
    }
}

func BenchmarkParseMultipleShort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ParseLine(":w00t TEST hello world")
    }
}


func BenchmarkParseMultipleAndLong(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ParseLine(":w00t TEST hello world :how are you today")
    }
}


