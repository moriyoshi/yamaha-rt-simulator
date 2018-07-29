package simulator

import "context"

var CmdShowStatusQacTm = &CommandSpec{
	[]Token{TShow, TStatus, TQacTm}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusQacTmServer = &CommandSpec{
	[]Token{TShow, TStatus, TQacTm, TServer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusQacTmClient = &CommandSpec{
	[]Token{TShow, TStatus, TQacTm, TClient}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusQacTmQualified = &CommandSpec{
	[]Token{TShow, TStatus, TQacTm, TQualified}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusQacTmUnqualified = &CommandSpec{
	[]Token{TShow, TStatus, TQacTm, TUnqualified}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
