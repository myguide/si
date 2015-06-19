// wrengo v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package wrengo
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

class Strings {
	// Foreign methods defined in Go
    foreign contains(s, substr)
    foreign hasPrefix(s, prefix)
    foreign hasSuffix(s, suffix)
    foreign index(s, sep)

    // Empty constructor
    new {

    }

    // Constructor with _string initializer
    new(string) {
    	_string = string
    }

    // _string setter
    string(string) {
    	_string = string
    }

    // _string getter
    string {
    	return _string
    }

    // Helper method for printing _string
    print {
    	IO.print(_string)
    }

    // Calls foreign contains(_,_) directly on
    // the _string property
    contains(substr) {
    	return this.contains(_string, substr)
    }

    // Calls foreign hasPrefix(_,_) directly on
    // the _string property
    hasPrefix(prefix) {
    	return this.hasPrefix(_string, prefix)
    }

    // Calls foreign hasSuffix(_,_) directly on
    // the _string property
    hasSuffix(suffix) {
    	return this.hasSuffix(_string, suffix)
    }

    // Calls foreign index(_,_) directly on
    // the _string property
    index(sep) {
    	return this.Index(_string, sep)
    }
}
