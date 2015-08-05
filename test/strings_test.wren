var errors = 0

var testString = Fn.new {
	var s = Strings.new()
	s.string("Hello")
	var res = s.string
	if ("Hello" != res) {
		IO.print("Error: testString expects Hello, got ", res)
		errors = errors + 1
	}
}

var testContains = Fn.new {
	var s = Strings.new()
	s.string("Hello, World!")
	var res = s.contains("ello,")
	if (true != res) {
		IO.print("Error: testContains expects true, got ", res)
		errors = errors + 1
	}
}

var testHasPrefix = Fn.new {
	var s = Strings.new()
	s.string("Hello, World!")
	var res = s.hasPrefix("Hello")
	if (true != res) {
		IO.print("Error: testHasPrefix expects true, got ", res)
		errors = errors + 1
	}
}

var testHasSuffix = Fn.new {
	var s = Strings.new()
	s.string("Hello, World!")
	var res = s.hasSuffix("orld!")
	if (true != res) {
		IO.print("Error: testHasSuffix expects true, got ", res)
		errors = errors + 1
	}
}

var testIndex = Fn.new {
	var s = Strings.new()
	s.string("Hello, World!")
	var res = s.index("o")
	if (4 != res) {
		IO.print("Error: testIndex expects 4, got ", res)
		errors = errors + 1
	}
}

var testLastIndex = Fn.new {
	var s = Strings.new()
	s.string("Hello, World!")
	var res = s.lastIndex("l")
	if (10 != res) {
		IO.print("Error: testLastIndex expects 10, got ", res)
		errors = errors + 1
	}
}

testString.call()
testContains.call()
testHasPrefix.call()
testHasSuffix.call()
testIndex.call()
testLastIndex.call()

if (errors == 0) {
	IO.print("All Strings tests passed!")
}

if (errors > 0) {
	IO.print(errors, " test(s) failing for Strings.")
}