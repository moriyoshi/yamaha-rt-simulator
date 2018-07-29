package simulator

import "context"

var CmdIpInterfaceIgmp = &CommandSpec{
	[]Token{TIp, TL2Interface, TIgmp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIgmp = &CommandSpec{
	[]Token{TIp, TPp, TIgmp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIgmp = &CommandSpec{
	[]Token{TIp, TTunnel, TIgmp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIgmp = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIgmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIgmp = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIgmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIgmp = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIgmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

