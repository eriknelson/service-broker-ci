package action

import (
	"fmt"
	"strings"
)

func Verify(repo string, args string) error {
	script, err := getTemplate(repo)
	if err != nil {
		return err
	}

	fmt.Printf("Running: %s %s\n", repo, args)
	combinedArgs := fmt.Sprintf("%s %s", script, args)
	output, err := RunCommand("bash", combinedArgs)
	if err != nil {
		if strings.Contains(string(output), "executable file not found in $PATH") ||
			strings.Contains(string(output), "No such file or directory") {
			output, err = RunCommand(repo, args)
			fmt.Println(string(output))
			fmt.Println(string(output))
			fmt.Println(string(output))
			if err != nil {
				return err
			}
		} else {
			fmt.Println(string(output))
			return err
		}
	} else {
		fmt.Println(string(output))
	}
	return nil
}
