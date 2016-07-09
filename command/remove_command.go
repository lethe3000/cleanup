package command

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/fsouza/go-dockerclient"
    "sort"
    "time"
)

func NewRemoveCommand() cli.Command {
    return cli.Command{
        Name:      "rm",
        Usage:     "remove images",
        ArgsUsage: " ",
        Flags: []cli.Flag{
            cli.StringFlag{Name: "name", Usage: "image name to remove"},
            cli.IntFlag{Name: "count", Usage: "remove images except latest count, sort by created date"},
            cli.Int64Flag{Name: "days", Usage: "remove images days before"},
        },
        Action: handleRemove,
    }
}

type ByCreated []MyImage
func (a ByCreated) Len() int           { return len(a) }
func (a ByCreated) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreated) Less(i, j int) bool { return a[i].Created > a[j].Created }

func handleRemove(c *cli.Context) error {
    fmt.Sprintln("Image name: %s", c.String("name"))
    if c.String("name") == "" {
        return nil
    }
    if c.Int("count") == 0 && c.Int64("days") == 0 {
        return nil
    }

    client, _ := getDockerClient(c.GlobalString("docker-entry"))
    images, _ := client.ListImages(docker.ListImagesOptions{All: false})

    myImages := Filter(restructAPIImage(images), c)

    fmt.Println("Images removed:")
    for _, image := range myImages {
        fmt.Printf("  %s  %s", blue(image.Name), time.Unix(image.Created, 0))
        err := client.RemoveImage(image.Name)
        if err != nil {
            fmt.Println("remove error: ", red(err))
        }
    }

    return nil
}

func Filter(mi []MyImage, c *cli.Context) []MyImage {
    filterd := make([]MyImage, 0)
    for _, i := range mi {
        if i.Repo == c.String("name") {
            filterd = append(filterd, i)
        }
    }
    // sort by image created
    sort.Sort(ByCreated(filterd))
    if c.Int("count") > 0 {
        fmt.Println("delete count", filterd[c.Int("count"):])
        return filterd[c.Int("count"):]
    }
    if c.Int64("days") > 0 {
        deadline := time.Now().Unix() - c.Int64("days") * 24 * 60 * 60
        // Always reserve latest image
        for idx, i := range filterd[1:] {
            if i.Created < deadline {
                return filterd[idx:]
            }
        }
        return filterd[1:]
    }
    return filterd
}