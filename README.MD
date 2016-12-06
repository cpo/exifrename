

# What is EXIF renamer

For cleaning up my photo collection, I wanted to create a small utility 
for renaming the files into a nicely structured directory layout.

The layout I use is the following:

```text
yyyy/mm/file.jpg
```

For example:

```text
    - 2015
        - 08
        - 09
        - 10
        - 11
        - 12
        - ...
    - 2016
        - 08
        - 09
        - 10
        - 11
        - 12
        - ...
```

When you invoke the program inside a specific folder, it will rename all the files it encounters recursively.

#Usage
The binary can be started with the following options:

```text
$ exifrename --help
Usage of exifrename:
  -template string
    	Template to use, e.g. {{.Format ("2006-01-02")}}-{{.Format ("2006-01-02 03:04:05"}} (default "{{.Format (\"2006/01\")}}/{{.Filename}}")
  -test
    	Do not execute changes (default true)
```

exifrename works in 'test mode' by default. It will not change / rename any files in this mode.

#Advanced usage: the template

The output filename / directory structure `exifrename` uses, can be overridden by specifying the `--template` commandline option.
The default value for the `--template` option is a golang text/template expression. 

The root of the context contains an object which contains:

```golang
type data struct {
	Filename string
	Fullname string
	Exif     *exif.Exif
}
```

These fields can be used inside the template string. The Exif field contains a structure which is specific for the camera which recorded the image.

To find out which information is inside your images, you can execute the following command:

```bash
$ exifrename --template="{{.}}"
```

The output will be something like:

```text
Filename: sample1.tif
Fullname: ./vendor/github.com/rwcarlsen/goexif/tiff/sample1.tif
Exif: {
    BitsPerSample: 1
    PhotometricInterpretation: 0
    ImageDescription: "converted PBM file"
    XResolution: "2000000/10000"
    PlanarConfiguration: 1
    ResolutionUnit: 2
    ImageWidth: 1728
    Compression: 4
    Orientation: 1
    SamplesPerPixel: 1
    YResolution: "2000000/10000"
    ImageLength: 2376
    DateTime: "2014:09:01 15:03:47"
}
```

Example to use these fields:

```bash
$ exifrename --template="{{.Exif.Orientation}}/{{.Filename}}"
```

This will create an output structure per orientation, so all the landscape images in one folder, all the portrait images in another folder.

