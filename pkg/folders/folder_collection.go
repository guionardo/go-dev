package folders

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/guionardo/go-dev/pkg/io"
	"github.com/guionardo/go-dev/pkg/logger"
)

type FolderCollection struct {
	Root     string             `yaml:"root"`
	Folders  map[string]*Folder `yaml:"folders"`
	MaxDepth int                `yaml:"maxSubLevel"`
}

func (fc *FolderCollection) FixPathsLoad() {
	for p, f := range fc.Folders {
		f.Path = p
	}
}

func CreateCollection(root string, maxSubLevel int) *FolderCollection {
	if maxSubLevel < 1 {
		maxSubLevel = 3
	}
	return &FolderCollection{
		Root:     root,
		Folders:  make(map[string]*Folder, 0),
		MaxDepth: maxSubLevel,
	}
}

func (fc *FolderCollection) Sync() error {
	existingFolders, err := io.ReadFolders(fc.Root, fc.MaxDepth)
	if err != nil {
		return err
	}
	for _, f := range existingFolders {
		_, err := fc.Get(f)
		if err != nil {
			fc.Folders[f] = &Folder{Path: f}
			logger.Info("Added new folder %s", f)
		}
	}
	var removed = make([]*Folder, 0, len(fc.Folders))
	for _, f := range fc.Folders {
		if stat, err := os.Stat(f.Path); err != nil || !stat.IsDir() {
			logger.Info("Removing missing folder %s", f.Path)
			removed = append(removed, f)
		}
	}
	if len(removed) == 0 {
		return nil
	}
	newFolders := make(map[string]*Folder, len(fc.Folders)-len(removed))

	for _, f := range fc.Folders {
		for _, r := range removed {
			if f.Path == r.Path {
				continue
			}
		}

		newFolders[f.Path] = f
	}
	fc.Folders = newFolders

	return nil
}

func (fc *FolderCollection) FixIgnored() {
	for _, f := range fc.Folders {
		children := fc.GetChildren(f)
		if len(children) == 0 {
			continue
		}
		for _, c := range children {
			if f.IgnoreSubFolders {
				c.IgnoreSubFolders = true
				c.Ignore = true
			}
		}
	}

}

func (fc *FolderCollection) Get(path string) (*Folder, error) {
	for _, f := range fc.Folders {
		if f.Path == path {
			return f, nil
		}
	}
	return nil, errors.New("Folder not found")
}

func (fc *FolderCollection) GetParent(folder *Folder) (*Folder, error) {
	return fc.GetNearestParent(folder.Path, false)
}

func (fc *FolderCollection) GetNearestParent(folder string, keepFinding bool) (*Folder, error) {
	parent := path.Dir(folder)
	for _, f := range fc.Folders {
		if f.Path == parent {
			return f, nil
		}
	}
	if keepFinding {
		return fc.GetNearestParent(parent, keepFinding)
	}
	return nil, errors.New("Folder not found")
}

func (fc *FolderCollection) GetChildren(folder *Folder) []*Folder {
	children := make([]*Folder, 0)
	for _, f := range fc.Folders {
		if f.Path != folder.Path && strings.HasPrefix(f.Path, folder.Path) {
			children = append(children, f)
		}
	}
	return children
}

func (fc *FolderCollection) Find(words []string) (folders []*Folder) {
	folders = make([]*Folder, 0)
	for _, f := range fc.Folders {
		if !f.Ignore && f.Match(words) {
			folders = append(folders, f)
		}
	}
	return
}
