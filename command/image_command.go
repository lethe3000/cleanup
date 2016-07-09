package command

import (
    "github.com/urfave/cli"
    "fmt"
    "github.com/fsouza/go-dockerclient"
    "strings"
)

func NewImageListCommand() cli.Command {
    return cli.Command{
        Name:      "ls",
        Usage:     "list images",
        ArgsUsage: " ",
        Action: handleImageList,
    }
}

// handleImage handles a request that list image.
func handleImageList(c *cli.Context) error {
    client, _ := getDockerClient(c.GlobalString("docker-entry"))
    images, _ := client.ListImages(docker.ListImagesOptions{All: false})
    repos := make(map[string]int)
    for _, img := range images {
        for _, repoTags := range img.RepoTags {
            lastIndex := strings.LastIndex(repoTags, ":")
            baseRepo := repoTags[0:lastIndex]
            repos[baseRepo] += 1
        }
    }

    printImageRepos(repos)
    return nil
}

func printImageRepos(repos map[string]int) {
    fmt.Println("Images:")

    for repo, count := range repos {
        fmt.Println("  ", blue(repo))
        fmt.Printf("                                            [%s]\n", red(count))
    }
}
