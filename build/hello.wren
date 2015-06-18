IO.print("Hello, World!")

var m = new Markdown
var c = m.Parse("#### Hello")

IO.print(c)

var f = new File
var t = f.Read("hello.wren")

IO.print(t)

var a = new Asink
a.DoLs