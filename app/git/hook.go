package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Hook(path, name, prefix string) {
	name = fmt.Sprintf("%s-%s", prefix, name)

	// Path to thehook script
	hook := filepath.Join(path, ".git", "hooks", name)

	// Check if the hook script exists and is executable
	_, err := os.Stat(hook)
	if os.IsNotExist(err) {
		//log.Fatal("hook not found")
	} else {
		// Execute the hook script
		cmd := exec.Command(hook)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("pre-commit hook executed successfully.")
	}
}
