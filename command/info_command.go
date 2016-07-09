package command

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/fsouza/go-dockerclient"
    "github.com/fatih/color"
    "time"
)

func NewImageInfoCommand() cli.Command {
    return cli.Command{
        Name:      "info",
        Usage:     "list all tags for given image repo",
        ArgsUsage: " ",
        Flags: []cli.Flag{
            cli.StringFlag{Name: "name", Usage: "name of image repo"},
        },
        Action: handleImageInfo,
    }
}

func handleImageInfo(c *cli.Context) error {
    fmt.Sprintln("Image name: %s", c.String("name"))
    if c.String("name") == "" {
        return nil
    }
    client, _ := getDockerClient(c.GlobalString("docker-entry"))
    images, _ := client.ListImages(docker.ListImagesOptions{All: false})

    myImages := Filter(restructAPIImage(images), c)

    for _, image := range myImages {
        fmt.Printf("    %s  %s  %s\n", blue(image.Name), image.ID, red(time.Unix(image.Created, 0)))
    }
    return nil
}
