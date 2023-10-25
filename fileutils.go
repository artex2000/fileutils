package fileutils

import (
                "os"
                "strings"
                "path/filepath"
       )

func ReadTextFile(FileName string) ([]string, error) {
        data, err := os.ReadFile(FileName)
        if err != nil {
                return nil, err
        }

        lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

        //Get rid of last line if there is nothing there
        tmp := lines[len(lines)-1]
        if len(tmp) == 0 {
                lines = lines[0:len(lines)-1]
        }
        return lines, nil
}

func ReadDirectory(DirName string) ([]os.FileInfo, error) {
        f, err := os.Open(DirName)
        if err != nil {
                return nil, err
        }
        defer f.Close()

        files, err := f.Readdir(-1)
        if err != nil {
                return nil, err
        }
        return files, nil 
}

func GetAllFiles(RootDir string, Recursive, StopOnError bool) ([]string, error) {
        var dirs  []string
        var outs  []string

        files, err := ReadDirectory(RootDir)
        if err != nil {
                return nil, err
        }

        for _, f := range (files) {
                n := f.Name()
                n = filepath.Join(RootDir, n)
                if f.IsDir() {
                        dirs = append (dirs, n)
                } else {
                        outs = append (outs, n)
                }
        }

        if !Recursive {
                return outs, nil
        }

        //Now do the same for subdirectories
        for {
                var next string
                if l := len (dirs); l > 0 {
                        next = dirs[0]
                        dirs = dirs[1:]
                } else {
                        break;
                }

                files, err = ReadDirectory(next)
                if err != nil {
                        if StopOnError {
                                return nil, err
                        } else {
                                continue
                        }
                }

                for _, f := range (files) {
                        n := f.Name()
                        n = filepath.Join(next, n)
                        if f.IsDir() {
                                dirs = append (dirs, n)
                        } else {
                                outs = append (outs, n)
                        }
                }
        }

        return outs, nil
}

func GetAllDirs(RootDir string, Recursive, StopOnError bool) ([]string, error) {
        var dirs  []string
        var outs  []string

        files, err := ReadDirectory(RootDir)
        if err != nil {
                return nil, err
        }

        for _, f := range (files) {
                n := f.Name()
                n = filepath.Join(RootDir, n)
                if f.IsDir() {
                        dirs = append (dirs, n)
                        outs = append (outs, n)
                }
        }

        if !Recursive {
                return outs, nil
        }

        //Now do the same for subdirectories
        for {
                var next string
                if l := len (dirs); l > 0 {
                        next = dirs[0]
                        dirs = dirs[1:]
                } else {
                        break;
                }

                files, err = ReadDirectory(next)
                if err != nil {
                        if StopOnError {
                                return nil, err
                        } else {
                                continue
                        }
                }

                for _, f := range (files) {
                        n := f.Name()
                        n = filepath.Join(next, n)
                        if f.IsDir() {
                                dirs = append (dirs, n)
                                outs = append (outs, n)
                        }
                }
        }

        return outs, nil
}

func GetFilesByExt(RootDir, Extension string, Recursive, StopOnError bool) ([]string, error) {
        var dirs  []string
        var outs  []string

        //Add dot to the extension to filter to be in line with filepath.Ext return values
        var filter string
	if Extension != "" {
		filter = "." + Extension
	} else {
                return GetAllFiles(RootDir, Recursive, StopOnError)
        }

        files, err := ReadDirectory(RootDir)
        if err != nil {
                return nil, err
        }

        for _, f := range (files) {
                n := f.Name()
                if f.IsDir() {
                        n = filepath.Join(RootDir, n)
                        dirs = append (dirs, n)
                } else if filter == filepath.Ext(n) {
                        n = filepath.Join(RootDir, n)
                        outs = append (outs, n)
                }
        }

        if !Recursive {
                return outs, nil
        }

        //Now do the same for subdirectories
        for {
                var next string
                if l := len (dirs); l > 0 {
                        next = dirs[0]
                        dirs = dirs[1:]
                } else {
                        break;
                }

                files, err = ReadDirectory(next)
                if err != nil {
                        if StopOnError {
                                return nil, err
                        } else {
                                continue
                        }
                }

                for _, f := range (files) {
                        n := f.Name()
                        if f.IsDir() {
                                n = filepath.Join(next, n)
                                dirs = append (dirs, n)
                        } else if filter == filepath.Ext(n) {
                                n = filepath.Join(next, n)
                                outs = append (outs, n)
                        }
                }
        }

        return outs, nil
}
