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
    "bc": {
        Id:          40,
        Name:        "B.C.",
        Description: "Johnny Hart's classic strip, B.C., puts a caveman twist on everything. From philosophical ants to punny bits of unconventional wisdom, you'll see why this strip has been a favorite for so many years.",
        Author:      "Johnny Hart, Mason Mastroianni",
        Url:         "https://www.gocomics.com/bc",
    },
    "calvinandhobbes": {
        Id:          50,
        Name:        "Calvin & Hobbes",
        Description: "Since its introduction 1985, Bill Watterson's comic follows the richly imaginative adventures of 6-year-old Calvin and his trusty tiger, Hobbes.",
        Author:      "Bill Watterson",
        Url:         "https://www.gocomics.com/calvinandhobbes",
    },
    "fowl-language": {
        Id:          60,
        Name:        "Fowl Language",
        Description: "Through Fowl Language, cartoonist Brian Gordon draws on his trials and tribulations of raising two small children. By poking fun at the daily tedium and frustrations of parenting, he hopes to give comfort to parents who are losing their minds just as quickly as he is.",
        Author:      "Brian Gordon",
        Url:         "https://www.gocomics.com/fowl-language",
    },
    "realitycheck": {
        Id:          70,
        Name:        "Reality Check",
        Description: "Cartoonist Dave Whamond offers an offbeat view of the world in Reality Check, his daily and Sunday comic panel that exposes the hidden hilarity in everyday situations. A thoroughly wacky look at life, Whamond says he just frames some of the silliness of everyday life in the comic and invites people to take a double-take - to look at life from another angle",
        Author:      "Dave Whamond",
        Url:         "https://www.gocomics.com/realitycheck",
    },
    "pearlsbeforeswine": {
        Id:          80,
        Name:        "Pearls Before Swine",
        Description: "At its heart, Pearls Before Swine is the comic strip tale of two friends: an arrogant Rat who thinks he knows it all and a slow-witted Pig who doesn't know any better. Together, this pair offers caustic commentary on humanity's quest for the unattainable.",
        Author:      "Stephan Pastis",
        Url:         "https://www.gocomics.com/pearlsbeforeswine",
    },
    "moderately-confused": {
        Id:          90,
        Name:        "Moderately Confused",
        Description: "Moderately Confused offers a gently absurd and playfully witty take on the vagaries of daily life, technology and politics.",
        Author:      "Jeff Stahler",
        Url:         "https://www.gocomics.com/moderately-confused",
    },
    "speedbump": {
        Id:          100,
        Name:        "Speed Bump",
        Description: "Dave Coverly admits there is no overriding theme, no tidy little philosophy that precisely describes what Speed Bump is about. Basically, he says, if life were a movie, these would be the outtakes.",
        Author:      "Dave Coverly",
        Url:         "https://www.gocomics.com/speedbump",
    },
    "garfield": {
        Id:          110,
        Name:        "Garfield",
        Description: "What a cat! A cat for all seasons. Sassy. Opinionated. This lasagna loving, mailman chasing, sarcastic cat is a classic that readers love. Garfield, Odie and Jon will leave you wanting a daily dose of this beloved bunch!",
        Author:      "Jim Davis",
        Url:         "https://www.gocomics.com/garfield",
    },
    "dicktracy": {
        Id:          120,
        Name:        "Dick Tracy",
        Description: "Dick Tracy is one of America's most-enduring pop-cultural icons, noteworthy for its steadfast, chisel-jawed hero and the gruesome gallery of villains he and his fearless team of Crimestoppers must outwit to put behind bars.",
        Author:      "Joe Staton, Mike Curtis",
        Url:         "https://www.gocomics.com/dicktracy",
    },
    "two-leaf-clover": {
        Id:          130,
        Name:        "Two Leaf Clover",
        Description: "A bittersweet take on many aspects of life. Jokes pretty much tackle anything from cubicle humour, pop culture to the quirkiest sides of human nature.",
        Author:      "Aaron Scott",
        Url:         "http://twoleafclover.net/",
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
