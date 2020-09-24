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
    "freerange": {
        Id:          140,
        Name:        "Free Range",
        Description: "Created by Bill Whitehead, a Midwestern resident, this comic is truly free range, plus gluten-free.",
        Author:      "Bill Whitehead",
        Url:         "https://gocomics.com/freerange",
    },
    "brevity": {
        Id:          150,
        Name:        "Brevity",
        Description: "Not available",
        Author:      "Dan Thompson",
        Url:         "https://gocomics.com/brevity",
    },
    "deflocked": {
        Id:          160,
        Name:        "DeFlocked",
        Description: "DeFlocked is about one sheep’s quest to pull the wool over the world’s eyes. Mamet is the fleecy larger-than-life star of this inter-species comedy, where four of society’s biggest misfits somehow try to become family",
        Author:      "Jeff Corriveau",
        Url:         "https://gocomics.com/deflocked",
    },
    "birdbrains": {
        Id:          170,
        Name:        "BirdBrains",
        Description: "This high-flying single-panel comic spotlights and savors the foibles, stupidity, and goofiness of our colorful unnatural world. It's all about attitude, as the denizens of Bluemel's realm devote themselves to surviving life's pitfalls, whether nesting in quicksand or coping with the technological advances of the tapeworm.",
        Author:      "Thom Bluemel",
        Url:         "https://gocomics.com/birdbrains",
    },
    "peanuts": {
        Id:          180,
        Name:        "Peanuts",
        Description: "If ever there is an iconic comic strip, it is Peanuts. What began in the funny pages in 1950 has developed into an enduring classic. Whether you're persnickety like Lucy, a philosopher like Linus, a joyous Flying Ace like Snoopy, or a lovable underdog like Charlie Brown, there is something to touch your heart or make you laugh in Peanuts.",
        Author:      "Charles Schulz",
        Url:         "https://gocomics.com/peanuts",
    },
    "duplex": {
        Id:          190,
        Name:        "The Duplex",
        Description: "Once upon a time there was a duplex where a young bachelor named Eno and his dog, Fang, shared an ultra-macho haven of beer snacks and male-bonding. Suddenly, their lives turned co-ed when Gina and her poodle, Mitzi, moved into the other half of their building...",
        Author:      "Glenn McCoy, Gary McCoy",
        Url:         "https://gocomics.com/duplex",
    },
    "jerryholbert": {
        Id:          200,
        Name:        "Jerry Holbert",
        Description: "A former staff cartoonist for the Boston Herald since 1986, Holbert serves up solid conservative commentary, delivered with a smile.",
        Author:      "Jerry Holbert",
        Url:         "https://gocomics.com/jerryholbert",
    },
    "life-on-earth": {
        Id:          210,
        Name:        "Life on Earth",
        Description: "Life On Earth is an absurdist take on the human condition and how it relates to the natural and not-so-natural world.",
        Author:      "Ham",
        Url:         "https://gocomics.com/life-on-earth",
    },
    "whyatt-cartoons": {
        Id:          220,
        Name:        "Whyatt Cartoons",
        Description: "Not available",
        Author:      "Tim Whyatt",
        Url:         "https://gocomics.com/whyatt-cartoons",
    },
    "pot-shots": {
        Id:          230,
        Name:        "Pot-Shots",
        Description: "Ashleigh Brilliant's POT-SHOTS are illustrated epigrams, never longer than 17 words, capable of standing alone without requiring any illustration. But the illustrations add a special dimension to the finished product.",
        Author:      "Ashleigh Brilliant",
        Url:         "https://gocomics.com/pot-shots",
    },
    "waynovision": {
        Id:          240,
        Name:        "WaynoVision",
        Description: "WaynoVision gives its readers a chance to view, comment and celebrate on the surrealism and absurdity of everyday life, in ways that are sometimes silly, sometimes smart, but always funny.",
        Author:      "Wayno",
        Url:         "https://gocomics.com/waynovision",
    },
    "dogeatdoug": {
        Id:          250,
        Name:        "Dog Eat Doug",
        Description: "Babies and puppies are both quite cute, but underneath the soft, cuddly exteriors lie the fearsome hearts of competitors. Well, not really. When a new baby joins the household, Sophie the dog is initially irritated, but eventually comes to see the baby, Doug, as the asset he is: a better way to get snacks.",
        Author:      "Brian Anderson",
        Url:         "https://gocomics.com/dogeatdoug",
    },
    "luann": {
        Id:          260,
        Name:        "Luann",
        Description: "LUANN is about the trials of becoming a young adult: the hilarity and drama, triumphs and flops, friendships and rivalries. Rich in character and intriguing 'what'll happen next?!' stories, LUANN is a compelling saga of life's most volatile stage.",
        Author:      "Greg Evans",
        Url:         "https://gocomics.com/luann",
    },
    "mikeluckovich": {
        Id:          270,
        Name:        "Mike Luckovich",
        Description: "Mike Luckovich, editorial cartoonist of The Atlanta Journal-Constitution, won the 1995 Pulitzer Prize for cartooning. His work also appears in Time, the New York Times and other media. He is distributed nationally by Creators Syndicate.",
        Author:      "Mike Luckovich",
        Url:         "https://gocomics.com/mikeluckovich",
    },
    "nonsequitur": {
        Id:          280,
        Name:        "Non Sequitur",
        Description: "This hilarious creation is not only creative but also clever. It tackles current cultural issues such as politics, celebrities, male-female relations, materialistic desires and society's obsession with weight. Non Sequitur will have you laughing at the controversy of everyday life.",
        Author:      "Wiley Miller",
        Url:         "https://gocomics.com/nonsequitur",
    },
    "bignate": {
        Id:          290,
        Name:        "Big Nate",
        Description: "Nate is 11 years old, four-and-a-half feet tall, and the all-time record holder for detentions in school history. He's a self-described genius and sixth grade Renaissance Man. Nate, who lives with his dad and older sister, enjoys pestering his family and teachers with his sarcasm.",
        Author:      "Lincoln Peirce",
        Url:         "https://gocomics.com/bignate",
    },
    "inthebleachers": {
        Id:          300,
        Name:        "In the Bleachers",
        Description: "Whatever your athletic interest, golf, baseball, running, or basketball and whether you haven’t picked up a ball since high school or you’re a serious sports fanatic, everyone can see the humor and irony highlighted by In the Bleachers.",
        Author:      "Ben Zaehringer",
        Url:         "https://gocomics.com/inthebleachers",
    },
    "imaginethis": {
        Id:          310,
        Name:        "Imagine This",
        Description: "Some adults wistfully wonder what became of the imaginary friends of their childhood. Not Darin. His are sitting on the couch, too-constant companions sharing every strange moment of his cheerfully frustrated life. There’s Dewey the Dinosaur, with heart of gold and head of fluff, and Clovis, the Teddy Bear With Serious Anger Issues (and some very bad habits). An unemployed 30-year-old graphic designer, Darin lives in his father’s basement, which he also shares with the enigmatic Robert the Plant.",
        Author:      "Lucas Turnbloom",
        Url:         "https://gocomics.com/imaginethis",
    },
    "bloomcounty": {
        Id:          320,
        Name:        "Bloom County",
        Description: "Bloom County, a 1980s cartoon-comic strip that dealt with socio-political issues as seen through the eyes of highly exaggerated characters (e.g. Bill the Cat and Opus the Penguin) and humorous analogies.",
        Author:      "Berkeley Breathed",
        Url:         "https://gocomics.com/bloomcounty",
    },
    "wizardofid": {
        Id:          330,
        Name:        "Wizard of Id",
        Description: "Not available",
        Author:      "Brant Parker, Johnny Hart",
        Url:         "https://gocomics.com/wizardofid",
    },
    "strangebrew": {
        Id:          340,
        Name:        "Strange Brew",
        Description: "With Strange Brew, John Deering - famous for his biting humor and political savvy as chief editorial cartoonist for the Arkansas Democrat-Gazette - has an outlet for his creative sense of humor and quirky view of life.",
        Author:      "John Deering",
        Url:         "https://gocomics.com/strangebrew",
    },
    "wumo": {
        Id:          350,
        Name:        "Wulff & Morgenthaler",
        Description: "WuMo celebrates life's absurdity and bittersweet ironies, holding up a funhouse mirror to our modern world and those who live in it. WuMo's inventiveness is reminiscent of their countryman Hans Christian Andersen - if Andersen's fairy tales had been populated by sadistic pandas, disgruntled office workers, crazy beavers, Albert Einstein, Snoop Dogg and Darth Vader.",
        Author:      "Mikael Wulff, Anders Morgenthaler",
        Url:         "https://gocomics.com/wumo",
    },
    "gasolinealley": {
        Id:          360,
        Name:        "Gasoline Alley",
        Description: "Gasoline Alley by Jim Scancarelli is a gentle, good-natured continuing story of four generations of Wallets. Readers return daily for this positive slice of life, with universal themes and commonplace situations.",
        Author:      "Jim Scancarelli",
        Url:         "https://gocomics.com/gasolinealley",
    },
    "getfuzzy": {
        Id:          370,
        Name:        "Get Fuzzy",
        Description: "Housecats are known to be aloof, but 'cat-titude' reaches new heights in Get Fuzzy, the bitingly hilarious comic strip from cartoonist Darby Conley.",
        Author:      "Darby Conley",
        Url:         "https://gocomics.com/getfuzzy",
    },
    "roseisrose": {
        Id:          380,
        Name:        "Rose is Rose",
        Description: "Rose is Rose presents the extraordinary nature of everyday life as seen through the eyes of the Gumbo family. The strip stars child-at-heart Rose and her ASD (Attentiveness Surplus Disorder) husband Jimbo. Their gentle son Pasquale is watched over by his Guardian Angel who morphs from tiny cherub into gargantuan protector. Family kitten Peekaboo boasts that her humans are the cutest in town.",
        Author:      "Don Wimmer, Pat Brady",
        Url:         "https://gocomics.com/roseisrose",
    },
    "robrogers": {
        Id:          390,
        Name:        "Rob Rogers",
        Description: "Rob Rogers is an award-winning freelance editorial cartoonist living in Pittsburgh. His cartoons have been vexing and entertaining readers there since 1984, first with the Pittsburgh Press (1984-93) and then the Pittsburgh Post-Gazette (1993-2018).",
        Author:      "Rob Rogers",
        Url:         "https://gocomics.com/robrogers",
    },
    "offthemark": {
        Id:          400,
        Name:        "Off the Mark",
        Description: "The off-the-wall humor of off the mark puts a refreshingly spin on the things we see everyday... from your favorite icons to your least favorite trends, from commercials to pets to computers. Slightly skewed and just a little twisted, off the mark scores a bull’s eye with readers looking for a laugh.",
        Author:      "Mark Parisi",
        Url:         "https://gocomics.com/offthemark",
    },
    "libertymeadows": {
        Id:          410,
        Name:        "Liberty Meadows",
        Description: "Featuring talking animals and dimwitted humans, Liberty Meadows is hilarious. While the humans worry about the development of the various animals, no one is having more fun than the animals themselves.",
        Author:      "Frank Cho",
        Url:         "https://gocomics.com/libertymeadows",
    },
    "heathcliff": {
        Id:          420,
        Name:        "Heathcliff",
        Description: "Called the 'Cat of the Century' by a major cat magazine, Heathcliff has been an incredible success. The cat is loved and recognized by millions in all corners of the globe.",
        Author:      "George Gately",
        Url:         "https://gocomics.com/heathcliff",
    },
    "herbandjamaal": {
        Id:          430,
        Name:        "Herb and Jamaal",
        Description: "Herb and Jamaal have seen each other through thick and thin. They are best friends. They get through everything in life together: household chores, their wives, work and even golf. These best friends will make you cherish yours!",
        Author:      "Stephen Bentley",
        Url:         "https://gocomics.com/herbandjamaal",
    },
    "harley": {
        Id:          440,
        Name:        "Harley",
        Description: "The open road, gasoline dripping from the pump, the smell of iced coffee with just a hint of tire rubber and the sound of an engine roaring like a panther whose foot has just been stepped on. Harley and his feline compadre ride the roads of life like champions.",
        Author:      "Dan Thompson",
        Url:         "https://gocomics.com/harley",
    },
    "9to5": {
        Id:          450,
        Name:        "9 to 5",
        Description: "Rancorous bosses, quirky workers, and an up-and-down stock market populate the world of 9 to 5. A cast of regular characters include J.B. Wells (the boss), Sims (office flunkey), and Ms. Forbes. While mainly a satire on business, the comic also pokes fun at technology, relationships, dogs and cats, and life in general.",
        Author:      "Harley Schwadron",
        Url:         "https://gocomics.com/9to5",
    },
    "herman": {
        Id:          460,
        Name:        "Herman",
        Description: "Herman, the hilarious groundbreaking cartoon feature that appears in hundreds of newspapers worldwide.",
        Author:      "Jim Unger",
        Url:         "https://gocomics.com/herman",
    },
    "pickles": {
        Id:          470,
        Name:        "Pickles",
        Description: "Pickles tells the story of Earl and Opal Pickles as they enjoy their golden years surrounded by friends and family.",
        Author:      "Brian Crane",
        Url:         "https://gocomics.com/pickles",
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
