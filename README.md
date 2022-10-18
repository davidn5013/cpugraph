# cpugraph 

Package for CpuUsage as asciigraph

Thanks to

[guptaohit](https://github.com/guptarohit/asciigraph)

[shirou](https://github.com/shirou/gopsutil)

## Usage as command

Command name is cpug and it look like this measure 30 times then start over to stop press ctrl-c

```
> cpug
 7.00 ┤ ╭╮
 6.00 ┤ ││
 5.00 ┤ │╰╮
 4.00 ┤ │ │           ╭╮        ╭╮
 3.00 ┤ │ │ ╭╮        ││        ││
 2.00 ┼╮│ │ ││        ││        │╰─╮
 1.00 ┤╰╯ ╰─╯╰───╮╭╮╭╮│╰╮ ╭──╮╭─╯  │
 0.00 ┤          ╰╯╰╯╰╯ ╰─╯  ╰╯    ╰
```

### Install

```bash
git clone https://github.com/davidn5013/cpugraph
cd cpugraph
go install ./...
```

After install the repo is not need cpug is go bin catalog

## Usage as module

```
go get github.com/davidn5013/cpugraph
```

then you add one or more of the func in your code like 

```
cpugraph.GetCpuUsage() 
cpugraph.PlotCGrap()
````

then run tidy in shell 

```
go mod tidy 
```

and then save the go file and the import pops up


