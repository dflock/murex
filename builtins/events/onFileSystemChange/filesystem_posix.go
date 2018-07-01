// +build !windows

package onFileSystemChange

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils/ansi"
)

const eventType = "onFileSystemChange"

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Path      string
	Operation string
}

func init() {
	evt := newWatch()
	events.AddEventType(eventType, evt)
	go evt.init()
}

type watch struct {
	watcher *fsnotify.Watcher
	error   error
	mutex   sync.Mutex
	paths   map[string]string // map of paths indexed by event name
	blocks  map[string][]rune // map of blocks indexed by path
}

func newWatch() (w *watch) {
	w = new(watch)
	w.watcher, w.error = fsnotify.NewWatcher()
	w.paths = make(map[string]string)
	w.blocks = make(map[string][]rune)

	return
}

// Callback returns the block to execute upon a triggered event
func (evt *watch) findCallbackBlock(path string) (block []rune) {
	evt.mutex.Lock()

	for {
		for len(path) > 1 && path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}

		block = evt.blocks[path]

		if len(block) > 0 {
			break
		}

		split := strings.Split(path, "/")
		switch len(split) {
		case 0:
			path = "/"
		case 1:
			path = strings.Join(split, "/")
		default:
			path = strings.Join(split[:len(split)-1], "/")
		}
		//debug.Log("path=" + path)
	}

	evt.mutex.Unlock()
	return
}

// Add a path to the watch event list
func (evt *watch) Add(name, path string, block []rune) error {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	pwd, err := os.Getwd()
	if err == nil && path[0] != '/' {
		path = pwd + "/" + path
	}

	err = evt.watcher.Add(path)
	if err == nil {
		evt.mutex.Lock()
		evt.paths[name] = path
		evt.blocks[path] = block
		evt.mutex.Unlock()
	}

	return err
}

// Remove a path to the watch event list
func (evt *watch) Remove(name string) error {
	path := evt.paths[name]
	if path == "" {
		return errors.New("No event found for this listener with the name `" + name + "`.")
	}

	err := evt.watcher.Remove(path)
	if err == nil {
		evt.mutex.Lock()
		delete(evt.paths, name)
		delete(evt.blocks, path)
		evt.mutex.Unlock()
	}

	return err
}

// Init starts a new watch event loop
func (evt *watch) init() {
	getName := func(path string) (name string) {
		var evtPath string
		evt.mutex.Lock()

		for name, evtPath = range evt.paths {
			if path == evtPath {
				evt.mutex.Unlock()
				return
			}
		}

		// code shouldn't hit this point anyway
		evt.mutex.Unlock()
		return
	}

	defer evt.watcher.Close()

	for {
		select {
		case event := <-evt.watcher.Events:
			events.Callback(
				getName(event.Name),
				Interrupt{
					Path:      event.Name,
					Operation: event.Op.String(),
				},
				evt.findCallbackBlock(event.Name),
				proc.ShellProcess.Stdout,
			)

		case err := <-evt.watcher.Errors:
			ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "error in watcher: "+err.Error())
		}
	}
}

// Dump returns all the events in fsWatch
func (evt *watch) Dump() interface{} {
	type jsonable struct {
		Path  string
		Block string
	}

	dump := make(map[string]jsonable)

	evt.mutex.Lock()

	for name, path := range evt.paths {
		dump[name] = jsonable{
			Path:  path,
			Block: string(evt.blocks[path]),
		}
	}

	evt.mutex.Unlock()

	return dump
}