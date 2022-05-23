package gofn

type Block struct {
	run     func()
	catch   func(err error)
	finally func()
}

func Try(fn func()) *Block {
	return &Block{run: fn}
}

func (b *Block) Catch(fn func(err error)) *Block {
	b.catch = fn
	return b
}

func (b *Block) Finally(fn func()) *Block {
	b.finally = fn
	return b
}

func (b *Block) Run() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if b.catch != nil {
				er := e.(error)
				err = er
				b.catch(er)
				if b.finally != nil {
					b.finally()
				}
			} else {
				if b.finally != nil {
					b.finally()
				}
			}
		}
	}()

	b.run()
	return nil
}
