package simulator

import "context"

var CmdConsolePrompt = &CommandSpec{
	[]Token{TConsole, TPrompt}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		sess.RunningConfig.SetPrompt(tis[2].Image)
		sess.ResetPrompt()
		return nil
	},
}
