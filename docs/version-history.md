1.2
---
* added several new comics from GoComics (one parser for many comics, yaay!)
* more fixings on the path concatenation, now /comic-hero and /comic-hero/ work the same
* better configuration organization
* on startup the program does an initial fetch all

1.1
---
* fixed the web path concatenation
* fixed the embedding of images in the feed descriptions
* removed the BASE_URL env property, as it was unnecessary (CONTEXT_PATH is enough)
* switched the IDs of feeds from INT32 hashes to SHA1, for better uniqueness guarantees
* embedded the IDs of feeds (SHA1) as feed's ID (Atom) and GUID (RSS 2.0)
* added versioning of the app and this version history document

1.0
---
* first release version, deployed in a private server ("production" env)
* three comics supported: Sinfest, Dilbert, Oglaf
* support for Atom and RSS 2.0 feeds
* support for env properties configuration (for easy Docker deployment)
