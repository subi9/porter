package exec

func (e *Exec) Install(commandFile string) error {
	err := e.LoadInstruction(commandFile)
	if err != nil {
		return err
	}
	return e.Execute()
}
