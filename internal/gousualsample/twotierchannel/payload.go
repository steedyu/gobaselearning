package twotierchannel

import "time"

type Payload struct {
	// [redacted]
}


func (p *Payload) UploadToS3() error {
	//fmt.Println("UploadToS3")
	time.Sleep(time.Second * 1)
	return nil
}
