// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"io"
	"os"
	"time"

	"github.com/kawana77b/go-clip/clipboard"
	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy from standard input to clipboard",
	Long:  `Copy from standard input to clipboard`,
	RunE:  runCopy,
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().DurationP("timeout", "t", 0, "Erase clipboard after timeout.  Durations are specified like \"20s\" or \"2h45m\".  0 (default) means never erase.")
}

func runCopy(cmd *cobra.Command, args []string) error {
	timeout, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		return err
	}

	out, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	if err := clipboard.WriteAll(string(out)); err != nil {
		return err
	}

	if timeout > 0 {
		<-time.After(timeout)
		text, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		if text == string(out) {
			err := clipboard.WriteAll("")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
