package datafile1

import (
	"os"
	"sync"
	"errors"
	"io"
)

// 数据的类型
type Data []byte

// 数据文件的接口类型。
type DataFile interface {
	// 读取一个数据块。
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块。
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号。
	Rsn() int64
	// 获取最后写入的数据块的序列号。
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

/*
1）
写解锁在进行的时候会试图唤醒所有因欲进行读锁定而被阻塞的Goroutine。
而读解锁在进行的时候智慧在已无任何读锁定的情况下试图唤醒一个因欲进行写锁定而被阻塞的Goroutine
2）
若对一个未被写锁定的读写锁进行写解锁，就会引发一个运行时恐慌，而对一个未被读锁定的读写锁进行读解锁却不会如此
 */

// 数据文件的实现类型。
type myDataFile struct {
	f       *os.File     // 文件。
	fmutex  sync.RWMutex // 被用于文件的读写锁。
	woffset int64        // 写操作需要用到的偏移量。
	roffset int64        // 读操作需要用到的偏移量。
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁。
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁。
	dataLen uint32       // 数据块长度。
}


func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	/*
	这里没有RLock之后，立马一个defer RUnlock是考虑到如下case:
	3个Groutine来并发执行Read方法，并有2个Goroutine来并发执行Write方法
	由于进行写操作的Goroutine比进行读操作的Goroutine少，所以过不了多久读偏移量roffset的值就会等于甚至大于写偏移量woffset的值
	这种情况会使上面的df.f.ReadAt方法返回的第二个结果值代表错误的非nil且会与io.EOF想等的值
	这种情况不应该看成错误，而应该把它看成边界情况，故如下处理
	 */
	//读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

func (df *myDataFile) Rsn() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) Wsn() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}


