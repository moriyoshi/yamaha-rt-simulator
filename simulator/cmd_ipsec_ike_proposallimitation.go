package simulator

import "context"

var CmdIpsecIkeProposalLimitation = &CommandSpec{
	[]Token{TIpsec, TIke, TProposalLimitation}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeProposalLimitation = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TProposalLimitation}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

