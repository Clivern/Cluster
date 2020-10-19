<p align="center">
    <img src="/assets/gopher.jpg" width="230" />
    <h3 align="center">Cluster</h3>
    <p align="center">Golang Package for System Clustering.</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/Cluster"><img src="https://travis-ci.com/Clivern/Cluster.svg?branch=main"></a>
        <a href="https://github.com/Clivern/Cluster/releases"><img src="https://img.shields.io/badge/Version-0.0.7-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Cluster"><img src="https://goreportcard.com/badge/github.com/Clivern/Cluster?v=0.0.6"></a>
        <a href="https://github.com/Clivern/Cluster/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>
<br/>

## Documentation

### Usage

To create a cluster

```golang
import (
    "fmt"
    "github.com/clivern/cluster"
)


clus := &cluster.Cluster{}

// Generate a unique name
nodeName := clus.GetNodeName()

// Get a default configs
config := clus.GetConfig()
config.Name = nodeName
config.BindPort = 0 // assign a free port
config.Events = &cluster.NodeEvents{}

// Override configs
clus.SetConfig(config)

clus.AddLocalNode([]string{}) // or []string{"x.x.x.x:port"} in case of the second, third ... node

fmt.Println(clus.GetLocalNode())

// 2020/10/18 20:44:19 [DEBUG] memberlist: Using dynamic bind port 52053
// A node has joined: Clivern-2.local--c5553465-2bc9-4ef1-8a83-384e5a0c4097
// Clivern-2.local--c5553465-2bc9-4ef1-8a83-384e5a0c4097
```

## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Cluster is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/cluster/releases) for changelogs for each release version of Cluster. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/cluster/issues


## Security Issues

If you discover a security vulnerability within Cluster, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Cluster** is authored and maintained by [@clivern](http://github.com/clivern).
