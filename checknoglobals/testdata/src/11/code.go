package code

import (
	"embed"
)

//go:embed embedfiles/*
var fileSystem embed.FS

//go:embed embedfiles/embedfile.txt
var data []byte

//go:embed embedfiles/*
//
var fileSystemExtraCommentLines embed.FS

//go:embed embedfiles/*

var fileSystemEmptyLines embed.FS

var fileSystemEmptyLinesNoComment embed.FS // want "fileSystemEmptyLinesNoComment is a global variable"

// go : embed that does not match
var fileSystemEmptyLinesOtherComment embed.FS // want "fileSystemEmptyLinesOtherComment is a global variable"

//go:embed embedfiles/*
//

var fileSystemExtraCommentLinesAndEmptyLines embed.FS

var (
	//go:embed embedfiles/*
	groupedFileSystem embed.FS

	//go:embed embedfiles/embedfile.txt
	groupedData []byte

	//go:embed embedfiles/*

	groupedFileSystemEmptyLines embed.FS

	groupedFileSystemEmptyLinesNoComment embed.FS // want "groupedFileSystemEmptyLinesNoComment is a global variable"

	// go : embed that does not match
	groupedFileSystemEmptyLinesOtherComment embed.FS // want "groupedFileSystemEmptyLinesOtherComment is a global variable"

	//go:embed embedfiles/*
	//

	groupedFileSystemExtraCommentLinesAndEmptyLines embed.FS
)
