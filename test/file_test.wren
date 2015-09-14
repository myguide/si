class TestFile {
	static read(t) {
		var f = File.new()
		var res = f.read("./build/hello.wren")

		if (!t.assertEquals("IO.print(\"Hello, World!\")\n", res)) {
			IO.print("Error: testRead expects IO.print(\"Hello, World!\")\n, got ", res)
		}
	}
}

TestFile.read(Testing.new())
