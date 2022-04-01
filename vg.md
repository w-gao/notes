# vg

Date: Mar 31, 2022 - xxx


vg (https://github.com/vgteam/vg) is a collection of tools for working with
genome variation graphs.

## Overview

A variation graph consists of a collection of nodes, edges, and paths:

```
 variation graph (vg)
┌────────────────────────────────┐
│ ┌───────┐ ┌───────┐ ┌───────┐  │
│ │ nodes │ │ edges │ │ paths │  │
│ └───────┘ └───────┘ └───────┘  │
└────────────────────────────────┘
```

- nodes "store sequence data"
- edges "describe bidirected linkages between nodes"
- paths are "walks through nodes defined by a series of `edit`s"
  - can represent haplotypes, etc


<!--
CPANG19 tutorial files

Material: https://gtpb.github.io/CPANG19/
Date: Mar 31, 2022 - xxx
-->

## Using vg (via docker)

I've built vg from source a while back, but it took a long time on my poor
machine. Now that the executable is a few versions behind, I decided to run vg
through its docker image for this, so I don't have to rebuild.

I pulled the latest (as of Mar 31, 2022) vg using:

```
$ docker pull quay.io/vgteam/vg:v1.39.0
```


Then ran a vg container as such:

```
$ docker run -d -it --name vgdocker -v ${PWD}:/vg/shared-files quay.io/vgteam/vg:v1.39.0
```

Notes:
- use `-d` to run in detach mode
- the `-it` options (`i`nterative, and `t`ty) are needed to keep the container
  running in the background
- use `--name` to provide a name for easier reference later
- set up a volume so that files can be easily shared between host and container


To stop the vg container, run:

```
$ docker stop vgdocker
```

(or use the container id returned from `docker run`)


If running in detach mode, we can run commands using:

```
$ docker exec -it vgdocker vg
```

Most of the time I will not be running in detach mode; I will be running the vg
commands interactively instead.



## Toy example

Agenda
- `vg construct` - construct vg graph
- `vg view` - visualize vg graph


### Construct a vg graph from FASTA and VCF

The FASTA file(s) serves as the reference and the VCF file(s) defines
variations of the sequence. To construct a vg graph, run:

```
$ vg construct -r tiny/tiny.fa -v tiny/tiny.vcf.gz -m 32 >tiny.vg
```

Notes
- `-r`: FASTA reference
- `-v`: VCF variants
- `-m`: max size for each node


### Visualize vg graphs

To visualize a vg graph, use the `vg view` command:

#### `-g`: GFA format (default)

```
$ vg view tiny.vg
H	VN:Z:1.0
S	5	C
S	7	A
S	12	ATAT
S	8	G
S	1	CAAATAAG
S	4	T
S	6	TTG
S	15	CCAACTCTCTG
S	2	A
S	10	A
S	9	AAATTTTCTGGAGTTCTAT
S	11	T
S	13	A
S	14	T
S	3	G
P	x	1+,3+,5+,6+,8+,9+,11+,12+,14+,15+	*
L	5	+	6	+	*
L	7	+	9	+	*
L	12	+	13	+	*
...
```

#### `-j`: JSON format

```
$ vg view -j tiny.vg | jq '.node[].sequence'
"C"
"A"
"ATAT"
"G"
"CAAATAAG"
"T"
"TTG"
"CCAACTCTCTG"
"A"
"A"
"AAATTTTCTGGAGTTCTAT"
"T"
"A"
"T"
"G"
```
	
Note	
- only node sequences are displayed here


#### `-d`: DOT format

```
$ vg view -dp tiny.vg | dot -Tsvg -o tiny.svg
```

![vg-view-dot](https://user-images.githubusercontent.com/20177171/161162117-493a4808-78e3-4b21-8d15-619aafc83f09.svg)

Notes
- The dot file can then be rendered into vector graphics (e.g.: PDFs or SVGs)
- `-p`: include the paths (this is the bottom path(s) that display the original
  FASTA sequences in the SVG)

## 1Mbp of 1000 Genomes data

Agenda
- `vg index`



## Resources

- overview
  - first paper on vg: https://dx.doi.org/10.1038/nbt.4227
  - great slides [\*link\*](https://gtpb.github.io/CPANG18/assets/CPANG18_Computational_Pangenomics_2018_(day1).pdf)
- setting up vg
  - wiki on docker: https://github.com/vgteam/vg/wiki/building-vg-(or-not-building-vg)#docker



<!--
* From JSON to vg

```
vg view -v -J graph.vg.json > graph.vg
```

* Generate xg index from vg

```
vg index -x graph.xg graph.vg
```
-->

