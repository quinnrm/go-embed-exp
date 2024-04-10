# go-embed-exp

Experiements with the Go embed package.

The `go-embed-exp` program is a CLI program written to test embedding files in various scenarios.

I needed to figure this out for work, and I'm going to make the experiemnts public.

## Dependencies

  - Go v1.21.8+
  - Fedora 40+

It was written on and tested on Fedora, but it should work on other operating systems or distros.

## Overview

The main objective is to figure out how to use the embed filesystem to store files. Files can be stored as strings or byte slices, but what I really need is a virtual filesystem to read files from.
