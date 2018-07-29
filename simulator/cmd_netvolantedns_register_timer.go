package simulator

import "context"

var CmdNetvolanteDnsRegisterTimer = &CommandSpec{
	[]Token{TNetvolanteDns, TRegister, TTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsRegisterTimer = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TRegister, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

