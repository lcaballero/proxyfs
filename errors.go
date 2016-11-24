package proxyfs

import "errors"

// ErrPathNotAbsolute occurs when a new FileSystem isn't rooted
// via an absolute path.
var ErrPathNotAbsolute = errors.New("asset mount not absolute path error")

// ErrOffsetBeyondBounds occurs when attempting to read outside the
// bounds of the backing file bytes.
var ErrOffsetBeyondBounds = errors.New("offset beyond bounds error")
