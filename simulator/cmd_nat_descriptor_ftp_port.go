package simulator

import "context"

var CmdNatDescriptorFtpPort = &CommandSpec{
	[]Token{TNat, TDescriptor, TFtp, TPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorFtpPort = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TFtp, TPort}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

