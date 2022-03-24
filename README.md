# RegexMatch

A library for parsing regex matches into `maps` and generic `T` structs when compiled with Go 1.18.

## Why?

This is really simple functionality. Like, stupid simple. So why create a library for it? The answer
lies with the DRY principal. Don't Repeat Yourself. I have copy/pasted and had to look up how to
associate named regex matches to their values countless times.

With the introduction of Generics in 1.18, this makes even more sense to implement now.

### Examples

The primary example I want to showcase is with Gorilla Mux. You can create custom path matchers
for awkward paths. This is one of the more frequent things I need (I know I can't be the only one)
and it isn't really simple or baked into the library.

But as a sample, here is how you would implement a custom matcher that provides path variables that
include unescaped forward slashes:

```go
func matcher(r *http.Request, rm *mux.RouteMatch) bool {
	// Handle awkward URL paths, such as:
	// - https://www.terraform.io/internals/provider-registry-protocol#list-available-versions
	regex := regexp.MustCompile(`api/(?P<namespace>(?:\w|-|\/)+)/(?P<type>(?:\w+|-)+)/versions/?`)
	matches, isMatch := regexstruct.RegexMatchMap(regex, r.URL.Path)
	if !isMatch {
		return false
	}
	rm.Vars = make(map[string]string)
	for key, value := range matches {
		rm.Vars[key] = value
	}
	return true
}
```
