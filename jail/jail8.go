package jail

import (
	"context"
	"os"
	"os/exec"
)

func CreateJail(ctx context.Context, confPath string) error {
	cmd := exec.CommandContext(ctx, "jail", "-cf", confPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}