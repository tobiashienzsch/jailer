package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NeoBSD/jailer"
)

// versionCmd represents the version sub command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version",
	RunE:  RunVersionCommand,
}

// RunVersionCommand prints the current jailer version
func RunVersionCommand(cmd *cobra.Command, args []string) error {
	// json
	if viper.GetBool("json") {
		version := struct {
			Version     string `json:"version"`
			BuildCommit string `json:"commit"`
			BuildDate   string `json:"date"`
			BuildOS     string `json:"build_os"`
		}{
			Version:     jailer.Version,
			BuildCommit: jailer.BuildCommit,
			BuildDate:   jailer.BuildDate,
			BuildOS:     jailer.BuildOS,
		}
		js, err := json.Marshal(version)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(js))
		return nil
	}

	// normal
	if viper.Get("verbose") == true {

		w := tabwriter.NewWriter(os.Stdout, 0, 0, tabWriterPadding, ' ', 0)
		fmt.Fprintf(w, "Version\t%s\t\n", jailer.Version)
		fmt.Fprintf(w, "Date\t%s\t\n", jailer.BuildDate)
		fmt.Fprintf(w, "Commit\t%s\t\n", jailer.BuildCommit)
		fmt.Fprintf(w, "BuildOS\t%s\t\n", jailer.BuildOS)

		return w.Flush()
	}

	fmt.Printf("%s-%s\n", jailer.Version, jailer.BuildCommit)
	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
