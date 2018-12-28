package config

import (
    "comic-hero/model"
    "strings"
)

var ComicDefs = map[string]model.ComicDef{
    "sinfest": {
        Id:          10,
        Name:        "Sinfest",
        Description: "Sinfest webcomic jumps from topic to topic and covers issues related to American politics, organized religion, and feminism.",
        Author:      "Tatsuya Ishida",
        Url:         "http://sinfest.net/",
    },
    "dilbert": {
        Id:          20,
        Name:        "Dilbert",
        Description: "The Dilbert strip is known for its satirical office humor about a white-collar, micromanaged office featuring engineer Dilbert as the title character.",
        Author:      "Scott Adams",
        Url:         "https://dilbert.com/",
    },
    "oglaf": {
        Id:          30,
        Name:        "Oglaf",
        Description: "Oglaf is a sexually explicit webcomic taking place in a fantasy realm with many recurring themes.",
        Author:      "Trudy Cooper, Doug Bayne",
        Url:         "https://www.oglaf.com/",
    },
}

func GetIdForComicName(comicName string) (int, bool) {
    comicName = strings.ToLower(comicName)
    for currentName, currentDef := range ComicDefs {
        if currentName == comicName {
            return currentDef.Id, true
        }
    }

    return 0, false
}

func GetComicDefForId(id int) (*model.ComicDef, bool) {
    for _, currentDef := range ComicDefs {
        if currentDef.Id == id {
            return &currentDef, true
        }
    }

    return nil, false
}
