// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/kawana77b/go-clip/clipboard"
	"github.com/spf13/cobra"
)

// pasteCmd represents the paste command
var pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Output from clipboard.",
	Long:  `Output from clipboard.`,
	RunE:  runPaste,
}

func init() {
	rootCmd.AddCommand(pasteCmd)
}

func runPaste(cmd *cobra.Command, args []string) error {
	text, err := clipboard.ReadAll()
	if err != nil {
		return err
	}
	fmt.Print(text)
	return nil
}
