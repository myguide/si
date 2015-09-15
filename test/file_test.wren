class TestFile {
	static read(t) {
		var f = File.new()
		var res = f.read("./build/hello.wren")

		if (!t.assertEquals("IO.print(\"Hello, World!\")\n", res)) {
			IO.print("Error: testRead expects IO.print(\"Hello, World!\")\n, got ", res)
		}
	}

	static write(t) {
		var f = File.new()
		var res = f.write("test.txt", "Test Content")
		if (!t.assertEquals(true, res)) {
			IO.print("Error: testWrite expects true, got ", res)
		}
	}
}

TestFile.read(Testing.new())
TestFile.write(Testing.new())
