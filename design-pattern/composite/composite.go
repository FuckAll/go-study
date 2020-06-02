package composite

import "os"

/*
	组合模式：将一组对象组织成树形结构，表示一种”整体-部分“的关系，让客户端（用户）能够统一单个对象和组合对象的处理逻辑。
	实现一个文件目录结构，能够增加、删除、统计文件数，文件大小等
*/

type FileSystemNode interface {
	getPath() string
	countNumberOfFiles() (int64, error)
	countSizeOfFiles() (int64, error)
}

type File struct {
	path string
}

func (f *File) File(path string) {
	f.path = path
}

func (f *File) getPath() string {
	return f.path
}

func (f *File) countNumberOfFiles() (int64, error) {
	return 1, nil
}

func (f *File) countSizeOfFiles() (int64, error) {
	fileInfo, err := os.Stat(f.path)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

type Directory struct {
	path    string
	subNode []FileSystemNode
}

func (d *Directory) Directory(path string) {
	d.path = path
}

func (d *Directory) addSubNode(node FileSystemNode) {
	d.subNode = append(d.subNode, node)
}

func (d *Directory) removeSubNode(node FileSystemNode) {
	length := len(d.subNode)
	removeIndex := -1
	for i := 0; i < length; i++ {
		if d.subNode[i] == node {
			removeIndex = i
			break
		}
	}
	if removeIndex > -1 {
		d.subNode = append(d.subNode[0:removeIndex], d.subNode[removeIndex+1:]...)
	}
}

func (d *Directory) getPath() string {
	return d.path
}

func (d *Directory) countNumberOfFiles() (int64, error) {
	var sum int64
	for _, sub := range d.subNode {
		x, err := sub.countNumberOfFiles()
		if err != nil {
			return 0, err
		}
		sum += x
	}
	return sum, nil
}

func (d *Directory) countSizeOfFiles() (int64, error) {
	var sum int64
	for _, sub := range d.subNode {
		x, err := sub.countSizeOfFiles()
		if err != nil {
			return 0, err
		}
		sum += x
	}
	return sum, nil
}
