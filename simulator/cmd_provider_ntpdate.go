package simulator

import "context"

var CmdProviderNtpdate = &CommandSpec{
	[]Token{TProvider, TNtpdate}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderNtpdate = &CommandSpec{
	[]Token{TNo, TProvider, TNtpdate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

