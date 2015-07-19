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

import "fmt"
import "strings"
import "regexp"

// :cameron.freenode.net NOTICE * :*** Looking up your hostname...
type Command struct {
    // cameron.freenode.net
    Prefix string

    // NOTICE
    Command string

    // [0]: *
    // [1]: *** Looking up your hostname...
    Parameters []string
}

func (this *Command) String() string {
    // TODO: this should use a loop instead of strings.Join for parameters, making sure that
    // there is nothing containing a space, and if there is, using a :.
    if (len(this.Prefix) > 0) {
        if (len(this.Parameters) > 0) {
            return fmt.Sprintf(":%s %s %s", this.Prefix, this.Command, strings.Join(this.Parameters, " "))
        } else {
            return fmt.Sprintf(":%s %s", this.Prefix, this.Command)
        }
    } else {
        if (len(this.Parameters) > 0) {
            return fmt.Sprintf("%s %s", this.Command, strings.Join(this.Parameters, " "))
        } else {
            return this.Command
        }
    }
}

var (
    spacesExpr = regexp.MustCompile(` +`)
)

func splitArg(line string) (arg string, rest string) {
    parts := spacesExpr.Split(line, 2)
    if len(parts) > 0 {
        arg = parts[0]
    }
    if len(parts) > 1 {
        rest = parts[1]
    }
    return
}

func ParseLine(line string) *Command {
    args := make([]string, 0)
    command := new(Command)

    if strings.HasPrefix(line, ":") {
        command.Prefix, line = splitArg(line)
        command.Prefix = command.Prefix[1:len(command.Prefix)]
    }
    arg, line := splitArg(line)
    command.Command = strings.ToUpper(arg)
    for len(line) > 0 {
        if strings.HasPrefix(line, ":") {
            args = append(args, line[len(":"):])
            break
        }
        arg, line = splitArg(line)
        args = append(args, arg)
    }
    command.Parameters = args
    return command
}

