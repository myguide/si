class Strings {
    foreign contains(s, substr)
    foreign hasPrefix(s, prefix)
    foreign hasSuffix(s, suffix)
    foreign index(s, sep)

    new {
    	return this
    }

    new(string) {
    	_string = string
    }

    print {
    	IO.print(_string)
    }

    contains(substr) {
    	return this.contains(_string, substr)
    }

    hasPrefix(prefix) {
    	return this.hasPrefix(_string, prefix)
    }

    hasSuffix(suffix) {
    	return this.hasSuffix(_string, suffix)
    }

    index(sep) {
    	return this.Index(_string, sep)
    }
}
