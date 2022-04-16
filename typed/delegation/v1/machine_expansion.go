package v1

type MachineExpansion interface {
	CheckPortConflict()
	ExportPort()
	ExecCommand()
}

func (c *machine) CheckPortConflict() {
	//TODO implement me
	panic("implement me")
}

func (c *machine) ExportPort() {
	//TODO implement me
	panic("implement me")
}

func (c *machine) ExecCommand() {
	//TODO implement me
	panic("implement me")
}
