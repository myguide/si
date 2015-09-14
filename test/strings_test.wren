class TestStrings {
	static string(t) {
		var s = Strings.new()
		s.string("Hello")
		var res = s.string

		if (!t.assertEquals("Hello", res)) {
			IO.print("Error: testString expects Hello, got ", res)
		}
	}

	static contains(t) {
		var s = Strings.new()
		s.string("Hello, World!")
		var res = s.contains("ello,")

		if (!t.assertEquals(true, res)) {
			IO.print("Error: testContains expects true, got ", res)
		}
	}

	static hasPrefix(t) {
		var s = Strings.new()
		s.string("Hello, World!")
		var res = s.hasPrefix("Hello")
		if (!t.assertEquals(true, res)) {
			IO.print("Error: testHasPrefix expects true, got ", res)
		}
	}

	static hasSuffix(t) {
		var s = Strings.new()
		s.string("Hello, World!")
		var res = s.hasSuffix("orld!")

		if (!t.assertEquals(true, res)) {
			IO.print("Error: testHasSuffix expects true, got ", res)
		}
	}

	static index(t) {
		var s = Strings.new()
		s.string("Hello, World!")
		var res = s.index("o")

		if (!t.assertEquals(4, res)) {
			IO.print("Error: testIndex expects 4, got ", res)
		}
	}

	static lastIndex(t) {
		var s = Strings.new()
		s.string("Hello, World!")
		var res = s.lastIndex("l")

		if (!t.assertEquals(10, res)) {
			IO.print("Error: testLastIndex expects 10, got ", res)
		}
	}
}

TestStrings.string(Testing.new())
TestStrings.contains(Testing.new())
TestStrings.hasPrefix(Testing.new())
TestStrings.hasSuffix(Testing.new())
TestStrings.index(Testing.new())
TestStrings.lastIndex(Testing.new())
