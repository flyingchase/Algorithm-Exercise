package exercises

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := bytes.NewBuffer([]byte("woshigezhu nizhidaoma?"))
	b := make([]byte, buf.Len())
	n, err := buf.Read(b)
	fmt.Printf("%s %v\n", b[:n], err)

	buf.WriteString("tayeshigezhu\n")
	buf.WriteTo(os.Stdout)

	n, err = buf.Write(b)
	fmt.Printf("%d %s %v\n", n, buf.String(), err)

	r, w := io.Pipe()
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {

			n, err = r.Read(buf)
			r.CloseWithError(errors.New("closed by reading pipe"))
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])

		}
	}()
	n, err = w.Write([]byte("hello world!"))
	fmt.Printf("write %d %v \n", n, err)

}

/*
------------------------------

// ReaderFrom 接口包装了基本的 ReadFrom 方法，用于从 r 中读取数据存入自身。
// 直到遇到 EOF 或读取出错为止，返回读取的字节数和遇到的错误。
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}

------------------------------

// WriterTo 接口包装了基本的 WriteTo 方法，用于将自身的数据写入 w 中。
// 直到数据全部写入完毕或遇到错误为止，返回写入的字节数和遇到的错误。
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}

------------------------------

// ReaderAt 接口包装了基本的 ReadAt 方法，用于将自身的数据写入 p 中。
// ReadAt 忽略之前的读写位置，从起始位置的 off 偏移处开始读取。
//
// 返回写入的字节数和遇到的错误，如果 p 被写满，则 err 会返回 nil。如果 p 没
// 有被写满，则会返回一个错误信息用于说明为什么没有写满（比如 io.EOF）。在这
// 方面 ReadAt 比 Read 更严格。如果 p 被写满的同时，自身的数据也刚好被读完，
// 则 err 即可以返回 nil 也可以返回 io.EOF。
//
// 即使不能将 p 填满，ReadAt 在被调用时也可能会使用整个 p 的空间作为缓存空间。
// 如果 ReadAt 自身的数据是从其它地方（比如网络）获取数的，那么在写入 p 的时
// 候，如果没有把 p 写满（比如网络延时），则 ReadAt 会阻塞，直到获取更多的数
// 据把 p 写满，或者所有数据都获取完毕，或者遇到读取错误（比如超时）时才返回。
// 在这方面，ReadAt 和 Read 是不同的。
//
// 如果 ReadAt 读取的对象是某个有偏移量的底层数据流时，则 ReadAt 方法既不能影
// 响底层的偏移量，也不应该被底层的偏移量影响。
//
// ReadAt 的调用者可以对同一数据流并行执行 ReadAt 方法。
//
// ReaderAt 的实现者不应该持有 p。
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}

------------------------------

// WriterAt 接口包装了基本的 WriteAt 方法，用于将 p 中的数据写入自身。
// ReadAt 忽略之前的读写位置，从起始位置的 off 偏移处开始写入。
//
// 返回写入的字节数和遇到的错误。如果 p 没有被读完，则必须返回一个 err 值来说
// 明为什么没有读完。
//
// 如果 WriterAt 写入的对象是某个有偏移量的底层数据流时，则 ReadAt 方法既不能
// 影响底层的偏移量，也不应该被底层的偏移量影响。
//
// WriterAt 的调用者可以对同一数据流的不同区段并行执行 WriteAt 方法。
//
// WriterAt 的实现者不应该持有 p。
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}

------------------------------

// ByteReader 接口包装了基本的 ReadByte 方法，用于从自身读出一个字节。
// 返回读出的字节和遇到的错误。
type ByteReader interface {
    ReadByte() (c byte, err error)
}

------------------------------

// ByteScanner 在 ByteReader 的基础上增加了一个 UnreadByte 方法，用于撤消最后
// 一次的 ReadByte 操作，以便下次的 ReadByte 操作可以读出与前一次一样的数据。
// UnreadByte 之前必须是 ReadByte 才能撤消成功，否则可能会返回一个错误信息（根
// 据不同的需求，UnreadByte 也可能返回 nil，允许随意调用 UnreadByte，但只有最
// 后一次的 ReadByte 可以被撤销，其它 UnreadByte 不执行任何操作）。
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}

------------------------------

// ByteWriter 接口包装了基本的 WriteByte 方法，用于将一个字节写入自身
// 返回遇到的错误
type ByteWriter interface {
    WriteByte(c byte) error
}

------------------------------

// RuneReader 接口包装了基本的 ReadRune 方法，用于从自身读取一个 UTF-8 编码的
// 字符到 r 中。
// 返回读取的字符、字符的编码长度和遇到的错误。
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}

------------------------------

// RuneScanner 在 RuneReader 的基础上增加了一个 UnreadRune 方法，用于撤消最后
// 一次的 ReadRune 操作，以便下次的 ReadRune 操作可以读出与前一次一样的数据。
// UnreadRune 之前必须是 ReadRune 才能撤消成功，否则可能会返回一个错误信息（根
// 据不同的需求，UnreadRune 也可能返回 nil，允许随意调用 UnreadRune，但只有最
// 后一次的 ReadRune 可以被撤销，其它 UnreadRune 不执行任何操作）。
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}

------------------------------

// bytes.NewBuffer 实现了很多基本的接口，可以通过 bytes 包学习接口的实现
func main() {
	buf := bytes.NewBuffer([]byte("Hello World!"))
	b := make([]byte, buf.Len())

	n, err := buf.Read(b)
	fmt.Printf("%s   %v\n", b[:n], err)
	// Hello World!   <nil>

	buf.WriteString("ABCDEFG\n")
	buf.WriteTo(os.Stdout)
	// ABCDEFG

	n, err = buf.Write(b)
	fmt.Printf("%d   %s   %v\n", n, buf.String(), err)
	// 12   Hello World!   <nil>

	c, err := buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err)
	// H   ello World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err)
	// e   llo World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   bytes.Buffer: UnreadByte: previous operation was not a read
}

------------------------------------------------------------

// WriteString 将字符串 s 写入到 w 中，返回写入的字节数和遇到的错误。
// 如果 w 实现了 WriteString 方法，则优先使用该方法将 s 写入 w 中。
// 否则，将 s 转换为 []byte，然后调用 w.Write 方法将数据写入 w 中。
func WriteString(w Writer, s string) (n int, err error)

// ReadAtLeast 从 r 中读取数据到 buf 中，要求至少读取 min 个字节。
// 返回读取的字节数和遇到的错误。
// 如果 min 超出了 buf 的容量，则 err 返回 io.ErrShortBuffer，否则：
// 1、读出的数据长度 == 0  ，则 err 返回 EOF。
// 2、读出的数据长度 <  min，则 err 返回 io.ErrUnexpectedEOF。
// 3、读出的数据长度 >= min，则 err 返回 nil。
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

// ReadFull 的功能和 ReadAtLeast 一样，只不过 min = len(buf)
func ReadFull(r Reader, buf []byte) (n int, err error)

// CopyN 从 src 中复制 n 个字节的数据到 dst 中，返回复制的字节数和遇到的错误。
// 只有当 written = n 时，err 才返回 nil。
// 如果 dst 实现了 ReadFrom 方法，则优先调用该方法执行复制操作。
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

// Copy 从 src 中复制数据到 dst 中，直到所有数据都复制完毕，返回复制的字节数和
// 遇到的错误。如果复制过程成功结束，则 err 返回 nil，而不是 EOF，因为 Copy 的
// 定义为“直到所有数据都复制完毕”，所以不会将 EOF 视为错误返回。
// 如果 src 实现了 WriteTo 方法，则调用 src.WriteTo(dst) 复制数据，否则
// 如果 dst 实现了 ReadeFrom 方法，则调用 dst.ReadeFrom(src) 复制数据
func Copy(dst Writer, src Reader) (written int64, err error)

// CopyBuffer 相当于 Copy，只不 Copy 在执行的过程中会创建一个临时的缓冲区来中
// 转数据，而 CopyBuffer 则可以单独提供一个缓冲区，让多个复制操作共用同一个缓
// 冲区，避免每次复制操作都创建新的缓冲区。如果 buf == nil，则 CopyBuffer 会
// 自动创建缓冲区。
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)

------------------------------

// 示例：WriteString、ReadAtLeast、ReadFull
func main() {
	io.WriteString(os.Stdout, "Hello World!\n")
	// Hello World!

	r := strings.NewReader("Hello World!")
	b := make([]byte, 15)

	n, err := io.ReadAtLeast(r, b, 20)
	fmt.Printf("%q   %d   %v\n", b[:n], n, err)
	// ""   0   short buffer

	r.Seek(0, 0)
	b = make([]byte, 15)

	n, err = io.ReadFull(r, b)
	fmt.Printf("%q   %d   %v\n", b[:n], n, err)
	// "Hello World!"   12   unexpected EOF
}

------------------------------

// 示例：CopyN、Copy、CopyBuffer
func main() {
	r := strings.NewReader("Hello World!")
	buf := make([]byte, 32)

	n, err := io.CopyN(os.Stdout, r, 5) // Hello
	fmt.Printf("\n%d   %v\n\n", n, err) // 5   <nil>

	r.Seek(0, 0)
	n, err = io.Copy(os.Stdout, r)      // Hello World!
	fmt.Printf("\n%d   %v\n\n", n, err) // 12   <nil>

	r.Seek(0, 0)
	r2 := strings.NewReader("ABCDEFG")
	r3 := strings.NewReader("abcdefg")

	n, err = io.CopyBuffer(os.Stdout, r, buf) // Hello World!
	fmt.Printf("\n%d   %v\n", n, err)         // 12   <nil>

	n, err = io.CopyBuffer(os.Stdout, r2, buf) // ABCDEFG
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>

	n, err = io.CopyBuffer(os.Stdout, r3, buf) // abcdefg
	fmt.Printf("\n%d   %v\n", n, err)          // 7   <nil>
}

------------------------------------------------------------

// LimitReader 对 r 进行封装，使其最多只能读取 n 个字节的数据。相当于对 r 做了
// 一个切片 r[:n] 返回。底层实现是一个 *LimitedReader（只有一个 Read 方法）。
func LimitReader(r Reader, n int64) Reader

// MultiReader 将多个 Reader 封装成一个单独的 Reader，多个 Reader 会按顺序读
// 取，当多个 Reader 都返回 EOF 之后，单独的 Reader 才返回 EOF，否则返回读取
// 过程中遇到的任何错误。
func MultiReader(readers ...Reader) Reader

// MultiReader 将向自身写入的数据同步写入到所有 writers 中。
func MultiWriter(writers ...Writer) Writer

// TeeReader 对 r 进行封装，使 r 在读取数据的同时，自动向 w 中写入数据。
// 它是一个无缓冲的 Reader，所以对 w 的写入操作必须在 r 的 Read 操作结束
// 之前完成。所有写入时遇到的错误都会被作为 Read 方法的 err 返回。
func TeeReader(r Reader, w Writer) Reader

------------------------------

// 示例 LimitReader
func main() {
	r := strings.NewReader("Hello World!")
	lr := io.LimitReader(r, 5)

	n, err := io.Copy(os.Stdout, lr)  // Hello
	fmt.Printf("\n%d   %v\n", n, err) // 5   <nil>
}

------------------------------

// 示例 MultiReader
func main() {
	r1 := strings.NewReader("Hello World!")
	r2 := strings.NewReader("ABCDEFG")
	r3 := strings.NewReader("abcdefg")
	b := make([]byte, 15)
	mr := io.MultiReader(r1, r2, r3)

	for n, err := 0, error(nil); err == nil; {
		n, err = mr.Read(b)
		fmt.Printf("%q\n", b[:n])
	}
	// "Hello World!"
	// "ABCDEFG"
	// "abcdefg"
	// ""

	r1.Seek(0, 0)
	r2.Seek(0, 0)
	r3.Seek(0, 0)
	mr = io.MultiReader(r1, r2, r3)
	io.Copy(os.Stdout, mr)
	// Hello World!ABCDEFGabcdefg
}

------------------------------

// 示例 MultiWriter
func main() {
	r := strings.NewReader("Hello World!\n")
	mw := io.MultiWriter(os.Stdout, os.Stdout, os.Stdout)

	r.WriteTo(mw)
	// Hello World!
	// Hello World!
	// Hello World!
}

// 示例 TeeReader
func main() {
	r := strings.NewReader("Hello World!")
	b := make([]byte, 15)
	tr := io.TeeReader(r, os.Stdout)

	n, err := tr.Read(b)                  // Hello World!
	fmt.Printf("\n%s   %v\n", b[:n], err) // Hello World!   <nil>
}

------------------------------------------------------------

// NewSectionReader 对 r 进行封装，使其只能从 off 位置开始读取，最多只能读取 n
// 个字节的的数据。相当于对 r 做了一个切片 r[off:off+n] 返回。
// 底层实现是一个 *SectionReader。
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

// SectionReader 实现了如下接口：
// io.Reader
// io.ReaderAt
// io.Seeker

// Size 返回允许读取部分的大小（即切片的长度 n）
func (s *SectionReader) Size()

------------------------------

// 示例 SectionReader
func main() {
	r := strings.NewReader("Hello World!")
	sr := io.NewSectionReader(r, 6, 5)

	n, err := io.Copy(os.Stdout, sr)                  // World
	fmt.Printf("\n%d   %d   %v\n", sr.Size(), n, err) // 5   5   <nil>
}

------------------------------------------------------------

// Pipe 在内存中创建一个同步管道，用于不同区域的代码之间相互传递数据。
// 返回的 *PipeReader 用于从管道中读取数据，*PipeWriter 用于向管道中写入数据。
// 管道没有缓冲区，读写操作可能会被阻塞。可以安全的对管道进行并行的读、写或关闭
// 操作，读写操作会依次执行，Close 会在被阻塞的 I/O 操作结束之后完成。
func Pipe() (*PipeReader, *PipeWriter)

// 从管道中读取数据，如果管道被关闭，则会返会一个错误信息：
// 1、如果写入端通过 CloseWithError 方法关闭了管道，则返回关闭时传入的错误信息。
// 2、如果写入端通过 Close 方法关闭了管道，则返回 io.EOF。
// 3、如果是读取端关闭了管道，则返回 io.ErrClosedPipe。
func (r *PipeReader) Read(data []byte) (n int, err error)

// 关闭管道
func (r *PipeReader) Close() error

// 关闭管道并传入错误信息。
func (r *PipeReader) CloseWithError(err error) error

// 向管道中写入数据，如果管道被关闭，则会返会一个错误信息：
// 1、如果读取端通过 CloseWithError 方法关闭了管道，则返回关闭时传入的错误信息。
// 2、如果读取端通过 Close 方法关闭了管道，则返回 io.ErrClosedPipe。
// 3、如果是写入端关闭了管道，则返回 io.ErrClosedPipe。
func (w *PipeWriter) Write(data []byte) (n int, err error)

// 关闭管道
func (w *PipeWriter) Close() error

// 关闭管道并传入错误信息。
func (w *PipeWriter) CloseWithError(err error) error

------------------------------

// 示例：管道（读取端关闭）
func main() {
	r, w := io.Pipe()
	// 启用一个例程进行读取
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			r.CloseWithError(errors.New("管道被读取端关闭"))
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])
		}
	}()
	// 主例程进行写入
	n, err := w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)
}

------------------------------

// 示例：管道（写入端关闭）
func main() {
	r, w := io.Pipe()
	// 启用一个例程进行读取
	go func() {
		buf := make([]byte, 5)
		for n, err := 0, error(nil); err == nil; {
			n, err = r.Read(buf)
			fmt.Printf("读取：%d, %v, %s\n", n, err, buf[:n])
		}
	}()
	// 主例程进行写入
	n, err := w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)

	w.CloseWithError(errors.New("管道被写入端关闭"))
	n, err = w.Write([]byte("Hello World !"))
	fmt.Printf("写入：%d, %v\n", n, err)
	time.Sleep(time.Second * 1)
}

------------------------------------------------------------
*/
