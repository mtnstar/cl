var rootCmd = &cobra.Command{
  Use:   "cl",
  Short: "command line client for cryptopus",
  Long: "command line client for cryptopus"
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
