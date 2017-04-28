package v2

import (
	"os"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	oldCmd "code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/v2/shared"
)

type SecurityGroupsCommand struct {
	usage           interface{} `usage:"CF_NAME security-groups"`
	relatedCommands interface{} `related_commands:"bind-security-group, bind-running-security-group, bind-staging-security-group, security-group"`

	SharedActor command.SharedActor
	Config      command.Config
	UI          command.UI
}

func (cmd *SecurityGroupsCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor()

	return nil
}

func (cmd SecurityGroupsCommand) Execute(args []string) error {
	if cmd.Config.Experimental() == false {
		oldCmd.Main(os.Getenv("CF_TRACE"), os.Args)
		return nil
	}
	cmd.UI.DisplayText(command.ExperimentalWarning)
	cmd.UI.DisplayNewline()

	_, err := cmd.Config.CurrentUser()
	if err != nil {
		return shared.HandleError(err)
	}

	err = cmd.SharedActor.CheckTarget(cmd.Config, false, false)
	if err != nil {
		return shared.HandleError(err)
	}

	return nil
}
