package command
import (
    "github.com/fsouza/go-dockerclient"
    "github.com/fatih/color"
    "strings"
)

type MyImage struct {
    ID          string
    Repo        string
    Name        string
    Created     int64
}

var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func getDockerClient(entryPoint string) (*docker.Client, error) {
    entryList := []string{"unix:///var/run/docker.sock", entryPoint}
    for _, e := range entryList {
        client, err := docker.NewClient(e)
        if err == nil {
            return client, err
        }
        continue
    }
    return nil, nil
}

func restructAPIImage(images []docker.APIImages) ([]MyImage) {
    var targetRepo []MyImage
    for _, image := range images {
        for _, repoTag := range image.RepoTags {
            lastIndex := strings.LastIndex(repoTag, ":")
            baseRepo := repoTag[0:lastIndex]
            targetRepo = append(targetRepo,
                                MyImage{ID: image.ID, Name: repoTag, Repo: baseRepo, Created: image.Created,})
        }
    }
    return targetRepo
}

//func getMyImage(images *docker.APIImages) (map [string]MyImage) {
//    repos := make(map[string]MyImage)
//    for _, image := range images {
//        myImage := MyImage{}
//        for _, repoTags := range image.RepoTags {
//            append(MyImage{ID: })
//            lastIndex := strings.LastIndex(repoTags, ":")
//            baseRepo := repoTags[0:lastIndex]
//        }
//
//    }
//    return repos
//}