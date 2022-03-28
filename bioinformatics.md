# Bioinformatics


## File formats


High-level overview

```
# ┌───────┐ align to  ┌───────────┐  variant  ┌─────┐
# │ FASTQ ├──────────►│ SAM / BAM ├──────────►│ VCF │
# └───────┘ reference └───────────┘  calling  └─────┘
#  sequence             alignments            variants
```
<!-- https://tinyurl.com/d373myk5 -->


### Sequence storage

- FASTA - `.fasta`, `.fa`
    * text-based format for storing nucleotide or peptide sequences
- FASTA index - `.fa.fai`
    * index file in conjuction with a FASTA file
    * allows "efficient access to arbitrary regions within those reference sequences" [\*link\*](http://manpages.ubuntu.com/manpages/bionic/man5/faidx.5.html)
- FASTQ - `.fastq`
    * text-based format for storing sequences + its corresponding quality scores
    * extension from FASTA


### Alignment/sequence

An aligner takes a number of reads (FASTQ file) and a reference sequence, and aligns the reads to the reference.


- Sequence alignment map (SAM) - `.sam`
    * tab-delimited, text-based format for storing sequence alignment information from an aligner
    * Contains same information as FASTQ, plus alignment information
    * FASTQ file --> align to a reference sequence using an aligner (e.g.: STAR) --> output SAM/BAM
- Binary alignment map (BAM) - `.bam`
    * binary format
    * Contains the exact same information as SAM except more compact
- SAM index `.sam.sai`
    * index file in conjunction with a SAM file
    * allows fast retrieval of data in a sorted SAM file
- BAM index `.bam.bai`
    * Same as SAM index for SAM, this is the index for a BAM file
- Useful links
    * SAM tools [\*link\*](https://en.wikipedia.org/wiki/SAMtools)
    * Difference between FASTA, FASTQ, and SAM/BAM [\*link\*](https://bioinformatics.stackexchange.com/a/385)


### Variants

- Variant call format (VCF) `.vcf`
    * tab-delimited, text-based format for storing gene sequence variations (against a reference)
    * variant calling identifies where the aligned reads (SAM/BAM) differ from the reference
    * Only variants (where the sequence differs from the ref) are stored; reads are not stored
- Useful links
    * Blog post about SAM/BAM and VCF [\*link\*](https://kaushikghose.wordpress.com/2014/03/26/sam-bam-vcf-what/)
    * Variant calling [\*link\*](https://www.ebi.ac.uk/training/online/courses/human-genetic-variation-introduction/variant-identification-and-analysis/)


### Genome variation graphs (vg)

- Graph alignment map (GAM) `.gam`
    * _


### UCSC Genome Browser

<!-- BME 110 Spring 2021 -->
- The UCSC genome browser works in a coordinate system; it needs an **alignment** instead of the raw DNA sequence

#### Data formats

- Browser Extensible Data (BED) `.bed`
    * Wikipedia [\*link\*](https://en.wikipedia.org/wiki/BED_(file_format))
- Wiggle Track (WIG) `.wig`
    * Docs [\*link\*](https://genome.ucsc.edu/goldenPath/help/wiggle.html)
    * Display of dense, continuous data such as GC content, probability scores, and transcriptome data
- Multiple Alignment Format (MAF) `.maf`
    * "stores multiple alignments at the DNA level between entire genomes"
- Other supported data formats [\*link\*](https://genome.ucsc.edu/FAQ/FAQformat.html)


### BLAST

To align and compare DNA/protein sequences

- NCBI quick start guide [\*link\*](https://www.ncbi.nlm.nih.gov/books/NBK1734/)



