package run

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/kubetrail/bip39/pkg/passphrases"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) error {
	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	if !prompt {
		return fmt.Errorf("this command is not designed to received piped input")
	}

	var mnemonic string
	if len(args) == 0 {
		if prompt {
			if err := mnemonics.Prompt(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}
		}

		mnemonic, err = mnemonics.Read(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read mnemonic from input: %w", err)
		}
	} else {
		mnemonic = mnemonics.NewFromFields(args)
	}

	fields := strings.Fields(mnemonic)

	if len(fields) == 0 {
		return fmt.Errorf("invalid mnemonic with zero fields")
	}

	_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", "Pl. enter mnemonic word index on prompt starting at 1")

	for {
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "> ")

		index, err := passphrases.Prompt(io.Discard)
		if err != nil {
			return fmt.Errorf("failed to read index value: %w", err)
		}

		v, err := strconv.ParseInt(index, 10, 64)
		if err != nil {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "failed to parse input as an integer: %v\n", err)
			continue
		}

		if int(v) < 1 || int(v) > len(fields) {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", "index out of bounds")
			continue
		}

		s := fields[int(v)-1]
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\r", s)
		time.Sleep(time.Second)
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\r", strings.Repeat(" ", len(s)+2))
	}
}
