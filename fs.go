//go:build tinygo

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/bgould/keyboard-firmware/keyboard/console"
)

func configureFilesystem() {
	if err := filesystem.Mount(); err != nil {
		println("Could not mount LittleFS filesystem: ", err.Error(), "\r\n")
	} else {
		println("Successfully mounted LittleFS filesystem.\r\n")
		fs_mounted = true

		if info, err := filesystem.Stat("saved.keymap"); err != nil {
			println("unable to load saved.keymap: ", err)
		} else {
			println("Attempting to load keymap file: ", info.Name())
			loadKeymap(info.Name())
		}
	}
}

func init() {

	// Keymap commands
	commands["save"] = console.CommandHandlerFunc(save)
	commands["load"] = console.CommandHandlerFunc(load)

	// Filesystem Commands
	commands["mount"] = console.CommandHandlerFunc(mount)
	commands["umount"] = console.CommandHandlerFunc(umount)
	commands["format"] = console.CommandHandlerFunc(format)
	commands["ls"] = console.CommandHandlerFunc(ls)
	commands["mv"] = console.CommandHandlerFunc(mv)
	commands["mkdir"] = console.CommandHandlerFunc(mkdir)
	commands["create"] = console.CommandHandlerFunc(create)
	commands["rm"] = console.CommandHandlerFunc(rm)
	commands["cat"] = console.CommandHandlerFunc(cat)
	// "samples": console.CommandHandlerFunc(samples),
	// "xxd":    console.CommandHandlerFunc(xxd),

}

const (
	savedKeymapFilename = "saved.keymap"
)

type FsErr uint8

const (
	ErrNotAFile FsErr = iota + 1
)

func (err FsErr) Error() string {
	switch err {
	case ErrNotAFile:
		return "not a file"
	default:
		return "unknown"
	}
}

// saveKeymap write the current in-memory keymap to the filesystem
func saveKeymap(filename string) (n int64, err error) {
	f, err := filesystem.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err = keymap.WriteTo(f)
	return n, err
}

// loadKeymap updates the current in-memory keymap from the filesystem
func loadKeymap(filename string) (n int64, err error) {
	f, err := filesystem.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	if f.IsDir() {
		return 0, ErrNotAFile
	}
	n, err = keymap.ReadFrom(f)
	return
}

// ########################### Commands ######################################/

func mount(cmd console.CommandInfo) int {
	if err := filesystem.Mount(); err != nil {
		println("Could not mount LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully mounted LittleFS filesystem.\r\n")
		return 0
	}
}

func format(cmd console.CommandInfo) int {
	if err := filesystem.Format(); err != nil {
		println("Could not format LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully formatted LittleFS filesystem.\r\n")
		return 0
	}
}

func umount(cmd console.CommandInfo) int {
	if err := filesystem.Unmount(); err != nil {
		println("Could not unmount LittleFS filesystem: ", err.Error(), "\r\n")
		return 1
	} else {
		println("Successfully unmounted LittleFS filesystem.\r\n")
		return 0
	}
}

func ls(cmd console.CommandInfo) int {
	path := "/"
	argv := cmd.Argv
	if len(argv) > 0 {
		path = strings.TrimSpace(argv[0])
	}
	dir, err := filesystem.Open(path)
	if err != nil {
		println("Could not open directory", path, ":", err.Error())
		return 1
	}
	defer dir.Close()
	infos, err := dir.Readdir(0)
	_ = infos
	if err != nil {
		println("Could not read directory", path, ":", err.Error())
		return 1
	}
	for _, info := range infos {
		s := "-rwxrwxrwx"
		if info.IsDir() {
			s = "drwxrwxrwx"
		}
		println(s, info.Size(), info.Name())
		//fmt.Printf("%s %5d %s\n", s, info.Size(), info.Name())
	}
	return 0
}

func mkdir(cmd console.CommandInfo) int {
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if debug {
		println("Trying mkdir to " + tgt)
	}
	if tgt == "" {
		println("Usage: mkdir <target dir>")
		return 1
	}
	err := filesystem.Mkdir(tgt, 0777)
	if err != nil {
		println("Could not mkdir " + tgt + ": " + err.Error())
	}
	return 0
}

func rm(cmd console.CommandInfo) int {
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if debug {
		println("Trying to rm ", tgt)
	}
	if tgt == "" {
		println("Usage: rm <target dir>")
		return 1
	}
	err := filesystem.Remove(tgt)
	if err != nil {
		println("Could not rm ", tgt, ":", err.Error())
		return 1
	}
	return 0
}

func mv(cmd console.CommandInfo) int {
	src, dst := "", ""
	argv := cmd.Argv
	if len(argv) == 2 {
		src = strings.TrimSpace(argv[0])
		dst = strings.TrimSpace(argv[1])
	}
	if debug {
		println("Trying mv from", src, "to", dst)
	}
	if src == "" || dst == "" {
		println("Usage: mv <srcfile> <destfile>")
		return 1
	}
	err := filesystem.Rename(src, dst)
	if err != nil {
		println("Could not mv ", src, "to", dst, ":", err.Error())
		return 1
	}
	return 0
}

// func samples(cmd console.CommandInfo) int {
// 	buf := make([]byte, 90)
// 	for i := 0; i < 5; i++ {
// 		name := fmt.Sprintf("file%d.txt", i)
// 		if bytes, err := createSampleFile(name, buf); err != nil {
// 			fmt.Printf("%s\r\n", err)
// 			return 1
// 		} else {
// 			fmt.Printf("wrote %d bytes to %s\r\n", bytes, name)
// 		}
// 	}
// 	return 0
// }

