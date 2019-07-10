package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/cmd/shipper/commands"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "shipper"
	app.Version = "1.0.0"
	app.Usage = "🚢 Deploys and runs scripts for you."
	app.Commands = []cli.Command{
		{
			Name:      "cql",
			Usage:     "🔨 Run a script against the database.",
			ArgsUsage: "./service.foo/config/cockroach.sql",
			Action:    cql,
		},
		{
			Name:   "connect",
			Usage:  "💽 connect directly to CockroachDB.",
			Action: connect,
		},
		{
			Name:   "status",
			Usage:  "🕵️‍♀️ Check the current versions of deploys in production",
			Action: commands.Status,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cql(ctx *cli.Context) error {
	filename := ctx.Args().First()
	if filename == "" {
		return fmt.Errorf("🙅‍♀️ no filename given")
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("🙅‍♀️ file '%s' doesn't exist", aurora.Green(filename))
		}

		return fmt.Errorf("⁉️ Something went wrong: %v", err)
	}

	fmt.Printf("⏳ Running %s against CockroachDB... ", aurora.Green(filename))

	// TODO: make this into an actual service 😂
	//  - Also log this to discord.
	cmd := cockroachExecute(false, "sql", "--execute", string(contents))
	if err := cmd.Run(); err != nil {
		fmt.Println("❌")
		fmt.Printf("  ⚠️ Failed to execute CockroachDB statement: %v\n\n", err)
		printOutput(cmd)

		os.Exit(1)
	}

	// check cockroach result

	fmt.Println("✅")

	return nil
}

func connect(ctx *cli.Context) error {
	cmd := cockroachExecute(true, "sql")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	fmt.Printf("⏳ Connecting to %v... \n", aurora.Green("running-kitten-cockroachdb-public"))

	if err := cmd.Run(); err != nil {
		fmt.Printf("  ⚠️ error from cockroach: %v", err)
		printOutput(cmd)

		os.Exit(1)
	}

	return nil
}

func printOutput(cmd *exec.Cmd) {
	// print output from process on failures
	if output, err := cmd.CombinedOutput(); err == nil {
		fmt.Println(string(output))
	}
}

func cockroachExecute(tty bool, commands ...string) *exec.Cmd {
	args := []string{
		"exec", "-i", "cockroachdb-client-secure",
		"--",
		"./cockroach", "--certs-dir=/cockroach-certs", "--host=running-kitten-cockroachdb-public", "--database=lolibrary",
	}

	// if tty, set that explicitly.
	if tty {
		args[1] = "-it"
	}

	args = append(args, commands...)

	return exec.Command("kubectl", args...)
}
