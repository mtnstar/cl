var rootCmd = &cobra.Command{
  Use:   "login",
  Short: "login",
  Long: "login"
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
