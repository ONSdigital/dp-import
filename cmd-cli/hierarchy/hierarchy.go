package hierarchy

import (
	"errors"
	"fmt"
	"github.com/ONSdigital/dp-import/cmd-cli/config"
	"github.com/ONSdigital/go-ns/log"
	"os"
	"os/exec"
)

var (
	Config                      config.Model
	hierarchyBuilderDir         = "%s/src/github.com/ONSdigital/dp-hierarchy-builder/"
	makeInstanceHierarchyCMDFMT = "INSTANCE_ID=\"%s\""
	workingDir                  = ""
)

func init() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Error(errors.New("could not find GOPATH env var"), nil)
		os.Exit(1)
	}
	workingDir = fmt.Sprintf(hierarchyBuilderDir, goPath)
}

func CreateFullHeirarchy() error {
	log.Info("executing hierarchy builder make full", nil)
	makeFull := exec.Command("make", "full")
	makeFull.Stdout = os.Stdout
	makeFull.Dir = workingDir

	err := makeFull.Run()
	if err != nil {
		log.ErrorC("failed to run makeFull", err, nil)
		return err
	}
	log.Info("hierarchy builder make full completed successfully", nil)
	return nil
}

func CreateInstanceHierarchy(instanceID string) error {
	log.Info("executing hierarchy builder make instance hierarchy", log.Data{"instanceID": instanceID})
	makeInstance := exec.Command("make", fmt.Sprintf(makeInstanceHierarchyCMDFMT, instanceID), "instance")
	makeInstance.Stdout = os.Stdout
	makeInstance.Dir = workingDir

	err := makeInstance.Run()
	if err != nil {
		log.ErrorC("failed to run makeInstance", err, nil)
		os.Exit(1)
	}

	log.Info("hierarchy builder make instance hierarchy complete successfully", log.Data{"instanceID": instanceID})
	return nil
}
