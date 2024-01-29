package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// by default, no flag is needed when writing out thoughts
// a subcommand is run when passing the -l flag
var rootCmd = &cobra.Command{
	Use: "",

	Long: `Write your thoughts after this. A new dir and file will be created in your $HOME.

Additionally, by passing in the -l flag, you can list your thoughts for the day.
By passing a date (YYYY-MM-DD) after, you can list your thoughts from a previous day.`,

	Example: `$ thoughts look up that cat video
$ thoughts l # this lists today's thoughts
$ thoughts l 2024-01-21 # this list a specific date's thoughts
$ thoughts l all # this lists all thoughts of all dates
`,

	Run: RunThoughtCommand,
}

func RunThoughtCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		return
	}

	homeDir := os.Getenv("HOME")
	thoughtsDir := path.Join(homeDir, ".thoughts_storage")
	todaysDate := time.Now().Format("2006-01-02")
	todayFile := path.Join(thoughtsDir, todaysDate+".txt")

	// create .thoughts_storage directory
	if _, err := os.Stat(thoughtsDir); os.IsNotExist(err) {
		os.Mkdir(path.Join(homeDir, thoughtsDir), 0755)
		return
	}

	if args[0] == "l" {
		err := ListThoughts(thoughtsDir, todaysDate, args)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	err := AddThought(todayFile, args)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ListThoughts(thoughtsDir, todaysDate string, args []string) error {
	// if the command is `thoughts l all` then read all files
	if len(args) > 1 && args[1] == "all" {
		err := ReadAllFiles(thoughtsDir)
		if err != nil {
			return err
		}
		return nil
	}

	dateFile := todaysDate
	if len(args) > 1 {
		dateFile = args[1]
	}

	if validateDateFormat(dateFile, time.DateOnly) == false {
		errString := "date in incorrect format, use YYYY-MM-DD"
		fmt.Println(errString)
		return errors.New(errString)
	}

	datePath := path.Join(thoughtsDir, dateFile+".txt")
	if _, err := os.Stat(datePath); os.IsNotExist(err) {
		fmt.Println("no thoughts for that day")
		return err
	}
	byteContents, err := os.ReadFile(datePath)
	if err != nil {
		fmt.Println("unable to open file: ", datePath)
		return err
	}
	fmt.Println(dateFile)
	fmt.Println(string(byteContents))

	return nil
}

func ReadAllFiles(thoughtsDir string) error {
	files, err := os.ReadDir(thoughtsDir)
	if err != nil {
		return err
	}

	for _, filePath := range files {
		if filePath.IsDir() {
			continue
		}
		fileDate := strings.Split(filePath.Name(), ".")[0]
		fullPath := path.Join(thoughtsDir, filePath.Name())
		byteContents, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println("unable to open file: ", fullPath)
			return err
		}
		fmt.Println(fileDate)
		fmt.Println(string(byteContents))
	}
	return nil
}

func validateDateFormat(dateString, layout string) bool {
	_, err := time.Parse(layout, dateString)
	return err == nil
}

func AddThought(todayFile string, args []string) error {
	// create file for today
	if _, err := os.Stat(todayFile); os.IsNotExist(err) {
		os.Create(todayFile)
	}

	f, err := os.OpenFile(todayFile, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("Could not open file: ", todayFile)
		return err
	}

	_, err = f.WriteString(strings.Join(args, " "))
	if err != nil {
		fmt.Println("Could not write to file: ", todayFile)
		return err
	}
	f.WriteString("\n")
	fmt.Println("Saved.")

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
