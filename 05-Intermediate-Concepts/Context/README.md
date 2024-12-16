# The Context
Servers need a way to handle metadata on individual requests. This metadata falls into two general categories: metadata that is required to correctly process the request, and metadata on when to stop processing the request. For example, an HTTP server might want to use a tracking ID to identify a chain of requests through a set of microservices. It also might want to set a timer that ends requests to other microservices if they take too long. Many languages use threadlocal variables to store this kind of information, associating data to a specific operating system thread of execution. This
does’t work in Go because goroutines don’t have unique identities that can be used to look up values. More importantly, threadlocals feel like magic; values go in one place and pop up somewhere else.
  
Go solves the request metadata problem with a construct called the context. Let’s see how to use it correctly.

## What Is the Context?
Rather than add a new feature to the language, a context is simply an instance that meets the Context interface defined in the context package. As you know, idiomatic Go encourages explicit data passing via function parameters. The same is true for the context. It is just another parameter to your function. Just like Go has a convention that the last return value from a function is an error, there is another Go convention that the context is explicitly passed through your program as the first parameter of a
function. The usual name for the context parameter is ctx:

```go
func logic(ctx context.Context, info string) (string, error) {
    // do some interesting stuff here
    return "", nil
}
```

In addition to defining the Context interface, the context package also contains several factory functions for creating and wrapping contexts. When you don’t have an existing context, such as at the entry point to a command-line program, create an empty initial context with the function context.Background. This returns a variable of type context.Context. (Yes, this is an exception to the usual pattern of returning a concrete type from a function call.)

An empty context is a starting point; each time you add metadata to the context, you do so by wrapping the existing context using one of the factory functions in the context package:

```go
ctx := context.Background()
result, err := logic(ctx, "a string")
```

> There is another function, context.TODO, that also creates an
empty context.Context. It is intended for temporary use during
development. If you aren’t sure where the context is going to come
from or how it’s going to be used, use context.TODO to put a place‐
holder in your code. Production code shouldn’t include
context.TODO.


When writing an HTTP server, you use a slightly different pattern for acquiring and passing the context through layers of middleware to the top-level http.Handler. Unfortunately, context was added to the Go APIs long after the net/http package was created. Due to the compatibility promise, there was no way to change the http.Handler interface to add a context.Context parameter.

The compatibility promise does allow new methods to be added to existing types, and that’s what the Go team did. There are two context-related methods on http.Request:

* Context returns the context.Context associated with the request.
* WithContext takes in a context.Context and returns a new http.Request with the old request’s state combined with the supplied context.Context.

Here’s the general pattern:
```go
func Middleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(
        func(rw http.ResponseWriter, req *http.Request) {
            ctx := req.Context()
            // wrap the context with stuff -- we'll see how soon!
            req = req.WithContext(ctx)
            handler.ServeHTTP(rw, req)
        },
    )
}
```

The first thing we do in our middleware is extract the existing context from the request using the Context method. After we put values into the context, we create a new request based on the old request and the now-populated context using the WithContext method. Finally, we call the handler and pass it our new request and the existing http.ResponseWriter.

When you get to the handler, you extract the context from the request using the Context method and call your business logic with the context as the first parameter, just like we saw previously:

```go
func handler(rw http.ResponseWriter, req *http.Request) {
    
    ctx := req.Context()

    err := req.ParseForm()
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        rw.Write([]byte(err.Error()))
        return
    }

    data := req.FormValue("data")
    result, err := logic(ctx, data)
    
    if err != nil {
        rw.WriteHeader(http.StatusInternalServerError)
        rw.Write([]byte(err.Error()))
        return
    }
    
    rw.Write([]byte(result))
}
```

There’s one more situation where you use the WithContext method: when making an HTTP call from your application to another HTTP service. Just like we did when passing a context through middleware, you set the context on the outgoing request using WithContext:

```go
type ServiceCaller struct {
    client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
    
    req, err := http.NewRequest(
        http.MethodGet,
        "http://example.com?data="+data, 
        nil
    )

    if err != nil {
        return "", err
    }
    
    req = req.WithContext(ctx)

    resp, err := sc.client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
    }

    // do the rest of the stuff to process the response
    id, err := processResponse(resp.Body)
    return id, err
}
```

Now that we know how to acquire and pass a context, let’s start making them useful.
We’ll begin with cancellation.

## Cancellation

Imagine that you have a request that spawns several goroutines, each one calling a
different HTTP service. If one service returns an error that prevents you from return‐
ing a valid result, there is no point in continuing to process the other goroutines. In
Go, this is called cancellation and the context provides the mechanism for
implementation.


To create a cancellable context, use the context.WithCancel function. It takes in a context.Context as a parameter and returns a context.Context and a context.CancelFunc. The returned context.Context is not the same context thatwas passed into the function. Instead, it is a child context that wraps the passed-in
parent context.Context. A context.CancelFunc is a function that cancels the context, telling all of the code that’s listening for potential cancellation that it’s time to stop processing.

> We’ll see this wrapping pattern several times. A context is treated as an immutable instance. Whenever we add information to a context, we do so by wrapping an existing parent context with a child context. >
>
> This allows us to use contexts to pass information into deeper layers of the code. The context is never used to pass information out of deeper layers to higher layers.