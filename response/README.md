### Response

Response helper package of [core](https://github.com/shkshariq/go-util/util) library

#### Usage

Render an error

```go
import "github.com/shkshariq/go-util/response"

response.HandleError(ctx, err, w)

```
Above will check the error type and if 
- Error is a [DomainError]() api response render helpers will generate following output

```json
    
```


