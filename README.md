# Test with Go
Link to course: [Test With Go](https://courses.calhoun.io/lessons/)

## Writing great tests
- Test with a purpose
- Don't overdo it
- Testing is a skill, skills evolve

### Testing with a purpose
```go
const SomeString = "This is my String"

func TestSomeString(t *testing.T) {
  want := "this is my string"
  if pkg.SomeString != want {
    t.Errorf("SomeString: %q but want: %q", pkg.SomeString, want)
  }
}
```
👆🏾 This is an example of a 💩 test. We are looking to make sure that the constant matches the test. The reality of this test asks that we verify that a engineer can copy and paste a string.  