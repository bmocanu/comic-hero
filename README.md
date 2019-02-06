# comic-hero
A daily comic scraper and RSS feed generator for some online comics.

Comic Hero is a server written in Go that periodically checks some of the most famous online 
comics websites for new comic issues. For each new issue, it collects the title, the date and the 
image URL and turns them into Atom and RSS 2.0 feeds. You can then point your RSS client to those
feed for a daily dose of online comics.

# Download
The latest version is 1.5, released Feb 7th, 2019.
Download the Windows and Linux binaries on the Release page.

# Run it
comic-hero needs to be run by you, on your premise, on your server. There is no online As A Service 
installation, as of now (Feb/06/2019). comic-hero is very Docker friendly, so it can also be easily 
run in a Docker container. See the `docker` folder in the repository for a sample of how to make
a Docker image and run it.  

# Configure it
comic-hero is configured using a configuration JSON file. The location of the config file is by
default assumed to be in a `config` folder in the same folder as the binary that you started, and 
the name of the file is assumed to be `config.json`.

The configuration folder and configuration file can be changed with:
```bash
comic-hero -config-dir=/custom-folder -config-file=local-config.json
```

Samples of configuration files can be found in the `config` folder in the repository.

While most of the attributes are self explanatory, here are some that are a bit special:
* `server.contextPath`: the optional name that comes after the host:port part. If you run the server
  in the root of the server web path, this attribute should be "/", otherwise set it to the name
  of your application path (e.g. "/comic-hero" will make the app accessible at http://host:port/comic-hero)
* `server.baseUrl`: the absolute external URL to your application. Usually it is protocol + host + port + contextPath
* `store.issuesStoredPerComic`: the number of comic issues to store per comic and include in the Atom or RSS 2.0 feed.
  After this number, the issues are rolled (oldest one is discarded to make room for the new one)
* `retrieve.issuesFetchingCronJobConfig`: the cron job configuration for checking ALL comics for new issues. 
  You can use https://crontab.guru/ to help you with cron job configuration patterns.
* `comics.enabled`: if comic-hero will scan this website or not. If this is false, the feed is still 
  published but it will have zero issues.
* `comics.proxyImage`: if comic-hero should proxy the actual image. Sometimes comics might be filtered
  by your company web protection filter (e.g. adult content is filtered out). comic-hero generates
  simple feed URLs and, if this option is _true_ it also generates custom URLs, acting as a proxy for 
  the issue image. Your browser will call comic-hero for getting the image, and comic-hero will, in 
  turn, fetch the image and return it.  

# Architecture 
comic-hero has a very clean architecture, with the main components highly decoupled and easy to 
change and extend. 

The `store` as a map of linked lists, each entry having the key the name of the comic and the value
a linked list with all the issues. New issues are added at the front of the linked list (queue), 
and periodically the last issue is discarded (as per configuration).

The `retrieve` functions are called periodically by a cron job. There is a retriever manager which 
calls all registered retrievers for getting new comic issues. Retrievers register dynamically to the 
manager. When a retriever manages to get a new issue, the manager sends the issue to the `store`.

The `serve` functions provide the HTML (+CSS +favicon) interface, and the Atom+RSS 2.0 feed controllers.
When called, the `serve` functions call the `store` to obtain the list of issues for a particular
comic, which are then turned into the respective feed (Atom or RSS 2.0).  

# Extend comic-hero
The application can be easily extended. New retrievers can be added easily and they are automatically
picked up, activated and published. The store can be changed, if desired, to something more fancy (e.g. Redis DB).

# Licensing
Copyright 2019+ Bogdan Mocanu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
> http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
