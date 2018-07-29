package simulator

import "context"

var CmdIpInterfacePimSparse = &CommandSpec{
	[]Token{TIp, TL2Interface, TPim, TSparse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpPimSparse = &CommandSpec{
	[]Token{TIp, TPp, TPim, TSparse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelPimSparse = &CommandSpec{
	[]Token{TIp, TTunnel, TPim, TSparse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfacePimSparse = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TPim, TSparse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpPimSparse = &CommandSpec{
	[]Token{TNo, TIp, TPp, TPim, TSparse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelPimSparse = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TPim, TSparse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

