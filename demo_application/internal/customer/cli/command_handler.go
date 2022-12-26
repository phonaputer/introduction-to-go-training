package cli

type CommandHandler interface {
	Handle(args []string)
}
