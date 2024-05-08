
<p align="center"><h1 align="center">
  go-synology-chat
</h1>

<p align="center">
  Simple wrapper to send Messages to Synology Chat in Go
</p>


# About

Simple wrapper to send Messages to [Synology Chat](https://www.synology.com/es-es/dsm/feature/chat) in Go

🌟 If you want a Node.js version, please check [synology-chat-communicator](https://www.npmjs.com/package/synology-chat-communicator) 🌟

## ❤️ Awesome Features:

- Out of the box simple Interface. 🔥
- Added support for multimedia messages 🍺
- Simple way to send direct messages to one or many users  🎉
- Direct support to list users and channels 🔊
- Zero dependencies 💪
- Easy to use and great test coverage ✅


## Installation

```bash
go get github.com/UlisesGascon/go-synology-chat@1.0.0
```

## Usage

### Simple example

```go
package main

import (
    "fmt"
    "github.com/UlisesGascon/go-synology-chat"
)

func main() {
    baseUrl := "https://<IP-OR-URL>:<PORT>"
    token := "<YOUR-TOKEN>"
    ignoreSSLErrors := true

    sc, err := synologychat.New(baseUrl, token, ignoreSSLErrors)
    if err != nil {
        fmt.Println("Error initializing SynologyChat:", err)
        return
    }

    users, err := sc.GetUsers()
    if err != nil {
        fmt.Println("Error getting users:", err)
        return
    }

    channels, err := sc.GetChannels()
    if err != nil {
        fmt.Println("Error getting channels:", err)
        return
    }

    const userId = 43451
    data, err := sc.SendDirectMessage([]int{userId}, "Hello, World!")
    if err != nil {
        fmt.Println("Error sending direct message:", err)
        return
    }

    fmt.Println(users)
    fmt.Println(channels)
    fmt.Println(data)
}
```

### Disable SSL validation

If you have a Synology NAS with invalid SSL certificates (due to expiration or other issues), you can disable the SSL validation in the requests generated by the library by using the configuration parameter `ignoreSSLErrors`.


## Additional Features

Please check out the official [DSM Documentation](https://kb.synology.com/en-uk/DSM/help/Chat/chat_integration?version=7) to include new features


## Contributing

Please read [CONTRIBUTING.md](https://github.com/UlisesGascon/.github/blob/main/contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/go-synology-chat/tags).

## Authors

- **Ulises Gascón** - Initial work - [@ulisesGascon](https://github.com/ulisesGascon)

See also the list of [contributors](https://github.com/go-synology-chat/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

- This project is under development, but you can help us to improve it! We :heart: FOSS!
