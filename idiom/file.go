package idiom

import (
	"fmt"
	"os"
)

type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Option func(*Options)

func UID(userID int) Option {
	return func(o *Options) {
		o.UID = userID
	}
}

func GID(groupID int) Option {
	return func(o *Options) {
		o.GID = groupID
	}
}

func Contents(c string) Option {
	return func(o *Options) {
		o.Contents = c
	}
}

func Permissions(perms os.FileMode) Option {
	return func(o *Options) {
		o.Permissions = perms
	}
}

// Constructor
func NewFile(filepath string, setters ...Option) error {
	// Default Options
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getuid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("file already exists: %s", filepath)
		}
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}

	return f.Chown(args.UID, args.GID)
}
