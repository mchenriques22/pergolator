package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"

	"github.com/antoninferrand/pergolator/codegen"
)

func main() {
	cmd := &cli.Command{
		Name:  "pergolator",
		Usage: "Generate percolator helpers",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dest-package",
				Value: os.Getenv("GOPACKAGE"),
				Usage: "Package of the generated files. It should be automatically set by the go:generate directive or the GOPACKAGE environment variable. If not, you must set it manually.",
			},
			&cli.StringFlag{
				Name:  "path",
				Value: "",
				Usage: "Destination of the generated files.",
			},
			&cli.StringFlag{
				Name:  "prefix",
				Value: "",
				Usage: "Prefix for the generated files. They will have the form <prefix>_queries_gen.go and <prefix>_percolator_gen.go.",
			},
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",
				Usage: "The log level.",
			},
			&cli.StringFlag{
				Name:  "descriptor",
				Value: "",
				Usage: "Path to the descriptor file.",
			},
			&cli.IntFlag{
				Name:  "max-depth",
				Value: 2,
				Usage: "Maximum depth of the recursion when generating the percolator. As a result, you can query only struct.field.field... up to max-depth. The root is considered as depth 0.",
			},
			&cli.BoolFlag{
				Name:  "rename-fields-to-snake-case",
				Value: false,
				Usage: "Rename the fields to snake case. This is useful if your query are using snake case. (Avoids to add tags, descriptor, or a modifier).",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     getSlogLevel(cmd.String("log-level")),
			})))

			goFile := os.Getenv("GOFILE")
			if goFile == "" {
				if cmd.String("path") == "" {
					return fmt.Errorf("-path is mandatory when GOFILE is not set")
				}

				if cmd.String("prefix") == "" {
					return fmt.Errorf("-prefix is mandatory when GOFILE is not set")
				}

				goFile = cmd.String("path") + "/" + cmd.String("prefix") + ".go"
			}

			var buffer bytes.Buffer
			err := codegen.Run(&buffer,
				cmd.String("dest-package"),
				cmd.String("descriptor"),
				cmd.Args().Slice(),
				cmd.Int("max-depth"),
				cmd.Bool("rename-fields-to-snake-case"),
			)
			if err != nil {
				return fmt.Errorf("failed to generate percolator helpers: %w", err)
			}

			ext := filepath.Ext(goFile)
			baseFilename := goFile[0 : len(goFile)-len(ext)]
			targetFilename := baseFilename + "_percolator_gen.go"

			// Open the file for writing
			file, err := os.OpenFile(targetFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return fmt.Errorf("failed to open file %s: %w", goFile, err)
			}
			defer func() {
				if err := file.Close(); err != nil {
					slog.Warn("failed to close file", slog.String("file", goFile), slog.String("error", err.Error()))
				}
			}()

			_, err = file.Write(buffer.Bytes())
			return err
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func getSlogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Warn("invalid log level, defaulting to info")
		return slog.LevelInfo
	}
}
