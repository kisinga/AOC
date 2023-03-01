package types

import "github.com/spf13/cobra"

type Solution interface {
	Runner() *cobra.Command
}
