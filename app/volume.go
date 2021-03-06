package app

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

type VolumeInfo struct {
	Name         string `json:"name"`
	ReplicaCount int    `json:"replicaCount"`
	Endpoint     string `json:"endpoint"`
}

func InfoCmd() cli.Command {
	return cli.Command{
		Name: "info",
		Action: func(c *cli.Context) {
			if err := info(c); err != nil {
				logrus.Fatalln("Error running journal command:", err)
			}
		},
	}
}

func info(c *cli.Context) error {
	cli := getCli(c)

	volume, err := cli.GetVolume()
	if err != nil {
		return err
	}

	info := VolumeInfo{
		Name:         volume.Name,
		ReplicaCount: volume.ReplicaCount,
		Endpoint:     volume.Endpoint,
	}

	output, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
