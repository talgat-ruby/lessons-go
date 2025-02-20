package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/db/internal/config"
)

type subcmdJet struct {
	cmd *flag.FlagSet
}

func newSubcmdJet() subcmd {
	cmd := flag.NewFlagSet(cmdNameJet, flag.ExitOnError)

	return &subcmdJet{
		cmd,
	}
}

func (s *subcmdJet) getCmdName() string {
	return s.cmd.Name()
}

func (s *subcmdJet) printDefaults() {
	s.cmd.PrintDefaults()
}

func (s *subcmdJet) parse(args []string) error {
	if err := s.cmd.Parse(args); err != nil {
		return err
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if !s.cmd.Parsed() {
		return fmt.Errorf("please provide correct arguments to %s command", s.cmd.Name())
	}

	return s.execJetCommand()
}

func (s *subcmdJet) execJetCommand() error {
	conf := config.New(context.Background())

	// Check if jet is installed
	if _, err := exec.LookPath("jet"); err != nil {
		return fmt.Errorf("jet command not found in PATH: %w", err)
	}

	// Create the command with arguments
	args := []string{
		fmt.Sprintf(
			"-dsn=%s",
			fmt.Sprintf(
				"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
				url.QueryEscape(conf.Postgres.User),
				url.QueryEscape(conf.Postgres.Password),
				url.QueryEscape(conf.Postgres.Host),
				conf.Postgres.Port,
				url.QueryEscape(conf.Postgres.Name),
			),
		),
		fmt.Sprintf("-schema=%s", "public"),
		fmt.Sprintf("-path=%s", "./internal/postgres/generated"),
	}

	cmd := exec.Command("jet", args...)

	// Create a pipe for capturing output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	// Read output asynchronously
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				log.Printf("Output: %s", strings.TrimSpace(string(buf[:n])))
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stderr.Read(buf)
			if n > 0 {
				log.Printf("Error: %s", strings.TrimSpace(string(buf[:n])))
			}
			if err != nil {
				break
			}
		}
	}()

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}
