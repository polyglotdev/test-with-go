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
👆🏾 This is an example of a 💩 test. We are looking to make sure that the constant matches the test. The reality of this test asks that we verify that a engineer can copy and paste a string. If we are verifying that the `const` not being set blows something up, we should test **that thing**. When we try and reverse engineer our code, we write poor tests that are not implementation agnostic.

```go
func DoStuff(inc Incrementor, n int)  {
  for i := 0; i < n; i++ {
    // ...
  }
  inc.Incrementor()
}

func TestDoStuff(t *testing.T) {
  inc := MockIncrementor{}
  initValue := inc.Value
  DoStuff(inc, 3)
  diff = inc.Value - initValue
  if diff != 3 {
    t.Errorf("Increment = %d; want %d" diff, 3)
  }
}
```
In this test we are doing a few different things here:
1. Store the initial value of the incrementor.
2. Call `DoStuff()` and pass the `inc` value along with a expected number of iterations
3. Take that initial value and whatever the current value is and assign it to the `diff` variable
4. Evaluate if `diff` is **NOT** equal to 3 which was our number of iterations from the `DoStuff()` and give an error message.

**Why is this a different type of test that what we wrote previously?**
1. We didn't care **HOW** it iterated, we care that the end goal was reached of `n` iterations resulting in the **expected outcome**.