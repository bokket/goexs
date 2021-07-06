package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
	"log"
)

func main()  {
	client,_:=hdfs.New("127.0.0.1:9000")
	file,_:=client.Open("/test/6.cc")

	fsInfo,_:=client.Stat("/test/6.cc")



	buf:=make([]byte,10240)
	file.ReadAt(buf,0)

	fsinfo,_:=client.StatFs()




	fmt.Println(string(buf))

	//buffer:=make([]byte,256)
	checksum,_:=file.Checksum()

	log.Println("hdfs file name",file.Name(),"file checksum",checksum,"file stat",file.Stat(),"file used",fsinfo.Used)

	log.Println("hdfs Fsinfo",fsInfo.Name(),"file isDir",fsInfo.IsDir(),"file mode",fsInfo.Mode(),
		"file modTime",fsInfo.ModTime(),"file size",fsInfo.Size(),"file sys",fsInfo.Sys())

	content,_:=client.GetContentSummary("/test/6.cc")

	log.Println("content summary ",content.Size(),"content directory",content.DirectoryCount(),
		"content filecount",content.FileCount(),"content namequota",content.NameQuota(),
		"content size afterReplication",content.SizeAfterReplication(),
		"content spacequota",content.SpaceQuota())

	err:=client.CopyToLocal("/test/6.cc","/home/bokket/下载/6.cc")
	if err!=nil {
		log.Println("copy to local error",err)
	}

	er:=client.CopyToRemote("/home/bokket/下载/6.cc","/test1")
	if er!=nil {
		log.Println("copy to remote error",er)
	}
}
