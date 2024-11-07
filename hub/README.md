# hub package
Downloads HuggingFace Hub files, a port of huggingFace_hub python library to Go. 

## Introduction

A simple, straight-forward port of [github.com/huggingface/huggingface_hub](https://github.com/huggingface/huggingface_hub) library for Go.

Features supported:

- Cache system that matches HuggingFace Hub, so the same cache can be shared with Python.
- Concurrency safe: only one download when multiple workers are trying to download simultaneously the same model.
- Allow arbitrary progress function to be called (for progress bar).
- Arbitrary revision.

TODOs:

- Add support for optional parameters.
- Authentication tokens: should be relatively easy.
- Resume downloads from interrupted connections.
- Check disk-space before starting to download.
