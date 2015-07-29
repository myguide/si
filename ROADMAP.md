# Roadmap for Wrengo

I'm still not very sure what direction I want this project to move in. I think
in the near future the name of it will change to something else. But in terms
of what it does (or will do) is still is not clear.

Some ideas I've had in terms of basic functionality have been:

#### File / directory watching
I'd like to be able to write a little script in Wren that watches a file or
a directory and then allows you to pass through a function as an action. So
having a public API similar to this:

```dart
var w = Watcher.new()

w.on('create', Fn.new{|event|
	IO.print("A new file has been added to the directory!")
})

w.attach('/path/to/my/files/')

w.start()
```

Still very rough. And I'm not sure what will actually go inside the event
just yet either.

Anything could be done within the action, in this example it's just a simple
`IO.print()` but the sort of stuff I'll be looking at would be code minifying,
file creation, deletion, execing commands (maybe) and other things similar to
that. Essentially, I want to be able to build a fairly basic project using
this.

#### Reading and writing to and from files
This should be kept fairly simple to start with. Being able to create a file
with permissions. The Go standard library API will be a fairly big influence
in relation to the API for this. But it could be even easier than that.

```dart
var f = File.new()

f.write("some content", "my-file.txt", 0775)
```

I like the idea of creating if not already existing, hence the permissions there.
Or indeed changing the permissions if it already exists.

#### Some sort of web interface
...
