package simulator

import "context"

var CmdIpInterfaceIgmpStatic = &CommandSpec{
	[]Token{TIp, TL2Interface, TIgmp, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIgmpStatic = &CommandSpec{
	[]Token{TIp, TPp, TIgmp, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIgmpStatic = &CommandSpec{
	[]Token{TIp, TTunnel, TIgmp, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIgmpStatic = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIgmp, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIgmpStatic = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIgmp, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIgmpStatic = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIgmp, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

