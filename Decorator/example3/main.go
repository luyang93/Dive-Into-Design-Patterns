package main

import "fmt"

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

type FileDataSource struct {
	filename string
	data     string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (f *FileDataSource) WriteData(data string) {
	fmt.Println("Writing data to", f.filename)
	f.data = data
}

func (f *FileDataSource) ReadData() string {
	fmt.Println("Reading data from", f.filename)
	return f.data
}

type DataSourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{wrappee: source}
}

func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() string {
	return d.wrappee.ReadData()
}

type EncryptionDecorator struct {
	*DataSourceDecorator
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{NewDataSourceDecorator(source)}
}

func (ed *EncryptionDecorator) WriteData(data string) {
	fmt.Println("encrypted")
	ed.DataSourceDecorator.WriteData(data)
}

func (ed *EncryptionDecorator) ReadData() string {
	fmt.Println("decrypted")
	return ed.DataSourceDecorator.ReadData()
}

type CompressionDecorator struct {
	*DataSourceDecorator
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{NewDataSourceDecorator(source)}
}

func (cd *CompressionDecorator) WriteData(data string) {
	fmt.Println("compressed")
	cd.DataSourceDecorator.WriteData(data)
}

func (cd *CompressionDecorator) ReadData() string {
	fmt.Println("decompressed")
	return cd.DataSourceDecorator.ReadData()
}

func main() {
	var source DataSource = NewFileDataSource("file.dat")
	source.WriteData("salaryRecords")
	fmt.Println(source.ReadData())
	fmt.Println("=======")

	source = NewDataSourceDecorator(source)
	source.WriteData("salaryRecords")
	fmt.Println(source.ReadData())
	fmt.Println("=======")

	source = NewCompressionDecorator(source)
	source.WriteData("salaryRecords")
	fmt.Println(source.ReadData())
	fmt.Println("=======")

	source = NewEncryptionDecorator(source)
	source.WriteData("salaryRecords")
	fmt.Println(source.ReadData())
	fmt.Println("=======")
}
