package meta

// 文件元信息
type FileMeta struct {
	FileSha1     string
	FileName     string
	FileSize     int64
	FileLocation string
	UploadAt     string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// 新增、更新文件元信息
func UpdateFileMeta(meta FileMeta)  {
	fileMetas[meta.FileSha1] = meta
}

// 根据sha1值获取文件的元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}