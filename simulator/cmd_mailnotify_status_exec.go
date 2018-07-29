package simulator

import "context"

func mailNotifyStatusExec(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
	if !sess.Enabled {
		return AdministratorUseOnly
	}
	return nil
}

var CmdMailNotifyStatusExec1 = &CommandSpec{
	[]Token{TMail, TNotify, TStatus, TExec}, 4,
	mailNotifyStatusExec,
}

var CmdMailNotifyStatusExec2 = &CommandSpec{
	[]Token{TMailNotify, TStatus, TExec}, 3,
	mailNotifyStatusExec,
}
