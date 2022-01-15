# Learn go with tests

https://github.com/quii/learn-go-with-tests

## Notes

### Select

- By prefixing a function call with defer it will now call that function at the end of the containing function.
- You can wait for values to be sent to a channel with `myVar := <-ch`. This is a blocking call, as you're waiting for a value.
- What select lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
- **Always make channels** Notice how we have to use make when creating a channel; rather than say `var ch chan struct{}`. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc. For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
- `time.After` can be used for timeouts `case <-time.After(10 * time.Second):`
- `httptest` A convenient way of creating test servers so you can have reliable and controllable tests.

```go
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
```
