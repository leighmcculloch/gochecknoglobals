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

	//go:embed embedfiles/*
	//

	groupedFileSystemExtraCommentLinesAndEmptyLines embed.FS
)
