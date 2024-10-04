package validatecmd

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/acronis/go-cti/internal/app/command"

	"github.com/acronis/go-cti/pkg/ctipackage"
	"github.com/acronis/go-cti/pkg/pacman"

	"github.com/spf13/cobra"
)

func New(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "validate cti",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			baseDir, err := command.GetWorkingDir(cmd)
			if err != nil {
				return fmt.Errorf("get base directory: %w", err)
			}

			return command.WrapError(execute(ctx, baseDir))
		},
	}
}

func execute(ctx context.Context, baseDir string) error {
	slog.Info("Validating package", slog.String("path", baseDir))
	pkg := ctipackage.New(baseDir)
	if err := pkg.Read(); err != nil {
		return fmt.Errorf("read package: %w", err)
	}

	// TODO: Validation for usage of indirect dependencies
	if err := pacman.Validate(pkg); err != nil {
		return fmt.Errorf("validate package: %w", err)
	}
	slog.Info("No errors found")
	return nil
}
