# Go Milestone
Is a module that allows for the user to check the speed of their go code between different "steps". This can be used to identify slow logic within the user's code base.

## Usage
You can import this module into to your project by referencing it as an `import` in the file you want to use it:
```go
import gomilestone "github.com/alexroden/go-milestone"
```
And run `go mod tidy` to add it to your `go.mod`.

Alternatively, you can explicitly add the module to you project:
```shell
go get -u github.com/alexroden/go-milestone
```

Once you have installed the module into your project you can use it.
Firstly, you will need to `Start` Milestone:
```go
gomilestone.Start()
```

After this you can then start to put `Step`s within your code base, which will mark in seconds the time between the `Start`, and each `Step`:
```go
gomilestone.Step()
```
You can also add additional messaging around these steps:
```go
gomilestone.Step(gomilestone.WithMessage("Hello there!"))
```

Once you have finished your tracking, you can then request your Milestone `Report`:
```go
gomilestone.Report()
```
This report is a Slice of report objects, containing the `Step` number, `Time` between the start, and any `Message` that has been set at a `Step`

You can also get the Milestone instance at anypoint by calling `GetInstance`.

## License
[MIT](https://choosealicense.com/licenses/mit/)