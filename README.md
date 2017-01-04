cligol
======

Simple and flexible CLI application written on GO.

### Features
- Read log files via SSH.

### Installation
```go get github.com/ashlinchak/cligol```

Go to the **cligol** source folder and build:

```go build *.go```

Prepare config files:

```cp ./config/servers.json.example ./config/servers.json```

Edit **servers.json** with your projects.

Add compiled file with the config path for searching. On Ubuntu you can do it by adding **cligol** directory to the PATH variable.
```export PATH=$PATH:/home/user/go/src/github.com/ashlinchak/cligol```

### Usage
```cligol server -log site.com -n 100```

### License

[MIT](LICENSE)
