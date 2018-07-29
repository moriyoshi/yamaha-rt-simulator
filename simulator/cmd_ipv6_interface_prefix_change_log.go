package simulator

import "context"

var CmdIpv6InterfacePrefixChangeLog = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TPrefix, TChange, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpPrefixChangeLog = &CommandSpec{
	[]Token{TIpv6, TPp, TPrefix, TChange, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelPrefixChangeLog = &CommandSpec{
	[]Token{TIpv6, TTunnel, TPrefix, TChange, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfacePrefixChangeLog = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TPrefix, TChange, TLog}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpPrefixChangeLog = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TPrefix, TChange, TLog}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelPrefixChangeLog = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TPrefix, TChange, TLog}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

