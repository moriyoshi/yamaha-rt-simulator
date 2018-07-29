package simulator

import "context"

var CmdSwitchControlFunctionExecute = &CommandSpec{
	[]Token{
		TSwitch, TControl, TFunction, TExecute,
		&EitherToken{
			TClearCounter,
			TClearMacaddressTable,
			TResetLoopdetect,
			TRestart,
			TRestartPoeSupply,
			TStartPoeSupply,
			TStopPoeSupply,
		},
	}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
