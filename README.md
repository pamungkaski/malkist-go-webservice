# Malkist Go Weeb Service

[![Build Status](https://travis-ci.com/pamungkaski/malkist-go-webservice.svg?branch=master)](https://travis-ci.com/pamungkaski/malkist-go-webservice)

A simple webservice implementing [malkist-go](https://github.com/pamungkaski/malkist-go)

## Instalation
### Prerequisites

- `go` (>= 1.5.1)
- `git`

**Install dep**

This project uses dep as a package manager for Go

It is strongly recommended that you use a released version. Release binaries are available on the [releases](https://github.com/golang/dep/releases) page.

On MacOS you can install or upgrade to the latest released version with Homebrew:

```sh
$ brew install dep
$ brew upgrade dep
```

On other platforms you can use the `install.sh` script:

```sh
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

It will install into your `$GOPATH/bin` directory by default or any other directory you specify using the `INSTALL_DIRECTORY` environment variable.

If your platform is not supported, you'll need to build it manually or let the team know and we'll consider adding your platform
to the release builds.

### Get the code

Assuming you have Go installed and configured (have $GOPATH setup and pointing to a single directory) here is how to build.

Check out the code

    $ go get github.com/pamungkaski/malkist-go-webservice

After that do

    $ dep ensure

### Running the app

After all the instalation, create `.env` file in the root folder to set the environment variable. What's need to be inside the 
`.env` is:

    KEY = [YOUR API KEY]
    PORT = [DESIRED PORT]
    
After that run:

    $ go instal ./...
    $ $GOPATH/bin/malkist-go-webservice

You can look for the endpoint documentation **[HERE](https://documenter.getpostman.com/view/3062890/malkist-web-service/RW89Jom2)**

### Testing the app

run :
    
    $ go test -v ./...

### Console Program
Run:

    $ cd cmd && cd src && go build -o ../binary/malkist -v && cd .. && cd ..
    $ ./cmd/binary/malkist
    
## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/pamungkaski/malkist-go-webservice . This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the Malkist project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/pamungkaski/Malkist-Ruby/blob/master/CODE_OF_CONDUCT.md).