func create(cmd console.CommandInfo) int {
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 2 {
		tgt = strings.TrimSpace(argv[1])
	}
	if debug {
		println("Trying create to " + tgt)
	}
	buf := make([]byte, 90)
	if bytes, err := createSampleFile(tgt, buf); err != nil {
		fmt.Printf("%s\r\n", err)
		return 1
	} else {
		println("wrote", bytes, "bytes to", tgt)
		// fmt.Printf("wrote %d bytes to %s\r\n", bytes, tgt)
		return 0
	}
}

// func write(argv []string) {
// 	tgt := ""
// 	if len(argv) == 2 {
// 		tgt = strings.TrimSpace(argv[1])
// 	}
// 	if debug {
// 		println("Trying receive to " + tgt)
// 	}
// 	buf := make([]byte, 1)
// 	f, err := fs.OpenFile(tgt, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
// 	if err != nil {
// 		fmt.Printf("error opening %s: %s\r\n", tgt, err.Error())
// 		return
// 	}
// 	defer f.Close()
// 	var n int
// 	for {
// 		if console.Buffered() > 0 {
// 			data, _ := console.ReadByte()
// 			switch data {
// 			case 0x04:
// 				fmt.Printf("wrote %d bytes to %s\r\n", n, tgt)
// 				return
// 			default:
// 				// anything else, just echo the character if it is printable
// 				if strconv.IsPrint(rune(data)) {
// 					console.WriteByte(data)
// 				}
// 				buf[0] = data
// 				if _, err := f.Write(buf); err != nil {
// 					fmt.Printf("\nerror writing: %s\r\n", err)
// 					return
// 				}
// 				n++
// 			}
// 		}
// 	}
// }

func createSampleFile(name string, buf []byte) (int, error) {
	for j := uint8(0); j < uint8(len(buf)); j++ {
		buf[j] = 0x20 + j
	}
	f, err := filesystem.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		// return 0, fmt.Errorf("error opening %s: %s", name, err.Error())
		return 0, err
	}
	defer f.Close()
	bytes, err := f.Write(buf)
	return bytes, err
}

func cat(cmd console.CommandInfo) int {
	tgt := ""
	argv := cmd.Argv
	if len(argv) == 1 {
		tgt = strings.TrimSpace(argv[0])
	}
	if debug {
		println("Trying to cat to " + tgt)
	}
	if tgt == "" {
		println("Usage: cat <target dir>")
		return 1
	}
	if debug {
		println("Getting entry")
	}
	f, err := filesystem.Open(tgt)
	if err != nil {
		println("Could not open: " + err.Error())
		return 1
	}
	defer f.Close()
	if f.IsDir() {
		println("Not a file: " + tgt)
		return 1
	}
	off := 0x0
	buf := make([]byte, 64)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			println("Error reading " + tgt + ": " + err.Error())
		}
		xxdfprint(os.Stdout, uint32(off), buf[:n])
		off += n
	}
	return 0
}

// func xxd(cmd console.CommandInfo) int {
// 	var err error
// 	var addr uint64 = 0x0
// 	var size int = 64
// 	argv := cmd.Argv
// 	switch len(argv) {
// 	case 3:
// 		if size, err = strconv.Atoi(argv[2]); err != nil {
// 			println("Invalid size argument: " + err.Error() + "\r\n")
// 			return 1
// 		}
// 		if size > 512 || size < 1 {
// 			fmt.Printf("Size of hexdump must be greater than 0 and less than %d\r\n", 512)
// 			return 1
// 		}
// 		fallthrough
// 	case 2:
// 		/*
// 			if argv[1][:2] != "0x" {
// 				println("Invalid hex address (should start with 0x)")
// 				return
// 			}
// 		*/
// 		if addr, err = strconv.ParseUint(argv[1], 16, 32); err != nil {
// 			println("Invalid address: " + err.Error() + "\r\n")
// 			return 1
// 		}
// 		fallthrough
// 	case 1:
// 		// no args supplied, so nothing to do here, just use the defaults
// 	default:
// 		println("usage: xxd <hex address, ex: 0xA0> <size of hexdump in bytes>\r\n")
// 		return 1
// 	}
// 	buf := make([]byte, size)
// 	//bsz := uint64(flash.SectorSize)
// 	//blockdev.ReadBlock(uint32(addr/bsz), uint32(addr%bsz), buf)
// 	blockdev.ReadAt(buf, int64(addr))
// 	xxdfprint(os.Stdout, uint32(addr), buf)
// 	return 0
// }

func xxdfprint(w io.Writer, offset uint32, b []byte) {
	var l int
	var addr = make([]byte, 16)
	var data = make([]byte, 48)
	var buf16 = make([]byte, 16)
	for i, c := 0, len(b); i < c; i += 16 {
		a := offset + uint32(i)
		bin2hex([]byte{byte(a >> 24), byte(a >> 16), byte(a >> 8), byte(a)}, addr)
		w.Write(addr)
		w.Write([]byte(": "))
		l = i + 16
		if l >= c {
			l = c
		}
		for j, n := 0, l-i; j < 16; j++ {
			data[j*3] = ' '
			data[j*3+1] = ' '
			data[j*3+2] = ' '
			if j >= n {
				buf16[j] = ' '
			} else {
				var buf [2]byte
				bin2hex([]byte{byte(b[i+j])}, buf[:])
				data[j*3+1] = buf[0]
				data[j*3+2] = buf[1]
				if !strconv.IsPrint(rune(b[i+j])) {
					buf16[j] = '.'
				} else {
					buf16[j] = b[i+j]
				}
			}
		}
		w.Write(data)
		w.Write([]byte("    "))
		w.Write(buf16)
		w.Write([]byte("\r\n"))
	}
}
