package simulator

import "context"

var CmdIpInterfaceOspfNeighbor = &CommandSpec{
	[]Token{TIp, TL2Interface, TOspf, TNeighbor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpOspfNeighbor = &CommandSpec{
	[]Token{TIp, TPp, TOspf, TNeighbor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelOspfNeighbor = &CommandSpec{
	[]Token{TIp, TTunnel, TOspf, TNeighbor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceOspfNeighbor = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TOspf, TNeighbor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpOspfNeighbor = &CommandSpec{
	[]Token{TNo, TIp, TPp, TOspf, TNeighbor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelOspfNeighbor = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TOspf, TNeighbor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

