package main

import (
	"context"
	"os"

	"github.com/bachue/qiniu-go-sdk/qiniupkg.com/api.v8/kodocli"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	FlashDownload  bool     `short:"f"`
	OutputFilePath string   `short:"o" long:"output"`
	IoHosts        []string `short:"h" long:"hosts"`
	Bucket         string   `short:"b" long:"bucket"`
	Uid            uint64   `short:"u" long:"uid"`
	Key            string   `short:"k" long:"key"`
	PartSize       uint64   `short:"p" long:"part"`
	Concurrency    int      `short:"c" long:"concurrency" default:"4"`
	Tries          uint64   `short:"t" long:"tries" default:"5"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	downloader := kodocli.NewDownloader(-1, &kodocli.DownloadConfig{
		IoHosts: opts.IoHosts,
		Tries:   opts.Tries,
	})
	outputFile, err := os.OpenFile(opts.OutputFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	if opts.FlashDownload {
		err = downloader.FlashDownloadFile(context.Background(), opts.Uid, opts.Bucket, opts.Key, outputFile, opts.PartSize, opts.Concurrency)
	} else {
		err = downloader.DownloadFile(context.Background(), opts.Uid, opts.Bucket, opts.Key, outputFile)
	}

	if err != nil {
		panic(err)
	}
}
