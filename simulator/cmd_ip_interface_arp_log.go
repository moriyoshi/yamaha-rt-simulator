package simulator

import "context"

var CmdIpInterfaceArpLog = &CommandSpec{
	[]Token{TIp, TL2Interface, TArp, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceArpLog = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TArp, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

