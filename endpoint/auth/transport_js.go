package auth

import (
	"io"
	"net/http"
	"strings"
	"syscall/js"
)

type fetchTransport struct{}

var transport http.RoundTripper = fetchTransport{}

func (fetchTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	opts := js.Global().Get("Object").New()
	opts.Set("method", req.Method)
	headers := js.Global().Get("Headers").New()
	for k, v := range req.Header {
		for _, v := range v {
			headers.Call("append", k, v)
		}
	}
	opts.Set("headers", headers)
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body.Close()
		if len(body) > 0 {
			buf := js.Global().Get("Uint8Array").New(len(body))
			js.CopyBytesToJS(buf, body)
			opts.Set("body", buf)
		}
	}
	promise := js.Global().Call("fetch", req.URL.String(), opts)
	var (
		cb1, cb2, success, failure js.Func
		res                        = &http.Response{}
		doneCh                     = make(chan struct{}, 1)
		errCh                      = make(chan error, 1)
	)
	defer func() {
		cb1.Release()
		cb2.Release()
		success.Release()
		failure.Release()
	}()
	cb1 := js.FuncOf(func(this js.Value, args []js.Value) any {
		r := args[0]
		p := js.Global().Get("Promise")
		if !r.Get("ok").Bool() {
			return p.Call("reject", r.Get("statusText"))
		}
		res.Status = r.Get("status").Int()
		headerIter := r.Get("headers").Call("entries")
		for {
			next := headerIter.call("next")
			if next.Get("done").Bool() {
				break
			}
			entry := next.Get("value")
			key, value := entry.Index(0).String(), entry.Index(1).String()
			res.Header.Add(key, value)
		}
		return r.Call("text")
	})
	cb2 := js.FuncOf(func(this js.Value, args []js.Value) any {
		body := args[0].String()
		res.Body = io.NopCloser(strings.NewReader(body))
		return nil
	})
	success := js.FuncOf(func(this js.Value, args []js.Value) any {
		doneCh <- struct{}{}
		return nil
	})
	failure := js.FuncOf(func(this js.Value, args []js.Value) any {
		errCh <- fmt.Errorf("fetch failed: %s", args[0].String())
		return nil
	})
	go func() {
		promise.Call("then", cb1).Call("then", cb2).Call("then", success).Call("catch", failure)
	}()
	select {
	case <-doneCh:
		return res, nil
	case err := <-errCh:
		return nil, err
	}
}
