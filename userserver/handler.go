Create(context.Context, *User, *Response) error
Get(context.Context, *User, *Response) error
GetAll(context.Context, *Request, *Response) error
Auth(context.Context, *User, *Token) error
ValidateToken(context.Context, *Token, *Token) error