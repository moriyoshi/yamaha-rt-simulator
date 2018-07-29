package simulator

import "context"

var CmdMailServerSmtp = &CommandSpec{
	[]Token{TMail, TServer, TSmtp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailServerSmtp = &CommandSpec{
	[]Token{TNo, TMail, TServer, TSmtp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

